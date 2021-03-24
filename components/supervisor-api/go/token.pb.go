// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: token.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenReuse int32

const (
	// REUSE_NEVER means the token can never be re-used.
	// This mode only makes sense when providing a token in response to a request.
	TokenReuse_REUSE_NEVER TokenReuse = 0
	// REUSE_EXACTLY means the token can only be reused when the requested scopes
	// exactly match those of the token.
	TokenReuse_REUSE_EXACTLY TokenReuse = 1
	// REUSE_WHEN_POSSIBLE means the token can be reused when the requested scopes
	// are a subset of the token's scopes.
	TokenReuse_REUSE_WHEN_POSSIBLE TokenReuse = 2
)

// Enum value maps for TokenReuse.
var (
	TokenReuse_name = map[int32]string{
		0: "REUSE_NEVER",
		1: "REUSE_EXACTLY",
		2: "REUSE_WHEN_POSSIBLE",
	}
	TokenReuse_value = map[string]int32{
		"REUSE_NEVER":         0,
		"REUSE_EXACTLY":       1,
		"REUSE_WHEN_POSSIBLE": 2,
	}
)

func (x TokenReuse) Enum() *TokenReuse {
	p := new(TokenReuse)
	*p = x
	return p
}

func (x TokenReuse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenReuse) Descriptor() protoreflect.EnumDescriptor {
	return file_token_proto_enumTypes[0].Descriptor()
}

func (TokenReuse) Type() protoreflect.EnumType {
	return &file_token_proto_enumTypes[0]
}

func (x TokenReuse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenReuse.Descriptor instead.
func (TokenReuse) EnumDescriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{0}
}

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host        string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Scope       []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Kind        string   `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GetTokenRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *GetTokenRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetTokenRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type GetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//* The username of the account associated with the token.
	User  string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Scope []string `protobuf:"bytes,3,rep,name=scope,proto3" json:"scope,omitempty"`
}

func (x *GetTokenResponse) Reset() {
	*x = GetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResponse) ProtoMessage() {}

func (x *GetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResponse.ProtoReflect.Descriptor instead.
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTokenResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GetTokenResponse) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

type SetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host       string                 `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Scope      []string               `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	Token      string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ExpiryDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	Reuse      TokenReuse             `protobuf:"varint,5,opt,name=reuse,proto3,enum=supervisor.TokenReuse" json:"reuse,omitempty"`
	Kind       string                 `protobuf:"bytes,6,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *SetTokenRequest) Reset() {
	*x = SetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTokenRequest) ProtoMessage() {}

func (x *SetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTokenRequest.ProtoReflect.Descriptor instead.
func (*SetTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{2}
}

func (x *SetTokenRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SetTokenRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *SetTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SetTokenRequest) GetExpiryDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiryDate
	}
	return nil
}

func (x *SetTokenRequest) GetReuse() TokenReuse {
	if x != nil {
		return x.Reuse
	}
	return TokenReuse_REUSE_NEVER
}

func (x *SetTokenRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type SetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetTokenResponse) Reset() {
	*x = SetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTokenResponse) ProtoMessage() {}

func (x *SetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTokenResponse.ProtoReflect.Descriptor instead.
func (*SetTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{3}
}

type ClearTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Token:
	//	*ClearTokenRequest_Value
	//	*ClearTokenRequest_All
	Token isClearTokenRequest_Token `protobuf_oneof:"token"`
	Kind  string                    `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *ClearTokenRequest) Reset() {
	*x = ClearTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearTokenRequest) ProtoMessage() {}

func (x *ClearTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearTokenRequest.ProtoReflect.Descriptor instead.
func (*ClearTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{4}
}

func (m *ClearTokenRequest) GetToken() isClearTokenRequest_Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (x *ClearTokenRequest) GetValue() string {
	if x, ok := x.GetToken().(*ClearTokenRequest_Value); ok {
		return x.Value
	}
	return ""
}

func (x *ClearTokenRequest) GetAll() bool {
	if x, ok := x.GetToken().(*ClearTokenRequest_All); ok {
		return x.All
	}
	return false
}

func (x *ClearTokenRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type isClearTokenRequest_Token interface {
	isClearTokenRequest_Token()
}

type ClearTokenRequest_Value struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3,oneof"`
}

type ClearTokenRequest_All struct {
	All bool `protobuf:"varint,2,opt,name=all,proto3,oneof"`
}

func (*ClearTokenRequest_Value) isClearTokenRequest_Token() {}

func (*ClearTokenRequest_All) isClearTokenRequest_Token() {}

type ClearTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearTokenResponse) Reset() {
	*x = ClearTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearTokenResponse) ProtoMessage() {}

func (x *ClearTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearTokenResponse.ProtoReflect.Descriptor instead.
func (*ClearTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{5}
}

type ProvideTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ProvideTokenRequest_Registration
	//	*ProvideTokenRequest_Answer
	Message isProvideTokenRequest_Message `protobuf_oneof:"message"`
}

func (x *ProvideTokenRequest) Reset() {
	*x = ProvideTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvideTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvideTokenRequest) ProtoMessage() {}

func (x *ProvideTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvideTokenRequest.ProtoReflect.Descriptor instead.
func (*ProvideTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{6}
}

func (m *ProvideTokenRequest) GetMessage() isProvideTokenRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ProvideTokenRequest) GetRegistration() *ProvideTokenRequest_RegisterProvider {
	if x, ok := x.GetMessage().(*ProvideTokenRequest_Registration); ok {
		return x.Registration
	}
	return nil
}

func (x *ProvideTokenRequest) GetAnswer() *SetTokenRequest {
	if x, ok := x.GetMessage().(*ProvideTokenRequest_Answer); ok {
		return x.Answer
	}
	return nil
}

type isProvideTokenRequest_Message interface {
	isProvideTokenRequest_Message()
}

type ProvideTokenRequest_Registration struct {
	Registration *ProvideTokenRequest_RegisterProvider `protobuf:"bytes,1,opt,name=registration,proto3,oneof"`
}

type ProvideTokenRequest_Answer struct {
	Answer *SetTokenRequest `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

func (*ProvideTokenRequest_Registration) isProvideTokenRequest_Message() {}

func (*ProvideTokenRequest_Answer) isProvideTokenRequest_Message() {}

type ProvideTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *GetTokenRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *ProvideTokenResponse) Reset() {
	*x = ProvideTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvideTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvideTokenResponse) ProtoMessage() {}

func (x *ProvideTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvideTokenResponse.ProtoReflect.Descriptor instead.
func (*ProvideTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{7}
}

func (x *ProvideTokenResponse) GetRequest() *GetTokenRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type ProvideTokenRequest_RegisterProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *ProvideTokenRequest_RegisterProvider) Reset() {
	*x = ProvideTokenRequest_RegisterProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvideTokenRequest_RegisterProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvideTokenRequest_RegisterProvider) ProtoMessage() {}

func (x *ProvideTokenRequest_RegisterProvider) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvideTokenRequest_RegisterProvider.ProtoReflect.Descriptor instead.
func (*ProvideTokenRequest_RegisterProvider) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ProvideTokenRequest_RegisterProvider) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_token_proto protoreflect.FileDescriptor

var file_token_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x52, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22,
	0xd0, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x0a, 0x05, 0x72, 0x65, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x75, 0x73, 0x65, 0x52, 0x05, 0x72, 0x65, 0x75, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x11, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x13, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x1a, 0x26, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2a, 0x49, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x75, 0x73,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x55, 0x53, 0x45, 0x5f, 0x4e, 0x45, 0x56, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x55, 0x53, 0x45, 0x5f, 0x45, 0x58, 0x41, 0x43,
	0x54, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x55, 0x53, 0x45, 0x5f, 0x57,
	0x48, 0x45, 0x4e, 0x5f, 0x50, 0x4f, 0x53, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x32, 0xdb,
	0x03, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x6b, 0x69, 0x6e, 0x64, 0x7d,
	0x2f, 0x7b, 0x68, 0x6f, 0x73, 0x74, 0x7d, 0x2f, 0x7b, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x7d, 0x12,
	0x69, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x6b, 0x69, 0x6e, 0x64, 0x7d,
	0x2f, 0x7b, 0x68, 0x6f, 0x73, 0x74, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x0a, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43,
	0x2a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x6b, 0x69, 0x6e,
	0x64, 0x7d, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x5a, 0x27, 0x2a, 0x25, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x6b, 0x69, 0x6e, 0x64, 0x7d, 0x2f, 0x63,
	0x6c, 0x65, 0x61, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x3d, 0x74, 0x72,
	0x75, 0x65, 0x7d, 0x12, 0x57, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_token_proto_rawDescOnce sync.Once
	file_token_proto_rawDescData = file_token_proto_rawDesc
)

func file_token_proto_rawDescGZIP() []byte {
	file_token_proto_rawDescOnce.Do(func() {
		file_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_proto_rawDescData)
	})
	return file_token_proto_rawDescData
}

var file_token_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_token_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_token_proto_goTypes = []interface{}{
	(TokenReuse)(0),                              // 0: supervisor.TokenReuse
	(*GetTokenRequest)(nil),                      // 1: supervisor.GetTokenRequest
	(*GetTokenResponse)(nil),                     // 2: supervisor.GetTokenResponse
	(*SetTokenRequest)(nil),                      // 3: supervisor.SetTokenRequest
	(*SetTokenResponse)(nil),                     // 4: supervisor.SetTokenResponse
	(*ClearTokenRequest)(nil),                    // 5: supervisor.ClearTokenRequest
	(*ClearTokenResponse)(nil),                   // 6: supervisor.ClearTokenResponse
	(*ProvideTokenRequest)(nil),                  // 7: supervisor.ProvideTokenRequest
	(*ProvideTokenResponse)(nil),                 // 8: supervisor.ProvideTokenResponse
	(*ProvideTokenRequest_RegisterProvider)(nil), // 9: supervisor.ProvideTokenRequest.RegisterProvider
	(*timestamppb.Timestamp)(nil),                // 10: google.protobuf.Timestamp
}
var file_token_proto_depIdxs = []int32{
	10, // 0: supervisor.SetTokenRequest.expiry_date:type_name -> google.protobuf.Timestamp
	0,  // 1: supervisor.SetTokenRequest.reuse:type_name -> supervisor.TokenReuse
	9,  // 2: supervisor.ProvideTokenRequest.registration:type_name -> supervisor.ProvideTokenRequest.RegisterProvider
	3,  // 3: supervisor.ProvideTokenRequest.answer:type_name -> supervisor.SetTokenRequest
	1,  // 4: supervisor.ProvideTokenResponse.request:type_name -> supervisor.GetTokenRequest
	1,  // 5: supervisor.TokenService.GetToken:input_type -> supervisor.GetTokenRequest
	3,  // 6: supervisor.TokenService.SetToken:input_type -> supervisor.SetTokenRequest
	5,  // 7: supervisor.TokenService.ClearToken:input_type -> supervisor.ClearTokenRequest
	7,  // 8: supervisor.TokenService.ProvideToken:input_type -> supervisor.ProvideTokenRequest
	2,  // 9: supervisor.TokenService.GetToken:output_type -> supervisor.GetTokenResponse
	4,  // 10: supervisor.TokenService.SetToken:output_type -> supervisor.SetTokenResponse
	6,  // 11: supervisor.TokenService.ClearToken:output_type -> supervisor.ClearTokenResponse
	8,  // 12: supervisor.TokenService.ProvideToken:output_type -> supervisor.ProvideTokenResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_token_proto_init() }
func file_token_proto_init() {
	if File_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenRequest); i {
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
		file_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenResponse); i {
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
		file_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTokenRequest); i {
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
		file_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTokenResponse); i {
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
		file_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearTokenRequest); i {
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
		file_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearTokenResponse); i {
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
		file_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvideTokenRequest); i {
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
		file_token_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvideTokenResponse); i {
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
		file_token_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvideTokenRequest_RegisterProvider); i {
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
	file_token_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ClearTokenRequest_Value)(nil),
		(*ClearTokenRequest_All)(nil),
	}
	file_token_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ProvideTokenRequest_Registration)(nil),
		(*ProvideTokenRequest_Answer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_token_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_proto_goTypes,
		DependencyIndexes: file_token_proto_depIdxs,
		EnumInfos:         file_token_proto_enumTypes,
		MessageInfos:      file_token_proto_msgTypes,
	}.Build()
	File_token_proto = out.File
	file_token_proto_rawDesc = nil
	file_token_proto_goTypes = nil
	file_token_proto_depIdxs = nil
}
