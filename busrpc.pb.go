// Code generated by protoc-gen-go.
// source: busrpc.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	busrpc.proto

It has these top-level messages:
	YunBusRequest
	YunBusResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type YunBusRequest struct {
	Qid       string `protobuf:"bytes,1,opt,name=qid" json:"qid,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Source    string `protobuf:"bytes,4,opt,name=source" json:"source,omitempty"`
}

func (m *YunBusRequest) Reset()                    { *m = YunBusRequest{} }
func (m *YunBusRequest) String() string            { return proto.CompactTextString(m) }
func (*YunBusRequest) ProtoMessage()               {}
func (*YunBusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *YunBusRequest) GetQid() string {
	if m != nil {
		return m.Qid
	}
	return ""
}

func (m *YunBusRequest) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *YunBusRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *YunBusRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type YunBusResponse struct {
	Qid       string `protobuf:"bytes,1,opt,name=qid" json:"qid,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Err       int32  `protobuf:"varint,3,opt,name=err" json:"err,omitempty"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *YunBusResponse) Reset()                    { *m = YunBusResponse{} }
func (m *YunBusResponse) String() string            { return proto.CompactTextString(m) }
func (*YunBusResponse) ProtoMessage()               {}
func (*YunBusResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *YunBusResponse) GetQid() string {
	if m != nil {
		return m.Qid
	}
	return ""
}

func (m *YunBusResponse) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *YunBusResponse) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *YunBusResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*YunBusRequest)(nil), "pb.YunBusRequest")
	proto.RegisterType((*YunBusResponse)(nil), "pb.YunBusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for YunBusRouter service

type YunBusRouterClient interface {
	ProcRequest(ctx context.Context, opts ...grpc.CallOption) (YunBusRouter_ProcRequestClient, error)
}

type yunBusRouterClient struct {
	cc *grpc.ClientConn
}

func NewYunBusRouterClient(cc *grpc.ClientConn) YunBusRouterClient {
	return &yunBusRouterClient{cc}
}

func (c *yunBusRouterClient) ProcRequest(ctx context.Context, opts ...grpc.CallOption) (YunBusRouter_ProcRequestClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_YunBusRouter_serviceDesc.Streams[0], c.cc, "/pb.YunBusRouter/ProcRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &yunBusRouterProcRequestClient{stream}
	return x, nil
}

type YunBusRouter_ProcRequestClient interface {
	Send(*YunBusRequest) error
	Recv() (*YunBusResponse, error)
	grpc.ClientStream
}

type yunBusRouterProcRequestClient struct {
	grpc.ClientStream
}

func (x *yunBusRouterProcRequestClient) Send(m *YunBusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yunBusRouterProcRequestClient) Recv() (*YunBusResponse, error) {
	m := new(YunBusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for YunBusRouter service

type YunBusRouterServer interface {
	ProcRequest(YunBusRouter_ProcRequestServer) error
}

func RegisterYunBusRouterServer(s *grpc.Server, srv YunBusRouterServer) {
	s.RegisterService(&_YunBusRouter_serviceDesc, srv)
}

func _YunBusRouter_ProcRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YunBusRouterServer).ProcRequest(&yunBusRouterProcRequestServer{stream})
}

type YunBusRouter_ProcRequestServer interface {
	Send(*YunBusResponse) error
	Recv() (*YunBusRequest, error)
	grpc.ServerStream
}

type yunBusRouterProcRequestServer struct {
	grpc.ServerStream
}

func (x *yunBusRouterProcRequestServer) Send(m *YunBusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yunBusRouterProcRequestServer) Recv() (*YunBusRequest, error) {
	m := new(YunBusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _YunBusRouter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.YunBusRouter",
	HandlerType: (*YunBusRouterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcRequest",
			Handler:       _YunBusRouter_ProcRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "busrpc.proto",
}

func init() { proto.RegisterFile("busrpc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2a, 0x2d, 0x2e,
	0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xca, 0xe6, 0xe2,
	0x8d, 0x2c, 0xcd, 0x73, 0x2a, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0,
	0x62, 0x2e, 0xcc, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x64, 0xb8,
	0x38, 0x4b, 0x32, 0x73, 0x81, 0x72, 0x89, 0xb9, 0x05, 0x12, 0x4c, 0x40, 0x71, 0x96, 0x20, 0x84,
	0x80, 0x90, 0x10, 0x17, 0x4b, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0x33, 0x50, 0x82, 0x27, 0x08, 0xcc,
	0x16, 0x12, 0xe3, 0x62, 0x2b, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x95, 0x60, 0x01, 0x1b, 0x03, 0xe5,
	0x29, 0xa5, 0x71, 0xf1, 0xc1, 0x2c, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x25, 0xd9, 0x36, 0xa0,
	0xfa, 0xd4, 0xa2, 0x22, 0xb0, 0x65, 0xac, 0x41, 0x20, 0x26, 0xdc, 0x7e, 0x16, 0x84, 0xfd, 0x46,
	0x5e, 0x5c, 0x3c, 0x50, 0x7b, 0xf2, 0x4b, 0x4b, 0x52, 0x8b, 0x84, 0xac, 0xb8, 0xb8, 0x03, 0x8a,
	0xf2, 0x93, 0x61, 0x5e, 0x14, 0xd4, 0x2b, 0x48, 0xd2, 0x43, 0xf1, 0xb5, 0x94, 0x10, 0xb2, 0x10,
	0xc4, 0x6d, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x49, 0x6c, 0xe0, 0xb0, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xca, 0x7d, 0xdf, 0xd9, 0x3b, 0x01, 0x00, 0x00,
}