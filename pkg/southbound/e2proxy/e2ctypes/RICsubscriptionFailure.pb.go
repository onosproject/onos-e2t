// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICsubscriptionFailure.proto

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

// Deprecated: Do not use.
type RICsubscriptionFailureT struct {
	ProtocolIEs          *ProtocolIE_Container_1544P2T `protobuf:"bytes,1,opt,name=protocolIEs,proto3" json:"protocolIEs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *RICsubscriptionFailureT) Reset()         { *m = RICsubscriptionFailureT{} }
func (m *RICsubscriptionFailureT) String() string { return proto.CompactTextString(m) }
func (*RICsubscriptionFailureT) ProtoMessage()    {}
func (*RICsubscriptionFailureT) Descriptor() ([]byte, []int) {
	return fileDescriptor_715e15f75edf8359, []int{0}
}

func (m *RICsubscriptionFailureT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICsubscriptionFailureT.Unmarshal(m, b)
}
func (m *RICsubscriptionFailureT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICsubscriptionFailureT.Marshal(b, m, deterministic)
}
func (m *RICsubscriptionFailureT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICsubscriptionFailureT.Merge(m, src)
}
func (m *RICsubscriptionFailureT) XXX_Size() int {
	return xxx_messageInfo_RICsubscriptionFailureT.Size(m)
}
func (m *RICsubscriptionFailureT) XXX_DiscardUnknown() {
	xxx_messageInfo_RICsubscriptionFailureT.DiscardUnknown(m)
}

var xxx_messageInfo_RICsubscriptionFailureT proto.InternalMessageInfo

func (m *RICsubscriptionFailureT) GetProtocolIEs() *ProtocolIE_Container_1544P2T {
	if m != nil {
		return m.ProtocolIEs
	}
	return nil
}

func init() {
	proto.RegisterType((*RICsubscriptionFailureT)(nil), "e2ctypes.RICsubscriptionFailure_t")
}

func init() { proto.RegisterFile("RICsubscriptionFailure.proto", fileDescriptor_715e15f75edf8359) }

var fileDescriptor_715e15f75edf8359 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x09, 0xf2, 0x74, 0x2e,
	0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x73, 0x4b, 0xcc, 0xcc, 0x29,
	0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c,
	0x48, 0x2d, 0x96, 0x92, 0x0a, 0x00, 0x09, 0x24, 0xe7, 0xe7, 0x78, 0xba, 0xea, 0x3a, 0xe7, 0xe7,
	0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0x41, 0x54, 0x29, 0x65, 0x72, 0x49, 0x60, 0x37, 0x25, 0xbe,
	0x44, 0xc8, 0x93, 0x8b, 0xbb, 0x00, 0xae, 0xb3, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48,
	0x5d, 0x0f, 0x66, 0xae, 0x1e, 0xc2, 0xd8, 0x78, 0xb8, 0xb1, 0xf1, 0x86, 0xa6, 0x26, 0x26, 0x01,
	0x46, 0xf1, 0x25, 0x41, 0xc8, 0x7a, 0xad, 0x98, 0x24, 0x18, 0x93, 0xd8, 0xc0, 0x02, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x84, 0x32, 0xa5, 0xb7, 0x00, 0x00, 0x00,
}
