// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ADVolumeByPrice.proto

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

//成交分价数据(ADVolumeByPrice)
type ADVolumeByPrice struct {
	HTSCSecurityID       string                   `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32                    `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32                    `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64                    `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	SecurityIDSource     ESecurityIDSource        `protobuf:"varint,5,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType            `protobuf:"varint,6,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	TotalVolumeTrade     int64                    `protobuf:"varint,7,opt,name=TotalVolumeTrade,proto3" json:"TotalVolumeTrade,omitempty"`
	Details              []*ADVolumeByPriceDetail `protobuf:"bytes,8,rep,name=Details,proto3" json:"Details,omitempty"`
	ExchangeDate         int32                    `protobuf:"varint,9,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32                    `protobuf:"varint,10,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ADVolumeByPrice) Reset()         { *m = ADVolumeByPrice{} }
func (m *ADVolumeByPrice) String() string { return proto.CompactTextString(m) }
func (*ADVolumeByPrice) ProtoMessage()    {}
func (*ADVolumeByPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_b18d28d45a112853, []int{0}
}

func (m *ADVolumeByPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADVolumeByPrice.Unmarshal(m, b)
}
func (m *ADVolumeByPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADVolumeByPrice.Marshal(b, m, deterministic)
}
func (m *ADVolumeByPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADVolumeByPrice.Merge(m, src)
}
func (m *ADVolumeByPrice) XXX_Size() int {
	return xxx_messageInfo_ADVolumeByPrice.Size(m)
}
func (m *ADVolumeByPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_ADVolumeByPrice.DiscardUnknown(m)
}

var xxx_messageInfo_ADVolumeByPrice proto.InternalMessageInfo

func (m *ADVolumeByPrice) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *ADVolumeByPrice) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *ADVolumeByPrice) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *ADVolumeByPrice) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *ADVolumeByPrice) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *ADVolumeByPrice) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *ADVolumeByPrice) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *ADVolumeByPrice) GetDetails() []*ADVolumeByPriceDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *ADVolumeByPrice) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *ADVolumeByPrice) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

//按价分量详情
type ADVolumeByPriceDetail struct {
	TradePrice           int64    `protobuf:"varint,1,opt,name=TradePrice,proto3" json:"TradePrice,omitempty"`
	TotalQty             int64    `protobuf:"varint,2,opt,name=TotalQty,proto3" json:"TotalQty,omitempty"`
	BuyQty               int64    `protobuf:"varint,3,opt,name=BuyQty,proto3" json:"BuyQty,omitempty"`
	SellQty              int64    `protobuf:"varint,4,opt,name=SellQty,proto3" json:"SellQty,omitempty"`
	TotalNumbers         int64    `protobuf:"varint,5,opt,name=TotalNumbers,proto3" json:"TotalNumbers,omitempty"`
	BuyNumbers           int64    `protobuf:"varint,6,opt,name=BuyNumbers,proto3" json:"BuyNumbers,omitempty"`
	SellNumbers          int64    `protobuf:"varint,7,opt,name=SellNumbers,proto3" json:"SellNumbers,omitempty"`
	VolumePerNumber      int64    `protobuf:"varint,8,opt,name=VolumePerNumber,proto3" json:"VolumePerNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ADVolumeByPriceDetail) Reset()         { *m = ADVolumeByPriceDetail{} }
func (m *ADVolumeByPriceDetail) String() string { return proto.CompactTextString(m) }
func (*ADVolumeByPriceDetail) ProtoMessage()    {}
func (*ADVolumeByPriceDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_b18d28d45a112853, []int{1}
}

func (m *ADVolumeByPriceDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADVolumeByPriceDetail.Unmarshal(m, b)
}
func (m *ADVolumeByPriceDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADVolumeByPriceDetail.Marshal(b, m, deterministic)
}
func (m *ADVolumeByPriceDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADVolumeByPriceDetail.Merge(m, src)
}
func (m *ADVolumeByPriceDetail) XXX_Size() int {
	return xxx_messageInfo_ADVolumeByPriceDetail.Size(m)
}
func (m *ADVolumeByPriceDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ADVolumeByPriceDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ADVolumeByPriceDetail proto.InternalMessageInfo

func (m *ADVolumeByPriceDetail) GetTradePrice() int64 {
	if m != nil {
		return m.TradePrice
	}
	return 0
}

func (m *ADVolumeByPriceDetail) GetTotalQty() int64 {
	if m != nil {
		return m.TotalQty
	}
	return 0
}

func (m *ADVolumeByPriceDetail) GetBuyQty() int64 {
	if m != nil {
		return m.BuyQty
	}
	return 0
}

func (m *ADVolumeByPriceDetail) GetSellQty() int64 {
	if m != nil {
		return m.SellQty
	}
	return 0
}

func (m *ADVolumeByPriceDetail) GetTotalNumbers() int64 {
	if m != nil {
		return m.TotalNumbers
	}
	return 0
}

func (m *ADVolumeByPriceDetail) GetBuyNumbers() int64 {
	if m != nil {
		return m.BuyNumbers
	}
	return 0
}

func (m *ADVolumeByPriceDetail) GetSellNumbers() int64 {
	if m != nil {
		return m.SellNumbers
	}
	return 0
}

func (m *ADVolumeByPriceDetail) GetVolumePerNumber() int64 {
	if m != nil {
		return m.VolumePerNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*ADVolumeByPrice)(nil), "com.htsc.mdc.insight.model.ADVolumeByPrice")
	proto.RegisterType((*ADVolumeByPriceDetail)(nil), "com.htsc.mdc.insight.model.ADVolumeByPriceDetail")
}

func init() { proto.RegisterFile("ADVolumeByPrice.proto", fileDescriptor_b18d28d45a112853) }

var fileDescriptor_b18d28d45a112853 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xe5, 0x9a, 0xee, 0x6e, 0xa7, 0xa5, 0xad, 0x8c, 0x00, 0x6b, 0x0f, 0x28, 0xac, 0x00,
	0x45, 0x1c, 0x2c, 0x51, 0x2e, 0x5c, 0x09, 0x59, 0xa9, 0x08, 0x81, 0xb6, 0xde, 0x88, 0xbb, 0x9b,
	0xb5, 0xba, 0x91, 0x92, 0x66, 0x15, 0x3b, 0x12, 0x79, 0x10, 0xee, 0x3c, 0x1f, 0x4f, 0x81, 0x3c,
	0xde, 0xb4, 0xf9, 0x23, 0xf6, 0x96, 0xf9, 0x79, 0xbe, 0xf1, 0xcc, 0x7c, 0x31, 0x3c, 0xff, 0x1c,
	0xff, 0x2c, 0xf3, 0xba, 0xd0, 0x51, 0xb3, 0xaa, 0xb2, 0x54, 0x8b, 0x5d, 0x55, 0xda, 0x92, 0xcd,
	0xd3, 0xb2, 0x10, 0x5b, 0x6b, 0x52, 0x51, 0x6c, 0x52, 0x91, 0xdd, 0x9b, 0xec, 0x6e, 0x6b, 0x45,
	0x51, 0x6e, 0x74, 0x3e, 0x7f, 0xb6, 0x5c, 0xeb, 0xb4, 0xae, 0x32, 0xdb, 0x24, 0xcd, 0x6e, 0x2f,
	0x98, 0xbf, 0x7c, 0x80, 0x5f, 0xe3, 0x75, 0x59, 0x57, 0x6d, 0xa5, 0xc5, 0x5f, 0x0a, 0x17, 0x83,
	0x3b, 0xd8, 0x3b, 0x38, 0xbf, 0x4e, 0xd6, 0x5f, 0x1e, 0x15, 0x9c, 0x04, 0x24, 0x3c, 0x91, 0x03,
	0xca, 0x5e, 0xc0, 0xe4, 0x7b, 0x1c, 0x2b, 0xab, 0xf9, 0x51, 0x40, 0xc2, 0x63, 0xb9, 0x8f, 0x3c,
	0x4f, 0xb2, 0x42, 0x73, 0xda, 0x72, 0x17, 0xb1, 0x37, 0xf0, 0x34, 0x56, 0x56, 0xb9, 0x6f, 0x63,
	0x55, 0xb1, 0xe3, 0x4f, 0x02, 0x12, 0x52, 0xd9, 0x87, 0xec, 0x06, 0x2e, 0xcd, 0xa0, 0x57, 0x7e,
	0x1c, 0x90, 0xf0, 0xfc, 0xea, 0xad, 0xe8, 0x8d, 0x8d, 0xe3, 0x8a, 0xd1, 0x60, 0x72, 0x24, 0x67,
	0x4b, 0x38, 0x33, 0x9d, 0x9d, 0xf0, 0x09, 0x96, 0x7b, 0x7d, 0xb0, 0x9c, 0x4b, 0x94, 0x3d, 0x19,
	0x7b, 0x0f, 0x97, 0x49, 0x69, 0x55, 0xee, 0xb7, 0x95, 0x54, 0x6a, 0xa3, 0xf9, 0x14, 0x47, 0x18,
	0x71, 0xf6, 0x0d, 0xa6, 0xb1, 0xb6, 0x2a, 0xcb, 0x0d, 0x9f, 0x05, 0x34, 0x3c, 0xbd, 0xfa, 0x20,
	0xfe, 0xef, 0x99, 0x18, 0x38, 0xe0, 0x95, 0xb2, 0xad, 0xc0, 0x16, 0x70, 0xb6, 0xfc, 0x95, 0x6e,
	0xd5, 0xfd, 0x9d, 0xc6, 0x75, 0x9f, 0xe0, 0x5a, 0x7b, 0xac, 0x9b, 0x83, 0xab, 0x87, 0x7e, 0x8e,
	0x63, 0x8b, 0xdf, 0x47, 0xa3, 0x1f, 0xca, 0x5f, 0xc1, 0x5e, 0x01, 0x60, 0xdf, 0xc8, 0xd0, 0x6e,
	0x2a, 0x3b, 0x84, 0xcd, 0x61, 0x86, 0x23, 0xde, 0xd8, 0x06, 0xcd, 0xa6, 0xf2, 0x21, 0x76, 0x76,
	0x47, 0x75, 0xe3, 0x4e, 0x28, 0x9e, 0xec, 0x23, 0xc6, 0x61, 0xba, 0xd6, 0x39, 0x4a, 0xbc, 0xd1,
	0x6d, 0xe8, 0x7a, 0x45, 0xf5, 0x8f, 0xba, 0xb8, 0xd5, 0x95, 0x41, 0x7b, 0xa9, 0xec, 0x31, 0xd7,
	0x51, 0x54, 0x37, 0x6d, 0xc6, 0xc4, 0x77, 0xf4, 0x48, 0x58, 0x00, 0xa7, 0xae, 0x5c, 0x9b, 0xe0,
	0x7d, 0xe8, 0x22, 0x16, 0xc2, 0x85, 0x1f, 0x75, 0xa5, 0x2b, 0xcf, 0xf8, 0x0c, 0xb3, 0x86, 0x38,
	0xfa, 0x04, 0x07, 0x1e, 0x54, 0x34, 0x5c, 0xd9, 0xca, 0x3d, 0x1c, 0x73, 0x4d, 0xfe, 0x10, 0x72,
	0x3b, 0xc1, 0x57, 0xf4, 0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x80, 0x1d, 0x2e, 0xa8,
	0x03, 0x00, 0x00,
}