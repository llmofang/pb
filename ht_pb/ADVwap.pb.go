// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ADVwap.proto

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

//VWAP指标
type ADVwap struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,5,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,6,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	PeriodType           EMDPeriodType     `protobuf:"varint,7,opt,name=PeriodType,proto3,enum=com.htsc.mdc.model.EMDPeriodType" json:"PeriodType,omitempty"`
	Volume               int64             `protobuf:"varint,8,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Vwap                 int64             `protobuf:"varint,9,opt,name=Vwap,proto3" json:"Vwap,omitempty"`
	ExchangeDate         int32             `protobuf:"varint,10,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32             `protobuf:"varint,11,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ADVwap) Reset()         { *m = ADVwap{} }
func (m *ADVwap) String() string { return proto.CompactTextString(m) }
func (*ADVwap) ProtoMessage()    {}
func (*ADVwap) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7a50b5160402c60, []int{0}
}

func (m *ADVwap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADVwap.Unmarshal(m, b)
}
func (m *ADVwap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADVwap.Marshal(b, m, deterministic)
}
func (m *ADVwap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADVwap.Merge(m, src)
}
func (m *ADVwap) XXX_Size() int {
	return xxx_messageInfo_ADVwap.Size(m)
}
func (m *ADVwap) XXX_DiscardUnknown() {
	xxx_messageInfo_ADVwap.DiscardUnknown(m)
}

var xxx_messageInfo_ADVwap proto.InternalMessageInfo

func (m *ADVwap) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *ADVwap) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *ADVwap) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *ADVwap) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *ADVwap) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *ADVwap) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *ADVwap) GetPeriodType() EMDPeriodType {
	if m != nil {
		return m.PeriodType
	}
	return EMDPeriodType_DefaultPeriod
}

func (m *ADVwap) GetVolume() int64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *ADVwap) GetVwap() int64 {
	if m != nil {
		return m.Vwap
	}
	return 0
}

func (m *ADVwap) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *ADVwap) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ADVwap)(nil), "com.htsc.mdc.insight.model.ADVwap")
}

func init() { proto.RegisterFile("ADVwap.proto", fileDescriptor_b7a50b5160402c60) }

var fileDescriptor_b7a50b5160402c60 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x4f, 0xc2, 0x30,
	0x10, 0x80, 0x53, 0x81, 0x29, 0x27, 0x12, 0x53, 0x12, 0x6d, 0x78, 0x9a, 0x44, 0xcd, 0x9e, 0x6a,
	0xa2, 0xbf, 0x00, 0xdc, 0x12, 0x7c, 0x20, 0xc1, 0x42, 0x78, 0xaf, 0xa5, 0x81, 0x25, 0x94, 0x2e,
	0x6b, 0x89, 0xf2, 0x6f, 0xfc, 0x57, 0xfe, 0x1d, 0xd3, 0x32, 0x65, 0x83, 0x84, 0xb7, 0xbb, 0x6f,
	0x77, 0xb7, 0xdd, 0xb7, 0x83, 0x56, 0x3f, 0x9e, 0x7d, 0xf2, 0x8c, 0x66, 0xb9, 0xb6, 0x1a, 0x77,
	0x85, 0x56, 0x74, 0x69, 0x8d, 0xa0, 0x6a, 0x2e, 0x68, 0xba, 0x36, 0xe9, 0x62, 0x69, 0xa9, 0xd2,
	0x73, 0xb9, 0xea, 0x76, 0x92, 0x51, 0x3c, 0x96, 0x79, 0xaa, 0xe7, 0xd3, 0x6d, 0x26, 0x77, 0x0d,
	0xdd, 0x4e, 0x32, 0x91, 0x62, 0x93, 0xa7, 0x76, 0x5b, 0x82, 0xb7, 0xff, 0xf0, 0x2d, 0x9e, 0xe8,
	0x4d, 0x2e, 0x8a, 0x07, 0xbd, 0x9f, 0x1a, 0x04, 0xbb, 0xf7, 0xe1, 0x47, 0x68, 0x0f, 0xa7, 0x93,
	0xd7, 0x7d, 0x21, 0x41, 0x21, 0x8a, 0x9a, 0xec, 0x80, 0xe2, 0x1b, 0x08, 0x46, 0x71, 0xcc, 0xad,
	0x24, 0x67, 0x21, 0x8a, 0x1a, 0xac, 0xc8, 0x76, 0x7c, 0x9a, 0x2a, 0x49, 0x6a, 0x7f, 0xdc, 0x65,
	0xf8, 0x1e, 0xae, 0x62, 0x6e, 0xb9, 0x8b, 0x8d, 0xe5, 0x2a, 0x23, 0xf5, 0x10, 0x45, 0x35, 0x56,
	0x85, 0xf8, 0x1d, 0xae, 0xcd, 0xc1, 0x27, 0x92, 0x46, 0x88, 0xa2, 0xf6, 0xf3, 0x03, 0xad, 0x28,
	0xf0, 0xab, 0xd3, 0xa3, 0x7d, 0xd8, 0x51, 0x3b, 0x4e, 0xa0, 0x65, 0x4a, 0x2a, 0x48, 0xe0, 0xc7,
	0xdd, 0x9d, 0x1c, 0xe7, 0x0a, 0x59, 0xa5, 0x0d, 0xf7, 0x01, 0xf6, 0x92, 0xc9, 0xf9, 0x89, 0x21,
	0xe5, 0xbf, 0xc1, 0x4a, 0x4d, 0x4e, 0xcd, 0x4c, 0xaf, 0x36, 0x4a, 0x92, 0x0b, 0xbf, 0x7b, 0x91,
	0x61, 0x0c, 0x75, 0xa7, 0x9e, 0x34, 0x3d, 0xf5, 0x31, 0xee, 0x41, 0x2b, 0xf9, 0x12, 0x4b, 0xbe,
	0x5e, 0x48, 0x2f, 0x19, 0xbc, 0xcc, 0x0a, 0x2b, 0xd7, 0x78, 0xe1, 0x97, 0xd5, 0x1a, 0xc7, 0x06,
	0x4f, 0x70, 0xe2, 0x74, 0x06, 0xc5, 0x91, 0x8d, 0xdd, 0x11, 0x98, 0x21, 0xfa, 0x46, 0xe8, 0x23,
	0xf0, 0x17, 0xf1, 0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x24, 0x5f, 0xb7, 0x38, 0x80, 0x02, 0x00,
	0x00,
}
