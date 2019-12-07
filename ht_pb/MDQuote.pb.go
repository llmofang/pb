// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDQuote.proto

package ht_pb

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

//外汇交易中心报价信息
type MDQuote struct {
	HTSCSecurityID       string             `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32              `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32              `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64              `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	TradingPhaseCode     string             `protobuf:"bytes,5,opt,name=TradingPhaseCode,proto3" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource     ESecurityIDSource  `protobuf:"varint,6,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType      `protobuf:"varint,7,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	ExchangeDate         int32              `protobuf:"varint,8,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32              `protobuf:"varint,9,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	MaxPx                int64              `protobuf:"varint,10,opt,name=MaxPx,proto3" json:"MaxPx,omitempty"`
	MinPx                int64              `protobuf:"varint,11,opt,name=MinPx,proto3" json:"MinPx,omitempty"`
	ChannelNo            int32              `protobuf:"varint,12,opt,name=ChannelNo,proto3" json:"ChannelNo,omitempty"`
	ApplSeqNum           int64              `protobuf:"varint,13,opt,name=ApplSeqNum,proto3" json:"ApplSeqNum,omitempty"`
	MDBookType           int32              `protobuf:"varint,14,opt,name=MDBookType,proto3" json:"MDBookType,omitempty"`
	MarketIndicator      string             `protobuf:"bytes,15,opt,name=MarketIndicator,proto3" json:"MarketIndicator,omitempty"`
	MarketDepth          int32              `protobuf:"varint,16,opt,name=MarketDepth,proto3" json:"MarketDepth,omitempty"`
	MDSubBookType        int32              `protobuf:"varint,17,opt,name=MDSubBookType,proto3" json:"MDSubBookType,omitempty"`
	MDCashBondQuotes     []*MDCashBondQuote `protobuf:"bytes,18,rep,name=MDCashBondQuotes,proto3" json:"MDCashBondQuotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MDQuote) Reset()         { *m = MDQuote{} }
func (m *MDQuote) String() string { return proto.CompactTextString(m) }
func (*MDQuote) ProtoMessage()    {}
func (*MDQuote) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a97af7097301890, []int{0}
}

func (m *MDQuote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDQuote.Unmarshal(m, b)
}
func (m *MDQuote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDQuote.Marshal(b, m, deterministic)
}
func (m *MDQuote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDQuote.Merge(m, src)
}
func (m *MDQuote) XXX_Size() int {
	return xxx_messageInfo_MDQuote.Size(m)
}
func (m *MDQuote) XXX_DiscardUnknown() {
	xxx_messageInfo_MDQuote.DiscardUnknown(m)
}

var xxx_messageInfo_MDQuote proto.InternalMessageInfo

func (m *MDQuote) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDQuote) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDQuote) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDQuote) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDQuote) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDQuote) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDQuote) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDQuote) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *MDQuote) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func (m *MDQuote) GetMaxPx() int64 {
	if m != nil {
		return m.MaxPx
	}
	return 0
}

func (m *MDQuote) GetMinPx() int64 {
	if m != nil {
		return m.MinPx
	}
	return 0
}

func (m *MDQuote) GetChannelNo() int32 {
	if m != nil {
		return m.ChannelNo
	}
	return 0
}

func (m *MDQuote) GetApplSeqNum() int64 {
	if m != nil {
		return m.ApplSeqNum
	}
	return 0
}

func (m *MDQuote) GetMDBookType() int32 {
	if m != nil {
		return m.MDBookType
	}
	return 0
}

func (m *MDQuote) GetMarketIndicator() string {
	if m != nil {
		return m.MarketIndicator
	}
	return ""
}

func (m *MDQuote) GetMarketDepth() int32 {
	if m != nil {
		return m.MarketDepth
	}
	return 0
}

func (m *MDQuote) GetMDSubBookType() int32 {
	if m != nil {
		return m.MDSubBookType
	}
	return 0
}

func (m *MDQuote) GetMDCashBondQuotes() []*MDCashBondQuote {
	if m != nil {
		return m.MDCashBondQuotes
	}
	return nil
}

//中国外汇交易中心的现券买卖做市报价信息
type MDCashBondQuote struct {
	QuoteType            int32    `protobuf:"varint,1,opt,name=QuoteType,proto3" json:"QuoteType,omitempty"`
	Side                 int32    `protobuf:"varint,2,opt,name=Side,proto3" json:"Side,omitempty"`
	PriceLevel           int32    `protobuf:"varint,3,opt,name=PriceLevel,proto3" json:"PriceLevel,omitempty"`
	QuoteID              string   `protobuf:"bytes,4,opt,name=QuoteID,proto3" json:"QuoteID,omitempty"`
	QuoteDate            int32    `protobuf:"varint,5,opt,name=QuoteDate,proto3" json:"QuoteDate,omitempty"`
	QuoteTime            int32    `protobuf:"varint,6,opt,name=QuoteTime,proto3" json:"QuoteTime,omitempty"`
	CleanPrice           int64    `protobuf:"varint,7,opt,name=CleanPrice,proto3" json:"CleanPrice,omitempty"`
	DirtyPrice           int64    `protobuf:"varint,8,opt,name=DirtyPrice,proto3" json:"DirtyPrice,omitempty"`
	TotalFaceValue       int64    `protobuf:"varint,9,opt,name=TotalFaceValue,proto3" json:"TotalFaceValue,omitempty"`
	ClearingMethod       int32    `protobuf:"varint,10,opt,name=ClearingMethod,proto3" json:"ClearingMethod,omitempty"`
	SettlType            string   `protobuf:"bytes,11,opt,name=SettlType,proto3" json:"SettlType,omitempty"`
	SettlDate            int32    `protobuf:"varint,12,opt,name=SettlDate,proto3" json:"SettlDate,omitempty"`
	SettlCurrency        string   `protobuf:"bytes,13,opt,name=SettlCurrency,proto3" json:"SettlCurrency,omitempty"`
	SettlCurrFxRate      int64    `protobuf:"varint,14,opt,name=SettlCurrFxRate,proto3" json:"SettlCurrFxRate,omitempty"`
	PartyRole            int32    `protobuf:"varint,15,opt,name=PartyRole,proto3" json:"PartyRole,omitempty"`
	TraderCode           string   `protobuf:"bytes,16,opt,name=TraderCode,proto3" json:"TraderCode,omitempty"`
	MaturityYield        int64    `protobuf:"varint,17,opt,name=MaturityYield,proto3" json:"MaturityYield,omitempty"`
	DeliveryType         int32    `protobuf:"varint,18,opt,name=DeliveryType,proto3" json:"DeliveryType,omitempty"`
	TraderAccountID      string   `protobuf:"bytes,19,opt,name=TraderAccountID,proto3" json:"TraderAccountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MDCashBondQuote) Reset()         { *m = MDCashBondQuote{} }
func (m *MDCashBondQuote) String() string { return proto.CompactTextString(m) }
func (*MDCashBondQuote) ProtoMessage()    {}
func (*MDCashBondQuote) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a97af7097301890, []int{1}
}

func (m *MDCashBondQuote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDCashBondQuote.Unmarshal(m, b)
}
func (m *MDCashBondQuote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDCashBondQuote.Marshal(b, m, deterministic)
}
func (m *MDCashBondQuote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDCashBondQuote.Merge(m, src)
}
func (m *MDCashBondQuote) XXX_Size() int {
	return xxx_messageInfo_MDCashBondQuote.Size(m)
}
func (m *MDCashBondQuote) XXX_DiscardUnknown() {
	xxx_messageInfo_MDCashBondQuote.DiscardUnknown(m)
}

var xxx_messageInfo_MDCashBondQuote proto.InternalMessageInfo

func (m *MDCashBondQuote) GetQuoteType() int32 {
	if m != nil {
		return m.QuoteType
	}
	return 0
}

func (m *MDCashBondQuote) GetSide() int32 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *MDCashBondQuote) GetPriceLevel() int32 {
	if m != nil {
		return m.PriceLevel
	}
	return 0
}

func (m *MDCashBondQuote) GetQuoteID() string {
	if m != nil {
		return m.QuoteID
	}
	return ""
}

func (m *MDCashBondQuote) GetQuoteDate() int32 {
	if m != nil {
		return m.QuoteDate
	}
	return 0
}

func (m *MDCashBondQuote) GetQuoteTime() int32 {
	if m != nil {
		return m.QuoteTime
	}
	return 0
}

func (m *MDCashBondQuote) GetCleanPrice() int64 {
	if m != nil {
		return m.CleanPrice
	}
	return 0
}

func (m *MDCashBondQuote) GetDirtyPrice() int64 {
	if m != nil {
		return m.DirtyPrice
	}
	return 0
}

func (m *MDCashBondQuote) GetTotalFaceValue() int64 {
	if m != nil {
		return m.TotalFaceValue
	}
	return 0
}

func (m *MDCashBondQuote) GetClearingMethod() int32 {
	if m != nil {
		return m.ClearingMethod
	}
	return 0
}

func (m *MDCashBondQuote) GetSettlType() string {
	if m != nil {
		return m.SettlType
	}
	return ""
}

func (m *MDCashBondQuote) GetSettlDate() int32 {
	if m != nil {
		return m.SettlDate
	}
	return 0
}

func (m *MDCashBondQuote) GetSettlCurrency() string {
	if m != nil {
		return m.SettlCurrency
	}
	return ""
}

func (m *MDCashBondQuote) GetSettlCurrFxRate() int64 {
	if m != nil {
		return m.SettlCurrFxRate
	}
	return 0
}

func (m *MDCashBondQuote) GetPartyRole() int32 {
	if m != nil {
		return m.PartyRole
	}
	return 0
}

func (m *MDCashBondQuote) GetTraderCode() string {
	if m != nil {
		return m.TraderCode
	}
	return ""
}

func (m *MDCashBondQuote) GetMaturityYield() int64 {
	if m != nil {
		return m.MaturityYield
	}
	return 0
}

func (m *MDCashBondQuote) GetDeliveryType() int32 {
	if m != nil {
		return m.DeliveryType
	}
	return 0
}

func (m *MDCashBondQuote) GetTraderAccountID() string {
	if m != nil {
		return m.TraderAccountID
	}
	return ""
}

func init() {
	proto.RegisterType((*MDQuote)(nil), "com.htsc.mdc.insight.model.MDQuote")
	proto.RegisterType((*MDCashBondQuote)(nil), "com.htsc.mdc.insight.model.MDCashBondQuote")
}

func init() { proto.RegisterFile("MDQuote.proto", fileDescriptor_3a97af7097301890) }

var fileDescriptor_3a97af7097301890 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x6f, 0x6e, 0xd3, 0x30,
	0x18, 0xc6, 0x15, 0xba, 0x74, 0xab, 0xbb, 0x75, 0xc5, 0x43, 0x60, 0x4d, 0x68, 0x2a, 0x13, 0xa0,
	0x0a, 0xa4, 0x08, 0x8d, 0x13, 0xac, 0x4d, 0xa7, 0x55, 0x22, 0x53, 0xe7, 0x54, 0x20, 0x3e, 0x7a,
	0x89, 0xd5, 0x58, 0x4b, 0xe3, 0xe2, 0x38, 0x53, 0x7b, 0x02, 0xae, 0xc1, 0xf9, 0x38, 0x05, 0xf2,
	0x9b, 0x34, 0x4d, 0x52, 0xb1, 0x6f, 0xf1, 0xcf, 0x7e, 0xff, 0xf8, 0xc9, 0xe3, 0x17, 0x9d, 0x78,
	0xee, 0x7d, 0x26, 0x35, 0x77, 0x56, 0x4a, 0x6a, 0x89, 0xcf, 0x03, 0xb9, 0x74, 0x22, 0x9d, 0x06,
	0xce, 0x32, 0x0c, 0x1c, 0x91, 0xa4, 0x62, 0x11, 0x69, 0x67, 0x29, 0x43, 0x1e, 0x9f, 0xbf, 0x99,
	0xf8, 0x3c, 0xc8, 0x94, 0xd0, 0x9b, 0xa9, 0xeb, 0xcb, 0x4c, 0x05, 0x45, 0xd0, 0xf9, 0x59, 0xb9,
	0x31, 0xdf, 0xac, 0x0a, 0x78, 0xf9, 0xd7, 0x46, 0x87, 0x45, 0x6e, 0xfc, 0x11, 0xf5, 0x6e, 0xe7,
	0xfe, 0x78, 0x17, 0x4e, 0xac, 0x81, 0x35, 0xec, 0xd0, 0x06, 0xc5, 0xaf, 0x51, 0xdb, 0x73, 0x5d,
	0xa6, 0x39, 0x79, 0x31, 0xb0, 0x86, 0x36, 0x2d, 0x56, 0x39, 0x9f, 0x8b, 0x25, 0x27, 0xad, 0x2d,
	0x37, 0x2b, 0xfc, 0x1e, 0x9d, 0xb8, 0x4c, 0x33, 0xf3, 0x9d, 0x6a, 0xb6, 0x5c, 0x91, 0x83, 0x81,
	0x35, 0x6c, 0xd1, 0x3a, 0xc4, 0x9f, 0x50, 0x7f, 0xae, 0x58, 0x28, 0x92, 0xc5, 0x2c, 0x62, 0x29,
	0x1f, 0xcb, 0x90, 0x13, 0x1b, 0xea, 0xef, 0x71, 0x7c, 0x8f, 0xfa, 0x69, 0xe3, 0x92, 0xa4, 0x3d,
	0xb0, 0x86, 0xbd, 0xab, 0x0f, 0x4e, 0x4d, 0x1a, 0x90, 0xc4, 0xd9, 0x53, 0x84, 0xee, 0x85, 0xe3,
	0x09, 0x3a, 0x4e, 0x2b, 0xf2, 0x90, 0x43, 0x48, 0xf7, 0xee, 0xd9, 0x74, 0xe6, 0x20, 0xad, 0x85,
	0xe1, 0x4b, 0x74, 0x3c, 0x59, 0x07, 0x11, 0x4b, 0x16, 0x1c, 0x14, 0x3a, 0x02, 0x25, 0x6a, 0xac,
	0x7a, 0x06, 0xd4, 0xea, 0xd4, 0xcf, 0x80, 0x66, 0xaf, 0x90, 0xed, 0xb1, 0xf5, 0x6c, 0x4d, 0x10,
	0x68, 0x95, 0x2f, 0x80, 0x8a, 0x64, 0xb6, 0x26, 0xdd, 0x82, 0x9a, 0x05, 0x7e, 0x8b, 0x3a, 0xe3,
	0x88, 0x25, 0x09, 0x8f, 0xef, 0x24, 0x39, 0x86, 0x64, 0x3b, 0x80, 0x2f, 0x10, 0xba, 0x5e, 0xad,
	0x62, 0x9f, 0xff, 0xba, 0xcb, 0x96, 0xe4, 0x04, 0x02, 0x2b, 0xc4, 0xec, 0x7b, 0xee, 0x48, 0xca,
	0x47, 0xb8, 0x76, 0x0f, 0xc2, 0x2b, 0x04, 0x0f, 0xd1, 0xa9, 0xc7, 0xd4, 0x23, 0xd7, 0xd3, 0x24,
	0x14, 0x01, 0xd3, 0x52, 0x91, 0x53, 0xf8, 0x2d, 0x4d, 0x8c, 0x07, 0xa8, 0x9b, 0x23, 0x97, 0xaf,
	0x74, 0x44, 0xfa, 0x90, 0xaa, 0x8a, 0x8c, 0x13, 0x3c, 0xd7, 0xcf, 0x1e, 0xca, 0x72, 0x2f, 0xe1,
	0x4c, 0x1d, 0xe2, 0x1f, 0xa8, 0xef, 0xb9, 0x63, 0x96, 0x46, 0x23, 0x99, 0x84, 0x60, 0xcd, 0x94,
	0xe0, 0x41, 0x6b, 0xd8, 0xbd, 0xfa, 0xec, 0xfc, 0xdf, 0xf8, 0x4e, 0x23, 0x86, 0xee, 0x25, 0xb9,
	0xfc, 0x6d, 0xa3, 0xd3, 0x06, 0x34, 0xe2, 0xc1, 0x07, 0xb4, 0x63, 0xe5, 0xe2, 0x95, 0x00, 0x63,
	0x74, 0xe0, 0x8b, 0x70, 0x6b, 0x74, 0xf8, 0x36, 0x82, 0xcd, 0x94, 0x08, 0xf8, 0x37, 0xfe, 0xc4,
	0xe3, 0xc2, 0xea, 0x15, 0x82, 0x09, 0x3a, 0x84, 0x04, 0x53, 0x17, 0x8c, 0xde, 0xa1, 0xdb, 0x65,
	0x59, 0x0b, 0x9c, 0x61, 0x57, 0x6a, 0x81, 0x2d, 0xca, 0x4e, 0x8c, 0x27, 0xda, 0xd5, 0x4e, 0x8c,
	0x21, 0x2e, 0x10, 0x1a, 0xc7, 0x9c, 0x25, 0x50, 0x08, 0xdc, 0xd9, 0xa2, 0x15, 0x62, 0xf6, 0x5d,
	0xa1, 0xf4, 0x26, 0xdf, 0x3f, 0xca, 0xf7, 0x77, 0xc4, 0x3c, 0xee, 0xb9, 0xd4, 0x2c, 0xbe, 0x61,
	0x01, 0xff, 0xce, 0xe2, 0x2c, 0xb7, 0x5d, 0x8b, 0x36, 0xa8, 0x39, 0x67, 0xb2, 0x2a, 0x91, 0x2c,
	0x3c, 0xae, 0x23, 0x19, 0x82, 0x03, 0x6d, 0xda, 0xa0, 0xa6, 0x5b, 0x9f, 0x6b, 0x1d, 0x83, 0x6e,
	0x5d, 0xb8, 0xe7, 0x0e, 0x94, 0xbb, 0x70, 0xd3, 0xc2, 0x92, 0x25, 0x30, 0x36, 0x80, 0xc5, 0x38,
	0x53, 0x8a, 0x27, 0xc1, 0x06, 0x5c, 0xd9, 0xa1, 0x75, 0x68, 0x8c, 0x57, 0x82, 0x9b, 0x35, 0x35,
	0x99, 0x7a, 0xd0, 0x72, 0x13, 0x9b, 0x6a, 0x33, 0xa6, 0xf4, 0x86, 0xca, 0x98, 0x83, 0x39, 0x6d,
	0xba, 0x03, 0x46, 0x19, 0x33, 0x40, 0xb8, 0x82, 0x91, 0xd2, 0x87, 0x52, 0x15, 0x02, 0xa6, 0x64,
	0x1a, 0x9e, 0xf0, 0x4f, 0xc1, 0xe3, 0x10, 0x4c, 0xd9, 0xa2, 0x75, 0x68, 0x1e, 0xad, 0xcb, 0x63,
	0xf1, 0xc4, 0x55, 0x3e, 0x1f, 0x70, 0xfe, 0x68, 0xab, 0xcc, 0x74, 0x9c, 0xe7, 0xbd, 0x0e, 0x02,
	0x99, 0x25, 0x7a, 0xea, 0x92, 0xb3, 0xfc, 0xa9, 0x34, 0xf0, 0xe8, 0x0b, 0x7a, 0x66, 0x84, 0x8f,
	0xb6, 0xd3, 0x7e, 0x66, 0x46, 0x74, 0x7a, 0x6b, 0xfd, 0xb1, 0xac, 0x87, 0x36, 0xcc, 0xeb, 0xaf,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x77, 0x7f, 0x22, 0xfd, 0x0a, 0x06, 0x00, 0x00,
}
