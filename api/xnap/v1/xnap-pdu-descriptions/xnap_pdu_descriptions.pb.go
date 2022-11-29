//
//SPDX-FileCopyrightText: 2022-present Intel Corporation
//
//SPDX-License-Identifier: Apache-2.0

////////////////////// xnap-pdu-descriptions.proto //////////////////////
// Protobuf generated from /xnap_v1.asn1 by asn1c-0.9.29
// XnAP-PDU-Descriptions { itu-t(0) identified-organization(4) etsi(0) mobileDomain(0) ngran-access(22) modules(3) xnap(2) version1(1) xnap-PDU-Descriptions(0) }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.4
// source: api/xnap/v1/xnap_pdu_descriptions.proto

package xnappdudescriptionsv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	xnap_commondatatypes "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-commondatatypes"
	_ "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-constants"
	_ "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-containers"
	xnap_pdu_contents "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-pdu-contents"
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

// sequence from xnap_v1.asn1:168
// @inject_tag: aper:"choiceExt"
// {XnAP-PDU}
type XnApPDu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from xnap_v1.asn1:168
	//
	// Types that are assignable to XnApPdu:
	//	*XnApPDu_InitiatingMessage
	//	*XnApPDu_SuccessfulOutcome
	//	*XnApPDu_UnsuccessfulOutcome
	XnApPdu isXnApPDu_XnApPdu `protobuf_oneof:"xn_ap_pdu"`
}

func (x *XnApPDu) Reset() {
	*x = XnApPDu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XnApPDu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XnApPDu) ProtoMessage() {}

func (x *XnApPDu) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XnApPDu.ProtoReflect.Descriptor instead.
func (*XnApPDu) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP(), []int{0}
}

func (m *XnApPDu) GetXnApPdu() isXnApPDu_XnApPdu {
	if m != nil {
		return m.XnApPdu
	}
	return nil
}

func (x *XnApPDu) GetInitiatingMessage() *InitiatingMessage {
	if x, ok := x.GetXnApPdu().(*XnApPDu_InitiatingMessage); ok {
		return x.InitiatingMessage
	}
	return nil
}

func (x *XnApPDu) GetSuccessfulOutcome() *SuccessfulOutcome {
	if x, ok := x.GetXnApPdu().(*XnApPDu_SuccessfulOutcome); ok {
		return x.SuccessfulOutcome
	}
	return nil
}

func (x *XnApPDu) GetUnsuccessfulOutcome() *UnsuccessfulOutcome {
	if x, ok := x.GetXnApPdu().(*XnApPDu_UnsuccessfulOutcome); ok {
		return x.UnsuccessfulOutcome
	}
	return nil
}

type isXnApPDu_XnApPdu interface {
	isXnApPDu_XnApPdu()
}

type XnApPDu_InitiatingMessage struct {
	// @inject_tag: aper:"choiceIdx:1,"
	InitiatingMessage *InitiatingMessage `protobuf:"bytes,1,opt,name=initiating_message,json=initiatingMessage,proto3,oneof"`
}

type XnApPDu_SuccessfulOutcome struct {
	// @inject_tag: aper:"choiceIdx:2,"
	SuccessfulOutcome *SuccessfulOutcome `protobuf:"bytes,2,opt,name=successful_outcome,json=successfulOutcome,proto3,oneof"`
}

type XnApPDu_UnsuccessfulOutcome struct {
	// @inject_tag: aper:"choiceIdx:3,"
	UnsuccessfulOutcome *UnsuccessfulOutcome `protobuf:"bytes,3,opt,name=unsuccessful_outcome,json=unsuccessfulOutcome,proto3,oneof"`
}

func (*XnApPDu_InitiatingMessage) isXnApPDu_XnApPdu() {}

func (*XnApPDu_SuccessfulOutcome) isXnApPDu_XnApPdu() {}

func (*XnApPDu_UnsuccessfulOutcome) isXnApPDu_XnApPdu() {}

// sequence from xnap_v1.asn1:175
// {InitiatingMessage}
type InitiatingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:0,valueUB:255,unique"
	ProcedureCode int32 `protobuf:"varint,1,opt,name=procedure_code,json=procedureCode,proto3" json:"procedure_code,omitempty"`
	// @inject_tag: aper:"valueLB:0,valueUB:2"
	Criticality xnap_commondatatypes.Criticality `protobuf:"varint,2,opt,name=criticality,proto3,enum=xnap.v1.Criticality" json:"criticality,omitempty"`
	// @inject_tag: aper:"canonicalOrder"
	Value *InitiatingMessageXnApElementaryProcedures `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InitiatingMessage) Reset() {
	*x = InitiatingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiatingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiatingMessage) ProtoMessage() {}

func (x *InitiatingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiatingMessage.ProtoReflect.Descriptor instead.
func (*InitiatingMessage) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP(), []int{1}
}

func (x *InitiatingMessage) GetProcedureCode() int32 {
	if x != nil {
		return x.ProcedureCode
	}
	return 0
}

func (x *InitiatingMessage) GetCriticality() xnap_commondatatypes.Criticality {
	if x != nil {
		return x.Criticality
	}
	return xnap_commondatatypes.Criticality(0)
}

func (x *InitiatingMessage) GetValue() *InitiatingMessageXnApElementaryProcedures {
	if x != nil {
		return x.Value
	}
	return nil
}

type InitiatingMessageXnApElementaryProcedures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ImValues:
	//	*InitiatingMessageXnApElementaryProcedures_XnSetupRequest
	ImValues isInitiatingMessageXnApElementaryProcedures_ImValues `protobuf_oneof:"im_values"`
}

func (x *InitiatingMessageXnApElementaryProcedures) Reset() {
	*x = InitiatingMessageXnApElementaryProcedures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiatingMessageXnApElementaryProcedures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiatingMessageXnApElementaryProcedures) ProtoMessage() {}

func (x *InitiatingMessageXnApElementaryProcedures) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiatingMessageXnApElementaryProcedures.ProtoReflect.Descriptor instead.
func (*InitiatingMessageXnApElementaryProcedures) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP(), []int{2}
}

func (m *InitiatingMessageXnApElementaryProcedures) GetImValues() isInitiatingMessageXnApElementaryProcedures_ImValues {
	if m != nil {
		return m.ImValues
	}
	return nil
}

func (x *InitiatingMessageXnApElementaryProcedures) GetXnSetupRequest() *xnap_pdu_contents.XnSetupRequest {
	if x, ok := x.GetImValues().(*InitiatingMessageXnApElementaryProcedures_XnSetupRequest); ok {
		return x.XnSetupRequest
	}
	return nil
}

type isInitiatingMessageXnApElementaryProcedures_ImValues interface {
	isInitiatingMessageXnApElementaryProcedures_ImValues()
}

type InitiatingMessageXnApElementaryProcedures_XnSetupRequest struct {
	// @inject_tag: aper:"valueExt"
	XnSetupRequest *xnap_pdu_contents.XnSetupRequest `protobuf:"bytes,1,opt,name=xn_setup_request,json=ric_subscription,proto3,oneof"` // ToDo - add the rest of the messages
}

func (*InitiatingMessageXnApElementaryProcedures_XnSetupRequest) isInitiatingMessageXnApElementaryProcedures_ImValues() {
}

// sequence from xnap_v1.asn1:181
// {SuccessfulOutcome}
type SuccessfulOutcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:0,valueUB:255,unique"
	ProcedureCode int32 `protobuf:"varint,1,opt,name=procedure_code,json=procedureCode,proto3" json:"procedure_code,omitempty"`
	// @inject_tag: aper:"valueLB:0,valueUB:2"
	Criticality xnap_commondatatypes.Criticality `protobuf:"varint,2,opt,name=criticality,proto3,enum=xnap.v1.Criticality" json:"criticality,omitempty"`
	// @inject_tag: aper:"canonicalOrder"
	Value *SuccessfulOutcomeXnApElementaryProcedures `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SuccessfulOutcome) Reset() {
	*x = SuccessfulOutcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessfulOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessfulOutcome) ProtoMessage() {}

func (x *SuccessfulOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessfulOutcome.ProtoReflect.Descriptor instead.
func (*SuccessfulOutcome) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP(), []int{3}
}

func (x *SuccessfulOutcome) GetProcedureCode() int32 {
	if x != nil {
		return x.ProcedureCode
	}
	return 0
}

func (x *SuccessfulOutcome) GetCriticality() xnap_commondatatypes.Criticality {
	if x != nil {
		return x.Criticality
	}
	return xnap_commondatatypes.Criticality(0)
}

func (x *SuccessfulOutcome) GetValue() *SuccessfulOutcomeXnApElementaryProcedures {
	if x != nil {
		return x.Value
	}
	return nil
}

type SuccessfulOutcomeXnApElementaryProcedures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SoValues:
	//	*SuccessfulOutcomeXnApElementaryProcedures_XnSetupResponse
	SoValues isSuccessfulOutcomeXnApElementaryProcedures_SoValues `protobuf_oneof:"so_values"`
}

func (x *SuccessfulOutcomeXnApElementaryProcedures) Reset() {
	*x = SuccessfulOutcomeXnApElementaryProcedures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessfulOutcomeXnApElementaryProcedures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessfulOutcomeXnApElementaryProcedures) ProtoMessage() {}

func (x *SuccessfulOutcomeXnApElementaryProcedures) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessfulOutcomeXnApElementaryProcedures.ProtoReflect.Descriptor instead.
func (*SuccessfulOutcomeXnApElementaryProcedures) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP(), []int{4}
}

func (m *SuccessfulOutcomeXnApElementaryProcedures) GetSoValues() isSuccessfulOutcomeXnApElementaryProcedures_SoValues {
	if m != nil {
		return m.SoValues
	}
	return nil
}

func (x *SuccessfulOutcomeXnApElementaryProcedures) GetXnSetupResponse() *xnap_pdu_contents.XnSetupResponse {
	if x, ok := x.GetSoValues().(*SuccessfulOutcomeXnApElementaryProcedures_XnSetupResponse); ok {
		return x.XnSetupResponse
	}
	return nil
}

type isSuccessfulOutcomeXnApElementaryProcedures_SoValues interface {
	isSuccessfulOutcomeXnApElementaryProcedures_SoValues()
}

type SuccessfulOutcomeXnApElementaryProcedures_XnSetupResponse struct {
	// @inject_tag: aper:"valueExt"
	XnSetupResponse *xnap_pdu_contents.XnSetupResponse `protobuf:"bytes,1,opt,name=xn_setup_response,json=ric_subscription,proto3,oneof"` // ToDo - add the rest of the messages
}

func (*SuccessfulOutcomeXnApElementaryProcedures_XnSetupResponse) isSuccessfulOutcomeXnApElementaryProcedures_SoValues() {
}

// sequence from xnap_v1.asn1:187
// {UnsuccessfulOutcome}
type UnsuccessfulOutcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:0,valueUB:255,unique"
	ProcedureCode int32 `protobuf:"varint,1,opt,name=procedure_code,json=procedureCode,proto3" json:"procedure_code,omitempty"`
	// @inject_tag: aper:"valueLB:0,valueUB:2"
	Criticality xnap_commondatatypes.Criticality `protobuf:"varint,2,opt,name=criticality,proto3,enum=xnap.v1.Criticality" json:"criticality,omitempty"`
	// @inject_tag: aper:"canonicalOrder"
	Value *UnsuccessfulOutcomeXnApElementaryProcedures `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UnsuccessfulOutcome) Reset() {
	*x = UnsuccessfulOutcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsuccessfulOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsuccessfulOutcome) ProtoMessage() {}

func (x *UnsuccessfulOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsuccessfulOutcome.ProtoReflect.Descriptor instead.
func (*UnsuccessfulOutcome) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP(), []int{5}
}

func (x *UnsuccessfulOutcome) GetProcedureCode() int32 {
	if x != nil {
		return x.ProcedureCode
	}
	return 0
}

func (x *UnsuccessfulOutcome) GetCriticality() xnap_commondatatypes.Criticality {
	if x != nil {
		return x.Criticality
	}
	return xnap_commondatatypes.Criticality(0)
}

func (x *UnsuccessfulOutcome) GetValue() *UnsuccessfulOutcomeXnApElementaryProcedures {
	if x != nil {
		return x.Value
	}
	return nil
}

type UnsuccessfulOutcomeXnApElementaryProcedures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UoValues:
	//	*UnsuccessfulOutcomeXnApElementaryProcedures_XnSetupFailure
	UoValues isUnsuccessfulOutcomeXnApElementaryProcedures_UoValues `protobuf_oneof:"uo_values"`
}

func (x *UnsuccessfulOutcomeXnApElementaryProcedures) Reset() {
	*x = UnsuccessfulOutcomeXnApElementaryProcedures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsuccessfulOutcomeXnApElementaryProcedures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsuccessfulOutcomeXnApElementaryProcedures) ProtoMessage() {}

func (x *UnsuccessfulOutcomeXnApElementaryProcedures) ProtoReflect() protoreflect.Message {
	mi := &file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsuccessfulOutcomeXnApElementaryProcedures.ProtoReflect.Descriptor instead.
func (*UnsuccessfulOutcomeXnApElementaryProcedures) Descriptor() ([]byte, []int) {
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP(), []int{6}
}

func (m *UnsuccessfulOutcomeXnApElementaryProcedures) GetUoValues() isUnsuccessfulOutcomeXnApElementaryProcedures_UoValues {
	if m != nil {
		return m.UoValues
	}
	return nil
}

func (x *UnsuccessfulOutcomeXnApElementaryProcedures) GetXnSetupFailure() *xnap_pdu_contents.XnSetupFailure {
	if x, ok := x.GetUoValues().(*UnsuccessfulOutcomeXnApElementaryProcedures_XnSetupFailure); ok {
		return x.XnSetupFailure
	}
	return nil
}

type isUnsuccessfulOutcomeXnApElementaryProcedures_UoValues interface {
	isUnsuccessfulOutcomeXnApElementaryProcedures_UoValues()
}

type UnsuccessfulOutcomeXnApElementaryProcedures_XnSetupFailure struct {
	// @inject_tag: aper:"valueExt"
	XnSetupFailure *xnap_pdu_contents.XnSetupFailure `protobuf:"bytes,1,opt,name=xn_setup_failure,json=f1_setup_response,proto3,oneof"` // ToDo - add the rest of the messages
}

func (*UnsuccessfulOutcomeXnApElementaryProcedures_XnSetupFailure) isUnsuccessfulOutcomeXnApElementaryProcedures_UoValues() {
}

var File_api_xnap_v1_xnap_pdu_descriptions_proto protoreflect.FileDescriptor

var file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x78, 0x6e,
	0x61, 0x70, 0x5f, 0x70, 0x64, 0x75, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x78, 0x6e, 0x61, 0x70, 0x2e,
	0x76, 0x31, 0x1a, 0x22, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x78, 0x6e, 0x61, 0x70,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f,
	0x78, 0x6e, 0x61, 0x70, 0x5f, 0x70, 0x64, 0x75, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76, 0x31,
	0x2f, 0x78, 0x6e, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x78,
	0x6e, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x61, 0x73, 0x6e, 0x31, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x83, 0x02, 0x0a, 0x07, 0x58, 0x6e, 0x41, 0x70, 0x50, 0x44, 0x75, 0x12, 0x4b, 0x0a,
	0x12, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x6e, 0x61, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x12, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x48, 0x00, 0x52, 0x11, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x75, 0x6e, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x13, 0x75, 0x6e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x78, 0x6e,
	0x5f, 0x61, 0x70, 0x5f, 0x70, 0x64, 0x75, 0x22, 0xbc, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x78, 0x6e, 0x61, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x78, 0x6e,
	0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x58, 0x6e, 0x41, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x29, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x58, 0x6e, 0x41, 0x70, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x10, 0x78, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x69, 0x6d,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x11, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x78, 0x6e, 0x61, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x78, 0x6e,
	0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x58, 0x6e, 0x41, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x29, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x58, 0x6e, 0x41, 0x70,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x78, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x6e, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x10, 0x72, 0x69, 0x63,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a,
	0x09, 0x73, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x13, 0x55,
	0x6e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x72, 0x69,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x4a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x58, 0x6e,
	0x41, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x82, 0x01,
	0x0a, 0x2b, 0x55, 0x6e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x4f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x58, 0x6e, 0x41, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x73, 0x12, 0x46, 0x0a,
	0x10, 0x78, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x78, 0x6e, 0x61, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x58, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x48, 0x00, 0x52, 0x11, 0x66, 0x31, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x75, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6f, 0x6e, 0x6f,
	0x73, 0x2d, 0x65, 0x32, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x78, 0x6e, 0x61, 0x70, 0x2f, 0x76,
	0x31, 0x2f, 0x78, 0x6e, 0x61, 0x70, 0x2d, 0x70, 0x64, 0x75, 0x2d, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x78, 0x6e, 0x61, 0x70, 0x70, 0x64, 0x75, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescOnce sync.Once
	file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescData = file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDesc
)

func file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescGZIP() []byte {
	file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescOnce.Do(func() {
		file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescData)
	})
	return file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDescData
}

var file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_xnap_v1_xnap_pdu_descriptions_proto_goTypes = []interface{}{
	(*XnApPDu)(nil),           // 0: xnap.v1.XnApPDu
	(*InitiatingMessage)(nil), // 1: xnap.v1.InitiatingMessage
	(*InitiatingMessageXnApElementaryProcedures)(nil),   // 2: xnap.v1.InitiatingMessageXnApElementaryProcedures
	(*SuccessfulOutcome)(nil),                           // 3: xnap.v1.SuccessfulOutcome
	(*SuccessfulOutcomeXnApElementaryProcedures)(nil),   // 4: xnap.v1.SuccessfulOutcomeXnApElementaryProcedures
	(*UnsuccessfulOutcome)(nil),                         // 5: xnap.v1.UnsuccessfulOutcome
	(*UnsuccessfulOutcomeXnApElementaryProcedures)(nil), // 6: xnap.v1.UnsuccessfulOutcomeXnApElementaryProcedures
	(xnap_commondatatypes.Criticality)(0),               // 7: xnap.v1.Criticality
	(*xnap_pdu_contents.XnSetupRequest)(nil),            // 8: xnap.v1.XnSetupRequest
	(*xnap_pdu_contents.XnSetupResponse)(nil),           // 9: xnap.v1.XnSetupResponse
	(*xnap_pdu_contents.XnSetupFailure)(nil),            // 10: xnap.v1.XnSetupFailure
}
var file_api_xnap_v1_xnap_pdu_descriptions_proto_depIdxs = []int32{
	1,  // 0: xnap.v1.XnApPDu.initiating_message:type_name -> xnap.v1.InitiatingMessage
	3,  // 1: xnap.v1.XnApPDu.successful_outcome:type_name -> xnap.v1.SuccessfulOutcome
	5,  // 2: xnap.v1.XnApPDu.unsuccessful_outcome:type_name -> xnap.v1.UnsuccessfulOutcome
	7,  // 3: xnap.v1.InitiatingMessage.criticality:type_name -> xnap.v1.Criticality
	2,  // 4: xnap.v1.InitiatingMessage.value:type_name -> xnap.v1.InitiatingMessageXnApElementaryProcedures
	8,  // 5: xnap.v1.InitiatingMessageXnApElementaryProcedures.xn_setup_request:type_name -> xnap.v1.XnSetupRequest
	7,  // 6: xnap.v1.SuccessfulOutcome.criticality:type_name -> xnap.v1.Criticality
	4,  // 7: xnap.v1.SuccessfulOutcome.value:type_name -> xnap.v1.SuccessfulOutcomeXnApElementaryProcedures
	9,  // 8: xnap.v1.SuccessfulOutcomeXnApElementaryProcedures.xn_setup_response:type_name -> xnap.v1.XnSetupResponse
	7,  // 9: xnap.v1.UnsuccessfulOutcome.criticality:type_name -> xnap.v1.Criticality
	6,  // 10: xnap.v1.UnsuccessfulOutcome.value:type_name -> xnap.v1.UnsuccessfulOutcomeXnApElementaryProcedures
	10, // 11: xnap.v1.UnsuccessfulOutcomeXnApElementaryProcedures.xn_setup_failure:type_name -> xnap.v1.XnSetupFailure
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_xnap_v1_xnap_pdu_descriptions_proto_init() }
func file_api_xnap_v1_xnap_pdu_descriptions_proto_init() {
	if File_api_xnap_v1_xnap_pdu_descriptions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XnApPDu); i {
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
		file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiatingMessage); i {
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
		file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiatingMessageXnApElementaryProcedures); i {
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
		file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessfulOutcome); i {
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
		file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessfulOutcomeXnApElementaryProcedures); i {
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
		file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsuccessfulOutcome); i {
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
		file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsuccessfulOutcomeXnApElementaryProcedures); i {
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
	file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*XnApPDu_InitiatingMessage)(nil),
		(*XnApPDu_SuccessfulOutcome)(nil),
		(*XnApPDu_UnsuccessfulOutcome)(nil),
	}
	file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*InitiatingMessageXnApElementaryProcedures_XnSetupRequest)(nil),
	}
	file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*SuccessfulOutcomeXnApElementaryProcedures_XnSetupResponse)(nil),
	}
	file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*UnsuccessfulOutcomeXnApElementaryProcedures_XnSetupFailure)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_xnap_v1_xnap_pdu_descriptions_proto_goTypes,
		DependencyIndexes: file_api_xnap_v1_xnap_pdu_descriptions_proto_depIdxs,
		MessageInfos:      file_api_xnap_v1_xnap_pdu_descriptions_proto_msgTypes,
	}.Build()
	File_api_xnap_v1_xnap_pdu_descriptions_proto = out.File
	file_api_xnap_v1_xnap_pdu_descriptions_proto_rawDesc = nil
	file_api_xnap_v1_xnap_pdu_descriptions_proto_goTypes = nil
	file_api_xnap_v1_xnap_pdu_descriptions_proto_depIdxs = nil
}
