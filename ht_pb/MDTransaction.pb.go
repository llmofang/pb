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
	HTSCSecurityID   string            `protobuf:"bytes,1,opt,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	MDDate           int32             `protobuf:"varint,2,opt,name=MDDate" json:"MDDate,omitempty"`
	MDTime           int32             `protobuf:"varint,3,opt,name=MDTime" json:"MDTime,omitempty"`
	DataTimestamp    int64             `protobuf:"varint,4,opt,name=DataTimestamp" json:"DataTimestamp,omitempty"`
	SecurityIDSource ESecurityIDSource `protobuf:"varint,5,opt,name=securityIDSource,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType     ESecurityType     `protobuf:"varint,6,opt,name=securityType,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	TradeIndex       int64             `protobuf:"varint,7,opt,name=TradeIndex" json:"TradeIndex,omitempty"`
	TradeBuyNo       int64             `protobuf:"varint,8,opt,name=TradeBuyNo" json:"TradeBuyNo,omitempty"`
	TradeSellNo      int64             `protobuf:"varint,9,opt,name=TradeSellNo" json:"TradeSellNo,omitempty"`
	TradeType        int32             `protobuf:"varint,10,opt,name=TradeType" json:"TradeType,omitempty"`
	TradeBSFlag      int32             `protobuf:"varint,11,opt,name=TradeBSFlag" json:"TradeBSFlag,omitempty"`
	TradePrice       int64             `protobuf:"varint,12,opt,name=TradePrice" json:"TradePrice,omitempty"`
	TradeQty         int64             `protobuf:"varint,13,opt,name=TradeQty" json:"TradeQty,omitempty"`
	TradeMoney       int64             `protobuf:"varint,14,opt,name=TradeMoney" json:"TradeMoney,omitempty"`
	NumTrades        int64             `protobuf:"varint,15,opt,name=NumTrades" json:"NumTrades,omitempty"`
}

func (m *MDTransaction) Reset()         { *m = MDTransaction{} }
func (m *MDTransaction) String() string { return proto.CompactTextString(m) }
func (*MDTransaction) ProtoMessage()    {}

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

func (m *MDTransaction) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func init() {
	proto.RegisterType((*MDTransaction)(nil), "com.htsc.mdc.insight.model.MDTransaction")
}