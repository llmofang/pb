// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ADStaringResult.proto

package com_htsc_mdc_model

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

//盯盘结果
type ADStaringResult struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,5,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,6,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	ExchangeDate         int32             `protobuf:"varint,7,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32             `protobuf:"varint,8,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	StaringResultID      string            `protobuf:"bytes,9,opt,name=StaringResultID,proto3" json:"StaringResultID,omitempty"`
	StrategyID           string            `protobuf:"bytes,10,opt,name=StrategyID,proto3" json:"StrategyID,omitempty"`
	AlgorithmID          string            `protobuf:"bytes,11,opt,name=AlgorithmID,proto3" json:"AlgorithmID,omitempty"`
	AlgorithmName        string            `protobuf:"bytes,12,opt,name=AlgorithmName,proto3" json:"AlgorithmName,omitempty"`
	CustomerID           string            `protobuf:"bytes,13,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	SystemID             string            `protobuf:"bytes,14,opt,name=SystemID,proto3" json:"SystemID,omitempty"`
	RmindValue           int64             `protobuf:"varint,15,opt,name=RmindValue,proto3" json:"RmindValue,omitempty"`
	RealCalValue         int64             `protobuf:"varint,16,opt,name=RealCalValue,proto3" json:"RealCalValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ADStaringResult) Reset()         { *m = ADStaringResult{} }
func (m *ADStaringResult) String() string { return proto.CompactTextString(m) }
func (*ADStaringResult) ProtoMessage()    {}
func (*ADStaringResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c2f448e8365b30f, []int{0}
}

func (m *ADStaringResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADStaringResult.Unmarshal(m, b)
}
func (m *ADStaringResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADStaringResult.Marshal(b, m, deterministic)
}
func (m *ADStaringResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADStaringResult.Merge(m, src)
}
func (m *ADStaringResult) XXX_Size() int {
	return xxx_messageInfo_ADStaringResult.Size(m)
}
func (m *ADStaringResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ADStaringResult.DiscardUnknown(m)
}

var xxx_messageInfo_ADStaringResult proto.InternalMessageInfo

func (m *ADStaringResult) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *ADStaringResult) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *ADStaringResult) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *ADStaringResult) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *ADStaringResult) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *ADStaringResult) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *ADStaringResult) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *ADStaringResult) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func (m *ADStaringResult) GetStaringResultID() string {
	if m != nil {
		return m.StaringResultID
	}
	return ""
}

func (m *ADStaringResult) GetStrategyID() string {
	if m != nil {
		return m.StrategyID
	}
	return ""
}

func (m *ADStaringResult) GetAlgorithmID() string {
	if m != nil {
		return m.AlgorithmID
	}
	return ""
}

func (m *ADStaringResult) GetAlgorithmName() string {
	if m != nil {
		return m.AlgorithmName
	}
	return ""
}

func (m *ADStaringResult) GetCustomerID() string {
	if m != nil {
		return m.CustomerID
	}
	return ""
}

func (m *ADStaringResult) GetSystemID() string {
	if m != nil {
		return m.SystemID
	}
	return ""
}

func (m *ADStaringResult) GetRmindValue() int64 {
	if m != nil {
		return m.RmindValue
	}
	return 0
}

func (m *ADStaringResult) GetRealCalValue() int64 {
	if m != nil {
		return m.RealCalValue
	}
	return 0
}

func init() {
	proto.RegisterType((*ADStaringResult)(nil), "com.htsc.mdc.insight.model.ADStaringResult")
}

func init() { proto.RegisterFile("ADStaringResult.proto", fileDescriptor_3c2f448e8365b30f) }

var fileDescriptor_3c2f448e8365b30f = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x6f, 0xda, 0x30,
	0x18, 0xc6, 0xe5, 0x31, 0x18, 0x98, 0xbf, 0xf2, 0xb4, 0xcd, 0xca, 0x61, 0xca, 0xd0, 0x36, 0xe5,
	0x94, 0x43, 0x7b, 0xe9, 0x15, 0x30, 0x12, 0x39, 0xb4, 0x6a, 0x1d, 0xd4, 0xbb, 0x1b, 0xac, 0x24,
	0x52, 0x9c, 0xa0, 0xd8, 0x91, 0xca, 0xb9, 0x5f, 0xa4, 0x1f, 0xb5, 0xb2, 0x43, 0x21, 0x0e, 0x12,
	0x37, 0xde, 0x9f, 0x9f, 0xd7, 0x2f, 0xcf, 0xe3, 0x37, 0xf0, 0xc7, 0x82, 0x84, 0x8a, 0x95, 0x69,
	0x1e, 0x53, 0x2e, 0xab, 0x4c, 0xf9, 0xfb, 0xb2, 0x50, 0x05, 0x72, 0xa2, 0x42, 0xf8, 0x89, 0x92,
	0x91, 0x2f, 0x76, 0x91, 0x9f, 0xe6, 0x32, 0x8d, 0x13, 0xe5, 0x8b, 0x62, 0xc7, 0x33, 0xe7, 0xfb,
	0x3a, 0xe4, 0x51, 0x55, 0xa6, 0xea, 0xb0, 0x3d, 0xec, 0x79, 0xdd, 0xe0, 0xfc, 0x3a, 0xc1, 0x80,
	0x84, 0x45, 0x55, 0x46, 0xc7, 0x83, 0xf9, 0x5b, 0x17, 0x4e, 0x5b, 0x33, 0xd0, 0x7f, 0x38, 0xd9,
	0x6c, 0xc3, 0xd5, 0xb9, 0x03, 0x03, 0x17, 0x78, 0x03, 0xda, 0xa2, 0xe8, 0x27, 0xec, 0xdd, 0x13,
	0xc2, 0x14, 0xc7, 0x5f, 0x5c, 0xe0, 0x75, 0xe9, 0xb1, 0xaa, 0xf9, 0x36, 0x15, 0x1c, 0x77, 0x3e,
	0xb9, 0xae, 0xd0, 0x5f, 0x38, 0x26, 0x4c, 0x31, 0xfd, 0x5b, 0x2a, 0x26, 0xf6, 0xf8, 0xab, 0x0b,
	0xbc, 0x0e, 0xb5, 0x21, 0x7a, 0x82, 0x33, 0xd9, 0xfa, 0xaf, 0xb8, 0xeb, 0x02, 0x6f, 0x72, 0xf3,
	0xcf, 0xb7, 0x6c, 0x1b, 0xbb, 0xfe, 0x85, 0x31, 0x7a, 0xd1, 0x8e, 0xd6, 0x70, 0x24, 0x1b, 0x99,
	0xe0, 0x9e, 0xb9, 0xee, 0xcf, 0xd5, 0xeb, 0xb4, 0x90, 0x5a, 0x6d, 0x68, 0x0e, 0x47, 0xeb, 0xd7,
	0x28, 0x61, 0x79, 0xcc, 0x8d, 0xeb, 0x6f, 0xc6, 0x9d, 0xc5, 0x9a, 0x1a, 0x93, 0x40, 0xdf, 0xd6,
	0x98, 0x1c, 0x3c, 0x38, 0xb5, 0x02, 0x0f, 0x08, 0x1e, 0x98, 0x80, 0xdb, 0x18, 0xfd, 0x86, 0x30,
	0x54, 0x25, 0x53, 0x3c, 0xd6, 0xaf, 0x00, 0x8d, 0xa8, 0x41, 0x90, 0x0b, 0x87, 0x8b, 0x2c, 0x2e,
	0xca, 0x54, 0x25, 0x22, 0x20, 0x78, 0x68, 0x04, 0x4d, 0xa4, 0x33, 0x3f, 0x95, 0x0f, 0x4c, 0x70,
	0x3c, 0x32, 0x1a, 0x1b, 0xea, 0x39, 0xab, 0x4a, 0xaa, 0x42, 0xf0, 0x32, 0x20, 0x78, 0x5c, 0xcf,
	0x39, 0x13, 0xe4, 0xc0, 0x7e, 0x78, 0x90, 0x8a, 0xeb, 0x21, 0x13, 0x73, 0x7a, 0xaa, 0x75, 0x2f,
	0x15, 0x69, 0xbe, 0x7b, 0x66, 0x59, 0xc5, 0xf1, 0xd4, 0x3c, 0x69, 0x83, 0xe8, 0x44, 0x28, 0x67,
	0xd9, 0x8a, 0x65, 0xb5, 0x62, 0x66, 0x14, 0x16, 0x5b, 0xde, 0xc1, 0x2b, 0x1b, 0xbd, 0x6c, 0x7f,
	0x04, 0x8f, 0x7a, 0x73, 0xe5, 0x06, 0xbc, 0x03, 0xf0, 0xd2, 0x33, 0x6b, 0x7c, 0xfb, 0x11, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0xfb, 0xc8, 0x1f, 0x29, 0x03, 0x00, 0x00,
}