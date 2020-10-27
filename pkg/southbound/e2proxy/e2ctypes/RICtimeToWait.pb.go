// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RICtimeToWait.proto

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

type RICtimeToWaitT int32 // Deprecated: Do not use.
const (
	RICtimeToWaitT_RICtimeToWait_zero   RICtimeToWaitT = 0
	RICtimeToWaitT_RICtimeToWait_w1ms   RICtimeToWaitT = 1
	RICtimeToWaitT_RICtimeToWait_w2ms   RICtimeToWaitT = 2
	RICtimeToWaitT_RICtimeToWait_w5ms   RICtimeToWaitT = 3
	RICtimeToWaitT_RICtimeToWait_w10ms  RICtimeToWaitT = 4
	RICtimeToWaitT_RICtimeToWait_w20ms  RICtimeToWaitT = 5
	RICtimeToWaitT_RICtimeToWait_w30ms  RICtimeToWaitT = 6
	RICtimeToWaitT_RICtimeToWait_w40ms  RICtimeToWaitT = 7
	RICtimeToWaitT_RICtimeToWait_w50ms  RICtimeToWaitT = 8
	RICtimeToWaitT_RICtimeToWait_w100ms RICtimeToWaitT = 9
	RICtimeToWaitT_RICtimeToWait_w200ms RICtimeToWaitT = 10
	RICtimeToWaitT_RICtimeToWait_w500ms RICtimeToWaitT = 11
	RICtimeToWaitT_RICtimeToWait_w1s    RICtimeToWaitT = 12
	RICtimeToWaitT_RICtimeToWait_w2s    RICtimeToWaitT = 13
	RICtimeToWaitT_RICtimeToWait_w5s    RICtimeToWaitT = 14
	RICtimeToWaitT_RICtimeToWait_w10s   RICtimeToWaitT = 15
	RICtimeToWaitT_RICtimeToWait_w20s   RICtimeToWaitT = 16
	RICtimeToWaitT_RICtimeToWait_w60s   RICtimeToWaitT = 17
)

var RICtimeToWaitT_name = map[int32]string{
	0:  "RICtimeToWait_zero",
	1:  "RICtimeToWait_w1ms",
	2:  "RICtimeToWait_w2ms",
	3:  "RICtimeToWait_w5ms",
	4:  "RICtimeToWait_w10ms",
	5:  "RICtimeToWait_w20ms",
	6:  "RICtimeToWait_w30ms",
	7:  "RICtimeToWait_w40ms",
	8:  "RICtimeToWait_w50ms",
	9:  "RICtimeToWait_w100ms",
	10: "RICtimeToWait_w200ms",
	11: "RICtimeToWait_w500ms",
	12: "RICtimeToWait_w1s",
	13: "RICtimeToWait_w2s",
	14: "RICtimeToWait_w5s",
	15: "RICtimeToWait_w10s",
	16: "RICtimeToWait_w20s",
	17: "RICtimeToWait_w60s",
}

var RICtimeToWaitT_value = map[string]int32{
	"RICtimeToWait_zero":   0,
	"RICtimeToWait_w1ms":   1,
	"RICtimeToWait_w2ms":   2,
	"RICtimeToWait_w5ms":   3,
	"RICtimeToWait_w10ms":  4,
	"RICtimeToWait_w20ms":  5,
	"RICtimeToWait_w30ms":  6,
	"RICtimeToWait_w40ms":  7,
	"RICtimeToWait_w50ms":  8,
	"RICtimeToWait_w100ms": 9,
	"RICtimeToWait_w200ms": 10,
	"RICtimeToWait_w500ms": 11,
	"RICtimeToWait_w1s":    12,
	"RICtimeToWait_w2s":    13,
	"RICtimeToWait_w5s":    14,
	"RICtimeToWait_w10s":   15,
	"RICtimeToWait_w20s":   16,
	"RICtimeToWait_w60s":   17,
}

func (x RICtimeToWaitT) String() string {
	return proto.EnumName(RICtimeToWaitT_name, int32(x))
}

func (RICtimeToWaitT) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c15599288c1fefe, []int{0}
}

func init() {
	proto.RegisterEnum("e2ctypes.RICtimeToWaitT", RICtimeToWaitT_name, RICtimeToWaitT_value)
}

func init() { proto.RegisterFile("RICtimeToWait.proto", fileDescriptor_3c15599288c1fefe) }

var fileDescriptor_3c15599288c1fefe = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xcd, 0x8e, 0x82, 0x30,
	0x10, 0xc0, 0xf1, 0x05, 0x76, 0x59, 0xb6, 0xfb, 0xc1, 0xd0, 0xf5, 0x83, 0xf8, 0x08, 0x1e, 0x4c,
	0x29, 0xd6, 0x17, 0xf0, 0xe4, 0xd5, 0x98, 0x78, 0x24, 0x6a, 0x7a, 0xe0, 0xd0, 0x94, 0x30, 0x4d,
	0x88, 0xbe, 0xa3, 0xef, 0x64, 0xca, 0x0d, 0x98, 0xeb, 0xff, 0xd7, 0xb4, 0xcd, 0x0c, 0xfb, 0x3f,
	0x1e, 0xf6, 0xae, 0x36, 0xfa, 0x64, 0xcf, 0x97, 0xda, 0x6d, 0x9a, 0xd6, 0x3a, 0xcb, 0x13, 0x2d,
	0x6f, 0xee, 0xde, 0x68, 0x5c, 0x3f, 0x23, 0x96, 0x0e, 0x4e, 0x54, 0x8e, 0x2f, 0x18, 0x1f, 0xa6,
	0x87, 0x6e, 0x2d, 0xbc, 0x4d, 0x7b, 0x57, 0x18, 0x84, 0x80, 0xe8, 0xd2, 0x20, 0x84, 0x44, 0x57,
	0x06, 0x21, 0xe2, 0xcb, 0xd1, 0xa7, 0xaa, 0xae, 0x10, 0x06, 0xe1, 0x9d, 0x00, 0xe9, 0xe1, 0x83,
	0x80, 0xd2, 0x43, 0x4c, 0xc0, 0xd6, 0xc3, 0x27, 0x01, 0xca, 0x43, 0xc2, 0x73, 0x36, 0x9b, 0x3c,
	0xee, 0xe5, 0x8b, 0x10, 0xd9, 0x0b, 0x23, 0x44, 0xf5, 0xf2, 0xcd, 0xe7, 0x2c, 0x1b, 0xdf, 0x86,
	0xf0, 0x43, 0x64, 0x89, 0xf0, 0x4b, 0x64, 0x85, 0xf0, 0x47, 0xcd, 0x55, 0x20, 0xa4, 0xd4, 0x5c,
	0x05, 0x02, 0x10, 0x7d, 0x27, 0x10, 0xb2, 0x55, 0x98, 0x07, 0xd7, 0xb8, 0x5f, 0x70, 0xf9, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x33, 0x0a, 0x14, 0x0c, 0xf7, 0x01, 0x00, 0x00,
}
