// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDSpot.proto

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

//现货
type MDSpot struct {
	HTSCSecurityID          string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate                  int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime                  int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp           int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	TradingPhaseCode        string            `protobuf:"bytes,5,opt,name=TradingPhaseCode,proto3" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource        ESecurityIDSource `protobuf:"varint,6,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType            ESecurityType     `protobuf:"varint,7,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	MaxPx                   int64             `protobuf:"varint,8,opt,name=MaxPx,proto3" json:"MaxPx,omitempty"`
	MinPx                   int64             `protobuf:"varint,9,opt,name=MinPx,proto3" json:"MinPx,omitempty"`
	PreClosePx              int64             `protobuf:"varint,10,opt,name=PreClosePx,proto3" json:"PreClosePx,omitempty"`
	NumTrades               int64             `protobuf:"varint,11,opt,name=NumTrades,proto3" json:"NumTrades,omitempty"`
	TotalVolumeTrade        int64             `protobuf:"varint,12,opt,name=TotalVolumeTrade,proto3" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade         int64             `protobuf:"varint,13,opt,name=TotalValueTrade,proto3" json:"TotalValueTrade,omitempty"`
	TotalWeightTrade        int64             `protobuf:"varint,14,opt,name=TotalWeightTrade,proto3" json:"TotalWeightTrade,omitempty"`
	LastPx                  int64             `protobuf:"varint,15,opt,name=LastPx,proto3" json:"LastPx,omitempty"`
	OpenPx                  int64             `protobuf:"varint,16,opt,name=OpenPx,proto3" json:"OpenPx,omitempty"`
	ClosePx                 int64             `protobuf:"varint,17,opt,name=ClosePx,proto3" json:"ClosePx,omitempty"`
	HighPx                  int64             `protobuf:"varint,18,opt,name=HighPx,proto3" json:"HighPx,omitempty"`
	LowPx                   int64             `protobuf:"varint,19,opt,name=LowPx,proto3" json:"LowPx,omitempty"`
	TradingDate             int32             `protobuf:"varint,20,opt,name=TradingDate,proto3" json:"TradingDate,omitempty"`
	PreOpenInterest         int64             `protobuf:"varint,21,opt,name=PreOpenInterest,proto3" json:"PreOpenInterest,omitempty"`
	PreSettlePrice          int64             `protobuf:"varint,22,opt,name=PreSettlePrice,proto3" json:"PreSettlePrice,omitempty"`
	OpenInterest            int64             `protobuf:"varint,23,opt,name=OpenInterest,proto3" json:"OpenInterest,omitempty"`
	SettlePrice             int64             `protobuf:"varint,24,opt,name=SettlePrice,proto3" json:"SettlePrice,omitempty"`
	InitOpenInterest        int64             `protobuf:"varint,25,opt,name=InitOpenInterest,proto3" json:"InitOpenInterest,omitempty"`
	InterestChg             int64             `protobuf:"varint,26,opt,name=InterestChg,proto3" json:"InterestChg,omitempty"`
	AveragePx               int64             `protobuf:"varint,27,opt,name=AveragePx,proto3" json:"AveragePx,omitempty"`
	LifeHighPx              int64             `protobuf:"varint,28,opt,name=LifeHighPx,proto3" json:"LifeHighPx,omitempty"`
	LifeLowPx               int64             `protobuf:"varint,29,opt,name=LifeLowPx,proto3" json:"LifeLowPx,omitempty"`
	BuyPx                   int64             `protobuf:"varint,30,opt,name=BuyPx,proto3" json:"BuyPx,omitempty"`
	BuyQty                  int64             `protobuf:"varint,31,opt,name=BuyQty,proto3" json:"BuyQty,omitempty"`
	BuyImplyQty             int64             `protobuf:"varint,32,opt,name=BuyImplyQty,proto3" json:"BuyImplyQty,omitempty"`
	SellPx                  int64             `protobuf:"varint,33,opt,name=SellPx,proto3" json:"SellPx,omitempty"`
	SellQty                 int64             `protobuf:"varint,34,opt,name=SellQty,proto3" json:"SellQty,omitempty"`
	SellImplyQty            int64             `protobuf:"varint,35,opt,name=SellImplyQty,proto3" json:"SellImplyQty,omitempty"`
	CommodityContractNumber string            `protobuf:"bytes,36,opt,name=CommodityContractNumber,proto3" json:"CommodityContractNumber,omitempty"`
	ExchangeDate            int32             `protobuf:"varint,37,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime            int32             `protobuf:"varint,38,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	BuyPriceQueue           []int64           `protobuf:"varint,51,rep,packed,name=BuyPriceQueue,proto3" json:"BuyPriceQueue,omitempty"`
	BuyOrderQtyQueue        []int64           `protobuf:"varint,52,rep,packed,name=BuyOrderQtyQueue,proto3" json:"BuyOrderQtyQueue,omitempty"`
	SellPriceQueue          []int64           `protobuf:"varint,53,rep,packed,name=SellPriceQueue,proto3" json:"SellPriceQueue,omitempty"`
	SellOrderQtyQueue       []int64           `protobuf:"varint,54,rep,packed,name=SellOrderQtyQueue,proto3" json:"SellOrderQtyQueue,omitempty"`
	BuyOrderQueue           []int64           `protobuf:"varint,55,rep,packed,name=BuyOrderQueue,proto3" json:"BuyOrderQueue,omitempty"`
	SellOrderQueue          []int64           `protobuf:"varint,56,rep,packed,name=SellOrderQueue,proto3" json:"SellOrderQueue,omitempty"`
	BuyNumOrdersQueue       []int64           `protobuf:"varint,57,rep,packed,name=BuyNumOrdersQueue,proto3" json:"BuyNumOrdersQueue,omitempty"`
	SellNumOrdersQueue      []int64           `protobuf:"varint,58,rep,packed,name=SellNumOrdersQueue,proto3" json:"SellNumOrdersQueue,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *MDSpot) Reset()         { *m = MDSpot{} }
func (m *MDSpot) String() string { return proto.CompactTextString(m) }
func (*MDSpot) ProtoMessage()    {}
func (*MDSpot) Descriptor() ([]byte, []int) {
	return fileDescriptor_605cc5696f2326ef, []int{0}
}

func (m *MDSpot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDSpot.Unmarshal(m, b)
}
func (m *MDSpot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDSpot.Marshal(b, m, deterministic)
}
func (m *MDSpot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDSpot.Merge(m, src)
}
func (m *MDSpot) XXX_Size() int {
	return xxx_messageInfo_MDSpot.Size(m)
}
func (m *MDSpot) XXX_DiscardUnknown() {
	xxx_messageInfo_MDSpot.DiscardUnknown(m)
}

var xxx_messageInfo_MDSpot proto.InternalMessageInfo

func (m *MDSpot) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDSpot) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDSpot) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDSpot) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDSpot) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDSpot) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDSpot) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDSpot) GetMaxPx() int64 {
	if m != nil {
		return m.MaxPx
	}
	return 0
}

func (m *MDSpot) GetMinPx() int64 {
	if m != nil {
		return m.MinPx
	}
	return 0
}

func (m *MDSpot) GetPreClosePx() int64 {
	if m != nil {
		return m.PreClosePx
	}
	return 0
}

func (m *MDSpot) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func (m *MDSpot) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *MDSpot) GetTotalValueTrade() int64 {
	if m != nil {
		return m.TotalValueTrade
	}
	return 0
}

func (m *MDSpot) GetTotalWeightTrade() int64 {
	if m != nil {
		return m.TotalWeightTrade
	}
	return 0
}

func (m *MDSpot) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *MDSpot) GetOpenPx() int64 {
	if m != nil {
		return m.OpenPx
	}
	return 0
}

func (m *MDSpot) GetClosePx() int64 {
	if m != nil {
		return m.ClosePx
	}
	return 0
}

func (m *MDSpot) GetHighPx() int64 {
	if m != nil {
		return m.HighPx
	}
	return 0
}

func (m *MDSpot) GetLowPx() int64 {
	if m != nil {
		return m.LowPx
	}
	return 0
}

func (m *MDSpot) GetTradingDate() int32 {
	if m != nil {
		return m.TradingDate
	}
	return 0
}

func (m *MDSpot) GetPreOpenInterest() int64 {
	if m != nil {
		return m.PreOpenInterest
	}
	return 0
}

func (m *MDSpot) GetPreSettlePrice() int64 {
	if m != nil {
		return m.PreSettlePrice
	}
	return 0
}

func (m *MDSpot) GetOpenInterest() int64 {
	if m != nil {
		return m.OpenInterest
	}
	return 0
}

func (m *MDSpot) GetSettlePrice() int64 {
	if m != nil {
		return m.SettlePrice
	}
	return 0
}

func (m *MDSpot) GetInitOpenInterest() int64 {
	if m != nil {
		return m.InitOpenInterest
	}
	return 0
}

func (m *MDSpot) GetInterestChg() int64 {
	if m != nil {
		return m.InterestChg
	}
	return 0
}

func (m *MDSpot) GetAveragePx() int64 {
	if m != nil {
		return m.AveragePx
	}
	return 0
}

func (m *MDSpot) GetLifeHighPx() int64 {
	if m != nil {
		return m.LifeHighPx
	}
	return 0
}

func (m *MDSpot) GetLifeLowPx() int64 {
	if m != nil {
		return m.LifeLowPx
	}
	return 0
}

func (m *MDSpot) GetBuyPx() int64 {
	if m != nil {
		return m.BuyPx
	}
	return 0
}

func (m *MDSpot) GetBuyQty() int64 {
	if m != nil {
		return m.BuyQty
	}
	return 0
}

func (m *MDSpot) GetBuyImplyQty() int64 {
	if m != nil {
		return m.BuyImplyQty
	}
	return 0
}

func (m *MDSpot) GetSellPx() int64 {
	if m != nil {
		return m.SellPx
	}
	return 0
}

func (m *MDSpot) GetSellQty() int64 {
	if m != nil {
		return m.SellQty
	}
	return 0
}

func (m *MDSpot) GetSellImplyQty() int64 {
	if m != nil {
		return m.SellImplyQty
	}
	return 0
}

func (m *MDSpot) GetCommodityContractNumber() string {
	if m != nil {
		return m.CommodityContractNumber
	}
	return ""
}

func (m *MDSpot) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *MDSpot) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func (m *MDSpot) GetBuyPriceQueue() []int64 {
	if m != nil {
		return m.BuyPriceQueue
	}
	return nil
}

func (m *MDSpot) GetBuyOrderQtyQueue() []int64 {
	if m != nil {
		return m.BuyOrderQtyQueue
	}
	return nil
}

func (m *MDSpot) GetSellPriceQueue() []int64 {
	if m != nil {
		return m.SellPriceQueue
	}
	return nil
}

func (m *MDSpot) GetSellOrderQtyQueue() []int64 {
	if m != nil {
		return m.SellOrderQtyQueue
	}
	return nil
}

func (m *MDSpot) GetBuyOrderQueue() []int64 {
	if m != nil {
		return m.BuyOrderQueue
	}
	return nil
}

func (m *MDSpot) GetSellOrderQueue() []int64 {
	if m != nil {
		return m.SellOrderQueue
	}
	return nil
}

func (m *MDSpot) GetBuyNumOrdersQueue() []int64 {
	if m != nil {
		return m.BuyNumOrdersQueue
	}
	return nil
}

func (m *MDSpot) GetSellNumOrdersQueue() []int64 {
	if m != nil {
		return m.SellNumOrdersQueue
	}
	return nil
}

func init() {
	proto.RegisterType((*MDSpot)(nil), "com.htsc.mdc.insight.model.MDSpot")
}

func init() { proto.RegisterFile("MDSpot.proto", fileDescriptor_605cc5696f2326ef) }

var fileDescriptor_605cc5696f2326ef = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0x6d, 0x4f, 0xdb, 0x48,
	0x10, 0xc7, 0x65, 0x72, 0xc0, 0xb1, 0x24, 0x21, 0x2c, 0x1c, 0xec, 0xe5, 0x38, 0x2e, 0x70, 0x80,
	0x22, 0x5e, 0xf8, 0x10, 0xdc, 0x5d, 0x69, 0xdf, 0xd5, 0x09, 0x12, 0x91, 0x78, 0x70, 0x1e, 0xd4,
	0xbe, 0x36, 0xce, 0x36, 0xb1, 0x64, 0xc7, 0x91, 0xbd, 0x6e, 0xed, 0x6f, 0xd3, 0x8f, 0xd0, 0x8f,
	0x58, 0xcd, 0x8c, 0xe3, 0xd8, 0x4e, 0xcb, 0x3b, 0xcf, 0x6f, 0xfe, 0x33, 0x9e, 0xfd, 0x7b, 0xe4,
	0x65, 0xd5, 0xc7, 0xee, 0x70, 0xee, 0x2b, 0x7d, 0x1e, 0xf8, 0xca, 0xe7, 0x4d, 0xdb, 0xf7, 0xf4,
	0xa9, 0x0a, 0x6d, 0xdd, 0x1b, 0xdb, 0xba, 0x33, 0x0b, 0x9d, 0xc9, 0x54, 0xe9, 0x9e, 0x3f, 0x96,
	0x6e, 0xf3, 0xf0, 0x6e, 0x28, 0xed, 0x28, 0x70, 0x54, 0xd2, 0xeb, 0x0e, 0xfd, 0x28, 0xb0, 0x25,
	0x15, 0x35, 0xf7, 0xb2, 0xc4, 0x28, 0x99, 0xa7, 0xf0, 0xf4, 0x5b, 0x8d, 0x6d, 0x50, 0x6b, 0x7e,
	0xc1, 0xea, 0xf7, 0xa3, 0x61, 0x67, 0x59, 0x2d, 0xb4, 0x96, 0xd6, 0xde, 0x1a, 0x94, 0x28, 0x3f,
	0x80, 0x8a, 0xae, 0xa5, 0xa4, 0x58, 0x6b, 0x69, 0xed, 0xf5, 0x41, 0x1a, 0x11, 0x1f, 0x39, 0x9e,
	0x14, 0x95, 0x05, 0x87, 0x88, 0x9f, 0xb1, 0x5a, 0xd7, 0x52, 0x16, 0x3c, 0x87, 0xca, 0xf2, 0xe6,
	0xe2, 0x97, 0x96, 0xd6, 0xae, 0x0c, 0x8a, 0x90, 0x5f, 0xb2, 0xc6, 0x28, 0xb0, 0xc6, 0xce, 0x6c,
	0x62, 0x4e, 0xad, 0x50, 0x76, 0xfc, 0xb1, 0x14, 0xeb, 0xf8, 0xfe, 0x15, 0xce, 0xfb, 0xac, 0x11,
	0x96, 0xce, 0x28, 0x36, 0x5a, 0x5a, 0xbb, 0x7e, 0x7d, 0xae, 0x17, 0x9c, 0x41, 0x47, 0xf4, 0x15,
	0x43, 0x06, 0x2b, 0xe5, 0xfc, 0x8e, 0x55, 0xc3, 0x9c, 0x3b, 0x62, 0x13, 0xdb, 0x9d, 0xbc, 0xda,
	0x0e, 0x84, 0x83, 0x42, 0x19, 0xdf, 0x67, 0xeb, 0x8f, 0x56, 0x6c, 0xc6, 0xe2, 0x57, 0x3c, 0x23,
	0x05, 0x48, 0x9d, 0x99, 0x19, 0x8b, 0xad, 0x94, 0x42, 0xc0, 0x8f, 0x19, 0x33, 0x03, 0xd9, 0x71,
	0xfd, 0x50, 0x9a, 0xb1, 0x60, 0x98, 0xca, 0x11, 0x7e, 0xc4, 0xb6, 0x9e, 0x22, 0x0f, 0x0e, 0x2f,
	0x43, 0xb1, 0x8d, 0xe9, 0x25, 0x40, 0xbf, 0x7c, 0x65, 0xb9, 0x1f, 0x7c, 0x37, 0xf2, 0x24, 0x42,
	0x51, 0x45, 0xd1, 0x0a, 0xe7, 0x6d, 0xb6, 0x43, 0xcc, 0x72, 0xa3, 0x54, 0x5a, 0x43, 0x69, 0x19,
	0x67, 0x5d, 0x3f, 0x4a, 0xd8, 0x28, 0x92, 0xd6, 0x73, 0x5d, 0x73, 0x1c, 0xbe, 0xf7, 0x83, 0x15,
	0x2a, 0x33, 0x16, 0x3b, 0xa8, 0x48, 0x23, 0xe0, 0xcf, 0x73, 0x09, 0xc7, 0x6d, 0x10, 0xa7, 0x88,
	0x0b, 0xb6, 0xb9, 0x38, 0xec, 0x2e, 0x26, 0x16, 0x21, 0x54, 0xdc, 0x3b, 0x93, 0xa9, 0x19, 0x0b,
	0x4e, 0x15, 0x14, 0x81, 0x6f, 0x0f, 0xfe, 0x17, 0x33, 0x16, 0x7b, 0xe4, 0x1b, 0x06, 0xbc, 0xc5,
	0xb6, 0xd3, 0x8d, 0xc0, 0x25, 0xdc, 0xc7, 0x65, 0xcb, 0x23, 0x38, 0xaf, 0x19, 0x48, 0x78, 0x6d,
	0x6f, 0xa6, 0x64, 0x20, 0x43, 0x25, 0x7e, 0xa3, 0xf3, 0x96, 0x30, 0xec, 0xbc, 0x19, 0xc8, 0xa1,
	0x54, 0xca, 0x95, 0x66, 0xe0, 0xd8, 0x52, 0x1c, 0xa0, 0xb0, 0x44, 0xf9, 0x29, 0xab, 0x16, 0xda,
	0x1d, 0xa2, 0xaa, 0xc0, 0x60, 0xae, 0x7c, 0x23, 0x81, 0x92, 0x3c, 0x02, 0x77, 0x7b, 0x33, 0x47,
	0x15, 0x3a, 0xfd, 0x4e, 0xee, 0x96, 0x39, 0x74, 0x5b, 0x3c, 0x77, 0xa6, 0x13, 0xd1, 0xa4, 0x6e,
	0x39, 0x04, 0xfb, 0xf1, 0xfe, 0xb3, 0x0c, 0xac, 0x09, 0x38, 0xfa, 0x07, 0xed, 0x47, 0x06, 0x60,
	0xbb, 0x1e, 0x9c, 0x4f, 0x32, 0xf5, 0xf5, 0x88, 0xb6, 0x6b, 0x49, 0xa0, 0x1a, 0x22, 0xf2, 0xf7,
	0x4f, 0xaa, 0xce, 0x00, 0x38, 0x6f, 0x44, 0x89, 0x19, 0x8b, 0x63, 0x72, 0x1e, 0x03, 0xf8, 0x4e,
	0x46, 0x94, 0xf4, 0x55, 0x22, 0xfe, 0xa2, 0xef, 0x44, 0x11, 0xcc, 0x6a, 0x44, 0x49, 0xcf, 0x9b,
	0xbb, 0x98, 0x6c, 0xd1, 0xac, 0x39, 0x04, 0x95, 0x43, 0xe9, 0xba, 0x66, 0x2c, 0x4e, 0xa8, 0x92,
	0x22, 0xd8, 0x09, 0x78, 0x82, 0xaa, 0x53, 0xda, 0x89, 0x34, 0x04, 0xc7, 0xe1, 0x31, 0x6b, 0xfa,
	0x37, 0x39, 0x9e, 0x67, 0xfc, 0x96, 0x1d, 0x76, 0x7c, 0xcf, 0xf3, 0xc7, 0x8e, 0x4a, 0x3a, 0xfe,
	0x4c, 0x05, 0x96, 0xad, 0x9e, 0x22, 0xef, 0x45, 0x06, 0xe2, 0x0c, 0x7f, 0x1d, 0x3f, 0x4b, 0x43,
	0xf7, 0xbb, 0xd8, 0x9e, 0x5a, 0xb3, 0x89, 0xc4, 0x25, 0x3a, 0xc7, 0x25, 0x2a, 0xb0, 0xbc, 0x06,
	0xff, 0x6a, 0x17, 0x45, 0x0d, 0xfe, 0xdb, 0xda, 0xac, 0x06, 0xd6, 0xc0, 0xd7, 0xed, 0x47, 0x32,
	0x92, 0xe2, 0xa6, 0x55, 0x69, 0x57, 0x8c, 0xb5, 0x86, 0x36, 0x28, 0x26, 0xb8, 0xce, 0x1a, 0x46,
	0x94, 0x3c, 0x07, 0x63, 0x19, 0xf4, 0x55, 0x42, 0xe2, 0x7f, 0x33, 0xf1, 0x4a, 0x8e, 0x5f, 0xb2,
	0x3a, 0x7a, 0xb4, 0x6c, 0xfd, 0x5f, 0xa6, 0x2e, 0x65, 0xf8, 0x15, 0xdb, 0x05, 0x52, 0x6c, 0xfe,
	0x7f, 0x26, 0x5f, 0x4d, 0xa6, 0x73, 0x13, 0x43, 0xf5, 0x9b, 0xc2, 0xdc, 0xcb, 0xc4, 0x62, 0x8e,
	0x9c, 0xf4, 0xb6, 0x38, 0x47, 0x4e, 0x7b, 0xc5, 0x76, 0x8d, 0x28, 0x79, 0x8a, 0x3c, 0x64, 0x21,
	0xc9, 0xdf, 0x2e, 0xe7, 0x58, 0x49, 0xf2, 0x6b, 0xc6, 0xa1, 0x47, 0xa9, 0xe4, 0x5d, 0x56, 0xf2,
	0x83, 0xac, 0xf1, 0x0f, 0x7b, 0xe5, 0xfa, 0x33, 0xd2, 0x8b, 0xd2, 0x84, 0xdb, 0x2d, 0xbc, 0xd7,
	0xbe, 0x6a, 0xda, 0xcb, 0x06, 0x5e, 0x75, 0x37, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0x9f,
	0x2d, 0x34, 0x44, 0x07, 0x00, 0x00,
}
