// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDFuture.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	MDFuture.proto

It has these top-level messages:
	MDFuture
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

// 期货
type MDFuture struct {
	HTSCSecurityID   string                               `protobuf:"bytes,1,opt,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	MDDate           int32                                `protobuf:"varint,2,opt,name=MDDate" json:"MDDate,omitempty"`
	MDTime           int32                                `protobuf:"varint,3,opt,name=MDTime" json:"MDTime,omitempty"`
	DataTimestamp    int64                                `protobuf:"varint,4,opt,name=DataTimestamp" json:"DataTimestamp,omitempty"`
	TradingPhaseCode string                               `protobuf:"bytes,5,opt,name=TradingPhaseCode" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource com_htsc_mdc_model.ESecurityIDSource `protobuf:"varint,6,opt,name=securityIDSource,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType     com_htsc_mdc_model1.ESecurityType    `protobuf:"varint,7,opt,name=securityType,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	MaxPx            int64                                `protobuf:"varint,8,opt,name=MaxPx" json:"MaxPx,omitempty"`
	MinPx            int64                                `protobuf:"varint,9,opt,name=MinPx" json:"MinPx,omitempty"`
	PreClosePx       int64                                `protobuf:"varint,10,opt,name=PreClosePx" json:"PreClosePx,omitempty"`
	NumTrades        int64                                `protobuf:"varint,11,opt,name=NumTrades" json:"NumTrades,omitempty"`
	TotalVolumeTrade int64                                `protobuf:"varint,12,opt,name=TotalVolumeTrade" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade  int64                                `protobuf:"varint,13,opt,name=TotalValueTrade" json:"TotalValueTrade,omitempty"`
	LastPx           int64                                `protobuf:"varint,14,opt,name=LastPx" json:"LastPx,omitempty"`
	OpenPx           int64                                `protobuf:"varint,15,opt,name=OpenPx" json:"OpenPx,omitempty"`
	ClosePx          int64                                `protobuf:"varint,16,opt,name=ClosePx" json:"ClosePx,omitempty"`
	HighPx           int64                                `protobuf:"varint,17,opt,name=HighPx" json:"HighPx,omitempty"`
	LowPx            int64                                `protobuf:"varint,18,opt,name=LowPx" json:"LowPx,omitempty"`
	TradingDate      int32                                `protobuf:"varint,19,opt,name=TradingDate" json:"TradingDate,omitempty"`
	PreOpenInterest  int64                                `protobuf:"varint,20,opt,name=PreOpenInterest" json:"PreOpenInterest,omitempty"`
	PreSettlePrice   int64                                `protobuf:"varint,21,opt,name=PreSettlePrice" json:"PreSettlePrice,omitempty"`
	OpenInterest     int64                                `protobuf:"varint,22,opt,name=OpenInterest" json:"OpenInterest,omitempty"`
	SettlePrice      int64                                `protobuf:"varint,23,opt,name=SettlePrice" json:"SettlePrice,omitempty"`
	PreDelta         int64                                `protobuf:"varint,24,opt,name=PreDelta" json:"PreDelta,omitempty"`
	CurrDelta        int64                                `protobuf:"varint,25,opt,name=CurrDelta" json:"CurrDelta,omitempty"`
	BuyOrderQueue    []int64                              `protobuf:"varint,100,rep,packed,name=BuyOrderQueue" json:"BuyOrderQueue,omitempty"`
	SellOrderQueue   []int64                              `protobuf:"varint,101,rep,packed,name=SellOrderQueue" json:"SellOrderQueue,omitempty"`
	Buy1Price        int64                                `protobuf:"varint,102,opt,name=Buy1Price" json:"Buy1Price,omitempty"`
	Buy1OrderQty     int64                                `protobuf:"varint,103,opt,name=Buy1OrderQty" json:"Buy1OrderQty,omitempty"`
	Buy1NumOrders    int64                                `protobuf:"varint,104,opt,name=Buy1NumOrders" json:"Buy1NumOrders,omitempty"`
	Sell1Price       int64                                `protobuf:"varint,105,opt,name=Sell1Price" json:"Sell1Price,omitempty"`
	Sell1OrderQty    int64                                `protobuf:"varint,106,opt,name=Sell1OrderQty" json:"Sell1OrderQty,omitempty"`
	Sell1NumOrders   int64                                `protobuf:"varint,107,opt,name=Sell1NumOrders" json:"Sell1NumOrders,omitempty"`
	Buy2Price        int64                                `protobuf:"varint,108,opt,name=Buy2Price" json:"Buy2Price,omitempty"`
	Buy2OrderQty     int64                                `protobuf:"varint,109,opt,name=Buy2OrderQty" json:"Buy2OrderQty,omitempty"`
	Buy2NumOrders    int64                                `protobuf:"varint,110,opt,name=Buy2NumOrders" json:"Buy2NumOrders,omitempty"`
	Sell2Price       int64                                `protobuf:"varint,111,opt,name=Sell2Price" json:"Sell2Price,omitempty"`
	Sell2OrderQty    int64                                `protobuf:"varint,112,opt,name=Sell2OrderQty" json:"Sell2OrderQty,omitempty"`
	Sell2NumOrders   int64                                `protobuf:"varint,113,opt,name=Sell2NumOrders" json:"Sell2NumOrders,omitempty"`
	Buy3Price        int64                                `protobuf:"varint,114,opt,name=Buy3Price" json:"Buy3Price,omitempty"`
	Buy3OrderQty     int64                                `protobuf:"varint,115,opt,name=Buy3OrderQty" json:"Buy3OrderQty,omitempty"`
	Buy3NumOrders    int64                                `protobuf:"varint,116,opt,name=Buy3NumOrders" json:"Buy3NumOrders,omitempty"`
	Sell3Price       int64                                `protobuf:"varint,117,opt,name=Sell3Price" json:"Sell3Price,omitempty"`
	Sell3OrderQty    int64                                `protobuf:"varint,118,opt,name=Sell3OrderQty" json:"Sell3OrderQty,omitempty"`
	Sell3NumOrders   int64                                `protobuf:"varint,119,opt,name=Sell3NumOrders" json:"Sell3NumOrders,omitempty"`
	Buy4Price        int64                                `protobuf:"varint,120,opt,name=Buy4Price" json:"Buy4Price,omitempty"`
	Buy4OrderQty     int64                                `protobuf:"varint,121,opt,name=Buy4OrderQty" json:"Buy4OrderQty,omitempty"`
	Buy4NumOrders    int64                                `protobuf:"varint,122,opt,name=Buy4NumOrders" json:"Buy4NumOrders,omitempty"`
	Sell4Price       int64                                `protobuf:"varint,123,opt,name=Sell4Price" json:"Sell4Price,omitempty"`
	Sell4OrderQty    int64                                `protobuf:"varint,124,opt,name=Sell4OrderQty" json:"Sell4OrderQty,omitempty"`
	Sell4NumOrders   int64                                `protobuf:"varint,125,opt,name=Sell4NumOrders" json:"Sell4NumOrders,omitempty"`
	Buy5Price        int64                                `protobuf:"varint,126,opt,name=Buy5Price" json:"Buy5Price,omitempty"`
	Buy5OrderQty     int64                                `protobuf:"varint,127,opt,name=Buy5OrderQty" json:"Buy5OrderQty,omitempty"`
	Buy5NumOrders    int64                                `protobuf:"varint,128,opt,name=Buy5NumOrders" json:"Buy5NumOrders,omitempty"`
	Sell5Price       int64                                `protobuf:"varint,129,opt,name=Sell5Price" json:"Sell5Price,omitempty"`
	Sell5OrderQty    int64                                `protobuf:"varint,130,opt,name=Sell5OrderQty" json:"Sell5OrderQty,omitempty"`
	Sell5NumOrders   int64                                `protobuf:"varint,131,opt,name=Sell5NumOrders" json:"Sell5NumOrders,omitempty"`
}

func (m *MDFuture) Reset()                    { *m = MDFuture{} }
func (m *MDFuture) String() string            { return proto.CompactTextString(m) }
func (*MDFuture) ProtoMessage()               {}
func (*MDFuture) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MDFuture) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDFuture) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDFuture) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDFuture) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDFuture) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDFuture) GetSecurityIDSource() com_htsc_mdc_model.ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return com_htsc_mdc_model.ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDFuture) GetSecurityType() com_htsc_mdc_model1.ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return com_htsc_mdc_model1.ESecurityType_DefaultSecurityType
}

func (m *MDFuture) GetMaxPx() int64 {
	if m != nil {
		return m.MaxPx
	}
	return 0
}

func (m *MDFuture) GetMinPx() int64 {
	if m != nil {
		return m.MinPx
	}
	return 0
}

func (m *MDFuture) GetPreClosePx() int64 {
	if m != nil {
		return m.PreClosePx
	}
	return 0
}

func (m *MDFuture) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func (m *MDFuture) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *MDFuture) GetTotalValueTrade() int64 {
	if m != nil {
		return m.TotalValueTrade
	}
	return 0
}

func (m *MDFuture) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *MDFuture) GetOpenPx() int64 {
	if m != nil {
		return m.OpenPx
	}
	return 0
}

func (m *MDFuture) GetClosePx() int64 {
	if m != nil {
		return m.ClosePx
	}
	return 0
}

func (m *MDFuture) GetHighPx() int64 {
	if m != nil {
		return m.HighPx
	}
	return 0
}

func (m *MDFuture) GetLowPx() int64 {
	if m != nil {
		return m.LowPx
	}
	return 0
}

func (m *MDFuture) GetTradingDate() int32 {
	if m != nil {
		return m.TradingDate
	}
	return 0
}

func (m *MDFuture) GetPreOpenInterest() int64 {
	if m != nil {
		return m.PreOpenInterest
	}
	return 0
}

func (m *MDFuture) GetPreSettlePrice() int64 {
	if m != nil {
		return m.PreSettlePrice
	}
	return 0
}

func (m *MDFuture) GetOpenInterest() int64 {
	if m != nil {
		return m.OpenInterest
	}
	return 0
}

func (m *MDFuture) GetSettlePrice() int64 {
	if m != nil {
		return m.SettlePrice
	}
	return 0
}

func (m *MDFuture) GetPreDelta() int64 {
	if m != nil {
		return m.PreDelta
	}
	return 0
}

func (m *MDFuture) GetCurrDelta() int64 {
	if m != nil {
		return m.CurrDelta
	}
	return 0
}

func (m *MDFuture) GetBuyOrderQueue() []int64 {
	if m != nil {
		return m.BuyOrderQueue
	}
	return nil
}

func (m *MDFuture) GetSellOrderQueue() []int64 {
	if m != nil {
		return m.SellOrderQueue
	}
	return nil
}

func (m *MDFuture) GetBuy1Price() int64 {
	if m != nil {
		return m.Buy1Price
	}
	return 0
}

func (m *MDFuture) GetBuy1OrderQty() int64 {
	if m != nil {
		return m.Buy1OrderQty
	}
	return 0
}

func (m *MDFuture) GetBuy1NumOrders() int64 {
	if m != nil {
		return m.Buy1NumOrders
	}
	return 0
}

func (m *MDFuture) GetSell1Price() int64 {
	if m != nil {
		return m.Sell1Price
	}
	return 0
}

func (m *MDFuture) GetSell1OrderQty() int64 {
	if m != nil {
		return m.Sell1OrderQty
	}
	return 0
}

func (m *MDFuture) GetSell1NumOrders() int64 {
	if m != nil {
		return m.Sell1NumOrders
	}
	return 0
}

func (m *MDFuture) GetBuy2Price() int64 {
	if m != nil {
		return m.Buy2Price
	}
	return 0
}

func (m *MDFuture) GetBuy2OrderQty() int64 {
	if m != nil {
		return m.Buy2OrderQty
	}
	return 0
}

func (m *MDFuture) GetBuy2NumOrders() int64 {
	if m != nil {
		return m.Buy2NumOrders
	}
	return 0
}

func (m *MDFuture) GetSell2Price() int64 {
	if m != nil {
		return m.Sell2Price
	}
	return 0
}

func (m *MDFuture) GetSell2OrderQty() int64 {
	if m != nil {
		return m.Sell2OrderQty
	}
	return 0
}

func (m *MDFuture) GetSell2NumOrders() int64 {
	if m != nil {
		return m.Sell2NumOrders
	}
	return 0
}

func (m *MDFuture) GetBuy3Price() int64 {
	if m != nil {
		return m.Buy3Price
	}
	return 0
}

func (m *MDFuture) GetBuy3OrderQty() int64 {
	if m != nil {
		return m.Buy3OrderQty
	}
	return 0
}

func (m *MDFuture) GetBuy3NumOrders() int64 {
	if m != nil {
		return m.Buy3NumOrders
	}
	return 0
}

func (m *MDFuture) GetSell3Price() int64 {
	if m != nil {
		return m.Sell3Price
	}
	return 0
}

func (m *MDFuture) GetSell3OrderQty() int64 {
	if m != nil {
		return m.Sell3OrderQty
	}
	return 0
}

func (m *MDFuture) GetSell3NumOrders() int64 {
	if m != nil {
		return m.Sell3NumOrders
	}
	return 0
}

func (m *MDFuture) GetBuy4Price() int64 {
	if m != nil {
		return m.Buy4Price
	}
	return 0
}

func (m *MDFuture) GetBuy4OrderQty() int64 {
	if m != nil {
		return m.Buy4OrderQty
	}
	return 0
}

func (m *MDFuture) GetBuy4NumOrders() int64 {
	if m != nil {
		return m.Buy4NumOrders
	}
	return 0
}

func (m *MDFuture) GetSell4Price() int64 {
	if m != nil {
		return m.Sell4Price
	}
	return 0
}

func (m *MDFuture) GetSell4OrderQty() int64 {
	if m != nil {
		return m.Sell4OrderQty
	}
	return 0
}

func (m *MDFuture) GetSell4NumOrders() int64 {
	if m != nil {
		return m.Sell4NumOrders
	}
	return 0
}

func (m *MDFuture) GetBuy5Price() int64 {
	if m != nil {
		return m.Buy5Price
	}
	return 0
}

func (m *MDFuture) GetBuy5OrderQty() int64 {
	if m != nil {
		return m.Buy5OrderQty
	}
	return 0
}

func (m *MDFuture) GetBuy5NumOrders() int64 {
	if m != nil {
		return m.Buy5NumOrders
	}
	return 0
}

func (m *MDFuture) GetSell5Price() int64 {
	if m != nil {
		return m.Sell5Price
	}
	return 0
}

func (m *MDFuture) GetSell5OrderQty() int64 {
	if m != nil {
		return m.Sell5OrderQty
	}
	return 0
}

func (m *MDFuture) GetSell5NumOrders() int64 {
	if m != nil {
		return m.Sell5NumOrders
	}
	return 0
}

func init() {
	proto.RegisterType((*MDFuture)(nil), "com.htsc.mdc.insight.model.MDFuture")
}

func init() { proto.RegisterFile("MDFuture.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0x5b, 0x6f, 0xda, 0x3c,
	0x18, 0xc7, 0x95, 0xb7, 0x6f, 0x4f, 0x6e, 0x4b, 0xfb, 0xba, 0x7d, 0x5b, 0x0f, 0x4d, 0x1b, 0xab,
	0xd6, 0x0d, 0xed, 0x02, 0xa9, 0x49, 0xf8, 0x02, 0x40, 0xa7, 0x56, 0xea, 0x21, 0x05, 0xb4, 0xfb,
	0x0c, 0x3c, 0xc8, 0x96, 0x10, 0xe6, 0x38, 0x6b, 0xd8, 0xf9, 0xf0, 0x45, 0xf6, 0x39, 0x77, 0x35,
	0xf9, 0x80, 0xed, 0x38, 0x52, 0xef, 0x78, 0xfe, 0x7e, 0xec, 0x1f, 0x3f, 0xdb, 0x91, 0x0c, 0x6a,
	0x57, 0xbd, 0x97, 0x39, 0xcd, 0x09, 0x6e, 0xcd, 0x49, 0x4a, 0x53, 0x58, 0x1f, 0xa5, 0x49, 0x6b,
	0x4a, 0xb3, 0x51, 0x2b, 0x19, 0x8f, 0x5a, 0xd1, 0x2c, 0x8b, 0x26, 0x53, 0xda, 0x4a, 0xd2, 0x31,
	0x8e, 0xeb, 0x47, 0x67, 0x03, 0x3c, 0xca, 0x49, 0x44, 0x17, 0x17, 0xbd, 0x41, 0x9a, 0x93, 0x91,
	0x9c, 0x54, 0xdf, 0x57, 0x03, 0xc3, 0xc5, 0x5c, 0x86, 0xc7, 0x7f, 0x76, 0xc1, 0xc6, 0x72, 0x71,
	0xf8, 0x0c, 0xd4, 0xce, 0x87, 0x83, 0xae, 0x9e, 0x8f, 0x9c, 0x86, 0xd3, 0xdc, 0xec, 0x5b, 0x29,
	0x3c, 0x04, 0x6b, 0x57, 0xbd, 0x5e, 0x48, 0x31, 0xfa, 0xa7, 0xe1, 0x34, 0x57, 0xfb, 0xb2, 0x12,
	0xf9, 0x30, 0x4a, 0x30, 0x5a, 0x59, 0xe6, 0xac, 0x82, 0x4f, 0xc1, 0x4e, 0x2f, 0xa4, 0x21, 0xfb,
	0x9d, 0xd1, 0x30, 0x99, 0xa3, 0x7f, 0x1b, 0x4e, 0x73, 0xa5, 0x5f, 0x0e, 0xe1, 0x0b, 0xb0, 0x37,
	0x24, 0xe1, 0x38, 0x9a, 0x4d, 0x82, 0x69, 0x98, 0xe1, 0x6e, 0x3a, 0xc6, 0x68, 0x95, 0xf3, 0x2b,
	0x39, 0xbc, 0x05, 0x7b, 0x99, 0x65, 0x89, 0xd6, 0x1a, 0x4e, 0xb3, 0xe6, 0x9e, 0xb4, 0x4a, 0x7b,
	0xc3, 0xf7, 0xa4, 0x55, 0xd9, 0x92, 0x7e, 0x65, 0x3a, 0x3c, 0x03, 0xdb, 0x99, 0xb1, 0x3f, 0x68,
	0x9d, 0x2f, 0xf7, 0xe4, 0xde, 0xe5, 0x58, 0x63, 0xbf, 0x34, 0x0d, 0x1e, 0x80, 0xd5, 0xab, 0xb0,
	0x08, 0x0a, 0xb4, 0xc1, 0x1d, 0x45, 0xc1, 0xd3, 0x68, 0x16, 0x14, 0x68, 0x53, 0xa6, 0xac, 0x80,
	0x8f, 0x00, 0x08, 0x08, 0xee, 0xc6, 0x69, 0x86, 0x83, 0x02, 0x01, 0x3e, 0x64, 0x24, 0xf0, 0x21,
	0xd8, 0xbc, 0xce, 0x13, 0x26, 0x8f, 0x33, 0xb4, 0xc5, 0x87, 0x75, 0xc0, 0xf7, 0x2b, 0xa5, 0x61,
	0xfc, 0x2a, 0x8d, 0xf3, 0x04, 0xf3, 0x10, 0x6d, 0xf3, 0xa6, 0x4a, 0x0e, 0x9b, 0x60, 0x57, 0x64,
	0x61, 0x9c, 0xcb, 0xd6, 0x1d, 0xde, 0x6a, 0xc7, 0xec, 0x0c, 0x2f, 0xc3, 0x8c, 0x06, 0x05, 0xaa,
	0xf1, 0x06, 0x59, 0xb1, 0xfc, 0x66, 0x8e, 0x99, 0xc2, 0xae, 0xc8, 0x45, 0x05, 0x11, 0x58, 0x5f,
	0x0a, 0xec, 0xf1, 0x81, 0x65, 0xc9, 0x66, 0x9c, 0x47, 0x93, 0x69, 0x50, 0xa0, 0xff, 0xc4, 0x0c,
	0x51, 0xb1, 0xbd, 0xb8, 0x4c, 0xef, 0x82, 0x02, 0x41, 0xb1, 0x17, 0xbc, 0x80, 0x0d, 0xb0, 0x25,
	0x4f, 0x99, 0x5f, 0xac, 0x7d, 0x7e, 0x81, 0xcc, 0x88, 0x39, 0x04, 0x04, 0x33, 0xec, 0xc5, 0x8c,
	0x62, 0x82, 0x33, 0x8a, 0x0e, 0x84, 0x83, 0x15, 0xb3, 0x7b, 0x1c, 0x10, 0x3c, 0xc0, 0x94, 0xc6,
	0x38, 0x20, 0xd1, 0x08, 0xa3, 0xff, 0x79, 0xa3, 0x95, 0xc2, 0x63, 0xb0, 0x5d, 0x5a, 0xee, 0x90,
	0x77, 0x95, 0x32, 0xf6, 0xbf, 0xcc, 0x85, 0x8e, 0x78, 0x8b, 0x19, 0xc1, 0x3a, 0xd8, 0x08, 0x08,
	0xee, 0xe1, 0x98, 0x86, 0x08, 0xf1, 0x61, 0x55, 0xb3, 0x13, 0xec, 0xe6, 0x84, 0x88, 0xc1, 0x07,
	0xe2, 0x04, 0x55, 0xc0, 0xbe, 0x8b, 0x4e, 0xbe, 0xb8, 0x21, 0x63, 0x4c, 0x6e, 0x73, 0x9c, 0x63,
	0x34, 0x6e, 0xac, 0xb0, 0xef, 0xa2, 0x14, 0x32, 0x9b, 0x01, 0x8e, 0x63, 0xa3, 0x0d, 0xf3, 0x36,
	0x2b, 0x65, 0xac, 0x4e, 0xbe, 0x38, 0x15, 0xff, 0xf3, 0x8d, 0x60, 0xa9, 0x80, 0xb9, 0xb2, 0x42,
	0xf4, 0xd3, 0x05, 0x9a, 0x08, 0x57, 0x33, 0x93, 0xff, 0xe7, 0xf4, 0x3a, 0x4f, 0x78, 0x94, 0xa1,
	0xa9, 0xf8, 0x4e, 0x4b, 0x21, 0xbb, 0xb5, 0x8c, 0x2c, 0x41, 0x91, 0xb8, 0xb5, 0x3a, 0x61, 0xab,
	0xf0, 0x4a, 0xa1, 0xde, 0x8a, 0x55, 0x4a, 0xe1, 0xd2, 0xca, 0x80, 0xbd, 0x13, 0x67, 0x54, 0x4e,
	0xa5, 0x95, 0x2b, 0x60, 0xb1, 0xb2, 0x72, 0x4d, 0x2b, 0x57, 0xa1, 0x12, 0x65, 0xe5, 0x5a, 0x56,
	0xae, 0x06, 0xcd, 0x94, 0x95, 0x5b, 0xb1, 0x92, 0xa0, 0x54, 0x5b, 0xb9, 0x25, 0x2b, 0x8d, 0x9a,
	0x6b, 0x2b, 0xd7, 0xb6, 0x32, 0x60, 0xef, 0xb5, 0x95, 0x6b, 0x5b, 0x79, 0x02, 0x46, 0x94, 0x95,
	0x67, 0x5a, 0x79, 0x0a, 0x95, 0x29, 0x2b, 0xcf, 0xb2, 0xf2, 0x34, 0x88, 0x2a, 0x2b, 0xaf, 0x62,
	0x25, 0x41, 0xb9, 0xb6, 0xf2, 0x4a, 0x56, 0x1a, 0xf5, 0x41, 0x5b, 0x79, 0xb6, 0x95, 0x01, 0xbb,
	0xd3, 0x56, 0x9e, 0x6d, 0xe5, 0x0b, 0x58, 0xa1, 0xac, 0x7c, 0xd3, 0xca, 0x57, 0xa8, 0x85, 0xb2,
	0xf2, 0x2d, 0x2b, 0x5f, 0x83, 0x3e, 0x2a, 0x2b, 0xbf, 0x62, 0x25, 0x41, 0x9f, 0xb4, 0x95, 0x5f,
	0xb2, 0xd2, 0xa8, 0xcf, 0xda, 0xca, 0xb7, 0xad, 0x0c, 0xd8, 0x17, 0x6d, 0xe5, 0xdb, 0x56, 0x6d,
	0x01, 0xfb, 0xaa, 0xac, 0xda, 0xa6, 0x55, 0x5b, 0xa1, 0xbe, 0x29, 0x2b, 0x95, 0xc1, 0x13, 0x6e,
	0xd5, 0xd6, 0xa0, 0xef, 0x8e, 0xd2, 0xd2, 0x29, 0x7c, 0x2c, 0xb4, 0x24, 0xe9, 0x87, 0xa3, 0xbd,
	0x24, 0xeb, 0x44, 0x78, 0x69, 0xd8, 0x4f, 0x47, 0x8b, 0x69, 0xdc, 0x73, 0x21, 0x66, 0xf0, 0x7e,
	0x39, 0xda, 0x4c, 0xc7, 0x9d, 0x53, 0x70, 0xcf, 0x43, 0xa2, 0xa3, 0x1e, 0x1d, 0x01, 0x7b, 0x29,
	0x64, 0xe7, 0xce, 0x6f, 0xc7, 0x79, 0xbd, 0xc6, 0x9f, 0x0d, 0xde, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0x12, 0xbb, 0xfa, 0x92, 0x08, 0x00, 0x00,
}
