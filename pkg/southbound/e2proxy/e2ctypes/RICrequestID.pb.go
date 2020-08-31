// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICrequestID.proto

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

type RICrequestIDT struct {
	RicRequestorID       int64    `protobuf:"varint,1,opt,name=ricRequestorID,proto3" json:"ricRequestorID,omitempty"`
	RicInstanceID        int64    `protobuf:"varint,2,opt,name=ricInstanceID,proto3" json:"ricInstanceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RICrequestIDT) Reset()         { *m = RICrequestIDT{} }
func (m *RICrequestIDT) String() string { return proto.CompactTextString(m) }
func (*RICrequestIDT) ProtoMessage()    {}
func (*RICrequestIDT) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc007813c583871, []int{0}
}

func (m *RICrequestIDT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICrequestIDT.Unmarshal(m, b)
}
func (m *RICrequestIDT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICrequestIDT.Marshal(b, m, deterministic)
}
func (m *RICrequestIDT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICrequestIDT.Merge(m, src)
}
func (m *RICrequestIDT) XXX_Size() int {
	return xxx_messageInfo_RICrequestIDT.Size(m)
}
func (m *RICrequestIDT) XXX_DiscardUnknown() {
	xxx_messageInfo_RICrequestIDT.DiscardUnknown(m)
}

var xxx_messageInfo_RICrequestIDT proto.InternalMessageInfo

func (m *RICrequestIDT) GetRicRequestorID() int64 {
	if m != nil {
		return m.RicRequestorID
	}
	return 0
}

func (m *RICrequestIDT) GetRicInstanceID() int64 {
	if m != nil {
		return m.RicInstanceID
	}
	return 0
}

func init() {
	proto.RegisterType((*RICrequestIDT)(nil), "e2ctypes.RICrequestID_t")
}

func init() { proto.RegisterFile("RICrequestID.proto", fileDescriptor_5cc007813c583871) }

var fileDescriptor_5cc007813c583871 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0a, 0xf2, 0x74, 0x2e,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xf1, 0x74, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x8a, 0xe3, 0xe2, 0x43, 0x96, 0x8f, 0x2f,
	0x11, 0x52, 0xe3, 0xe2, 0x2b, 0xca, 0x4c, 0x0e, 0x82, 0x88, 0xe4, 0x17, 0x79, 0xba, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x30, 0x07, 0xa1, 0x89, 0x0a, 0xa9, 0x70, 0xf1, 0x16, 0x65, 0x26, 0x7b, 0xe6,
	0x15, 0x97, 0x24, 0xe6, 0x25, 0xa7, 0x7a, 0xba, 0x48, 0x30, 0x81, 0x95, 0xa1, 0x0a, 0x26, 0xb1,
	0x81, 0x2d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x96, 0x23, 0x1e, 0x0e, 0x86, 0x00, 0x00,
	0x00,
}
