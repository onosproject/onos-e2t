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

// AddSubscribtionResponse a subscription response
type AddSubscribtionResponse struct {
	Header               *v1beta1.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                  `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AddSubscribtionResponse) Reset()         { *m = AddSubscribtionResponse{} }
func (m *AddSubscribtionResponse) String() string { return proto.CompactTextString(m) }
func (*AddSubscribtionResponse) ProtoMessage()    {}
func (*AddSubscribtionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{0}
}
func (m *AddSubscribtionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSubscribtionResponse.Unmarshal(m, b)
}
func (m *AddSubscribtionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSubscribtionResponse.Marshal(b, m, deterministic)
}
func (m *AddSubscribtionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSubscribtionResponse.Merge(m, src)
}
func (m *AddSubscribtionResponse) XXX_Size() int {
	return xxx_messageInfo_AddSubscribtionResponse.Size(m)
}
func (m *AddSubscribtionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSubscribtionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddSubscribtionResponse proto.InternalMessageInfo

func (m *AddSubscribtionResponse) GetHeader() *v1beta1.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AddSubscribtionResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// DeleteSubscriptionRequest a subscription delete request
type DeleteSubscriptionRequest struct {
	Header               *v1beta1.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                 `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DeleteSubscriptionRequest) Reset()         { *m = DeleteSubscriptionRequest{} }
func (m *DeleteSubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSubscriptionRequest) ProtoMessage()    {}
func (*DeleteSubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{1}
}
func (m *DeleteSubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSubscriptionRequest.Unmarshal(m, b)
}
func (m *DeleteSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSubscriptionRequest.Merge(m, src)
}
func (m *DeleteSubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSubscriptionRequest.Size(m)
}
func (m *DeleteSubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSubscriptionRequest proto.InternalMessageInfo

func (m *DeleteSubscriptionRequest) GetHeader() *v1beta1.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DeleteSubscriptionRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// DeleteSubscriptionResponse a subscription delete response
type DeleteSubscriptionResponse struct {
	Header               *v1beta1.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                  `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DeleteSubscriptionResponse) Reset()         { *m = DeleteSubscriptionResponse{} }
func (m *DeleteSubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteSubscriptionResponse) ProtoMessage()    {}
func (*DeleteSubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{2}
}
func (m *DeleteSubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSubscriptionResponse.Unmarshal(m, b)
}
func (m *DeleteSubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *DeleteSubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSubscriptionResponse.Merge(m, src)
}
func (m *DeleteSubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteSubscriptionResponse.Size(m)
}
func (m *DeleteSubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSubscriptionResponse proto.InternalMessageInfo

func (m *DeleteSubscriptionResponse) GetHeader() *v1beta1.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DeleteSubscriptionResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// AddSubscriptionRequest a subscription request
type AddSubscriptionRequest struct {
	Header               *v1beta1.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                 `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AddSubscriptionRequest) Reset()         { *m = AddSubscriptionRequest{} }
func (m *AddSubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*AddSubscriptionRequest) ProtoMessage()    {}
func (*AddSubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8732d0de81f42f54, []int{3}
}
func (m *AddSubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSubscriptionRequest.Unmarshal(m, b)
}
func (m *AddSubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *AddSubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSubscriptionRequest.Merge(m, src)
}
func (m *AddSubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_AddSubscriptionRequest.Size(m)
}
func (m *AddSubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSubscriptionRequest proto.InternalMessageInfo

func (m *AddSubscriptionRequest) GetHeader() *v1beta1.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AddSubscriptionRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*AddSubscribtionResponse)(nil), "ricapi.e2.subscription.v1beta1.AddSubscribtionResponse")
	proto.RegisterType((*DeleteSubscriptionRequest)(nil), "ricapi.e2.subscription.v1beta1.DeleteSubscriptionRequest")
	proto.RegisterType((*DeleteSubscriptionResponse)(nil), "ricapi.e2.subscription.v1beta1.DeleteSubscriptionResponse")
	proto.RegisterType((*AddSubscriptionRequest)(nil), "ricapi.e2.subscription.v1beta1.AddSubscriptionRequest")
}

func init() {
	proto.RegisterFile("api/ricapi/e2/subscription/v1beta1/subscription.proto", fileDescriptor_8732d0de81f42f54)
}

var fileDescriptor_8732d0de81f42f54 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0x2c, 0xc8, 0xd4,
	0x2f, 0xca, 0x4c, 0x06, 0x51, 0xa9, 0x46, 0xfa, 0xc5, 0xa5, 0x49, 0xc5, 0xc9, 0x45, 0x99, 0x05,
	0x25, 0x99, 0xf9, 0x79, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0x28, 0x82, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x72, 0x10, 0x2d, 0x7a, 0xa9, 0x46, 0x7a, 0x28, 0xb2, 0x50, 0x2d,
	0x52, 0xea, 0x08, 0x23, 0x33, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x8a, 0xe1, 0xa6, 0x41, 0xf9, 0x10,
	0x83, 0x94, 0xca, 0xb8, 0xc4, 0x1d, 0x53, 0x52, 0x82, 0x21, 0x66, 0x24, 0x81, 0xcc, 0x08, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x72, 0xe4, 0x62, 0x83, 0xa8, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x36, 0xd2, 0xd4, 0x43, 0x58, 0x0a, 0x33, 0x04, 0x6a, 0xa8, 0x1e, 0x4c, 0x93, 0x07,
	0x58, 0x3c, 0x08, 0xaa, 0x51, 0x48, 0x82, 0x8b, 0xbd, 0x20, 0xb1, 0x32, 0x27, 0x3f, 0x31, 0x45,
	0x82, 0x55, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc6, 0x55, 0x2a, 0xe7, 0x92, 0x74, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x0d, 0x46, 0x72, 0x7e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x03, 0x9a,
	0xcd, 0x1a, 0x78, 0x6d, 0x06, 0xeb, 0x21, 0xda, 0xe2, 0x4a, 0x2e, 0x29, 0x6c, 0x16, 0xd3, 0xc3,
	0xcf, 0x25, 0x5c, 0x62, 0x88, 0xb0, 0xa6, 0x97, 0x87, 0x8d, 0x56, 0x30, 0x71, 0x09, 0x23, 0xdb,
	0x19, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xd4, 0xc4, 0xc8, 0xc5, 0x8f, 0xe6, 0x1c, 0x21,
	0x33, 0x3d, 0xfc, 0xe9, 0x4a, 0x0f, 0xbb, 0xfb, 0xa5, 0xcc, 0x89, 0xd7, 0x87, 0x9a, 0xc6, 0xba,
	0x19, 0xb9, 0x84, 0x30, 0xa3, 0x43, 0xc8, 0x92, 0x90, 0x79, 0x38, 0xd3, 0x8e, 0x94, 0x15, 0x39,
	0x5a, 0x21, 0xae, 0x71, 0x72, 0x89, 0x72, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0xcf, 0xcf, 0xcb, 0x2f, 0x2e, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x01, 0xb3, 0x75,
	0x53, 0x8d, 0x4a, 0xf4, 0x09, 0x67, 0xd7, 0x24, 0x36, 0x70, 0xce, 0x32, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xac, 0x75, 0xae, 0xf5, 0xdb, 0x03, 0x00, 0x00,
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
	AddSubscription(ctx context.Context, in *AddSubscriptionRequest, opts ...grpc.CallOption) (*AddSubscribtionResponse, error)
	// DeleteSubscription delete E2 subscriptions on E2 Node.
	DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*DeleteSubscriptionResponse, error)
}

type subscriptionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionServiceClient(cc *grpc.ClientConn) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) AddSubscription(ctx context.Context, in *AddSubscriptionRequest, opts ...grpc.CallOption) (*AddSubscribtionResponse, error) {
	out := new(AddSubscribtionResponse)
	err := c.cc.Invoke(ctx, "/ricapi.e2.subscription.v1beta1.SubscriptionService/AddSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*DeleteSubscriptionResponse, error) {
	out := new(DeleteSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/ricapi.e2.subscription.v1beta1.SubscriptionService/DeleteSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
type SubscriptionServiceServer interface {
	// AddSubscription establish E2 subscriptions on E2 Node.
	AddSubscription(context.Context, *AddSubscriptionRequest) (*AddSubscribtionResponse, error)
	// DeleteSubscription delete E2 subscriptions on E2 Node.
	DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*DeleteSubscriptionResponse, error)
}

// UnimplementedSubscriptionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (*UnimplementedSubscriptionServiceServer) AddSubscription(ctx context.Context, req *AddSubscriptionRequest) (*AddSubscribtionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (*UnimplementedSubscriptionServiceServer) DeleteSubscription(ctx context.Context, req *DeleteSubscriptionRequest) (*DeleteSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}

func RegisterSubscriptionServiceServer(s *grpc.Server, srv SubscriptionServiceServer) {
	s.RegisterService(&_SubscriptionService_serviceDesc, srv)
}

func _SubscriptionService_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubscriptionRequest)
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
		return srv.(SubscriptionServiceServer).AddSubscription(ctx, req.(*AddSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriptionRequest)
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
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, req.(*DeleteSubscriptionRequest))
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
