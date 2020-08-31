// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ErrorIndication.proto

package e2ctypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorIndicationT struct {
	ProtocolIEs          *ProtocolIE_Container_1544P10T `protobuf:"bytes,1,opt,name=protocolIEs,proto3" json:"protocolIEs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ErrorIndicationT) Reset()         { *m = ErrorIndicationT{} }
func (m *ErrorIndicationT) String() string { return proto.CompactTextString(m) }
func (*ErrorIndicationT) ProtoMessage()    {}
func (*ErrorIndicationT) Descriptor() ([]byte, []int) {
	return fileDescriptor_6673d2c5d89c5d72, []int{0}
}

func (m *ErrorIndicationT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorIndicationT.Unmarshal(m, b)
}
func (m *ErrorIndicationT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorIndicationT.Marshal(b, m, deterministic)
}
func (m *ErrorIndicationT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorIndicationT.Merge(m, src)
}
func (m *ErrorIndicationT) XXX_Size() int {
	return xxx_messageInfo_ErrorIndicationT.Size(m)
}
func (m *ErrorIndicationT) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorIndicationT.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorIndicationT proto.InternalMessageInfo

func (m *ErrorIndicationT) GetProtocolIEs() *ProtocolIE_Container_1544P10T {
	if m != nil {
		return m.ProtocolIEs
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorIndicationT)(nil), "e2ctypes.ErrorIndication_t")
}

func init() { proto.RegisterFile("ErrorIndication.proto", fileDescriptor_6673d2c5d89c5d72) }

var fileDescriptor_6673d2c5d89c5d72 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x75, 0x2d, 0x2a, 0xca,
	0x2f, 0xf2, 0xcc, 0x4b, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x92, 0x0a, 0x00, 0x09,
	0x24, 0xe7, 0xe7, 0x78, 0xba, 0xea, 0x3a, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0x41,
	0x54, 0x29, 0xc5, 0x73, 0x09, 0xa2, 0x69, 0x8f, 0x2f, 0x11, 0xf2, 0xe2, 0xe2, 0x2e, 0x80, 0x6b,
	0x29, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xd2, 0xd0, 0x83, 0x19, 0xa8, 0x87, 0x30, 0x2f,
	0x1e, 0x6e, 0x5e, 0xbc, 0xa1, 0xa9, 0x89, 0x49, 0x80, 0xa1, 0x41, 0x7c, 0x49, 0x10, 0xb2, 0xe6,
	0x24, 0x36, 0x30, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x46, 0x99, 0xbd, 0x16, 0xa6, 0x00,
	0x00, 0x00,
}
