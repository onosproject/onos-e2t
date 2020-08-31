// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RANfunctionsID-List.proto

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

type RANfunctionsID_ListT struct {
	List                 []*RANfunctionID_ItemIEsT `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RANfunctionsID_ListT) Reset()         { *m = RANfunctionsID_ListT{} }
func (m *RANfunctionsID_ListT) String() string { return proto.CompactTextString(m) }
func (*RANfunctionsID_ListT) ProtoMessage()    {}
func (*RANfunctionsID_ListT) Descriptor() ([]byte, []int) {
	return fileDescriptor_929afe113fac455e, []int{0}
}

func (m *RANfunctionsID_ListT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RANfunctionsID_ListT.Unmarshal(m, b)
}
func (m *RANfunctionsID_ListT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RANfunctionsID_ListT.Marshal(b, m, deterministic)
}
func (m *RANfunctionsID_ListT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RANfunctionsID_ListT.Merge(m, src)
}
func (m *RANfunctionsID_ListT) XXX_Size() int {
	return xxx_messageInfo_RANfunctionsID_ListT.Size(m)
}
func (m *RANfunctionsID_ListT) XXX_DiscardUnknown() {
	xxx_messageInfo_RANfunctionsID_ListT.DiscardUnknown(m)
}

var xxx_messageInfo_RANfunctionsID_ListT proto.InternalMessageInfo

func (m *RANfunctionsID_ListT) GetList() []*RANfunctionID_ItemIEsT {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*RANfunctionsID_ListT)(nil), "e2ctypes.RANfunctionsID_List_t")
}

func init() { proto.RegisterFile("RANfunctionsID-List.proto", fileDescriptor_929afe113fac455e) }

var fileDescriptor_929afe113fac455e = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x0c, 0x72, 0xf4, 0x4b,
	0x2b, 0xcd, 0x4b, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xf6, 0x74, 0xd1, 0xf5, 0xc9, 0x2c, 0x2e, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x96,
	0x52, 0x08, 0x00, 0x09, 0x24, 0xe7, 0xe7, 0x78, 0xba, 0xea, 0x06, 0x67, 0xe6, 0xa5, 0xe7, 0xa4,
	0x3a, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0xa5, 0x16, 0x41, 0xd4, 0x2a, 0xf9, 0x71, 0x89, 0xa2,
	0x1a, 0x14, 0x0f, 0x32, 0x28, 0xbe, 0x44, 0xc8, 0x94, 0x8b, 0x25, 0x27, 0xb3, 0xb8, 0x44, 0x82,
	0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x51, 0x0f, 0x66, 0xa6, 0x1e, 0x92, 0x72, 0x4f, 0x97, 0x78,
	0xcf, 0x92, 0xd4, 0x5c, 0x4f, 0xd7, 0xe2, 0xf8, 0x92, 0x20, 0xb0, 0xf2, 0x24, 0x36, 0xb0, 0xb1,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xd6, 0xb0, 0x1a, 0x9f, 0x00, 0x00, 0x00,
}
