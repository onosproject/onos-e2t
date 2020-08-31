// Code generated by protoc-gen-go. DO NOT EDIT.
// source: e2ap-elementary-procedures.proto

package e2ctypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type EmptyMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyMessage) Reset()         { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()    {}
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_59b3e78f6e914a2b, []int{0}
}

func (m *EmptyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyMessage.Unmarshal(m, b)
}
func (m *EmptyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyMessage.Marshal(b, m, deterministic)
}
func (m *EmptyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyMessage.Merge(m, src)
}
func (m *EmptyMessage) XXX_Size() int {
	return xxx_messageInfo_EmptyMessage.Size(m)
}
func (m *EmptyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyMessage proto.InternalMessageInfo

type E2SetupResponseOutcome struct {
	// Types that are valid to be assigned to Outcome:
	//	*E2SetupResponseOutcome_SuccessfulOutcome
	//	*E2SetupResponseOutcome_UnsuccessfulOutcome
	Outcome              isE2SetupResponseOutcome_Outcome `protobuf_oneof:"outcome"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *E2SetupResponseOutcome) Reset()         { *m = E2SetupResponseOutcome{} }
func (m *E2SetupResponseOutcome) String() string { return proto.CompactTextString(m) }
func (*E2SetupResponseOutcome) ProtoMessage()    {}
func (*E2SetupResponseOutcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_59b3e78f6e914a2b, []int{1}
}

func (m *E2SetupResponseOutcome) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_E2SetupResponseOutcome.Unmarshal(m, b)
}
func (m *E2SetupResponseOutcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_E2SetupResponseOutcome.Marshal(b, m, deterministic)
}
func (m *E2SetupResponseOutcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E2SetupResponseOutcome.Merge(m, src)
}
func (m *E2SetupResponseOutcome) XXX_Size() int {
	return xxx_messageInfo_E2SetupResponseOutcome.Size(m)
}
func (m *E2SetupResponseOutcome) XXX_DiscardUnknown() {
	xxx_messageInfo_E2SetupResponseOutcome.DiscardUnknown(m)
}

var xxx_messageInfo_E2SetupResponseOutcome proto.InternalMessageInfo

type isE2SetupResponseOutcome_Outcome interface {
	isE2SetupResponseOutcome_Outcome()
}

type E2SetupResponseOutcome_SuccessfulOutcome struct {
	SuccessfulOutcome *E2SetupResponseT `protobuf:"bytes,1,opt,name=successfulOutcome,proto3,oneof"`
}

type E2SetupResponseOutcome_UnsuccessfulOutcome struct {
	UnsuccessfulOutcome *E2SetupFailureT `protobuf:"bytes,2,opt,name=unsuccessfulOutcome,proto3,oneof"`
}

func (*E2SetupResponseOutcome_SuccessfulOutcome) isE2SetupResponseOutcome_Outcome() {}

func (*E2SetupResponseOutcome_UnsuccessfulOutcome) isE2SetupResponseOutcome_Outcome() {}

func (m *E2SetupResponseOutcome) GetOutcome() isE2SetupResponseOutcome_Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (m *E2SetupResponseOutcome) GetSuccessfulOutcome() *E2SetupResponseT {
	if x, ok := m.GetOutcome().(*E2SetupResponseOutcome_SuccessfulOutcome); ok {
		return x.SuccessfulOutcome
	}
	return nil
}

func (m *E2SetupResponseOutcome) GetUnsuccessfulOutcome() *E2SetupFailureT {
	if x, ok := m.GetOutcome().(*E2SetupResponseOutcome_UnsuccessfulOutcome); ok {
		return x.UnsuccessfulOutcome
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*E2SetupResponseOutcome) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*E2SetupResponseOutcome_SuccessfulOutcome)(nil),
		(*E2SetupResponseOutcome_UnsuccessfulOutcome)(nil),
	}
}

type RICsubscriptionOutcome struct {
	// Types that are valid to be assigned to Outcome:
	//	*RICsubscriptionOutcome_SuccessfulOutcome
	//	*RICsubscriptionOutcome_UnsuccessfulOutcome
	Outcome              isRICsubscriptionOutcome_Outcome `protobuf_oneof:"outcome"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RICsubscriptionOutcome) Reset()         { *m = RICsubscriptionOutcome{} }
func (m *RICsubscriptionOutcome) String() string { return proto.CompactTextString(m) }
func (*RICsubscriptionOutcome) ProtoMessage()    {}
func (*RICsubscriptionOutcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_59b3e78f6e914a2b, []int{2}
}

func (m *RICsubscriptionOutcome) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICsubscriptionOutcome.Unmarshal(m, b)
}
func (m *RICsubscriptionOutcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICsubscriptionOutcome.Marshal(b, m, deterministic)
}
func (m *RICsubscriptionOutcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICsubscriptionOutcome.Merge(m, src)
}
func (m *RICsubscriptionOutcome) XXX_Size() int {
	return xxx_messageInfo_RICsubscriptionOutcome.Size(m)
}
func (m *RICsubscriptionOutcome) XXX_DiscardUnknown() {
	xxx_messageInfo_RICsubscriptionOutcome.DiscardUnknown(m)
}

var xxx_messageInfo_RICsubscriptionOutcome proto.InternalMessageInfo

type isRICsubscriptionOutcome_Outcome interface {
	isRICsubscriptionOutcome_Outcome()
}

type RICsubscriptionOutcome_SuccessfulOutcome struct {
	SuccessfulOutcome *RICsubscriptionResponseT `protobuf:"bytes,1,opt,name=successfulOutcome,proto3,oneof"`
}

type RICsubscriptionOutcome_UnsuccessfulOutcome struct {
	UnsuccessfulOutcome *RICsubscriptionFailureT `protobuf:"bytes,2,opt,name=unsuccessfulOutcome,proto3,oneof"`
}

func (*RICsubscriptionOutcome_SuccessfulOutcome) isRICsubscriptionOutcome_Outcome() {}

func (*RICsubscriptionOutcome_UnsuccessfulOutcome) isRICsubscriptionOutcome_Outcome() {}

func (m *RICsubscriptionOutcome) GetOutcome() isRICsubscriptionOutcome_Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (m *RICsubscriptionOutcome) GetSuccessfulOutcome() *RICsubscriptionResponseT {
	if x, ok := m.GetOutcome().(*RICsubscriptionOutcome_SuccessfulOutcome); ok {
		return x.SuccessfulOutcome
	}
	return nil
}

func (m *RICsubscriptionOutcome) GetUnsuccessfulOutcome() *RICsubscriptionFailureT {
	if x, ok := m.GetOutcome().(*RICsubscriptionOutcome_UnsuccessfulOutcome); ok {
		return x.UnsuccessfulOutcome
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RICsubscriptionOutcome) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RICsubscriptionOutcome_SuccessfulOutcome)(nil),
		(*RICsubscriptionOutcome_UnsuccessfulOutcome)(nil),
	}
}

type RICsubscriptionDeleteOutcome struct {
	// Types that are valid to be assigned to Outcome:
	//	*RICsubscriptionDeleteOutcome_SuccessfulOutcome
	//	*RICsubscriptionDeleteOutcome_UnsuccessfulOutcome
	Outcome              isRICsubscriptionDeleteOutcome_Outcome `protobuf_oneof:"outcome"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *RICsubscriptionDeleteOutcome) Reset()         { *m = RICsubscriptionDeleteOutcome{} }
func (m *RICsubscriptionDeleteOutcome) String() string { return proto.CompactTextString(m) }
func (*RICsubscriptionDeleteOutcome) ProtoMessage()    {}
func (*RICsubscriptionDeleteOutcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_59b3e78f6e914a2b, []int{3}
}

func (m *RICsubscriptionDeleteOutcome) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICsubscriptionDeleteOutcome.Unmarshal(m, b)
}
func (m *RICsubscriptionDeleteOutcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICsubscriptionDeleteOutcome.Marshal(b, m, deterministic)
}
func (m *RICsubscriptionDeleteOutcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICsubscriptionDeleteOutcome.Merge(m, src)
}
func (m *RICsubscriptionDeleteOutcome) XXX_Size() int {
	return xxx_messageInfo_RICsubscriptionDeleteOutcome.Size(m)
}
func (m *RICsubscriptionDeleteOutcome) XXX_DiscardUnknown() {
	xxx_messageInfo_RICsubscriptionDeleteOutcome.DiscardUnknown(m)
}

var xxx_messageInfo_RICsubscriptionDeleteOutcome proto.InternalMessageInfo

type isRICsubscriptionDeleteOutcome_Outcome interface {
	isRICsubscriptionDeleteOutcome_Outcome()
}

type RICsubscriptionDeleteOutcome_SuccessfulOutcome struct {
	SuccessfulOutcome *RICsubscriptionDeleteResponseT `protobuf:"bytes,1,opt,name=successfulOutcome,proto3,oneof"`
}

type RICsubscriptionDeleteOutcome_UnsuccessfulOutcome struct {
	UnsuccessfulOutcome *RICsubscriptionDeleteFailureT `protobuf:"bytes,2,opt,name=unsuccessfulOutcome,proto3,oneof"`
}

func (*RICsubscriptionDeleteOutcome_SuccessfulOutcome) isRICsubscriptionDeleteOutcome_Outcome() {}

func (*RICsubscriptionDeleteOutcome_UnsuccessfulOutcome) isRICsubscriptionDeleteOutcome_Outcome() {}

func (m *RICsubscriptionDeleteOutcome) GetOutcome() isRICsubscriptionDeleteOutcome_Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (m *RICsubscriptionDeleteOutcome) GetSuccessfulOutcome() *RICsubscriptionDeleteResponseT {
	if x, ok := m.GetOutcome().(*RICsubscriptionDeleteOutcome_SuccessfulOutcome); ok {
		return x.SuccessfulOutcome
	}
	return nil
}

func (m *RICsubscriptionDeleteOutcome) GetUnsuccessfulOutcome() *RICsubscriptionDeleteFailureT {
	if x, ok := m.GetOutcome().(*RICsubscriptionDeleteOutcome_UnsuccessfulOutcome); ok {
		return x.UnsuccessfulOutcome
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RICsubscriptionDeleteOutcome) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RICsubscriptionDeleteOutcome_SuccessfulOutcome)(nil),
		(*RICsubscriptionDeleteOutcome_UnsuccessfulOutcome)(nil),
	}
}

type RICcontrolOutcome struct {
	// Types that are valid to be assigned to Outcome:
	//	*RICcontrolOutcome_SuccessfulOutcome
	//	*RICcontrolOutcome_UnsuccessfulOutcome
	Outcome              isRICcontrolOutcome_Outcome `protobuf_oneof:"outcome"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RICcontrolOutcome) Reset()         { *m = RICcontrolOutcome{} }
func (m *RICcontrolOutcome) String() string { return proto.CompactTextString(m) }
func (*RICcontrolOutcome) ProtoMessage()    {}
func (*RICcontrolOutcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_59b3e78f6e914a2b, []int{4}
}

func (m *RICcontrolOutcome) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RICcontrolOutcome.Unmarshal(m, b)
}
func (m *RICcontrolOutcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RICcontrolOutcome.Marshal(b, m, deterministic)
}
func (m *RICcontrolOutcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RICcontrolOutcome.Merge(m, src)
}
func (m *RICcontrolOutcome) XXX_Size() int {
	return xxx_messageInfo_RICcontrolOutcome.Size(m)
}
func (m *RICcontrolOutcome) XXX_DiscardUnknown() {
	xxx_messageInfo_RICcontrolOutcome.DiscardUnknown(m)
}

var xxx_messageInfo_RICcontrolOutcome proto.InternalMessageInfo

type isRICcontrolOutcome_Outcome interface {
	isRICcontrolOutcome_Outcome()
}

type RICcontrolOutcome_SuccessfulOutcome struct {
	SuccessfulOutcome *RICcontrolAcknowledgeT `protobuf:"bytes,1,opt,name=successfulOutcome,proto3,oneof"`
}

type RICcontrolOutcome_UnsuccessfulOutcome struct {
	UnsuccessfulOutcome *RICcontrolFailureT `protobuf:"bytes,2,opt,name=unsuccessfulOutcome,proto3,oneof"`
}

func (*RICcontrolOutcome_SuccessfulOutcome) isRICcontrolOutcome_Outcome() {}

func (*RICcontrolOutcome_UnsuccessfulOutcome) isRICcontrolOutcome_Outcome() {}

func (m *RICcontrolOutcome) GetOutcome() isRICcontrolOutcome_Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (m *RICcontrolOutcome) GetSuccessfulOutcome() *RICcontrolAcknowledgeT {
	if x, ok := m.GetOutcome().(*RICcontrolOutcome_SuccessfulOutcome); ok {
		return x.SuccessfulOutcome
	}
	return nil
}

func (m *RICcontrolOutcome) GetUnsuccessfulOutcome() *RICcontrolFailureT {
	if x, ok := m.GetOutcome().(*RICcontrolOutcome_UnsuccessfulOutcome); ok {
		return x.UnsuccessfulOutcome
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RICcontrolOutcome) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RICcontrolOutcome_SuccessfulOutcome)(nil),
		(*RICcontrolOutcome_UnsuccessfulOutcome)(nil),
	}
}

var E_Criticality = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51234,
	Name:          "e2ctypes.criticality",
	Tag:           "bytes,51234,opt,name=criticality",
	Filename:      "e2ap-elementary-procedures.proto",
}

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "e2ctypes.EmptyMessage")
	proto.RegisterType((*E2SetupResponseOutcome)(nil), "e2ctypes.E2setupResponseOutcome")
	proto.RegisterType((*RICsubscriptionOutcome)(nil), "e2ctypes.RICsubscriptionOutcome")
	proto.RegisterType((*RICsubscriptionDeleteOutcome)(nil), "e2ctypes.RICsubscriptionDeleteOutcome")
	proto.RegisterType((*RICcontrolOutcome)(nil), "e2ctypes.RICcontrolOutcome")
	proto.RegisterExtension(E_Criticality)
}

func init() { proto.RegisterFile("e2ap-elementary-procedures.proto", fileDescriptor_59b3e78f6e914a2b) }

var fileDescriptor_59b3e78f6e914a2b = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0x3c, 0x5a, 0x3a, 0x45, 0x54, 0x9d, 0xd2, 0x92, 0x3a, 0x2d, 0x0a, 0xae, 0x84, 0xc2,
	0x47, 0x5d, 0x64, 0xfe, 0xf8, 0xa3, 0x21, 0x88, 0x80, 0x4a, 0xa9, 0x23, 0x55, 0x42, 0x42, 0x8a,
	0x9c, 0xc9, 0x6d, 0x30, 0x38, 0x1e, 0x33, 0x0f, 0xa1, 0xfc, 0xb2, 0x02, 0xc4, 0x12, 0x58, 0x09,
	0x12, 0x2c, 0x80, 0x35, 0xb0, 0x12, 0x24, 0x3f, 0x12, 0x8f, 0x3d, 0x93, 0x28, 0x7f, 0xd1, 0x7d,
	0x9c, 0x73, 0xcf, 0xcd, 0xf5, 0x19, 0xd4, 0x02, 0x2f, 0x48, 0x8e, 0x21, 0x82, 0x09, 0xc4, 0x22,
	0x60, 0xd3, 0xe3, 0x84, 0x51, 0x02, 0x23, 0xc9, 0x80, 0xbb, 0x09, 0xa3, 0x82, 0xe2, 0xdb, 0xe0,
	0x11, 0x31, 0x4d, 0x80, 0xdb, 0xad, 0x31, 0xa5, 0xe3, 0x08, 0x4e, 0xd2, 0xf8, 0x50, 0x5e, 0x9d,
	0x8c, 0x80, 0x13, 0x16, 0x26, 0x82, 0xb2, 0xac, 0xd6, 0xde, 0xed, 0x32, 0x46, 0x59, 0x2f, 0x1e,
	0x85, 0x24, 0x10, 0x21, 0x8d, 0xf3, 0xf0, 0xbd, 0xae, 0xc7, 0x41, 0xc8, 0xc4, 0x87, 0x2f, 0x12,
	0xb8, 0x98, 0x15, 0x17, 0x51, 0x9e, 0xd0, 0x98, 0x43, 0xa5, 0xf8, 0x65, 0x10, 0x46, 0x92, 0x15,
	0xd1, 0x03, 0xbf, 0xd7, 0xe1, 0x72, 0x98, 0xf1, 0x85, 0x34, 0x56, 0xa1, 0x0e, 0x6b, 0x59, 0x05,
	0xb2, 0xda, 0xac, 0x42, 0x3b, 0x95, 0xec, 0x0b, 0x88, 0x40, 0x80, 0x4a, 0x70, 0x64, 0xa8, 0x51,
	0x68, 0xf4, 0x40, 0x2a, 0x19, 0xf6, 0x81, 0x83, 0x50, 0xc1, 0x77, 0xf2, 0x98, 0x02, 0x76, 0xdf,
	0xef, 0x75, 0x08, 0x8d, 0x05, 0xa3, 0x91, 0x5a, 0xdd, 0x9c, 0x27, 0x9e, 0x93, 0xcf, 0x31, 0xfd,
	0x1a, 0xc1, 0x68, 0xac, 0xe9, 0x52, 0x78, 0x9d, 0xbb, 0xe8, 0x4e, 0x77, 0x92, 0x88, 0xe9, 0x19,
	0x70, 0x1e, 0x8c, 0xc1, 0xf9, 0x65, 0xa1, 0xbd, 0xca, 0xfe, 0xcf, 0xa5, 0x20, 0x74, 0x02, 0xf8,
	0x0d, 0xda, 0xe6, 0x92, 0x10, 0xe0, 0xfc, 0x4a, 0x46, 0x79, 0xb0, 0x61, 0xb5, 0xac, 0xf6, 0xa6,
	0xd7, 0x74, 0x8b, 0x63, 0x70, 0x2b, 0xcd, 0x03, 0xf1, 0xea, 0x9a, 0x5f, 0xef, 0xc3, 0x6f, 0xd1,
	0x8e, 0x8c, 0xeb, 0x70, 0xd7, 0x53, 0x38, 0xbb, 0x06, 0x97, 0x0f, 0x9d, 0xa2, 0xe9, 0x1a, 0x4f,
	0x37, 0xd0, 0x3a, 0xcd, 0x7e, 0x3a, 0x7f, 0x2d, 0xb4, 0x57, 0xd9, 0x78, 0xc1, 0xda, 0x37, 0x4b,
	0x38, 0x9a, 0x73, 0x1a, 0x8e, 0xc6, 0x24, 0xe5, 0x72, 0x91, 0x14, 0xc7, 0x08, 0xbb, 0x8a, 0xa4,
	0x7f, 0x16, 0x3a, 0xd0, 0x1e, 0x51, 0x31, 0xc3, 0x7b, 0xb3, 0xb0, 0xc7, 0xc6, 0x09, 0xd4, 0x63,
	0x35, 0xc9, 0xfb, 0xb0, 0x48, 0x5e, 0x7b, 0x09, 0xf8, 0x2a, 0x22, 0xff, 0x58, 0x68, 0x7b, 0x7e,
	0xa6, 0x05, 0xfd, 0x85, 0x59, 0xd9, 0x43, 0x85, 0xbc, 0x7e, 0xfb, 0x26, 0x45, 0x17, 0x8b, 0x14,
	0x1d, 0xea, 0x40, 0x57, 0x90, 0xe1, 0xfd, 0xbe, 0x89, 0x1a, 0x5d, 0x2f, 0x48, 0xba, 0x33, 0xef,
	0x7c, 0x37, 0xb3, 0x4e, 0x4c, 0xd0, 0x96, 0x1f, 0x92, 0x7e, 0x69, 0x4f, 0xd8, 0x59, 0x70, 0x78,
	0xe9, 0xf7, 0x3d, 0x10, 0x76, 0xcb, 0x58, 0x93, 0x0f, 0xe0, 0xa0, 0x1f, 0xdf, 0xf6, 0xd7, 0x18,
	0x7c, 0x02, 0x22, 0xb0, 0x44, 0xbb, 0x15, 0x92, 0xec, 0xcf, 0xc0, 0xed, 0xa5, 0xa7, 0x50, 0x10,
	0x3e, 0x5a, 0x52, 0xa9, 0xa3, 0xbd, 0x44, 0xc8, 0x0f, 0x49, 0x27, 0xdb, 0x18, 0xd6, 0xee, 0x71,
	0x4e, 0xd0, 0xd4, 0xa5, 0x35, 0xa8, 0x4f, 0x2c, 0xdc, 0x47, 0xeb, 0xb9, 0x0b, 0x60, 0x5b, 0xe3,
	0x33, 0x9a, 0x1d, 0xe9, 0x0d, 0x4c, 0x19, 0xf6, 0x35, 0xba, 0x95, 0xba, 0x2b, 0x6e, 0x94, 0x06,
	0x29, 0x59, 0xf0, 0x40, 0xd8, 0xfb, 0xb5, 0x4c, 0xf1, 0xa1, 0x54, 0x06, 0xf4, 0xd1, 0x56, 0xe5,
	0x7d, 0xc3, 0x65, 0x43, 0x54, 0x53, 0x03, 0x61, 0xef, 0x95, 0x92, 0x65, 0xef, 0x4d, 0x51, 0xc3,
	0x71, 0x4c, 0x19, 0x3c, 0x3b, 0x45, 0x9b, 0x84, 0x85, 0x22, 0x24, 0x41, 0x14, 0x8a, 0x29, 0x7e,
	0xe0, 0x66, 0x6f, 0xac, 0x5b, 0xbc, 0xb1, 0xee, 0x19, 0x88, 0x8f, 0x74, 0x74, 0x9e, 0xfe, 0x27,
	0xbc, 0xf1, 0xf3, 0xfb, 0x8d, 0x96, 0xd5, 0xde, 0xf0, 0xcb, 0x4d, 0xc3, 0xb5, 0xb4, 0xf8, 0xe9,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x04, 0xb4, 0x36, 0xcc, 0x07, 0x00, 0x00,
}
