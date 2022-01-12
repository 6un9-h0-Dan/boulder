package responder

import (
	"bytes"
	"context"
	"crypto"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/letsencrypt/boulder/core"
	"github.com/letsencrypt/boulder/issuance"
	blog "github.com/letsencrypt/boulder/log"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ocsp"
)

type responderID struct {
	nameHash []byte
	keyHash  []byte
}

type filterSource struct {
	wrapped        Source
	hashAlgorithm  crypto.Hash
	issuers        map[issuance.IssuerNameID]responderID
	serialPrefixes []string
	counter        *prometheus.CounterVec
	log            blog.Logger
}

// NewFilterSource returns a filterSource which performs various checks on the
// OCSP requests sent to the wrapped Source, and the OCSP responses returned
// by it.
func NewFilterSource(issuerCerts []*issuance.Certificate, serialPrefixes []string, wrapped Source, stats prometheus.Registerer, log blog.Logger) (Source, error) {
	if len(issuerCerts) < 1 {
		return nil, errors.New("Filter must include at least 1 issuer cert")
	}
	var issuersByNameId map[issuance.IssuerNameID]responderID
	for _, issuerCert := range issuerCerts {
		keyHash := issuerCert.KeyHash()
		nameHash := issuerCert.NameHash()
		rid := responderID{
			keyHash:  keyHash[:],
			nameHash: nameHash[:],
		}
		issuersByNameId[issuerCert.NameID()] = rid
	}
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "ocsp_filter_responses",
		Help: "Count of OCSP requests/responses by action taken by the filter",
	}, []string{"result"})
	return &filterSource{
		wrapped:        wrapped,
		hashAlgorithm:  crypto.SHA1,
		issuers:        issuersByNameId,
		serialPrefixes: serialPrefixes,
		counter:        counter,
		log:            log,
	}, nil
}

// Response implements the Source interface. It checks the incoming request
// to ensure that we want to handle it, fetches the response from the wrapped
// Source, and checks that the response matches the request.
func (src *filterSource) Response(ctx context.Context, req *ocsp.Request) (*Response, error) {
	iss, err := src.checkRequest(req)
	if err != nil {
		src.log.Debugf("Not responding to filtered OCSP request: %s", err.Error())
		src.counter.WithLabelValues("request_filtered").Inc()
		return nil, err
	}

	resp, err := src.wrapped.Response(ctx, req)
	if err != nil {
		src.counter.WithLabelValues("wrapped_error").Inc()
		return nil, err
	}

	err = src.checkResponse(iss, resp)
	if err != nil {
		src.log.Warningf("OCSP Response not sent (issuer and serial mismatch) for CA=%s, Serial=%s", hex.EncodeToString(req.IssuerKeyHash), core.SerialToString(req.SerialNumber))
		src.counter.WithLabelValues("response_filtered").Inc()
		return nil, err
	}

	src.counter.WithLabelValues("success").Inc()
	return resp, nil
}

// checkRequest returns a descriptive error if the request does not satisfy any of
// the requirements of an OCSP request, or nil if the request should be handled.
// If the request passes all checks, then checkRequest returns the unique id of
// the issuer cert specified in the request.
func (src *filterSource) checkRequest(req *ocsp.Request) (issuance.IssuerNameID, error) {
	if req.HashAlgorithm != src.hashAlgorithm {
		return 0, fmt.Errorf("unsupported issuer key/name hash algorithm %s: %w", req.HashAlgorithm, ErrNotFound)
	}

	if len(src.serialPrefixes) > 0 {
		serialString := core.SerialToString(req.SerialNumber)
		match := false
		for _, prefix := range src.serialPrefixes {
			if strings.HasPrefix(serialString, prefix) {
				match = true
				break
			}
		}
		if !match {
			return 0, fmt.Errorf("unrecognized serial prefix: %w", ErrNotFound)
		}
	}

	for nameID, rid := range src.issuers {
		if bytes.Equal(req.IssuerKeyHash, rid.keyHash) {
			return nameID, nil
		}
	}
	return 0, fmt.Errorf("unrecognized issuer key hash %s: %w", hex.EncodeToString(req.IssuerKeyHash), ErrNotFound)
}

// checkResponse returns nil if the ocsp response was generated by the same
// issuer as was identified in the request, and an error otherwise. This filters
// out, for example, responses which are for a serial that we issued, but from a
// different issuer than that contained in the request.
func (src *filterSource) checkResponse(iss issuance.IssuerNameID, resp *Response) error {
	rid, ok := src.issuers[iss]
	if !ok {
		return fmt.Errorf("unrecognized issuer name ID: %d", iss)
	}
	if !bytes.Equal(rid.keyHash, resp.ResponderKeyHash) {
		return fmt.Errorf("responder does not match requested issuer")
	}
	return nil
}
