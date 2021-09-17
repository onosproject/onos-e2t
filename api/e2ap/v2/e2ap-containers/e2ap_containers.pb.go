//
//SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
//SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

////////////////////// e2ap-containers.proto //////////////////////
// Protobuf generated from /e2ap_v2.asn by asn1c-0.9.29
// E2AP-Containers { iso(1) identified-organization(3) dod(6) internet(1) private(4) enterprise(1) 53148 e2(1) version2(2) e2ap(1) e2ap-Containers(5) }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.1
// source: api/e2ap/v2/e2ap_containers.proto

package e2ap_containers

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-constants"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// sequence from e2ap_v2.asn:1908
// Param E2AP-PROTOCOL-IES:IEsSetParam
// {ProtocolIE-Container001}
type ProtocolIeContainer001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []*ProtocolIeField001 `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeContainer001) Reset() {
	*x = ProtocolIeContainer001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeContainer001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeContainer001) ProtoMessage() {}

func (x *ProtocolIeContainer001) ProtoReflect() protoreflect.Message {
	mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeContainer001.ProtoReflect.Descriptor instead.
func (*ProtocolIeContainer001) Descriptor() ([]byte, []int) {
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP(), []int{0}
}

func (x *ProtocolIeContainer001) GetValue() []*ProtocolIeField001 {
	if x != nil {
		return x.Value
	}
	return nil
}

// reference from e2ap_v2.asn:1911
// Param E2AP-PROTOCOL-IES:IEsSetParam
// {ProtocolIE-SingleContainer001}
type ProtocolIeSingleContainer001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *ProtocolIeField001 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeSingleContainer001) Reset() {
	*x = ProtocolIeSingleContainer001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeSingleContainer001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeSingleContainer001) ProtoMessage() {}

func (x *ProtocolIeSingleContainer001) ProtoReflect() protoreflect.Message {
	mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeSingleContainer001.ProtoReflect.Descriptor instead.
func (*ProtocolIeSingleContainer001) Descriptor() ([]byte, []int) {
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP(), []int{1}
}

func (x *ProtocolIeSingleContainer001) GetValue() *ProtocolIeField001 {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from e2ap_v2.asn:1914
// Param E2AP-PROTOCOL-IES:IEsSetParam
// {ProtocolIE-Field001}
type ProtocolIeField001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *e2ap_constants.IdRicrequestId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Criticality e2ap_commondatatypes.Criticality `protobuf:"varint,2,opt,name=criticality,proto3,enum=e2ap.v2.Criticality" json:"criticality,omitempty"` //    @id value = 3 [ json_name="value"];
}

func (x *ProtocolIeField001) Reset() {
	*x = ProtocolIeField001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeField001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeField001) ProtoMessage() {}

func (x *ProtocolIeField001) ProtoReflect() protoreflect.Message {
	mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeField001.ProtoReflect.Descriptor instead.
func (*ProtocolIeField001) Descriptor() ([]byte, []int) {
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP(), []int{2}
}

func (x *ProtocolIeField001) GetId() *e2ap_constants.IdRicrequestId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ProtocolIeField001) GetCriticality() e2ap_commondatatypes.Criticality {
	if x != nil {
		return x.Criticality
	}
	return e2ap_commondatatypes.Criticality(0)
}

// sequence from e2ap_v2.asn:1927
// Param E2AP-PROTOCOL-IES-PAIR:IEsSetParam
// {ProtocolIE-ContainerPair}
type ProtocolIeContainerPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []*ProtocolIeFieldPair `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeContainerPair) Reset() {
	*x = ProtocolIeContainerPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeContainerPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeContainerPair) ProtoMessage() {}

func (x *ProtocolIeContainerPair) ProtoReflect() protoreflect.Message {
	mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeContainerPair.ProtoReflect.Descriptor instead.
func (*ProtocolIeContainerPair) Descriptor() ([]byte, []int) {
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP(), []int{3}
}

func (x *ProtocolIeContainerPair) GetValue() []*ProtocolIeFieldPair {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from e2ap_v2.asn:1930
// Param E2AP-PROTOCOL-IES-PAIR:IEsSetParam
// {ProtocolIE-FieldPair}
type ProtocolIeFieldPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProtocolIeFieldPair) Reset() {
	*x = ProtocolIeFieldPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeFieldPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeFieldPair) ProtoMessage() {}

func (x *ProtocolIeFieldPair) ProtoReflect() protoreflect.Message {
	mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeFieldPair.ProtoReflect.Descriptor instead.
func (*ProtocolIeFieldPair) Descriptor() ([]byte, []int) {
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP(), []int{4}
}

// sequence from e2ap_v2.asn:1945
// Param INTEGER:lowerBound
// Param INTEGER:upperBound
// Param E2AP-PROTOCOL-IES:IEsSetParam
// {ProtocolIE-ContainerList}
type ProtocolIeContainerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []*ProtocolIeSingleContainer001 `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeContainerList) Reset() {
	*x = ProtocolIeContainerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeContainerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeContainerList) ProtoMessage() {}

func (x *ProtocolIeContainerList) ProtoReflect() protoreflect.Message {
	mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeContainerList.ProtoReflect.Descriptor instead.
func (*ProtocolIeContainerList) Descriptor() ([]byte, []int) {
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP(), []int{5}
}

func (x *ProtocolIeContainerList) GetValue() []*ProtocolIeSingleContainer001 {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from e2ap_v2.asn:1949
// Param INTEGER:lowerBound
// Param INTEGER:upperBound
// Param E2AP-PROTOCOL-IES-PAIR:IEsSetParam
// {ProtocolIE-ContainerPairList}
type ProtocolIeContainerPairList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:""
	Value []*ProtocolIeContainerPair `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ProtocolIeContainerPairList) Reset() {
	*x = ProtocolIeContainerPairList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolIeContainerPairList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolIeContainerPairList) ProtoMessage() {}

func (x *ProtocolIeContainerPairList) ProtoReflect() protoreflect.Message {
	mi := &file_api_e2ap_v2_e2ap_containers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolIeContainerPairList.ProtoReflect.Descriptor instead.
func (*ProtocolIeContainerPairList) Descriptor() ([]byte, []int) {
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP(), []int{6}
}

func (x *ProtocolIeContainerPairList) GetValue() []*ProtocolIeContainerPair {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_api_e2ap_v2_e2ap_containers_proto protoreflect.FileDescriptor

var file_api_e2ap_v2_e2ap_containers_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x32, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x32,
	0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x32, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x22, 0x65, 0x32,
	0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x32, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x65, 0x32, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x32, 0x61, 0x70, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x16, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x30, 0x30, 0x31, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x32, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x30, 0x31,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x49, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x30, 0x30, 0x31, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x32, 0x61, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x30, 0x30, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x75, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x30, 0x30, 0x31,
	0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65,
	0x32, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x52, 0x69, 0x63, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x72, 0x69,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x65, 0x32, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x22, 0x4d, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x32, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x32,
	0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x15, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x50, 0x61, 0x69, 0x72, 0x22, 0x56, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x49, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x65, 0x32, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x30, 0x30, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x55, 0x0a, 0x1b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x32, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x49, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x2d, 0x65, 0x32, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x32,
	0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x32, 0x61, 0x70, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_e2ap_v2_e2ap_containers_proto_rawDescOnce sync.Once
	file_api_e2ap_v2_e2ap_containers_proto_rawDescData = file_api_e2ap_v2_e2ap_containers_proto_rawDesc
)

func file_api_e2ap_v2_e2ap_containers_proto_rawDescGZIP() []byte {
	file_api_e2ap_v2_e2ap_containers_proto_rawDescOnce.Do(func() {
		file_api_e2ap_v2_e2ap_containers_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_e2ap_v2_e2ap_containers_proto_rawDescData)
	})
	return file_api_e2ap_v2_e2ap_containers_proto_rawDescData
}

var file_api_e2ap_v2_e2ap_containers_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_e2ap_v2_e2ap_containers_proto_goTypes = []interface{}{
	(*ProtocolIeContainer001)(nil),        // 0: e2ap.v2.ProtocolIeContainer001
	(*ProtocolIeSingleContainer001)(nil),  // 1: e2ap.v2.ProtocolIeSingleContainer001
	(*ProtocolIeField001)(nil),            // 2: e2ap.v2.ProtocolIeField001
	(*ProtocolIeContainerPair)(nil),       // 3: e2ap.v2.ProtocolIeContainerPair
	(*ProtocolIeFieldPair)(nil),           // 4: e2ap.v2.ProtocolIeFieldPair
	(*ProtocolIeContainerList)(nil),       // 5: e2ap.v2.ProtocolIeContainerList
	(*ProtocolIeContainerPairList)(nil),   // 6: e2ap.v2.ProtocolIeContainerPairList
	(*e2ap_constants.IdRicrequestId)(nil), // 7: e2ap.v2.IdRicrequestId
	(e2ap_commondatatypes.Criticality)(0), // 8: e2ap.v2.Criticality
}
var file_api_e2ap_v2_e2ap_containers_proto_depIdxs = []int32{
	2, // 0: e2ap.v2.ProtocolIeContainer001.value:type_name -> e2ap.v2.ProtocolIeField001
	2, // 1: e2ap.v2.ProtocolIeSingleContainer001.value:type_name -> e2ap.v2.ProtocolIeField001
	7, // 2: e2ap.v2.ProtocolIeField001.id:type_name -> e2ap.v2.IdRicrequestId
	8, // 3: e2ap.v2.ProtocolIeField001.criticality:type_name -> e2ap.v2.Criticality
	4, // 4: e2ap.v2.ProtocolIeContainerPair.value:type_name -> e2ap.v2.ProtocolIeFieldPair
	1, // 5: e2ap.v2.ProtocolIeContainerList.value:type_name -> e2ap.v2.ProtocolIeSingleContainer001
	3, // 6: e2ap.v2.ProtocolIeContainerPairList.value:type_name -> e2ap.v2.ProtocolIeContainerPair
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_e2ap_v2_e2ap_containers_proto_init() }
func file_api_e2ap_v2_e2ap_containers_proto_init() {
	if File_api_e2ap_v2_e2ap_containers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_e2ap_v2_e2ap_containers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeContainer001); i {
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
		file_api_e2ap_v2_e2ap_containers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeSingleContainer001); i {
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
		file_api_e2ap_v2_e2ap_containers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeField001); i {
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
		file_api_e2ap_v2_e2ap_containers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeContainerPair); i {
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
		file_api_e2ap_v2_e2ap_containers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeFieldPair); i {
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
		file_api_e2ap_v2_e2ap_containers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeContainerList); i {
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
		file_api_e2ap_v2_e2ap_containers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolIeContainerPairList); i {
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
			RawDescriptor: file_api_e2ap_v2_e2ap_containers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_e2ap_v2_e2ap_containers_proto_goTypes,
		DependencyIndexes: file_api_e2ap_v2_e2ap_containers_proto_depIdxs,
		MessageInfos:      file_api_e2ap_v2_e2ap_containers_proto_msgTypes,
	}.Build()
	File_api_e2ap_v2_e2ap_containers_proto = out.File
	file_api_e2ap_v2_e2ap_containers_proto_rawDesc = nil
	file_api_e2ap_v2_e2ap_containers_proto_goTypes = nil
	file_api_e2ap_v2_e2ap_containers_proto_depIdxs = nil
}
