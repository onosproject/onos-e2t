//
//SPDX-FileCopyrightText: 2022-present Intel Corporation
//
//SPDX-License-Identifier: Apache-2.0

////////////////////// xnap-commondatatypes.proto //////////////////////
// Protobuf generated from /xnap_v1.asn1 by asn1c-0.9.29
// XnAP-CommonDataTypes { itu-t(0) identified-organization(4) etsi(0) mobileDomain(0) ngran-access(22) modules(3) xnap(2) version1(1) xnap-CommonDataTypes(3) }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.4
// source: api/xnap/v1/xnap_commondatatypes.proto

package xnapcommondatatypesv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// enumerated from xnap_v1.asn1:8675
type Criticality int32

const (
	Criticality_CRITICALITY_REJECT Criticality = 0
	Criticality_CRITICALITY_IGNORE Criticality = 1
	Criticality_CRITICALITY_NOTIFY Criticality = 2
)

// Enum value maps for Criticality.
var (
	Criticality_name = map[int32]string{
		0: "CRITICALITY_REJECT",
		1: "CRITICALITY_IGNORE",
		2: "CRITICALITY_NOTIFY",
	}
	Criticality_value = map[string]int32{
		"CRITICALITY_REJECT": 0,
		"CRITICALITY_IGNORE": 1,
		"CRITICALITY_NOTIFY": 2,
	}
)

func (x Criticality) Enum() *Criticality {
	p := new(Criticality)
	*p = x
	return p
}

func (x Criticality) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Criticality) Descriptor() protoreflect.EnumDescriptor {
	return file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes[0].Descriptor()
}

func (Criticality) Type() protoreflect.EnumType {
	return &file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes[0]
}

func (x Criticality) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Criticality.Descriptor instead.
func (Criticality) EnumDescriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{0}
}

// enumerated from xnap_v1.asn1:8677
type Presence int32

const (
	Presence_PRESENCE_OPTIONAL    Presence = 0
	Presence_PRESENCE_CONDITIONAL Presence = 1
	Presence_PRESENCE_MANDATORY   Presence = 2
)

// Enum value maps for Presence.
var (
	Presence_name = map[int32]string{
		0: "PRESENCE_OPTIONAL",
		1: "PRESENCE_CONDITIONAL",
		2: "PRESENCE_MANDATORY",
	}
	Presence_value = map[string]int32{
		"PRESENCE_OPTIONAL":    0,
		"PRESENCE_CONDITIONAL": 1,
		"PRESENCE_MANDATORY":   2,
	}
)

func (x Presence) Enum() *Presence {
	p := new(Presence)
	*p = x
	return p
}

func (x Presence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Presence) Descriptor() protoreflect.EnumDescriptor {
	return file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes[1].Descriptor()
}

func (Presence) Type() protoreflect.EnumType {
	return &file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes[1]
}

func (x Presence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Presence.Descriptor instead.
func (Presence) EnumDescriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{1}
}

// enumerated from xnap_v1.asn1:8690
type TriggeringMessage int32

const (
	TriggeringMessage_TRIGGERING_MESSAGE_INITIATING_MESSAGE   TriggeringMessage = 0
	TriggeringMessage_TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME   TriggeringMessage = 1
	TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME TriggeringMessage = 2
)

// Enum value maps for TriggeringMessage.
var (
	TriggeringMessage_name = map[int32]string{
		0: "TRIGGERING_MESSAGE_INITIATING_MESSAGE",
		1: "TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME",
		2: "TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME",
	}
	TriggeringMessage_value = map[string]int32{
		"TRIGGERING_MESSAGE_INITIATING_MESSAGE":   0,
		"TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME":   1,
		"TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME": 2,
	}
)

func (x TriggeringMessage) Enum() *TriggeringMessage {
	p := new(TriggeringMessage)
	*p = x
	return p
}

func (x TriggeringMessage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TriggeringMessage) Descriptor() protoreflect.EnumDescriptor {
	return file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes[2].Descriptor()
}

func (TriggeringMessage) Type() protoreflect.EnumType {
	return &file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes[2]
}

func (x TriggeringMessage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TriggeringMessage.Descriptor instead.
func (TriggeringMessage) EnumDescriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{2}
}

// constant Integer from xnap_v1.asn1:8665
// {-}
type MaxPrivateIes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:65535,valueUB:65535,"
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MaxPrivateIes) Reset() {
	*x = MaxPrivateIes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxPrivateIes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxPrivateIes) ProtoMessage() {}

func (x *MaxPrivateIes) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxPrivateIes.ProtoReflect.Descriptor instead.
func (*MaxPrivateIes) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{0}
}

func (x *MaxPrivateIes) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// constant Integer from xnap_v1.asn1:8666
// {-}
type MaxProtocolExtensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:65535,valueUB:65535,"
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MaxProtocolExtensions) Reset() {
	*x = MaxProtocolExtensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxProtocolExtensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxProtocolExtensions) ProtoMessage() {}

func (x *MaxProtocolExtensions) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxProtocolExtensions.ProtoReflect.Descriptor instead.
func (*MaxProtocolExtensions) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{1}
}

func (x *MaxProtocolExtensions) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// constant Integer from xnap_v1.asn1:8667
// {-}
type MaxProtocolIes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:65535,valueUB:65535,"
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MaxProtocolIes) Reset() {
	*x = MaxProtocolIes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxProtocolIes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxProtocolIes) ProtoMessage() {}

func (x *MaxProtocolIes) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxProtocolIes.ProtoReflect.Descriptor instead.
func (*MaxProtocolIes) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{2}
}

func (x *MaxProtocolIes) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// sequence from xnap_v1.asn1:8680
// {PrivateIE-ID}
type PrivateIeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from xnap_v1.asn1:8680
	//
	// Types that are assignable to PrivateIeId:
	//	*PrivateIeID_Local
	//	*PrivateIeID_Global
	PrivateIeId isPrivateIeID_PrivateIeId `protobuf_oneof:"private_ie_id"`
}

func (x *PrivateIeID) Reset() {
	*x = PrivateIeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateIeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateIeID) ProtoMessage() {}

func (x *PrivateIeID) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateIeID.ProtoReflect.Descriptor instead.
func (*PrivateIeID) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{3}
}

func (m *PrivateIeID) GetPrivateIeId() isPrivateIeID_PrivateIeId {
	if m != nil {
		return m.PrivateIeId
	}
	return nil
}

func (x *PrivateIeID) GetLocal() int32 {
	if x, ok := x.GetPrivateIeId().(*PrivateIeID_Local); ok {
		return x.Local
	}
	return 0
}

func (x *PrivateIeID) GetGlobal() string {
	if x, ok := x.GetPrivateIeId().(*PrivateIeID_Global); ok {
		return x.Global
	}
	return ""
}

type isPrivateIeID_PrivateIeId interface {
	isPrivateIeID_PrivateIeId()
}

type PrivateIeID_Local struct {
	// @inject_tag: aper:"choiceIdx:1,valueLB:0,valueUB:65535,"
	Local int32 `protobuf:"varint,1,opt,name=local,proto3,oneof"`
}

type PrivateIeID_Global struct {
	// @inject_tag: aper:"choiceIdx:2,"
	Global string `protobuf:"bytes,2,opt,name=global,proto3,oneof"`
}

func (*PrivateIeID_Local) isPrivateIeID_PrivateIeId() {}

func (*PrivateIeID_Global) isPrivateIeID_PrivateIeId() {}

// range of Integer from xnap_v1.asn1:8684
// {ProcedureCode}
type ProcedureCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:0,valueUB:255,"
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProcedureCode) Reset() {
	*x = ProcedureCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureCode) ProtoMessage() {}

func (x *ProcedureCode) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureCode.ProtoReflect.Descriptor instead.
func (*ProcedureCode) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{4}
}

func (x *ProcedureCode) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// range of Integer from xnap_v1.asn1:8687
// {ProtocolIE-ID}
type ProtocolIeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:0,valueUB:65535,"
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeID) Reset() {
	*x = ProtocolIeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeID) ProtoMessage() {}

func (x *ProtocolIeID) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeID.ProtoReflect.Descriptor instead.
func (*ProtocolIeID) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP(), []int{5}
}

func (x *ProtocolIeID) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_api_xnap_v1_xnap_commondatatypes_proto protoreflect.FileDescriptor

var file_api_xnap_v1_xnap_commondatatypes_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x78, 0x6e,
	0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76,
	0x31, 0x1a, 0x1a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61,
	0x73, 0x6e, 0x31, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30,
	0x0a, 0x0d, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x1a, 0x04, 0x08, 0xff, 0xff, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x38, 0x0a, 0x15, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x08,
	0xff, 0xff, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x0e, 0x4d, 0x61,
	0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x1a, 0x04, 0x08, 0xff, 0xff, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x50, 0x0a,
	0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x05,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x42, 0x0f,
	0x0a, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x22,
	0x31, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x1a, 0x05, 0x18, 0xff, 0x01, 0x28, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x31, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65,
	0x49, 0x44, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x1a, 0x06, 0x18, 0xff, 0xff, 0x03, 0x28, 0x00, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x55, 0x0a, 0x0b, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x47, 0x4e, 0x4f,
	0x52, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x10, 0x02, 0x2a, 0x53, 0x0a, 0x08,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x45, 0x53,
	0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x44,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x45,
	0x53, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x44, 0x41, 0x54, 0x4f, 0x52, 0x59, 0x10,
	0x02, 0x2a, 0x96, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x54, 0x52, 0x49, 0x47, 0x47,
	0x45, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x4e,
	0x49, 0x54, 0x49, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x49, 0x4e, 0x47,
	0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x46, 0x55, 0x4c, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x2b, 0x0a,
	0x27, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c,
	0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x10, 0x02, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x2d, 0x65, 0x32, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x78, 0x6e, 0x61, 0x70, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x78,
	0x6e, 0x61, 0x70, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_xnap_v1_xnap_commondatatypes_proto_rawDescOnce sync.Once
	file_api_xnap_v1_xnap_commondatatypes_proto_rawDescData = file_api_xnap_v1_xnap_commondatatypes_proto_rawDesc
)

func file_api_xnap_v1_xnap_commondatatypes_proto_rawDescGZIP() []byte {
	file_api_xnap_v1_xnap_commondatatypes_proto_rawDescOnce.Do(func() {
		file_api_xnap_v1_xnap_commondatatypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_xnap_v1_xnap_commondatatypes_proto_rawDescData)
	})
	return file_api_xnap_v1_xnap_commondatatypes_proto_rawDescData
}

var file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_xnap_v1_xnap_commondatatypes_proto_goTypes = []interface{}{
	(Criticality)(0),              // 0: xnap.v1.Criticality
	(Presence)(0),                 // 1: xnap.v1.Presence
	(TriggeringMessage)(0),        // 2: xnap.v1.TriggeringMessage
	(*MaxPrivateIes)(nil),         // 3: xnap.v1.MaxPrivateIes
	(*MaxProtocolExtensions)(nil), // 4: xnap.v1.MaxProtocolExtensions
	(*MaxProtocolIes)(nil),        // 5: xnap.v1.MaxProtocolIes
	(*PrivateIeID)(nil),           // 6: xnap.v1.PrivateIeID
	(*ProcedureCode)(nil),         // 7: xnap.v1.ProcedureCode
	(*ProtocolIeID)(nil),          // 8: xnap.v1.ProtocolIeID
}
var file_api_xnap_v1_xnap_commondatatypes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_xnap_v1_xnap_commondatatypes_proto_init() }
func file_api_xnap_v1_xnap_commondatatypes_proto_init() {
	if File_api_xnap_v1_xnap_commondatatypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxPrivateIes); i {
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
		file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxProtocolExtensions); i {
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
		file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxProtocolIes); i {
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
		file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateIeID); i {
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
		file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureCode); i {
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
		file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeID); i {
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
	file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*PrivateIeID_Local)(nil),
		(*PrivateIeID_Global)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_xnap_v1_xnap_commondatatypes_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_xnap_v1_xnap_commondatatypes_proto_goTypes,
		DependencyIndexes: file_api_xnap_v1_xnap_commondatatypes_proto_depIdxs,
		EnumInfos:         file_api_xnap_v1_xnap_commondatatypes_proto_enumTypes,
		MessageInfos:      file_api_xnap_v1_xnap_commondatatypes_proto_msgTypes,
	}.Build()
	File_api_xnap_v1_xnap_commondatatypes_proto = out.File
	file_api_xnap_v1_xnap_commondatatypes_proto_rawDesc = nil
	file_api_xnap_v1_xnap_commondatatypes_proto_goTypes = nil
	file_api_xnap_v1_xnap_commondatatypes_proto_depIdxs = nil
}
