// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDTransaction.proto

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

//逐笔成交数据，存储和用户模型
type MDTransaction struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,5,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,6,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	TradeIndex           int64             `protobuf:"varint,7,opt,name=TradeIndex,proto3" json:"TradeIndex,omitempty"`
	TradeBuyNo           int64             `protobuf:"varint,8,opt,name=TradeBuyNo,proto3" json:"TradeBuyNo,omitempty"`
	TradeSellNo          int64             `protobuf:"varint,9,opt,name=TradeSellNo,proto3" json:"TradeSellNo,omitempty"`
	TradeType            int32             `protobuf:"varint,10,opt,name=TradeType,proto3" json:"TradeType,omitempty"`
	TradeBSFlag          int32             `protobuf:"varint,11,opt,name=TradeBSFlag,proto3" json:"TradeBSFlag,omitempty"`
	TradePrice           int64             `protobuf:"varint,12,opt,name=TradePrice,proto3" json:"TradePrice,omitempty"`
	TradeQty             int64             `protobuf:"varint,13,opt,name=TradeQty,proto3" json:"TradeQty,omitempty"`
	TradeMoney           int64             `protobuf:"varint,14,opt,name=TradeMoney,proto3" json:"TradeMoney,omitempty"`
	ChannelNo            int32             `protobuf:"varint,16,opt,name=ChannelNo,proto3" json:"ChannelNo,omitempty"`
	ExchangeDate         int32             `protobuf:"varint,17,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32             `protobuf:"varint,18,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	TradeCleanPrice      int64             `protobuf:"varint,19,opt,name=TradeCleanPrice,proto3" json:"TradeCleanPrice,omitempty"`
	AccruedInterestAmt   int64             `protobuf:"varint,20,opt,name=AccruedInterestAmt,proto3" json:"AccruedInterestAmt,omitempty"`
	TradeDirtyPrice      int64             `protobuf:"varint,21,opt,name=TradeDirtyPrice,proto3" json:"TradeDirtyPrice,omitempty"`
	MaturityYield        int64             `protobuf:"varint,22,opt,name=MaturityYield,proto3" json:"MaturityYield,omitempty"`
	FITradingMethod      string            `protobuf:"bytes,23,opt,name=FITradingMethod,proto3" json:"FITradingMethod,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MDTransaction) Reset()         { *m = MDTransaction{} }
func (m *MDTransaction) String() string { return proto.CompactTextString(m) }
func (*MDTransaction) ProtoMessage()    {}
func (*MDTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_18576041e048e6dd, []int{0}
}

func (m *MDTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDTransaction.Unmarshal(m, b)
}
func (m *MDTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDTransaction.Marshal(b, m, deterministic)
}
func (m *MDTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDTransaction.Merge(m, src)
}
func (m *MDTransaction) XXX_Size() int {
	return xxx_messageInfo_MDTransaction.Size(m)
}
func (m *MDTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MDTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MDTransaction proto.InternalMessageInfo

func (m *MDTransaction) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDTransaction) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDTransaction) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDTransaction) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDTransaction) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDTransaction) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDTransaction) GetTradeIndex() int64 {
	if m != nil {
		return m.TradeIndex
	}
	return 0
}

func (m *MDTransaction) GetTradeBuyNo() int64 {
	if m != nil {
		return m.TradeBuyNo
	}
	return 0
}

func (m *MDTransaction) GetTradeSellNo() int64 {
	if m != nil {
		return m.TradeSellNo
	}
	return 0
}

func (m *MDTransaction) GetTradeType() int32 {
	if m != nil {
		return m.TradeType
	}
	return 0
}

func (m *MDTransaction) GetTradeBSFlag() int32 {
	if m != nil {
		return m.TradeBSFlag
	}
	return 0
}

func (m *MDTransaction) GetTradePrice() int64 {
	if m != nil {
		return m.TradePrice
	}
	return 0
}

func (m *MDTransaction) GetTradeQty() int64 {
	if m != nil {
		return m.TradeQty
	}
	return 0
}

func (m *MDTransaction) GetTradeMoney() int64 {
	if m != nil {
		return m.TradeMoney
	}
	return 0
}

func (m *MDTransaction) GetChannelNo() int32 {
	if m != nil {
		return m.ChannelNo
	}
	return 0
}

func (m *MDTransaction) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *MDTransaction) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func (m *MDTransaction) GetTradeCleanPrice() int64 {
	if m != nil {
		return m.TradeCleanPrice
	}
	return 0
}

func (m *MDTransaction) GetAccruedInterestAmt() int64 {
	if m != nil {
		return m.AccruedInterestAmt
	}
	return 0
}

func (m *MDTransaction) GetTradeDirtyPrice() int64 {
	if m != nil {
		return m.TradeDirtyPrice
	}
	return 0
}

func (m *MDTransaction) GetMaturityYield() int64 {
	if m != nil {
		return m.MaturityYield
	}
	return 0
}

func (m *MDTransaction) GetFITradingMethod() string {
	if m != nil {
		return m.FITradingMethod
	}
	return ""
}

func init() {
	proto.RegisterType((*MDTransaction)(nil), "com.htsc.mdc.insight.model.MDTransaction")
}

func init() { proto.RegisterFile("MDTransaction.proto", fileDescriptor_18576041e048e6dd) }

var fileDescriptor_18576041e048e6dd = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x4a, 0x43, 0xb3, 0x4d, 0x42, 0xd9, 0x40, 0xbb, 0x8a, 0x10, 0x0a, 0x15, 0x20,
	0x9f, 0x7c, 0x00, 0x89, 0x7b, 0x13, 0xa7, 0x6a, 0x0e, 0xae, 0x5a, 0xdb, 0x17, 0x8e, 0xcb, 0x7a,
	0x14, 0x5b, 0xb2, 0x77, 0xa3, 0xf5, 0x46, 0xaa, 0xdf, 0x86, 0x37, 0xe3, 0x55, 0xd0, 0x8e, 0x8b,
	0xb3, 0x4e, 0xab, 0xde, 0x32, 0xdf, 0xcc, 0xfc, 0xfb, 0x4f, 0x3c, 0x43, 0xa6, 0x51, 0x98, 0x6a,
	0x2e, 0x6b, 0x2e, 0x4c, 0xa1, 0x64, 0xb0, 0xd5, 0xca, 0x28, 0x3a, 0x13, 0xaa, 0x0a, 0x72, 0x53,
	0x8b, 0xa0, 0xca, 0x44, 0x50, 0xc8, 0xba, 0xd8, 0xe4, 0x26, 0xa8, 0x54, 0x06, 0xe5, 0xec, 0x62,
	0x95, 0x80, 0xd8, 0xe9, 0xc2, 0x34, 0xeb, 0x30, 0x51, 0x3b, 0x2d, 0xa0, 0x6d, 0x9a, 0x4d, 0xbb,
	0x44, 0xda, 0x6c, 0x1f, 0xe1, 0xe5, 0xdf, 0x01, 0x19, 0xf7, 0x5e, 0xa0, 0xdf, 0xc8, 0xe4, 0x26,
	0x4d, 0x96, 0x7b, 0x11, 0xe6, 0xcd, 0x3d, 0x7f, 0x18, 0x1f, 0x50, 0x7a, 0x4e, 0x06, 0x51, 0x18,
	0x72, 0x03, 0xec, 0xd5, 0xdc, 0xf3, 0x8f, 0xe3, 0xc7, 0xa8, 0xe5, 0x69, 0x51, 0x01, 0x3b, 0xfa,
	0xcf, 0x6d, 0x44, 0xbf, 0x90, 0x71, 0xc8, 0x0d, 0xb7, 0xbf, 0x6b, 0xc3, 0xab, 0x2d, 0x7b, 0x3d,
	0xf7, 0xfc, 0xa3, 0xb8, 0x0f, 0xe9, 0x3d, 0x39, 0xab, 0x0f, 0xec, 0xb3, 0xe3, 0xb9, 0xe7, 0x4f,
	0xbe, 0x7f, 0x0d, 0x7a, 0x43, 0xe3, 0xb0, 0xc1, 0x93, 0x59, 0xe3, 0x27, 0xed, 0x74, 0x45, 0x46,
	0xb5, 0x33, 0x38, 0x1b, 0xa0, 0xdc, 0xe7, 0x17, 0xe5, 0x6c, 0x61, 0xdc, 0x6b, 0xa3, 0x9f, 0x08,
	0x49, 0x35, 0xcf, 0x60, 0x2d, 0x33, 0x78, 0x60, 0x6f, 0xd0, 0xbc, 0x43, 0xba, 0xfc, 0x62, 0xd7,
	0xdc, 0x2a, 0x76, 0xe2, 0xe4, 0x91, 0xd0, 0x39, 0x39, 0xc5, 0x28, 0x81, 0xb2, 0xbc, 0x55, 0x6c,
	0x88, 0x05, 0x2e, 0xa2, 0x1f, 0xc9, 0x10, 0x43, 0x74, 0x49, 0xf0, 0xcf, 0xdb, 0x83, 0xae, 0x7f,
	0x91, 0x5c, 0x97, 0x7c, 0xc3, 0x4e, 0x31, 0xef, 0xa2, 0xce, 0xc1, 0x9d, 0x2e, 0x04, 0xb0, 0x91,
	0xe3, 0x00, 0x09, 0x9d, 0x91, 0x13, 0x8c, 0xee, 0x4d, 0xc3, 0xc6, 0x98, 0xed, 0xe2, 0xae, 0x37,
	0x52, 0x12, 0x1a, 0x36, 0x71, 0x7a, 0x91, 0x58, 0x6f, 0xcb, 0x9c, 0x4b, 0x09, 0xd6, 0xfb, 0x59,
	0xeb, 0xad, 0x03, 0xf4, 0x92, 0x8c, 0x56, 0x0f, 0x22, 0xe7, 0x72, 0x03, 0xb8, 0x11, 0xef, 0xb0,
	0xa0, 0xc7, 0xdc, 0x1a, 0xdc, 0x0e, 0xda, 0xaf, 0xc1, 0x1d, 0xf1, 0xc9, 0x5b, 0x7c, 0x73, 0x59,
	0x02, 0x97, 0xed, 0x18, 0x53, 0xb4, 0x72, 0x88, 0x69, 0x40, 0xe8, 0x95, 0x10, 0x7a, 0x07, 0xd9,
	0x5a, 0x1a, 0xd0, 0x50, 0x9b, 0xab, 0xca, 0xb0, 0xf7, 0x58, 0xfc, 0x4c, 0xa6, 0x53, 0x0e, 0x0b,
	0x6d, 0x9a, 0x56, 0xf9, 0x83, 0xa3, 0xbc, 0xc7, 0x76, 0x4f, 0x23, 0x6e, 0xf0, 0xbb, 0xff, 0x2a,
	0xa0, 0xcc, 0xd8, 0x79, 0xbb, 0xa7, 0x3d, 0x68, 0xf5, 0xae, 0xd7, 0xb6, 0xb5, 0x90, 0x9b, 0x08,
	0x4c, 0xae, 0x32, 0x76, 0x81, 0x67, 0x72, 0x88, 0x17, 0x3f, 0xc9, 0x0b, 0xd7, 0xba, 0xe8, 0x9f,
	0xf7, 0x9d, 0xbd, 0xc9, 0xfa, 0xc6, 0xfb, 0xe3, 0x79, 0xbf, 0x07, 0x78, 0xa0, 0x3f, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x79, 0xcc, 0x17, 0x3c, 0x01, 0x04, 0x00, 0x00,
}
