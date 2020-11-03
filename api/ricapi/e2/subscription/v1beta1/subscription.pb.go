// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/ricapi/e2/subscription/v1beta1/subscription.proto

// Package ricapi.subscription.v1beta1 defines the interior gRPC interface for subscription service

package v1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1beta1 "github.com/onosproject/onos-e2t/api/ricapi/e2/headers/v1beta1"
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

// SubscribeResponse a subscription response
type SubscribeResponse struct {
	Header               *v1beta1.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                  `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{0}
}
func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetHeader() *v1beta1.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SubscribeResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// SubscribeDeleteRequest a subscription delete request
type SubscribeDeleteRequest struct {
	Header               *v1beta1.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                 `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SubscribeDeleteRequest) Reset()         { *m = SubscribeDeleteRequest{} }
func (m *SubscribeDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeDeleteRequest) ProtoMessage()    {}
func (*SubscribeDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{1}
}
func (m *SubscribeDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeDeleteRequest.Unmarshal(m, b)
}
func (m *SubscribeDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeDeleteRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeDeleteRequest.Merge(m, src)
}
func (m *SubscribeDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeDeleteRequest.Size(m)
}
func (m *SubscribeDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeDeleteRequest proto.InternalMessageInfo

func (m *SubscribeDeleteRequest) GetHeader() *v1beta1.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SubscribeDeleteRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// SubscribeDeleteResponse a subscription delete response
type SubscribeDeleteResponse struct {
	Header               *v1beta1.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                  `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SubscribeDeleteResponse) Reset()         { *m = SubscribeDeleteResponse{} }
func (m *SubscribeDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeDeleteResponse) ProtoMessage()    {}
func (*SubscribeDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{2}
}
func (m *SubscribeDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeDeleteResponse.Unmarshal(m, b)
}
func (m *SubscribeDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeDeleteResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeDeleteResponse.Merge(m, src)
}
func (m *SubscribeDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeDeleteResponse.Size(m)
}
func (m *SubscribeDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeDeleteResponse proto.InternalMessageInfo

func (m *SubscribeDeleteResponse) GetHeader() *v1beta1.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SubscribeDeleteResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// SubscribeRequest a subscription request
type SubscribeRequest struct {
	Header               *v1beta1.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                 `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{3}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetHeader() *v1beta1.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SubscribeRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeResponse)(nil), "ricapi.e2.subscription.v1beta1.SubscribeResponse")
	proto.RegisterType((*SubscribeDeleteRequest)(nil), "ricapi.e2.subscription.v1beta1.SubscribeDeleteRequest")
	proto.RegisterType((*SubscribeDeleteResponse)(nil), "ricapi.e2.subscription.v1beta1.SubscribeDeleteResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "ricapi.e2.subscription.v1beta1.SubscribeRequest")
}

func init() {
	proto.RegisterFile("api/ricapi/e2/subscription/v1beta1/subscription.proto", fileDescriptor_8732d0de81f42f54)
}

var fileDescriptor_8732d0de81f42f54 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x15, 0x24, 0x8a, 0x74, 0x20, 0x01, 0x46, 0x82, 0x28, 0x03, 0xaa, 0xba, 0x10, 0x06,
	0x1c, 0x12, 0x04, 0xac, 0xb4, 0xea, 0xc0, 0x9c, 0x6c, 0x6c, 0xf9, 0x73, 0xa2, 0x46, 0x25, 0x36,
	0xb6, 0x13, 0x89, 0x07, 0xe0, 0x11, 0x78, 0x5f, 0x44, 0xec, 0xb4, 0x49, 0x91, 0x4a, 0x3b, 0xc0,
	0x94, 0xd8, 0xe7, 0xdf, 0xf7, 0xdd, 0x9d, 0xee, 0xe0, 0x36, 0x15, 0x2c, 0x90, 0x2c, 0xff, 0xfe,
	0x60, 0x14, 0xa8, 0x2a, 0x53, 0xb9, 0x64, 0x42, 0x33, 0x5e, 0x06, 0x75, 0x98, 0xa1, 0x4e, 0xc3,
	0xde, 0x25, 0x15, 0x92, 0x6b, 0x4e, 0xce, 0x0d, 0x42, 0x31, 0xa2, 0xbd, 0xa8, 0x45, 0xbc, 0x8b,
	0xa5, 0xe4, 0x0c, 0xd3, 0x02, 0xa5, 0x5a, 0xa8, 0xd9, 0xb3, 0x11, 0x1a, 0x09, 0x38, 0x4e, 0x8c,
	0x40, 0x86, 0x31, 0x2a, 0xc1, 0x4b, 0x85, 0x64, 0x0c, 0x03, 0xf3, 0xca, 0x75, 0x86, 0x8e, 0xbf,
	0x1f, 0x5d, 0xd2, 0xa5, 0x5d, 0x8b, 0x5b, 0x39, 0xda, 0x42, 0x8f, 0xcd, 0x7d, 0x6c, 0x41, 0xe2,
	0xc2, 0x9e, 0x48, 0xdf, 0xe7, 0x3c, 0x2d, 0xdc, 0xdd, 0xa1, 0xe3, 0x1f, 0xc4, 0xed, 0x71, 0xa4,
	0xe1, 0x74, 0xe1, 0x38, 0xc5, 0x39, 0x6a, 0x8c, 0xf1, 0xad, 0x42, 0xa5, 0xc9, 0xc3, 0x8a, 0xad,
	0xbf, 0xd6, 0xb6, 0x61, 0x36, 0x76, 0xad, 0xe1, 0xec, 0x87, 0xeb, 0x7f, 0x54, 0x5b, 0xc2, 0x51,
	0xa7, 0xbf, 0x7f, 0x5e, 0x67, 0xf4, 0xb9, 0x03, 0x27, 0x49, 0x67, 0x22, 0x12, 0x94, 0x35, 0xcb,
	0x91, 0xd4, 0x70, 0x38, 0x2e, 0x8a, 0x6e, 0x84, 0x5c, 0xd3, 0xf5, 0x43, 0x44, 0x57, 0x13, 0xf7,
	0xc2, 0x2d, 0x08, 0xdb, 0xdc, 0x0f, 0x07, 0x88, 0xe9, 0x77, 0xcf, 0xfb, 0x6e, 0x63, 0xa5, 0xde,
	0x88, 0x78, 0xf7, 0x5b, 0x73, 0x26, 0x8f, 0xc9, 0xf4, 0x69, 0xf2, 0xcc, 0xf4, 0xac, 0xca, 0x68,
	0xce, 0x5f, 0x03, 0x5e, 0x72, 0x25, 0x24, 0x7f, 0xc1, 0x5c, 0x37, 0xff, 0x57, 0x18, 0xe9, 0xe0,
	0xf7, 0x4d, 0xcc, 0x06, 0xcd, 0xd2, 0xdc, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x33, 0x17,
	0x30, 0xb6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	// AddSubscription establish E2 subscriptions on E2 Node.
	AddSubscription(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	// DeleteSubscription delete E2 subscriptions on E2 Node.
	DeleteSubscription(ctx context.Context, in *SubscribeDeleteRequest, opts ...grpc.CallOption) (*SubscribeDeleteResponse, error)
}

type subscriptionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionServiceClient(cc *grpc.ClientConn) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) AddSubscription(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, "/ricapi.e2.subscription.v1beta1.SubscriptionService/AddSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) DeleteSubscription(ctx context.Context, in *SubscribeDeleteRequest, opts ...grpc.CallOption) (*SubscribeDeleteResponse, error) {
	out := new(SubscribeDeleteResponse)
	err := c.cc.Invoke(ctx, "/ricapi.e2.subscription.v1beta1.SubscriptionService/DeleteSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
type SubscriptionServiceServer interface {
	// AddSubscription establish E2 subscriptions on E2 Node.
	AddSubscription(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	// DeleteSubscription delete E2 subscriptions on E2 Node.
	DeleteSubscription(context.Context, *SubscribeDeleteRequest) (*SubscribeDeleteResponse, error)
}

// UnimplementedSubscriptionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (*UnimplementedSubscriptionServiceServer) AddSubscription(ctx context.Context, req *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) DeleteSubscription(ctx context.Context, req *SubscribeDeleteRequest) (*SubscribeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}

func RegisterSubscriptionServiceServer(s *grpc.Server, srv SubscriptionServiceServer) {
	s.RegisterService(&_SubscriptionService_serviceDesc, srv)
}

func _SubscriptionService_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).AddSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ricapi.e2.subscription.v1beta1.SubscriptionService/AddSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).AddSubscription(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ricapi.e2.subscription.v1beta1.SubscriptionService/DeleteSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, req.(*SubscribeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubscriptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ricapi.e2.subscription.v1beta1.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSubscription",
			Handler:    _SubscriptionService_AddSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _SubscriptionService_DeleteSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ricapi/e2/subscription/v1beta1/subscription.proto",
}
