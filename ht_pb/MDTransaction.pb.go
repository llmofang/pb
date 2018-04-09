// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDTransaction.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	MDTransaction.proto

It has these top-level messages:
	MDTransaction
*/
package com_htsc_mdc_insight_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_htsc_mdc_model "."
import com_htsc_mdc_model1 "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 逐笔成交数据，存储和用户模型
type MDTransaction struct {
	HTSCSecurityID   string                               `protobuf:"bytes,1,opt,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	MDDate           int32                                `protobuf:"varint,2,opt,name=MDDate" json:"MDDate,omitempty"`
	MDTime           int32                                `protobuf:"varint,3,opt,name=MDTime" json:"MDTime,omitempty"`
	DataTimestamp    int64                                `protobuf:"varint,4,opt,name=DataTimestamp" json:"DataTimestamp,omitempty"`
	SecurityIDSource com_htsc_mdc_model.ESecurityIDSource `protobuf:"varint,5,opt,name=securityIDSource,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType     com_htsc_mdc_model1.ESecurityType    `protobuf:"varint,6,opt,name=securityType,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	TradeIndex       int64                                `protobuf:"varint,7,opt,name=TradeIndex" json:"TradeIndex,omitempty"`
	TradeBuyNo       int64                                `protobuf:"varint,8,opt,name=TradeBuyNo" json:"TradeBuyNo,omitempty"`
	TradeSellNo      int64                                `protobuf:"varint,9,opt,name=TradeSellNo" json:"TradeSellNo,omitempty"`
	TradeType        int32                                `protobuf:"varint,10,opt,name=TradeType" json:"TradeType,omitempty"`
	TradeBSFlag      int32                                `protobuf:"varint,11,opt,name=TradeBSFlag" json:"TradeBSFlag,omitempty"`
	TradePrice       int64                                `protobuf:"varint,12,opt,name=TradePrice" json:"TradePrice,omitempty"`
	TradeQty         int64                                `protobuf:"varint,13,opt,name=TradeQty" json:"TradeQty,omitempty"`
	TradeMoney       int64                                `protobuf:"varint,14,opt,name=TradeMoney" json:"TradeMoney,omitempty"`
	NumTrades        int64                                `protobuf:"varint,15,opt,name=NumTrades" json:"NumTrades,omitempty"`
}

func (m *MDTransaction) Reset()                    { *m = MDTransaction{} }
func (m *MDTransaction) String() string            { return proto.CompactTextString(m) }
func (*MDTransaction) ProtoMessage()               {}
func (*MDTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *MDTransaction) GetSecurityIDSource() com_htsc_mdc_model.ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return com_htsc_mdc_model.ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDTransaction) GetSecurityType() com_htsc_mdc_model1.ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return com_htsc_mdc_model1.ESecurityType_DefaultSecurityType
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

func (m *MDTransaction) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func init() {
	proto.RegisterType((*MDTransaction)(nil), "com.htsc.mdc.insight.model.MDTransaction")
}

func init() { proto.RegisterFile("MDTransaction.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdd, 0x6a, 0xea, 0x40,
	0x14, 0x85, 0x99, 0xe3, 0xcf, 0xd1, 0xf1, 0xe7, 0x1c, 0x46, 0x68, 0x87, 0x50, 0x4a, 0x5a, 0xda,
	0x92, 0xab, 0x5c, 0xb4, 0xd0, 0x07, 0x48, 0x63, 0xd1, 0x0b, 0x45, 0x93, 0xbc, 0xc0, 0x34, 0x19,
	0x34, 0x90, 0x64, 0x24, 0x33, 0x81, 0xe6, 0x6d, 0xfa, 0x2e, 0x7d, 0xb1, 0x92, 0xad, 0x8d, 0x13,
	0x05, 0xef, 0x5c, 0xdf, 0x9a, 0xb5, 0x5d, 0x3b, 0x6c, 0x3c, 0x59, 0xb8, 0x41, 0xce, 0x32, 0xc9,
	0x42, 0x15, 0x8b, 0xcc, 0xde, 0xe5, 0x42, 0x09, 0x62, 0x84, 0x22, 0xb5, 0xb7, 0x4a, 0x86, 0x76,
	0x1a, 0x85, 0x76, 0x9c, 0xc9, 0x78, 0xb3, 0x55, 0x76, 0x2a, 0x22, 0x9e, 0x18, 0xd7, 0x53, 0x9f,
	0x87, 0x45, 0x1e, 0xab, 0x72, 0xee, 0xfa, 0xa2, 0xc8, 0x43, 0xbe, 0x0f, 0x19, 0x93, 0xda, 0x08,
	0xca, 0xdd, 0x01, 0xde, 0x7f, 0xb7, 0xf1, 0xa8, 0xf1, 0x0f, 0xe4, 0x09, 0x8f, 0x67, 0x81, 0xff,
	0x76, 0x1c, 0x42, 0x91, 0x89, 0xac, 0xbe, 0x77, 0x42, 0xc9, 0x15, 0xee, 0x2e, 0x5c, 0x97, 0x29,
	0x4e, 0xff, 0x98, 0xc8, 0xea, 0x78, 0x07, 0xb5, 0xe7, 0x41, 0x9c, 0x72, 0xda, 0xfa, 0xe5, 0x95,
	0x22, 0x0f, 0x78, 0xe4, 0x32, 0xc5, 0xaa, 0xdf, 0x52, 0xb1, 0x74, 0x47, 0xdb, 0x26, 0xb2, 0x5a,
	0x5e, 0x13, 0x92, 0x35, 0xfe, 0x2f, 0x4f, 0xea, 0xd3, 0x8e, 0x89, 0xac, 0xf1, 0xf3, 0xa3, 0xdd,
	0x58, 0x1a, 0x96, 0xb5, 0xcf, 0x76, 0xf5, 0xce, 0xe2, 0x64, 0x8a, 0x87, 0x52, 0x5b, 0x9c, 0x76,
	0x61, 0xdc, 0xdd, 0xc5, 0x71, 0xd5, 0x43, 0xaf, 0x11, 0x23, 0xb7, 0x18, 0x07, 0x39, 0x8b, 0xf8,
	0x3c, 0x8b, 0xf8, 0x27, 0xfd, 0x0b, 0xe5, 0x35, 0x52, 0xfb, 0x4e, 0x51, 0x2e, 0x05, 0xed, 0x69,
	0x3e, 0x10, 0x62, 0xe2, 0x01, 0x28, 0x9f, 0x27, 0xc9, 0x52, 0xd0, 0x3e, 0x3c, 0xd0, 0x11, 0xb9,
	0xc1, 0x7d, 0x90, 0xd0, 0x12, 0xc3, 0xc7, 0x3b, 0x82, 0x3a, 0xef, 0xf8, 0xef, 0x09, 0xdb, 0xd0,
	0x01, 0xf8, 0x3a, 0xaa, 0x1b, 0xac, 0xf2, 0x38, 0xe4, 0x74, 0xa8, 0x35, 0x00, 0x42, 0x0c, 0xdc,
	0x03, 0xb5, 0x56, 0x25, 0x1d, 0x81, 0x5b, 0xeb, 0x3a, 0xbb, 0x10, 0x19, 0x2f, 0xe9, 0x58, 0xcb,
	0x02, 0xa9, 0xba, 0x2d, 0x8b, 0x14, 0x80, 0xa4, 0xff, 0xc0, 0x3e, 0x02, 0xe7, 0x15, 0x5f, 0xb8,
	0x48, 0xa7, 0x79, 0xc2, 0xab, 0xea, 0xee, 0xe4, 0x0c, 0x7d, 0x21, 0xf4, 0xd1, 0x85, 0x23, 0x7c,
	0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0xee, 0x13, 0x27, 0x87, 0xe5, 0x02, 0x00, 0x00,
}
