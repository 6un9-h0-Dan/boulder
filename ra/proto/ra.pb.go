// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: ra.proto

package proto

import (
	proto "github.com/letsencrypt/boulder/core/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewAuthorizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authz *proto.Authorization `protobuf:"bytes,1,opt,name=authz,proto3" json:"authz,omitempty"`
	RegID int64                `protobuf:"varint,2,opt,name=regID,proto3" json:"regID,omitempty"`
}

func (x *NewAuthorizationRequest) Reset() {
	*x = NewAuthorizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAuthorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAuthorizationRequest) ProtoMessage() {}

func (x *NewAuthorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAuthorizationRequest.ProtoReflect.Descriptor instead.
func (*NewAuthorizationRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{0}
}

func (x *NewAuthorizationRequest) GetAuthz() *proto.Authorization {
	if x != nil {
		return x.Authz
	}
	return nil
}

func (x *NewAuthorizationRequest) GetRegID() int64 {
	if x != nil {
		return x.RegID
	}
	return 0
}

type NewCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Csr          []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	RegID        int64  `protobuf:"varint,2,opt,name=regID,proto3" json:"regID,omitempty"`
	IssuerNameID int64  `protobuf:"varint,3,opt,name=issuerNameID,proto3" json:"issuerNameID,omitempty"`
}

func (x *NewCertificateRequest) Reset() {
	*x = NewCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCertificateRequest) ProtoMessage() {}

func (x *NewCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCertificateRequest.ProtoReflect.Descriptor instead.
func (*NewCertificateRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{1}
}

func (x *NewCertificateRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

func (x *NewCertificateRequest) GetRegID() int64 {
	if x != nil {
		return x.RegID
	}
	return 0
}

func (x *NewCertificateRequest) GetIssuerNameID() int64 {
	if x != nil {
		return x.IssuerNameID
	}
	return 0
}

type UpdateRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base   *proto.Registration `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Update *proto.Registration `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
}

func (x *UpdateRegistrationRequest) Reset() {
	*x = UpdateRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegistrationRequest) ProtoMessage() {}

func (x *UpdateRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegistrationRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRegistrationRequest) GetBase() *proto.Registration {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UpdateRegistrationRequest) GetUpdate() *proto.Registration {
	if x != nil {
		return x.Update
	}
	return nil
}

type UpdateAuthorizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authz          *proto.Authorization `protobuf:"bytes,1,opt,name=authz,proto3" json:"authz,omitempty"`
	ChallengeIndex int64                `protobuf:"varint,2,opt,name=challengeIndex,proto3" json:"challengeIndex,omitempty"`
	Response       *proto.Challenge     `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateAuthorizationRequest) Reset() {
	*x = UpdateAuthorizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorizationRequest) ProtoMessage() {}

func (x *UpdateAuthorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorizationRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthorizationRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAuthorizationRequest) GetAuthz() *proto.Authorization {
	if x != nil {
		return x.Authz
	}
	return nil
}

func (x *UpdateAuthorizationRequest) GetChallengeIndex() int64 {
	if x != nil {
		return x.ChallengeIndex
	}
	return 0
}

func (x *UpdateAuthorizationRequest) GetResponse() *proto.Challenge {
	if x != nil {
		return x.Response
	}
	return nil
}

type PerformValidationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authz          *proto.Authorization `protobuf:"bytes,1,opt,name=authz,proto3" json:"authz,omitempty"`
	ChallengeIndex int64                `protobuf:"varint,2,opt,name=challengeIndex,proto3" json:"challengeIndex,omitempty"`
}

func (x *PerformValidationRequest) Reset() {
	*x = PerformValidationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerformValidationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformValidationRequest) ProtoMessage() {}

func (x *PerformValidationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformValidationRequest.ProtoReflect.Descriptor instead.
func (*PerformValidationRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{4}
}

func (x *PerformValidationRequest) GetAuthz() *proto.Authorization {
	if x != nil {
		return x.Authz
	}
	return nil
}

func (x *PerformValidationRequest) GetChallengeIndex() int64 {
	if x != nil {
		return x.ChallengeIndex
	}
	return 0
}

type RevokeCertificateWithRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert  []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	Code  int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	RegID int64  `protobuf:"varint,3,opt,name=regID,proto3" json:"regID,omitempty"`
}

func (x *RevokeCertificateWithRegRequest) Reset() {
	*x = RevokeCertificateWithRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeCertificateWithRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeCertificateWithRegRequest) ProtoMessage() {}

func (x *RevokeCertificateWithRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeCertificateWithRegRequest.ProtoReflect.Descriptor instead.
func (*RevokeCertificateWithRegRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{5}
}

func (x *RevokeCertificateWithRegRequest) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *RevokeCertificateWithRegRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RevokeCertificateWithRegRequest) GetRegID() int64 {
	if x != nil {
		return x.RegID
	}
	return 0
}

type AdministrativelyRevokeCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert         []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	Serial       string `protobuf:"bytes,4,opt,name=serial,proto3" json:"serial,omitempty"`
	Code         int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	AdminName    string `protobuf:"bytes,3,opt,name=adminName,proto3" json:"adminName,omitempty"`
	SkipBlockKey bool   `protobuf:"varint,5,opt,name=skipBlockKey,proto3" json:"skipBlockKey,omitempty"`
}

func (x *AdministrativelyRevokeCertificateRequest) Reset() {
	*x = AdministrativelyRevokeCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdministrativelyRevokeCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdministrativelyRevokeCertificateRequest) ProtoMessage() {}

func (x *AdministrativelyRevokeCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdministrativelyRevokeCertificateRequest.ProtoReflect.Descriptor instead.
func (*AdministrativelyRevokeCertificateRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{6}
}

func (x *AdministrativelyRevokeCertificateRequest) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *AdministrativelyRevokeCertificateRequest) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *AdministrativelyRevokeCertificateRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AdministrativelyRevokeCertificateRequest) GetAdminName() string {
	if x != nil {
		return x.AdminName
	}
	return ""
}

func (x *AdministrativelyRevokeCertificateRequest) GetSkipBlockKey() bool {
	if x != nil {
		return x.SkipBlockKey
	}
	return false
}

type NewOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrationID int64    `protobuf:"varint,1,opt,name=registrationID,proto3" json:"registrationID,omitempty"`
	Names          []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *NewOrderRequest) Reset() {
	*x = NewOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrderRequest) ProtoMessage() {}

func (x *NewOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrderRequest.ProtoReflect.Descriptor instead.
func (*NewOrderRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{7}
}

func (x *NewOrderRequest) GetRegistrationID() int64 {
	if x != nil {
		return x.RegistrationID
	}
	return 0
}

func (x *NewOrderRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type FinalizeOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *proto.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Csr   []byte       `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (x *FinalizeOrderRequest) Reset() {
	*x = FinalizeOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ra_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeOrderRequest) ProtoMessage() {}

func (x *FinalizeOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ra_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeOrderRequest.ProtoReflect.Descriptor instead.
func (*FinalizeOrderRequest) Descriptor() ([]byte, []int) {
	return file_ra_proto_rawDescGZIP(), []int{8}
}

func (x *FinalizeOrderRequest) GetOrder() *proto.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *FinalizeOrderRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

var File_ra_proto protoreflect.FileDescriptor

var file_ra_proto_rawDesc = []byte{
	0x0a, 0x08, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x72, 0x61, 0x1a, 0x15,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x17, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x67, 0x49, 0x44, 0x22, 0x63,
	0x0a, 0x15, 0x4e, 0x65, 0x77, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x67, 0x49, 0x44, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x22, 0x6f, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x26,
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6d, 0x0a, 0x18, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0x5f, 0x0a, 0x1f, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x67, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65,
	0x67, 0x49, 0x44, 0x22, 0xac, 0x01, 0x0a, 0x28, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x79, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x63, 0x65, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4b,
	0x65, 0x79, 0x22, 0x4f, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x73, 0x72,
	0x32, 0xb7, 0x06, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0f, 0x4e, 0x65,
	0x77, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x72, 0x61,
	0x2e, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x72, 0x61, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x11,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x2e, 0x72, 0x61, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x18, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52,
	0x65, 0x67, 0x12, 0x23, 0x2e, 0x72, 0x61, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x16, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x17, 0x44, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x21, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x79, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x72, 0x61, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x79, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x2e, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x72,
	0x61, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x74, 0x73, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x2f, 0x62, 0x6f, 0x75, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ra_proto_rawDescOnce sync.Once
	file_ra_proto_rawDescData = file_ra_proto_rawDesc
)

func file_ra_proto_rawDescGZIP() []byte {
	file_ra_proto_rawDescOnce.Do(func() {
		file_ra_proto_rawDescData = protoimpl.X.CompressGZIP(file_ra_proto_rawDescData)
	})
	return file_ra_proto_rawDescData
}

var file_ra_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ra_proto_goTypes = []interface{}{
	(*NewAuthorizationRequest)(nil),                  // 0: ra.NewAuthorizationRequest
	(*NewCertificateRequest)(nil),                    // 1: ra.NewCertificateRequest
	(*UpdateRegistrationRequest)(nil),                // 2: ra.UpdateRegistrationRequest
	(*UpdateAuthorizationRequest)(nil),               // 3: ra.UpdateAuthorizationRequest
	(*PerformValidationRequest)(nil),                 // 4: ra.PerformValidationRequest
	(*RevokeCertificateWithRegRequest)(nil),          // 5: ra.RevokeCertificateWithRegRequest
	(*AdministrativelyRevokeCertificateRequest)(nil), // 6: ra.AdministrativelyRevokeCertificateRequest
	(*NewOrderRequest)(nil),                          // 7: ra.NewOrderRequest
	(*FinalizeOrderRequest)(nil),                     // 8: ra.FinalizeOrderRequest
	(*proto.Authorization)(nil),                      // 9: core.Authorization
	(*proto.Registration)(nil),                       // 10: core.Registration
	(*proto.Challenge)(nil),                          // 11: core.Challenge
	(*proto.Order)(nil),                              // 12: core.Order
	(*proto.Certificate)(nil),                        // 13: core.Certificate
	(*emptypb.Empty)(nil),                            // 14: google.protobuf.Empty
}
var file_ra_proto_depIdxs = []int32{
	9,  // 0: ra.NewAuthorizationRequest.authz:type_name -> core.Authorization
	10, // 1: ra.UpdateRegistrationRequest.base:type_name -> core.Registration
	10, // 2: ra.UpdateRegistrationRequest.update:type_name -> core.Registration
	9,  // 3: ra.UpdateAuthorizationRequest.authz:type_name -> core.Authorization
	11, // 4: ra.UpdateAuthorizationRequest.response:type_name -> core.Challenge
	9,  // 5: ra.PerformValidationRequest.authz:type_name -> core.Authorization
	12, // 6: ra.FinalizeOrderRequest.order:type_name -> core.Order
	10, // 7: ra.RegistrationAuthority.NewRegistration:input_type -> core.Registration
	0,  // 8: ra.RegistrationAuthority.NewAuthorization:input_type -> ra.NewAuthorizationRequest
	1,  // 9: ra.RegistrationAuthority.NewCertificate:input_type -> ra.NewCertificateRequest
	2,  // 10: ra.RegistrationAuthority.UpdateRegistration:input_type -> ra.UpdateRegistrationRequest
	4,  // 11: ra.RegistrationAuthority.PerformValidation:input_type -> ra.PerformValidationRequest
	5,  // 12: ra.RegistrationAuthority.RevokeCertificateWithReg:input_type -> ra.RevokeCertificateWithRegRequest
	10, // 13: ra.RegistrationAuthority.DeactivateRegistration:input_type -> core.Registration
	9,  // 14: ra.RegistrationAuthority.DeactivateAuthorization:input_type -> core.Authorization
	6,  // 15: ra.RegistrationAuthority.AdministrativelyRevokeCertificate:input_type -> ra.AdministrativelyRevokeCertificateRequest
	7,  // 16: ra.RegistrationAuthority.NewOrder:input_type -> ra.NewOrderRequest
	8,  // 17: ra.RegistrationAuthority.FinalizeOrder:input_type -> ra.FinalizeOrderRequest
	10, // 18: ra.RegistrationAuthority.NewRegistration:output_type -> core.Registration
	9,  // 19: ra.RegistrationAuthority.NewAuthorization:output_type -> core.Authorization
	13, // 20: ra.RegistrationAuthority.NewCertificate:output_type -> core.Certificate
	10, // 21: ra.RegistrationAuthority.UpdateRegistration:output_type -> core.Registration
	9,  // 22: ra.RegistrationAuthority.PerformValidation:output_type -> core.Authorization
	14, // 23: ra.RegistrationAuthority.RevokeCertificateWithReg:output_type -> google.protobuf.Empty
	14, // 24: ra.RegistrationAuthority.DeactivateRegistration:output_type -> google.protobuf.Empty
	14, // 25: ra.RegistrationAuthority.DeactivateAuthorization:output_type -> google.protobuf.Empty
	14, // 26: ra.RegistrationAuthority.AdministrativelyRevokeCertificate:output_type -> google.protobuf.Empty
	12, // 27: ra.RegistrationAuthority.NewOrder:output_type -> core.Order
	12, // 28: ra.RegistrationAuthority.FinalizeOrder:output_type -> core.Order
	18, // [18:29] is the sub-list for method output_type
	7,  // [7:18] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_ra_proto_init() }
func file_ra_proto_init() {
	if File_ra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAuthorizationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegistrationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthorizationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerformValidationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeCertificateWithRegRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdministrativelyRevokeCertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ra_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ra_proto_goTypes,
		DependencyIndexes: file_ra_proto_depIdxs,
		MessageInfos:      file_ra_proto_msgTypes,
	}.Build()
	File_ra_proto = out.File
	file_ra_proto_rawDesc = nil
	file_ra_proto_goTypes = nil
	file_ra_proto_depIdxs = nil
}
