// Code generated by protoc-gen-go. DO NOT EDIT.
// GlobalE2node-gNB-ID.proto is a deprecated file.

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

type GlobalE2NodeGNB_ID struct {
	GlobalGNB_ID         *GlobalgNB_IDT `protobuf:"bytes,1,opt,name=global_gNB_ID,json=globalGNBID,proto3" json:"global_gNB_ID,omitempty"` // Deprecated: Do not use.
	GNB_CU_UP_ID         int32          `protobuf:"varint,2,opt,name=gNB_CU_UP_ID,json=gNBCUUPID,proto3" json:"gNB_CU_UP_ID,omitempty"`    // Deprecated: Do not use.
	GNB_DU_ID            int32          `protobuf:"varint,3,opt,name=gNB_DU_ID,json=gNBDUID,proto3" json:"gNB_DU_ID,omitempty"`            // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GlobalE2NodeGNB_ID) Reset()         { *m = GlobalE2NodeGNB_ID{} }
func (m *GlobalE2NodeGNB_ID) String() string { return proto.CompactTextString(m) }
func (*GlobalE2NodeGNB_ID) ProtoMessage()    {}
func (*GlobalE2NodeGNB_ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1856d676a8aeb2c5, []int{0}
}

func (m *GlobalE2NodeGNB_ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalE2NodeGNB_ID.Unmarshal(m, b)
}
func (m *GlobalE2NodeGNB_ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalE2NodeGNB_ID.Marshal(b, m, deterministic)
}
func (m *GlobalE2NodeGNB_ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalE2NodeGNB_ID.Merge(m, src)
}
func (m *GlobalE2NodeGNB_ID) XXX_Size() int {
	return xxx_messageInfo_GlobalE2NodeGNB_ID.Size(m)
}
func (m *GlobalE2NodeGNB_ID) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalE2NodeGNB_ID.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalE2NodeGNB_ID proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *GlobalE2NodeGNB_ID) GetGlobalGNB_ID() *GlobalgNB_IDT {
	if m != nil {
		return m.GlobalGNB_ID
	}
	return nil
}

// Deprecated: Do not use.
func (m *GlobalE2NodeGNB_ID) GetGNB_CU_UP_ID() int32 {
	if m != nil {
		return m.GNB_CU_UP_ID
	}
	return 0
}

// Deprecated: Do not use.
func (m *GlobalE2NodeGNB_ID) GetGNB_DU_ID() int32 {
	if m != nil {
		return m.GNB_DU_ID
	}
	return 0
}

func init() {
	proto.RegisterType((*GlobalE2NodeGNB_ID)(nil), "e2ctypes.GlobalE2node_gNB_ID")
}

func init() { proto.RegisterFile("GlobalE2node-gNB-ID.proto", fileDescriptor_1856d676a8aeb2c5) }

var fileDescriptor_1856d676a8aeb2c5 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x74, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0x71, 0x35, 0xca, 0xcb, 0x4f, 0x49, 0xd5, 0x4d, 0xf7, 0x73, 0xd2, 0xf5, 0x74, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96,
	0x12, 0x82, 0x28, 0x42, 0x96, 0x55, 0x9a, 0xc3, 0xc8, 0x25, 0x8c, 0xac, 0x37, 0x3e, 0xdd, 0xcf,
	0x29, 0xde, 0xd3, 0x45, 0xc8, 0x81, 0x8b, 0x37, 0x1d, 0x2c, 0x0c, 0x15, 0x90, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x36, 0x92, 0xd0, 0x83, 0x99, 0xa6, 0x07, 0x37, 0x2c, 0xde, 0xd3, 0x25, 0xbe, 0xc4,
	0x89, 0x49, 0x82, 0x31, 0x88, 0x1b, 0xa2, 0xc5, 0xdd, 0xcf, 0xc9, 0xd3, 0x45, 0x48, 0x99, 0x8b,
	0x07, 0x24, 0xe9, 0x1c, 0x1a, 0x1f, 0x1a, 0x00, 0x32, 0x80, 0x49, 0x81, 0x51, 0x83, 0x15, 0xac,
	0x8c, 0x33, 0xdd, 0xcf, 0xc9, 0x39, 0x34, 0x34, 0xc0, 0xd3, 0x45, 0x48, 0x8e, 0x0b, 0xc4, 0x89,
	0x77, 0x09, 0x05, 0xa9, 0x60, 0x86, 0xab, 0x60, 0x4f, 0xf7, 0x73, 0x72, 0x09, 0xf5, 0x74, 0x71,
	0x62, 0xde, 0xc1, 0xc8, 0x98, 0xc4, 0x06, 0x76, 0xaa, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe1,
	0x5c, 0x56, 0x8b, 0xe5, 0x00, 0x00, 0x00,
}
