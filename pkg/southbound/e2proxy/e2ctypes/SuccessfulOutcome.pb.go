// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SuccessfulOutcome.proto

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

type SuccessfulOutcomeT struct {
	ProcedureCode ProcedureCodeT `protobuf:"varint,1,opt,name=procedureCode,proto3,enum=e2ctypes.ProcedureCodeT" json:"procedureCode,omitempty"`
	Criticality   CriticalityT   `protobuf:"varint,2,opt,name=criticality,proto3,enum=e2ctypes.CriticalityT" json:"criticality,omitempty"`
	// Types that are valid to be assigned to Choice:
	//	*SuccessfulOutcomeT_RICsubscriptionResponse
	//	*SuccessfulOutcomeT_RICsubscriptionDeleteResponse
	//	*SuccessfulOutcomeT_RICserviceUpdateAcknowledge
	//	*SuccessfulOutcomeT_RICcontrolAcknowledge
	//	*SuccessfulOutcomeT_E2SetupResponse
	//	*SuccessfulOutcomeT_ResetResponse
	Choice               isSuccessfulOutcomeT_Choice `protobuf_oneof:"choice"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SuccessfulOutcomeT) Reset()         { *m = SuccessfulOutcomeT{} }
func (m *SuccessfulOutcomeT) String() string { return proto.CompactTextString(m) }
func (*SuccessfulOutcomeT) ProtoMessage()    {}
func (*SuccessfulOutcomeT) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d8b79d17aef3d40, []int{0}
}

func (m *SuccessfulOutcomeT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessfulOutcomeT.Unmarshal(m, b)
}
func (m *SuccessfulOutcomeT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessfulOutcomeT.Marshal(b, m, deterministic)
}
func (m *SuccessfulOutcomeT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessfulOutcomeT.Merge(m, src)
}
func (m *SuccessfulOutcomeT) XXX_Size() int {
	return xxx_messageInfo_SuccessfulOutcomeT.Size(m)
}
func (m *SuccessfulOutcomeT) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessfulOutcomeT.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessfulOutcomeT proto.InternalMessageInfo

func (m *SuccessfulOutcomeT) GetProcedureCode() ProcedureCodeT {
	if m != nil {
		return m.ProcedureCode
	}
	return ProcedureCodeT_ProcedureCode_id_dummy
}

func (m *SuccessfulOutcomeT) GetCriticality() CriticalityT {
	if m != nil {
		return m.Criticality
	}
	return CriticalityT_Criticality_reject
}

type isSuccessfulOutcomeT_Choice interface {
	isSuccessfulOutcomeT_Choice()
}

type SuccessfulOutcomeT_RICsubscriptionResponse struct {
	RICsubscriptionResponse *RICsubscriptionResponseT `protobuf:"bytes,3,opt,name=RICsubscriptionResponse,proto3,oneof"`
}

type SuccessfulOutcomeT_RICsubscriptionDeleteResponse struct {
	RICsubscriptionDeleteResponse *RICsubscriptionDeleteResponseT `protobuf:"bytes,4,opt,name=RICsubscriptionDeleteResponse,proto3,oneof"`
}

type SuccessfulOutcomeT_RICserviceUpdateAcknowledge struct {
	RICserviceUpdateAcknowledge *RICserviceUpdateAcknowledgeT `protobuf:"bytes,5,opt,name=RICserviceUpdateAcknowledge,proto3,oneof"`
}

type SuccessfulOutcomeT_RICcontrolAcknowledge struct {
	RICcontrolAcknowledge *RICcontrolAcknowledgeT `protobuf:"bytes,6,opt,name=RICcontrolAcknowledge,proto3,oneof"`
}

type SuccessfulOutcomeT_E2SetupResponse struct {
	E2SetupResponse *E2SetupResponseT `protobuf:"bytes,7,opt,name=E2setupResponse,proto3,oneof"`
}

type SuccessfulOutcomeT_ResetResponse struct {
	ResetResponse *ResetResponseT `protobuf:"bytes,8,opt,name=ResetResponse,proto3,oneof"`
}

func (*SuccessfulOutcomeT_RICsubscriptionResponse) isSuccessfulOutcomeT_Choice() {}

func (*SuccessfulOutcomeT_RICsubscriptionDeleteResponse) isSuccessfulOutcomeT_Choice() {}

func (*SuccessfulOutcomeT_RICserviceUpdateAcknowledge) isSuccessfulOutcomeT_Choice() {}

func (*SuccessfulOutcomeT_RICcontrolAcknowledge) isSuccessfulOutcomeT_Choice() {}

func (*SuccessfulOutcomeT_E2SetupResponse) isSuccessfulOutcomeT_Choice() {}

func (*SuccessfulOutcomeT_ResetResponse) isSuccessfulOutcomeT_Choice() {}

func (m *SuccessfulOutcomeT) GetChoice() isSuccessfulOutcomeT_Choice {
	if m != nil {
		return m.Choice
	}
	return nil
}

func (m *SuccessfulOutcomeT) GetRICsubscriptionResponse() *RICsubscriptionResponseT {
	if x, ok := m.GetChoice().(*SuccessfulOutcomeT_RICsubscriptionResponse); ok {
		return x.RICsubscriptionResponse
	}
	return nil
}

func (m *SuccessfulOutcomeT) GetRICsubscriptionDeleteResponse() *RICsubscriptionDeleteResponseT {
	if x, ok := m.GetChoice().(*SuccessfulOutcomeT_RICsubscriptionDeleteResponse); ok {
		return x.RICsubscriptionDeleteResponse
	}
	return nil
}

func (m *SuccessfulOutcomeT) GetRICserviceUpdateAcknowledge() *RICserviceUpdateAcknowledgeT {
	if x, ok := m.GetChoice().(*SuccessfulOutcomeT_RICserviceUpdateAcknowledge); ok {
		return x.RICserviceUpdateAcknowledge
	}
	return nil
}

func (m *SuccessfulOutcomeT) GetRICcontrolAcknowledge() *RICcontrolAcknowledgeT {
	if x, ok := m.GetChoice().(*SuccessfulOutcomeT_RICcontrolAcknowledge); ok {
		return x.RICcontrolAcknowledge
	}
	return nil
}

func (m *SuccessfulOutcomeT) GetE2SetupResponse() *E2SetupResponseT {
	if x, ok := m.GetChoice().(*SuccessfulOutcomeT_E2SetupResponse); ok {
		return x.E2SetupResponse
	}
	return nil
}

func (m *SuccessfulOutcomeT) GetResetResponse() *ResetResponseT {
	if x, ok := m.GetChoice().(*SuccessfulOutcomeT_ResetResponse); ok {
		return x.ResetResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SuccessfulOutcomeT) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SuccessfulOutcomeT_RICsubscriptionResponse)(nil),
		(*SuccessfulOutcomeT_RICsubscriptionDeleteResponse)(nil),
		(*SuccessfulOutcomeT_RICserviceUpdateAcknowledge)(nil),
		(*SuccessfulOutcomeT_RICcontrolAcknowledge)(nil),
		(*SuccessfulOutcomeT_E2SetupResponse)(nil),
		(*SuccessfulOutcomeT_ResetResponse)(nil),
	}
}

func init() {
	proto.RegisterType((*SuccessfulOutcomeT)(nil), "e2ctypes.SuccessfulOutcome_t")
}

func init() { proto.RegisterFile("SuccessfulOutcome.proto", fileDescriptor_1d8b79d17aef3d40) }

var fileDescriptor_1d8b79d17aef3d40 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x86, 0x3f, 0x3e, 0x05, 0xc9, 0x12, 0x34, 0x2e, 0x21, 0xac, 0x10, 0x12, 0x90, 0x83, 0x78,
	0xe1, 0x50, 0x4f, 0x9e, 0x0c, 0x56, 0xa3, 0x9c, 0x34, 0x35, 0x1e, 0x3c, 0x35, 0xb0, 0x1d, 0xb5,
	0xa1, 0x76, 0xd7, 0xdd, 0xa9, 0x86, 0xdf, 0xe7, 0x1f, 0x33, 0x6e, 0x20, 0xed, 0x16, 0xda, 0xeb,
	0xce, 0x33, 0xcf, 0x9b, 0x9d, 0x19, 0xd2, 0x79, 0x4a, 0x38, 0x07, 0xad, 0x5f, 0x93, 0xe8, 0x21,
	0x41, 0x2e, 0x3e, 0x60, 0x22, 0x95, 0x40, 0x41, 0xeb, 0xe0, 0x70, 0x5c, 0x49, 0xd0, 0xdd, 0xd6,
	0xa3, 0x12, 0x1c, 0x82, 0x44, 0x81, 0x2b, 0x82, 0x75, 0xb9, 0x7b, 0xec, 0xaa, 0x10, 0x43, 0x3e,
	0x8f, 0x42, 0x5c, 0xad, 0x9f, 0xfa, 0xde, 0xcc, 0xd5, 0xc9, 0x42, 0x73, 0x15, 0x4a, 0x0c, 0x45,
	0xec, 0x81, 0x96, 0x22, 0xd6, 0x9b, 0x8e, 0x51, 0xae, 0x7c, 0x03, 0x11, 0x20, 0xe4, 0xa0, 0xe1,
	0x1f, 0x04, 0xea, 0x2b, 0xe4, 0xf0, 0x2c, 0x83, 0x39, 0xc2, 0x94, 0x2f, 0x63, 0xf1, 0x1d, 0x41,
	0xf0, 0xb6, 0x41, 0x7a, 0xde, 0xcc, 0xe5, 0x22, 0x46, 0x25, 0xa2, 0xed, 0x62, 0xfb, 0xd6, 0xd1,
	0x80, 0x89, 0xcc, 0x69, 0x5b, 0x1e, 0x68, 0x40, 0xfb, 0xf1, 0xf4, 0xa7, 0x4a, 0x5a, 0x5b, 0xbf,
	0xf7, 0x91, 0x5e, 0x91, 0xa6, 0xcc, 0xfe, 0x98, 0x55, 0x06, 0x95, 0xf1, 0xa1, 0x73, 0x32, 0xd9,
	0x4c, 0x64, 0x62, 0x0d, 0xc4, 0x47, 0xcf, 0xe6, 0xe9, 0x25, 0x69, 0xf0, 0x74, 0x3a, 0xec, 0xbf,
	0x69, 0xef, 0xa4, 0xed, 0x99, 0xd1, 0xf9, 0xe8, 0x65, 0x59, 0xea, 0x93, 0x4e, 0xc1, 0x14, 0xd9,
	0xde, 0xa0, 0x32, 0x6e, 0x38, 0xa3, 0x54, 0x53, 0x00, 0xfa, 0x78, 0xff, 0xcf, 0x2b, 0xb2, 0xd0,
	0x4f, 0xd2, 0x2f, 0xdd, 0x03, 0xdb, 0x37, 0x31, 0xe7, 0x85, 0x31, 0x36, 0x6e, 0xc2, 0xca, 0x8d,
	0x74, 0x49, 0x7a, 0x25, 0x5b, 0x65, 0x55, 0x13, 0x78, 0x66, 0x07, 0x16, 0xc0, 0x26, 0xae, 0xcc,
	0x46, 0x5f, 0x48, 0x7b, 0xe7, 0x7d, 0xb0, 0x9a, 0x89, 0x19, 0x5a, 0x31, 0xdb, 0x98, 0x09, 0xd8,
	0x6d, 0xa0, 0x77, 0xe4, 0x28, 0x77, 0x5d, 0xec, 0xc0, 0x48, 0x7b, 0xa9, 0x34, 0x07, 0x18, 0x5d,
	0xbe, 0x8b, 0x4e, 0x49, 0xd3, 0xba, 0x47, 0x56, 0x37, 0x9a, 0xcc, 0x81, 0x59, 0x65, 0x23, 0xb1,
	0x3b, 0xae, 0xeb, 0xa4, 0xc6, 0xdf, 0x45, 0xc8, 0x61, 0x51, 0x33, 0xc7, 0x7c, 0xf1, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xe9, 0xb7, 0x4d, 0xe0, 0xc9, 0x03, 0x00, 0x00,
}
