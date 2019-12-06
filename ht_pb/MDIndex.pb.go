// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDIndex.proto

package com_htsc_mdc_insight_model

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

//MDIndexRecord指数，存储和用户模型
type MDIndex struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	TradingPhaseCode     string            `protobuf:"bytes,5,opt,name=TradingPhaseCode,proto3" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,6,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,7,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	MaxPx                int64             `protobuf:"varint,8,opt,name=MaxPx,proto3" json:"MaxPx,omitempty"`
	MinPx                int64             `protobuf:"varint,9,opt,name=MinPx,proto3" json:"MinPx,omitempty"`
	PreClosePx           int64             `protobuf:"varint,10,opt,name=PreClosePx,proto3" json:"PreClosePx,omitempty"`
	NumTrades            int64             `protobuf:"varint,11,opt,name=NumTrades,proto3" json:"NumTrades,omitempty"`
	TotalVolumeTrade     int64             `protobuf:"varint,12,opt,name=TotalVolumeTrade,proto3" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade      int64             `protobuf:"varint,13,opt,name=TotalValueTrade,proto3" json:"TotalValueTrade,omitempty"`
	LastPx               int64             `protobuf:"varint,14,opt,name=LastPx,proto3" json:"LastPx,omitempty"`
	OpenPx               int64             `protobuf:"varint,15,opt,name=OpenPx,proto3" json:"OpenPx,omitempty"`
	ClosePx              int64             `protobuf:"varint,16,opt,name=ClosePx,proto3" json:"ClosePx,omitempty"`
	HighPx               int64             `protobuf:"varint,17,opt,name=HighPx,proto3" json:"HighPx,omitempty"`
	LowPx                int64             `protobuf:"varint,18,opt,name=LowPx,proto3" json:"LowPx,omitempty"`
	ChannelNo            int32             `protobuf:"varint,19,opt,name=ChannelNo,proto3" json:"ChannelNo,omitempty"`
	ExchangeDate         int32             `protobuf:"varint,20,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32             `protobuf:"varint,21,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MDIndex) Reset()         { *m = MDIndex{} }
func (m *MDIndex) String() string { return proto.CompactTextString(m) }
func (*MDIndex) ProtoMessage()    {}
func (*MDIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7201a10acce0b70, []int{0}
}

func (m *MDIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDIndex.Unmarshal(m, b)
}
func (m *MDIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDIndex.Marshal(b, m, deterministic)
}
func (m *MDIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDIndex.Merge(m, src)
}
func (m *MDIndex) XXX_Size() int {
	return xxx_messageInfo_MDIndex.Size(m)
}
func (m *MDIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_MDIndex.DiscardUnknown(m)
}

var xxx_messageInfo_MDIndex proto.InternalMessageInfo

func (m *MDIndex) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDIndex) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDIndex) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDIndex) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDIndex) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDIndex) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDIndex) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDIndex) GetMaxPx() int64 {
	if m != nil {
		return m.MaxPx
	}
	return 0
}

func (m *MDIndex) GetMinPx() int64 {
	if m != nil {
		return m.MinPx
	}
	return 0
}

func (m *MDIndex) GetPreClosePx() int64 {
	if m != nil {
		return m.PreClosePx
	}
	return 0
}

func (m *MDIndex) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func (m *MDIndex) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *MDIndex) GetTotalValueTrade() int64 {
	if m != nil {
		return m.TotalValueTrade
	}
	return 0
}

func (m *MDIndex) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *MDIndex) GetOpenPx() int64 {
	if m != nil {
		return m.OpenPx
	}
	return 0
}

func (m *MDIndex) GetClosePx() int64 {
	if m != nil {
		return m.ClosePx
	}
	return 0
}

func (m *MDIndex) GetHighPx() int64 {
	if m != nil {
		return m.HighPx
	}
	return 0
}

func (m *MDIndex) GetLowPx() int64 {
	if m != nil {
		return m.LowPx
	}
	return 0
}

func (m *MDIndex) GetChannelNo() int32 {
	if m != nil {
		return m.ChannelNo
	}
	return 0
}

func (m *MDIndex) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *MDIndex) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MDIndex)(nil), "com.htsc.mdc.insight.model.MDIndex")
}

func init() { proto.RegisterFile("MDIndex.proto", fileDescriptor_d7201a10acce0b70) }

var fileDescriptor_d7201a10acce0b70 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xe1, 0x8e, 0xd2, 0x40,
	0x10, 0xc7, 0x53, 0x4f, 0x40, 0x56, 0xe0, 0x70, 0xef, 0xd4, 0x09, 0x31, 0x06, 0x2f, 0x6a, 0x1a,
	0x3f, 0x34, 0x46, 0xdf, 0xe0, 0x28, 0x09, 0x97, 0x1c, 0xe7, 0x5a, 0x88, 0xdf, 0xd7, 0x76, 0x43,
	0x9b, 0xb4, 0x5d, 0xd2, 0xdd, 0xc6, 0xbd, 0xb7, 0xf1, 0xfd, 0x7c, 0x09, 0xb3, 0xb3, 0xe5, 0xa0,
	0x90, 0xf0, 0x8d, 0xff, 0x6f, 0xfe, 0xb3, 0x64, 0xfe, 0x9d, 0x21, 0xc3, 0x65, 0x78, 0x57, 0x26,
	0xc2, 0x04, 0xdb, 0x4a, 0x6a, 0x49, 0x27, 0xb1, 0x2c, 0x82, 0x54, 0xab, 0x38, 0x28, 0x92, 0x38,
	0xc8, 0x4a, 0x95, 0x6d, 0x52, 0x1d, 0x14, 0x32, 0x11, 0xf9, 0xe4, 0xed, 0x7c, 0x25, 0xe2, 0xba,
	0xca, 0xf4, 0xe3, 0x5d, 0xb8, 0x92, 0x75, 0x15, 0x0b, 0xd7, 0x34, 0xb9, 0x7a, 0x2a, 0xac, 0x1f,
	0xb7, 0x0d, 0xbc, 0xf9, 0xd7, 0x21, 0xbd, 0xe6, 0x6d, 0xfa, 0x99, 0x8c, 0x16, 0xeb, 0xd5, 0x6c,
	0xdf, 0x0e, 0xde, 0xd4, 0xf3, 0xfb, 0xd1, 0x11, 0xa5, 0x6f, 0x48, 0x77, 0x19, 0x86, 0x5c, 0x0b,
	0x78, 0x36, 0xf5, 0xfc, 0x4e, 0xd4, 0x28, 0xc7, 0xd7, 0x59, 0x21, 0xe0, 0x62, 0xc7, 0xad, 0xa2,
	0x1f, 0xc9, 0x30, 0xe4, 0x9a, 0xdb, 0xdf, 0x4a, 0xf3, 0x62, 0x0b, 0xcf, 0xa7, 0x9e, 0x7f, 0x11,
	0xb5, 0x21, 0xfd, 0x42, 0xc6, 0xeb, 0x8a, 0x27, 0x59, 0xb9, 0x61, 0x29, 0x57, 0x62, 0x26, 0x13,
	0x01, 0x1d, 0xfc, 0xff, 0x13, 0x4e, 0x7f, 0x92, 0xb1, 0x3a, 0x1a, 0x12, 0xba, 0x53, 0xcf, 0x1f,
	0x7d, 0xfb, 0x14, 0xb4, 0xa2, 0xc1, 0x48, 0x82, 0x93, 0x44, 0xa2, 0x93, 0x76, 0x3a, 0x27, 0x03,
	0x75, 0x10, 0x0f, 0xf4, 0xf0, 0xb9, 0x0f, 0x67, 0x9f, 0xb3, 0xc6, 0xa8, 0xd5, 0x46, 0xaf, 0x49,
	0x67, 0xc9, 0x0d, 0x33, 0xf0, 0x02, 0x67, 0x74, 0x02, 0x69, 0x56, 0x32, 0x03, 0xfd, 0x86, 0x5a,
	0x41, 0xdf, 0x13, 0xc2, 0x2a, 0x31, 0xcb, 0xa5, 0x12, 0xcc, 0x00, 0xc1, 0xd2, 0x01, 0xa1, 0xef,
	0x48, 0xff, 0xa1, 0x2e, 0xec, 0xf0, 0x42, 0xc1, 0x4b, 0x2c, 0xef, 0x01, 0xe6, 0x25, 0x35, 0xcf,
	0x7f, 0xc9, 0xbc, 0x2e, 0x04, 0x42, 0x18, 0xa0, 0xe9, 0x84, 0x53, 0x9f, 0x5c, 0x3a, 0xc6, 0xf3,
	0xba, 0xb1, 0x0e, 0xd1, 0x7a, 0x8c, 0xed, 0x37, 0xbc, 0xe7, 0x4a, 0x33, 0x03, 0x23, 0x34, 0x34,
	0xca, 0xf2, 0x1f, 0x5b, 0x61, 0x47, 0xb8, 0x74, 0xdc, 0x29, 0x0a, 0xa4, 0xb7, 0x1b, 0x60, 0x8c,
	0x85, 0x9d, 0xb4, 0x1d, 0x8b, 0x6c, 0x93, 0x32, 0x03, 0xaf, 0x5c, 0x87, 0x53, 0x36, 0x8b, 0x7b,
	0xf9, 0x87, 0x19, 0xa0, 0x2e, 0x0b, 0x14, 0x76, 0xd6, 0x59, 0xca, 0xcb, 0x52, 0xe4, 0x0f, 0x12,
	0xae, 0x70, 0x7d, 0xf6, 0x80, 0xde, 0x90, 0xc1, 0xdc, 0xc4, 0x29, 0x2f, 0x37, 0x02, 0xf7, 0xee,
	0x1a, 0x0d, 0x2d, 0x76, 0xe8, 0xc1, 0x1d, 0x7c, 0xdd, 0xf6, 0x58, 0x76, 0xfb, 0x95, 0x9c, 0xb9,
	0x9c, 0xdb, 0xdd, 0x91, 0x31, 0x7b, 0x19, 0x6a, 0xe1, 0xfd, 0xf5, 0xbc, 0xdf, 0x5d, 0x3c, 0x93,
	0xef, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x26, 0x3f, 0x2b, 0x81, 0x03, 0x00, 0x00,
}
