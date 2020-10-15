// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/xapp/e2t.proto

// Package xapp defines the interior gRPC interfaces for xApps to interact with E2T.

package xapp

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type E2Control struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *E2Control) Reset()         { *m = E2Control{} }
func (m *E2Control) String() string { return proto.CompactTextString(m) }
func (*E2Control) ProtoMessage()    {}
func (*E2Control) Descriptor() ([]byte, []int) {
	return fileDescriptor_31980201a64b57f1, []int{0}
}
func (m *E2Control) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_E2Control.Unmarshal(m, b)
}
func (m *E2Control) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_E2Control.Marshal(b, m, deterministic)
}
func (m *E2Control) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E2Control.Merge(m, src)
}
func (m *E2Control) XXX_Size() int {
	return xxx_messageInfo_E2Control.Size(m)
}
func (m *E2Control) XXX_DiscardUnknown() {
	xxx_messageInfo_E2Control.DiscardUnknown(m)
}

var xxx_messageInfo_E2Control proto.InternalMessageInfo

type E2Message struct {
	// ID of E2 node that sent the message
	E2Node string `protobuf:"bytes,1,opt,name=e2node,proto3" json:"e2node,omitempty"`
	// Service model ID
	ServiceModel string `protobuf:"bytes,2,opt,name=service_model,json=serviceModel,proto3" json:"service_model,omitempty"`
	// Message data (encoded as ASN.1 or Protobuf)
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *E2Message) Reset()         { *m = E2Message{} }
func (m *E2Message) String() string { return proto.CompactTextString(m) }
func (*E2Message) ProtoMessage()    {}
func (*E2Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_31980201a64b57f1, []int{1}
}
func (m *E2Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_E2Message.Unmarshal(m, b)
}
func (m *E2Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_E2Message.Marshal(b, m, deterministic)
}
func (m *E2Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E2Message.Merge(m, src)
}
func (m *E2Message) XXX_Size() int {
	return xxx_messageInfo_E2Message.Size(m)
}
func (m *E2Message) XXX_DiscardUnknown() {
	xxx_messageInfo_E2Message.DiscardUnknown(m)
}

var xxx_messageInfo_E2Message proto.InternalMessageInfo

func (m *E2Message) GetE2Node() string {
	if m != nil {
		return m.E2Node
	}
	return ""
}

func (m *E2Message) GetServiceModel() string {
	if m != nil {
		return m.ServiceModel
	}
	return ""
}

func (m *E2Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*E2Control)(nil), "onos.e2t.xapp.E2Control")
	proto.RegisterType((*E2Message)(nil), "onos.e2t.xapp.E2Message")
}

func init() { proto.RegisterFile("api/xapp/e2t.proto", fileDescriptor_31980201a64b57f1) }

var fileDescriptor_31980201a64b57f1 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x46, 0x89, 0x42, 0xa5, 0xd3, 0xf6, 0x92, 0x83, 0x84, 0x1e, 0xa4, 0xd4, 0xcb, 0x9e, 0x12,
	0x89, 0xbf, 0x40, 0x64, 0x8f, 0xbd, 0x6c, 0xbd, 0x4b, 0xd6, 0x8c, 0x21, 0x10, 0x33, 0x21, 0x49,
	0x45, 0xff, 0xbd, 0x64, 0xbb, 0x0a, 0xd2, 0xe3, 0x9b, 0xf7, 0x1d, 0xde, 0x00, 0x37, 0xc9, 0xab,
	0x2f, 0x93, 0x92, 0x42, 0x5d, 0x65, 0xca, 0x54, 0x89, 0x6f, 0x28, 0x52, 0x91, 0x8d, 0x9b, 0xd8,
	0xde, 0x39, 0x22, 0x17, 0x50, 0x4d, 0x72, 0x3c, 0xbd, 0x2b, 0x7b, 0xca, 0xa6, 0x7a, 0x8a, 0xe7,
	0xf9, 0x7e, 0x05, 0xcb, 0x5e, 0x3f, 0x53, 0xac, 0x99, 0xc2, 0x7e, 0x6c, 0x70, 0xc0, 0x52, 0x8c,
	0x43, 0x7e, 0x0b, 0x0b, 0xd4, 0x91, 0x2c, 0x0a, 0xb6, 0x63, 0xdd, 0x72, 0x98, 0x89, 0xdf, 0xc3,
	0xa6, 0x60, 0xfe, 0xf4, 0x6f, 0xf8, 0xfa, 0x41, 0x16, 0x83, 0xb8, 0x9a, 0xf4, 0x7a, 0x3e, 0x1e,
	0xda, 0x8d, 0x0b, 0xb8, 0x49, 0xe6, 0x3b, 0x90, 0xb1, 0xe2, 0x7a, 0xc7, 0xba, 0xf5, 0xf0, 0x8b,
	0xfa, 0x08, 0xd0, 0xeb, 0x97, 0xe3, 0x79, 0xcc, 0x7b, 0x58, 0x0d, 0xe8, 0x7c, 0xa9, 0x98, 0x9f,
	0x52, 0xe2, 0x42, 0xfe, 0xab, 0x97, 0x7f, 0x69, 0xdb, 0x4b, 0x33, 0x77, 0x76, 0xec, 0x81, 0x8d,
	0x8b, 0xe9, 0x99, 0xc7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x23, 0x6e, 0x55, 0x11, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// E2TServiceClient is the client API for E2TService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type E2TServiceClient interface {
	// RegisterApp establishes a bi-directional stream for conducting interactions with the E2 nodes in the RAN environment.
	RegisterApp(ctx context.Context, opts ...grpc.CallOption) (E2TService_RegisterAppClient, error)
}

type e2TServiceClient struct {
	cc *grpc.ClientConn
}

func NewE2TServiceClient(cc *grpc.ClientConn) E2TServiceClient {
	return &e2TServiceClient{cc}
}

func (c *e2TServiceClient) RegisterApp(ctx context.Context, opts ...grpc.CallOption) (E2TService_RegisterAppClient, error) {
	stream, err := c.cc.NewStream(ctx, &_E2TService_serviceDesc.Streams[0], "/onos.e2t.xapp.E2TService/RegisterApp", opts...)
	if err != nil {
		return nil, err
	}
	x := &e2TServiceRegisterAppClient{stream}
	return x, nil
}

type E2TService_RegisterAppClient interface {
	Send(*E2Control) error
	Recv() (*E2Message, error)
	grpc.ClientStream
}

type e2TServiceRegisterAppClient struct {
	grpc.ClientStream
}

func (x *e2TServiceRegisterAppClient) Send(m *E2Control) error {
	return x.ClientStream.SendMsg(m)
}

func (x *e2TServiceRegisterAppClient) Recv() (*E2Message, error) {
	m := new(E2Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// E2TServiceServer is the server API for E2TService service.
type E2TServiceServer interface {
	// RegisterApp establishes a bi-directional stream for conducting interactions with the E2 nodes in the RAN environment.
	RegisterApp(E2TService_RegisterAppServer) error
}

// UnimplementedE2TServiceServer can be embedded to have forward compatible implementations.
type UnimplementedE2TServiceServer struct {
}

func (*UnimplementedE2TServiceServer) RegisterApp(srv E2TService_RegisterAppServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterApp not implemented")
}

func RegisterE2TServiceServer(s *grpc.Server, srv E2TServiceServer) {
	s.RegisterService(&_E2TService_serviceDesc, srv)
}

func _E2TService_RegisterApp_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(E2TServiceServer).RegisterApp(&e2TServiceRegisterAppServer{stream})
}

type E2TService_RegisterAppServer interface {
	Send(*E2Message) error
	Recv() (*E2Control, error)
	grpc.ServerStream
}

type e2TServiceRegisterAppServer struct {
	grpc.ServerStream
}

func (x *e2TServiceRegisterAppServer) Send(m *E2Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *e2TServiceRegisterAppServer) Recv() (*E2Control, error) {
	m := new(E2Control)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _E2TService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.e2t.xapp.E2TService",
	HandlerType: (*E2TServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterApp",
			Handler:       _E2TService_RegisterApp_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/xapp/e2t.proto",
}
