// Code generated by protoc-gen-go. DO NOT EDIT.
// GlobalE2node-eNB-ID.proto is a deprecated file.

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

type GlobalE2NodeENB_IDT struct {
	GlobalENB_ID         *GlobalENB_IDT `protobuf:"bytes,1,opt,name=global_eNB_ID,json=globalENBID,proto3" json:"global_eNB_ID,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GlobalE2NodeENB_IDT) Reset()         { *m = GlobalE2NodeENB_IDT{} }
func (m *GlobalE2NodeENB_IDT) String() string { return proto.CompactTextString(m) }
func (*GlobalE2NodeENB_IDT) ProtoMessage()    {}
func (*GlobalE2NodeENB_IDT) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc0e2720ccfa433e, []int{0}
}

func (m *GlobalE2NodeENB_IDT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalE2NodeENB_IDT.Unmarshal(m, b)
}
func (m *GlobalE2NodeENB_IDT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalE2NodeENB_IDT.Marshal(b, m, deterministic)
}
func (m *GlobalE2NodeENB_IDT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalE2NodeENB_IDT.Merge(m, src)
}
func (m *GlobalE2NodeENB_IDT) XXX_Size() int {
	return xxx_messageInfo_GlobalE2NodeENB_IDT.Size(m)
}
func (m *GlobalE2NodeENB_IDT) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalE2NodeENB_IDT.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalE2NodeENB_IDT proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *GlobalE2NodeENB_IDT) GetGlobalENB_ID() *GlobalENB_IDT {
	if m != nil {
		return m.GlobalENB_ID
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalE2NodeENB_IDT)(nil), "e2ctypes.GlobalE2node_eNB_ID_t")
}

func init() { proto.RegisterFile("GlobalE2node-eNB-ID.proto", fileDescriptor_bc0e2720ccfa433e) }

var fileDescriptor_bc0e2720ccfa433e = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x74, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0x71, 0x35, 0xca, 0xcb, 0x4f, 0x49, 0xd5, 0x4d, 0xf5, 0x73, 0xd2, 0xf5, 0x74, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96,
	0x12, 0x82, 0x2a, 0x42, 0x92, 0x55, 0x8a, 0xe4, 0x12, 0x45, 0xd6, 0x1a, 0x9f, 0xea, 0xe7, 0x14,
	0xef, 0xe9, 0x12, 0x5f, 0x22, 0xe4, 0xc0, 0xc5, 0x9b, 0x0e, 0x96, 0x80, 0x0a, 0x49, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x1b, 0x49, 0xe8, 0xc1, 0x8c, 0xd3, 0x83, 0x9b, 0x06, 0xd6, 0xe0, 0xc4, 0x24,
	0xc1, 0x18, 0xc4, 0x9d, 0x0e, 0x13, 0xf3, 0x74, 0x71, 0x62, 0xde, 0xc1, 0xc8, 0x98, 0xc4, 0x06,
	0xb6, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x24, 0xae, 0xac, 0x54, 0xa1, 0x00, 0x00, 0x00,
}
