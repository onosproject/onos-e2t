// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICsubscriptionDeleteFailure.proto

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
type RICsubscriptionDeleteFailureT struct {
	ProtocolIEs          *ProtocolIE_Container_1544P5T `protobuf:"bytes,1,opt,name=protocolIEs,proto3" json:"protocolIEs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *RICsubscriptionDeleteFailureT) Reset()         { *m = RICsubscriptionDeleteFailureT{} }
func (m *RICsubscriptionDeleteFailureT) String() string { return proto.CompactTextString(m) }
func (*RICsubscriptionDeleteFailureT) ProtoMessage()    {}
func (*RICsubscriptionDeleteFailureT) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b8ccb854f2c04d6, []int{0}
}

func (m *RICsubscriptionDeleteFailureT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICsubscriptionDeleteFailureT.Unmarshal(m, b)
}
func (m *RICsubscriptionDeleteFailureT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICsubscriptionDeleteFailureT.Marshal(b, m, deterministic)
}
func (m *RICsubscriptionDeleteFailureT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICsubscriptionDeleteFailureT.Merge(m, src)
}
func (m *RICsubscriptionDeleteFailureT) XXX_Size() int {
	return xxx_messageInfo_RICsubscriptionDeleteFailureT.Size(m)
}
func (m *RICsubscriptionDeleteFailureT) XXX_DiscardUnknown() {
	xxx_messageInfo_RICsubscriptionDeleteFailureT.DiscardUnknown(m)
}

var xxx_messageInfo_RICsubscriptionDeleteFailureT proto.InternalMessageInfo

func (m *RICsubscriptionDeleteFailureT) GetProtocolIEs() *ProtocolIE_Container_1544P5T {
	if m != nil {
		return m.ProtocolIEs
	}
	return nil
}

func init() {
	proto.RegisterType((*RICsubscriptionDeleteFailureT)(nil), "e2ctypes.RICsubscriptionDeleteFailure_t")
}

func init() { proto.RegisterFile("RICsubscriptionDeleteFailure.proto", fileDescriptor_6b8ccb854f2c04d6) }

var fileDescriptor_6b8ccb854f2c04d6 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x0a, 0xf2, 0x74, 0x2e,
	0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x73, 0x49, 0xcd, 0x49, 0x2d,
	0x49, 0x75, 0x4b, 0xcc, 0xcc, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x92, 0x0a, 0x00, 0x09, 0x24, 0xe7, 0xe7,
	0x78, 0xba, 0xea, 0x3a, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0x41, 0x54, 0x29, 0xe5,
	0x73, 0xc9, 0xe1, 0x33, 0x2b, 0xbe, 0x44, 0xc8, 0x93, 0x8b, 0xbb, 0x00, 0xae, 0xbf, 0x58, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5d, 0x0f, 0x66, 0xba, 0x1e, 0xc2, 0xf0, 0x78, 0xb8, 0xe1,
	0xf1, 0x86, 0xa6, 0x26, 0x26, 0x01, 0xa6, 0xf1, 0x25, 0x41, 0xc8, 0x7a, 0xad, 0x98, 0x24, 0x18,
	0x93, 0xd8, 0xc0, 0x02, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x11, 0x5a, 0xa2, 0xc3,
	0x00, 0x00, 0x00,
}
