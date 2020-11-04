// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/ricapi/e2/headers/v1beta1/headers.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EncodingType determines encoding type for the response messages
type EncodingType int32

const (
	EncodingType_ENCODING_TYPE_ASN1  EncodingType = 0
	EncodingType_ENCODING_TYPE_PROTO EncodingType = 1
)

var EncodingType_name = map[int32]string{
	0: "ENCODING_TYPE_ASN1",
	1: "ENCODING_TYPE_PROTO",
}

var EncodingType_value = map[string]int32{
	"ENCODING_TYPE_ASN1":  0,
	"ENCODING_TYPE_PROTO": 1,
}

func (x EncodingType) String() string {
	return proto.EnumName(EncodingType_name, int32(x))
}

func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{0}
}

// ResponseStatus
type ResponseStatus int32

const (
	ResponseStatus_RESPONSE_STATUS_FAILED    ResponseStatus = 0
	ResponseStatus_RESPONSE_STATUS_SUCCEEDED ResponseStatus = 1
)

var ResponseStatus_name = map[int32]string{
	0: "RESPONSE_STATUS_FAILED",
	1: "RESPONSE_STATUS_SUCCEEDED",
}

var ResponseStatus_value = map[string]int32{
	"RESPONSE_STATUS_FAILED":    0,
	"RESPONSE_STATUS_SUCCEEDED": 1,
}

func (x ResponseStatus) String() string {
	return proto.EnumName(ResponseStatus_name, int32(x))
}

func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{1}
}

type ServiceModelInfo struct {
	ServiceModelId       string   `protobuf:"bytes,1,opt,name=service_model_id,json=serviceModelId,proto3" json:"service_model_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceModelInfo) Reset()         { *m = ServiceModelInfo{} }
func (m *ServiceModelInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceModelInfo) ProtoMessage()    {}
func (*ServiceModelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{0}
}
func (m *ServiceModelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceModelInfo.Unmarshal(m, b)
}
func (m *ServiceModelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceModelInfo.Marshal(b, m, deterministic)
}
func (m *ServiceModelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceModelInfo.Merge(m, src)
}
func (m *ServiceModelInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceModelInfo.Size(m)
}
func (m *ServiceModelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceModelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceModelInfo proto.InternalMessageInfo

func (m *ServiceModelInfo) GetServiceModelId() string {
	if m != nil {
		return m.ServiceModelId
	}
	return ""
}

// RequestHeader a common request header for all requests including encoding type, client/xApp/session info, ordering info, etc
type RequestHeader struct {
	EncodingType         EncodingType      `protobuf:"varint,1,opt,name=encoding_type,json=encodingType,proto3,enum=ricapi.e2.headers.v1beta1.EncodingType" json:"encoding_type,omitempty"`
	ServiceModelInfo     *ServiceModelInfo `protobuf:"bytes,2,opt,name=service_model_info,json=serviceModelInfo,proto3" json:"service_model_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{1}
}
func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetEncodingType() EncodingType {
	if m != nil {
		return m.EncodingType
	}
	return EncodingType_ENCODING_TYPE_ASN1
}

func (m *RequestHeader) GetServiceModelInfo() *ServiceModelInfo {
	if m != nil {
		return m.ServiceModelInfo
	}
	return nil
}

// ResponseHeader a common response header for all responses including encoding type, client/xApp/session info, ordering info, etc
type ResponseHeader struct {
	EncodingType         EncodingType      `protobuf:"varint,1,opt,name=encoding_type,json=encodingType,proto3,enum=ricapi.e2.headers.v1beta1.EncodingType" json:"encoding_type,omitempty"`
	ServiceModelInfo     *ServiceModelInfo `protobuf:"bytes,2,opt,name=service_model_info,json=serviceModelInfo,proto3" json:"service_model_info,omitempty"`
	ResponseStatus       ResponseStatus    `protobuf:"varint,3,opt,name=response_status,json=responseStatus,proto3,enum=ricapi.e2.headers.v1beta1.ResponseStatus" json:"response_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f29ad563ec8c20e, []int{2}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetEncodingType() EncodingType {
	if m != nil {
		return m.EncodingType
	}
	return EncodingType_ENCODING_TYPE_ASN1
}

func (m *ResponseHeader) GetServiceModelInfo() *ServiceModelInfo {
	if m != nil {
		return m.ServiceModelInfo
	}
	return nil
}

func (m *ResponseHeader) GetResponseStatus() ResponseStatus {
	if m != nil {
		return m.ResponseStatus
	}
	return ResponseStatus_RESPONSE_STATUS_FAILED
}

func init() {
	proto.RegisterEnum("ricapi.e2.headers.v1beta1.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("ricapi.e2.headers.v1beta1.ResponseStatus", ResponseStatus_name, ResponseStatus_value)
	proto.RegisterType((*ServiceModelInfo)(nil), "ricapi.e2.headers.v1beta1.ServiceModelInfo")
	proto.RegisterType((*RequestHeader)(nil), "ricapi.e2.headers.v1beta1.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "ricapi.e2.headers.v1beta1.ResponseHeader")
}

func init() {
	proto.RegisterFile("api/ricapi/e2/headers/v1beta1/headers.proto", fileDescriptor_6f29ad563ec8c20e)
}

var fileDescriptor_6f29ad563ec8c20e = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0x41, 0x8b, 0xda, 0x40,
	0x18, 0x35, 0x16, 0x0a, 0x9d, 0x6a, 0x1a, 0xa6, 0x60, 0xb5, 0x50, 0x10, 0x2f, 0xb5, 0x4a, 0x13,
	0x4c, 0xaf, 0x2d, 0x62, 0xcd, 0xb4, 0x95, 0xda, 0x44, 0x66, 0xe2, 0xc1, 0x5e, 0x86, 0x98, 0x7c,
	0x6a, 0x96, 0x35, 0x93, 0xcd, 0x8c, 0x82, 0x3f, 0x62, 0xff, 0xcf, 0xfe, 0xbc, 0xc5, 0xa8, 0xac,
	0x06, 0xd6, 0xfb, 0x9e, 0x92, 0xf7, 0xbe, 0xf7, 0x3d, 0xde, 0x1b, 0x66, 0x50, 0x37, 0x48, 0x63,
	0x2b, 0x8b, 0xc3, 0xfd, 0x07, 0x6c, 0x6b, 0x05, 0x41, 0x04, 0x99, 0xb4, 0xb6, 0xbd, 0x39, 0xa8,
	0xa0, 0x77, 0xc2, 0x66, 0x9a, 0x09, 0x25, 0x70, 0xe3, 0x20, 0x34, 0xc1, 0x36, 0x4f, 0x83, 0xa3,
	0xb0, 0xf5, 0x1d, 0x19, 0x0c, 0xb2, 0x6d, 0x1c, 0xc2, 0x3f, 0x11, 0xc1, 0xed, 0x28, 0x59, 0x08,
	0xdc, 0x46, 0x86, 0x3c, 0x70, 0x7c, 0xbd, 0x27, 0x79, 0x1c, 0xd5, 0xb5, 0xa6, 0xd6, 0x7e, 0x43,
	0x75, 0x79, 0xae, 0x8d, 0x5a, 0x0f, 0x1a, 0xaa, 0x52, 0xb8, 0xdb, 0x80, 0x54, 0x7f, 0x72, 0x63,
	0x3c, 0x46, 0x55, 0x48, 0x42, 0x11, 0xc5, 0xc9, 0x92, 0xab, 0x5d, 0x0a, 0xf9, 0xa2, 0x6e, 0x7f,
	0x36, 0x9f, 0x8d, 0x60, 0x92, 0xa3, 0xde, 0xdf, 0xa5, 0x40, 0x2b, 0x70, 0x86, 0xf0, 0x0c, 0xe1,
	0x42, 0x92, 0x64, 0x21, 0xea, 0xe5, 0xa6, 0xd6, 0x7e, 0x6b, 0x77, 0xaf, 0x58, 0x16, 0x2b, 0x51,
	0x43, 0x16, 0x98, 0xd6, 0x7d, 0x19, 0xe9, 0x14, 0x64, 0x2a, 0x12, 0x09, 0x2f, 0x2c, 0x3b, 0xa6,
	0xe8, 0x5d, 0x76, 0x8c, 0xce, 0xa5, 0x0a, 0xd4, 0x46, 0xd6, 0x5f, 0xe5, 0x51, 0xbf, 0x5c, 0xf1,
	0x3d, 0x95, 0x65, 0xf9, 0x02, 0xd5, 0xb3, 0x0b, 0xdc, 0xe9, 0xa3, 0xca, 0x79, 0x19, 0x5c, 0x43,
	0x98, 0xb8, 0x43, 0xcf, 0x19, 0xb9, 0xbf, 0xb9, 0x3f, 0x9b, 0x10, 0x3e, 0x60, 0x6e, 0xcf, 0x28,
	0xe1, 0x0f, 0xe8, 0xfd, 0x25, 0x3f, 0xa1, 0x9e, 0xef, 0x19, 0x5a, 0xe7, 0xef, 0xd3, 0x79, 0x1e,
	0x2c, 0xf1, 0x47, 0x54, 0xa3, 0x84, 0x4d, 0x3c, 0x97, 0x11, 0xce, 0xfc, 0x81, 0x3f, 0x65, 0xfc,
	0xd7, 0x60, 0x34, 0x26, 0x8e, 0x51, 0xc2, 0x9f, 0x50, 0xa3, 0x38, 0x63, 0xd3, 0xe1, 0x90, 0x10,
	0x87, 0x38, 0x86, 0xf6, 0xb3, 0xff, 0xff, 0xc7, 0x32, 0x56, 0xab, 0xcd, 0xdc, 0x0c, 0xc5, 0xda,
	0x12, 0x89, 0x90, 0x69, 0x26, 0x6e, 0x20, 0x54, 0xf9, 0xff, 0x57, 0xb0, 0x95, 0x75, 0xf5, 0x01,
	0xcc, 0x5f, 0xe7, 0x37, 0xff, 0xdb, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x6e, 0x0d, 0x7a,
	0x28, 0x03, 0x00, 0x00,
}
