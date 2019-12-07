// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDFIQuote.proto

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

//固定收益（Fixed Income）确定报价数据
type MDFIQuote struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,5,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,6,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	ExchangeDate         int32             `protobuf:"varint,7,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32             `protobuf:"varint,8,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	MessageNumber        int64             `protobuf:"varint,9,opt,name=MessageNumber,proto3" json:"MessageNumber,omitempty"`
	BrokerDataType       int32             `protobuf:"varint,10,opt,name=BrokerDataType,proto3" json:"BrokerDataType,omitempty"`
	AccruedInterest      int64             `protobuf:"varint,11,opt,name=AccruedInterest,proto3" json:"AccruedInterest,omitempty"`
	BuyQuoteID           string            `protobuf:"bytes,20,opt,name=BuyQuoteID,proto3" json:"BuyQuoteID,omitempty"`
	BuyQuoteTime         int32             `protobuf:"varint,21,opt,name=BuyQuoteTime,proto3" json:"BuyQuoteTime,omitempty"`
	BuyQuoter            string            `protobuf:"bytes,22,opt,name=BuyQuoter,proto3" json:"BuyQuoter,omitempty"`
	BuyCleanPrice        int64             `protobuf:"varint,23,opt,name=BuyCleanPrice,proto3" json:"BuyCleanPrice,omitempty"`
	BuyVolume            int64             `protobuf:"varint,24,opt,name=BuyVolume,proto3" json:"BuyVolume,omitempty"`
	BuyDirtyPrice        int64             `protobuf:"varint,25,opt,name=BuyDirtyPrice,proto3" json:"BuyDirtyPrice,omitempty"`
	BuyMaturityYield     int64             `protobuf:"varint,26,opt,name=BuyMaturityYield,proto3" json:"BuyMaturityYield,omitempty"`
	BuyQuoteComment      string            `protobuf:"bytes,27,opt,name=BuyQuoteComment,proto3" json:"BuyQuoteComment,omitempty"`
	BuyQuoteType         int32             `protobuf:"varint,28,opt,name=BuyQuoteType,proto3" json:"BuyQuoteType,omitempty"`
	BuyBargainType       int32             `protobuf:"varint,29,opt,name=BuyBargainType,proto3" json:"BuyBargainType,omitempty"`
	BuyRelationType      int32             `protobuf:"varint,30,opt,name=BuyRelationType,proto3" json:"BuyRelationType,omitempty"`
	BuyExerciseType      int32             `protobuf:"varint,31,opt,name=BuyExerciseType,proto3" json:"BuyExerciseType,omitempty"`
	BuyYieldType         int32             `protobuf:"varint,32,opt,name=BuyYieldType,proto3" json:"BuyYieldType,omitempty"`
	SellQuoteID          string            `protobuf:"bytes,33,opt,name=SellQuoteID,proto3" json:"SellQuoteID,omitempty"`
	SellQuoteTime        int32             `protobuf:"varint,34,opt,name=SellQuoteTime,proto3" json:"SellQuoteTime,omitempty"`
	SellQuoter           string            `protobuf:"bytes,35,opt,name=SellQuoter,proto3" json:"SellQuoter,omitempty"`
	SellCleanPrice       int64             `protobuf:"varint,36,opt,name=SellCleanPrice,proto3" json:"SellCleanPrice,omitempty"`
	SellVolume           int64             `protobuf:"varint,37,opt,name=SellVolume,proto3" json:"SellVolume,omitempty"`
	SellDirtyPrice       int64             `protobuf:"varint,38,opt,name=SellDirtyPrice,proto3" json:"SellDirtyPrice,omitempty"`
	SellMaturityYield    int64             `protobuf:"varint,39,opt,name=SellMaturityYield,proto3" json:"SellMaturityYield,omitempty"`
	SellQuoteComment     string            `protobuf:"bytes,40,opt,name=SellQuoteComment,proto3" json:"SellQuoteComment,omitempty"`
	SellQuoteType        int32             `protobuf:"varint,41,opt,name=SellQuoteType,proto3" json:"SellQuoteType,omitempty"`
	SellBargainType      int32             `protobuf:"varint,42,opt,name=SellBargainType,proto3" json:"SellBargainType,omitempty"`
	SellRelationType     int32             `protobuf:"varint,43,opt,name=SellRelationType,proto3" json:"SellRelationType,omitempty"`
	SellExerciseType     int32             `protobuf:"varint,44,opt,name=SellExerciseType,proto3" json:"SellExerciseType,omitempty"`
	SellYieldType        int32             `protobuf:"varint,45,opt,name=SellYieldType,proto3" json:"SellYieldType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MDFIQuote) Reset()         { *m = MDFIQuote{} }
func (m *MDFIQuote) String() string { return proto.CompactTextString(m) }
func (*MDFIQuote) ProtoMessage()    {}
func (*MDFIQuote) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ab23b79765a8fc8, []int{0}
}

func (m *MDFIQuote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDFIQuote.Unmarshal(m, b)
}
func (m *MDFIQuote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDFIQuote.Marshal(b, m, deterministic)
}
func (m *MDFIQuote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDFIQuote.Merge(m, src)
}
func (m *MDFIQuote) XXX_Size() int {
	return xxx_messageInfo_MDFIQuote.Size(m)
}
func (m *MDFIQuote) XXX_DiscardUnknown() {
	xxx_messageInfo_MDFIQuote.DiscardUnknown(m)
}

var xxx_messageInfo_MDFIQuote proto.InternalMessageInfo

func (m *MDFIQuote) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDFIQuote) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDFIQuote) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDFIQuote) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDFIQuote) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDFIQuote) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDFIQuote) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *MDFIQuote) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func (m *MDFIQuote) GetMessageNumber() int64 {
	if m != nil {
		return m.MessageNumber
	}
	return 0
}

func (m *MDFIQuote) GetBrokerDataType() int32 {
	if m != nil {
		return m.BrokerDataType
	}
	return 0
}

func (m *MDFIQuote) GetAccruedInterest() int64 {
	if m != nil {
		return m.AccruedInterest
	}
	return 0
}

func (m *MDFIQuote) GetBuyQuoteID() string {
	if m != nil {
		return m.BuyQuoteID
	}
	return ""
}

func (m *MDFIQuote) GetBuyQuoteTime() int32 {
	if m != nil {
		return m.BuyQuoteTime
	}
	return 0
}

func (m *MDFIQuote) GetBuyQuoter() string {
	if m != nil {
		return m.BuyQuoter
	}
	return ""
}

func (m *MDFIQuote) GetBuyCleanPrice() int64 {
	if m != nil {
		return m.BuyCleanPrice
	}
	return 0
}

func (m *MDFIQuote) GetBuyVolume() int64 {
	if m != nil {
		return m.BuyVolume
	}
	return 0
}

func (m *MDFIQuote) GetBuyDirtyPrice() int64 {
	if m != nil {
		return m.BuyDirtyPrice
	}
	return 0
}

func (m *MDFIQuote) GetBuyMaturityYield() int64 {
	if m != nil {
		return m.BuyMaturityYield
	}
	return 0
}

func (m *MDFIQuote) GetBuyQuoteComment() string {
	if m != nil {
		return m.BuyQuoteComment
	}
	return ""
}

func (m *MDFIQuote) GetBuyQuoteType() int32 {
	if m != nil {
		return m.BuyQuoteType
	}
	return 0
}

func (m *MDFIQuote) GetBuyBargainType() int32 {
	if m != nil {
		return m.BuyBargainType
	}
	return 0
}

func (m *MDFIQuote) GetBuyRelationType() int32 {
	if m != nil {
		return m.BuyRelationType
	}
	return 0
}

func (m *MDFIQuote) GetBuyExerciseType() int32 {
	if m != nil {
		return m.BuyExerciseType
	}
	return 0
}

func (m *MDFIQuote) GetBuyYieldType() int32 {
	if m != nil {
		return m.BuyYieldType
	}
	return 0
}

func (m *MDFIQuote) GetSellQuoteID() string {
	if m != nil {
		return m.SellQuoteID
	}
	return ""
}

func (m *MDFIQuote) GetSellQuoteTime() int32 {
	if m != nil {
		return m.SellQuoteTime
	}
	return 0
}

func (m *MDFIQuote) GetSellQuoter() string {
	if m != nil {
		return m.SellQuoter
	}
	return ""
}

func (m *MDFIQuote) GetSellCleanPrice() int64 {
	if m != nil {
		return m.SellCleanPrice
	}
	return 0
}

func (m *MDFIQuote) GetSellVolume() int64 {
	if m != nil {
		return m.SellVolume
	}
	return 0
}

func (m *MDFIQuote) GetSellDirtyPrice() int64 {
	if m != nil {
		return m.SellDirtyPrice
	}
	return 0
}

func (m *MDFIQuote) GetSellMaturityYield() int64 {
	if m != nil {
		return m.SellMaturityYield
	}
	return 0
}

func (m *MDFIQuote) GetSellQuoteComment() string {
	if m != nil {
		return m.SellQuoteComment
	}
	return ""
}

func (m *MDFIQuote) GetSellQuoteType() int32 {
	if m != nil {
		return m.SellQuoteType
	}
	return 0
}

func (m *MDFIQuote) GetSellBargainType() int32 {
	if m != nil {
		return m.SellBargainType
	}
	return 0
}

func (m *MDFIQuote) GetSellRelationType() int32 {
	if m != nil {
		return m.SellRelationType
	}
	return 0
}

func (m *MDFIQuote) GetSellExerciseType() int32 {
	if m != nil {
		return m.SellExerciseType
	}
	return 0
}

func (m *MDFIQuote) GetSellYieldType() int32 {
	if m != nil {
		return m.SellYieldType
	}
	return 0
}

func init() {
	proto.RegisterType((*MDFIQuote)(nil), "com.htsc.mdc.insight.model.MDFIQuote")
}

func init() { proto.RegisterFile("MDFIQuote.proto", fileDescriptor_2ab23b79765a8fc8) }

var fileDescriptor_2ab23b79765a8fc8 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xdf, 0x4e, 0xd4, 0x40,
	0x18, 0xc5, 0x53, 0x11, 0x74, 0x07, 0x04, 0x1c, 0x15, 0xc6, 0x15, 0xb1, 0x20, 0x60, 0x45, 0xec,
	0x05, 0x3e, 0x81, 0xdd, 0xae, 0x61, 0x2f, 0xd6, 0x40, 0x77, 0x63, 0xe2, 0x65, 0xe9, 0x4e, 0x76,
	0x1b, 0xfb, 0x67, 0x33, 0x9d, 0x26, 0xf4, 0x6d, 0x7c, 0x1d, 0xdf, 0xca, 0xcc, 0xd7, 0x99, 0x76,
	0xa6, 0x4d, 0xb8, 0xdb, 0xf9, 0xf5, 0xf4, 0xf4, 0xfb, 0xce, 0x69, 0xb6, 0x68, 0x6f, 0xea, 0xff,
	0x98, 0xdc, 0x95, 0x39, 0xa7, 0xee, 0x9a, 0xe5, 0x3c, 0xc7, 0xc3, 0x28, 0x4f, 0xdd, 0x15, 0x2f,
	0x22, 0x37, 0x5d, 0x44, 0x6e, 0x9c, 0x15, 0xf1, 0x72, 0xc5, 0xdd, 0x34, 0x5f, 0xd0, 0x64, 0x78,
	0x38, 0x9e, 0xd1, 0xa8, 0x64, 0x31, 0xaf, 0x26, 0xfe, 0x2c, 0x2f, 0x59, 0x24, 0x6f, 0x1a, 0xbe,
	0x6a, 0x2e, 0xcc, 0xab, 0xb5, 0x84, 0xa7, 0xff, 0xb6, 0xd1, 0xa0, 0x71, 0xc7, 0x17, 0x68, 0xf7,
	0x66, 0x3e, 0x1b, 0xb5, 0x06, 0xc4, 0xb2, 0x2d, 0x67, 0x10, 0x74, 0x28, 0x3e, 0x40, 0x5b, 0x53,
	0xdf, 0x0f, 0x39, 0x25, 0x4f, 0x6c, 0xcb, 0xd9, 0x0c, 0xe4, 0xa9, 0xe6, 0xf3, 0x38, 0xa5, 0x64,
	0x43, 0x71, 0x71, 0xc2, 0x67, 0xe8, 0x85, 0x1f, 0xf2, 0x50, 0xfc, 0x2e, 0x78, 0x98, 0xae, 0xc9,
	0x53, 0xdb, 0x72, 0x36, 0x02, 0x13, 0xe2, 0x3b, 0xb4, 0x5f, 0x74, 0x46, 0x27, 0x9b, 0xb6, 0xe5,
	0xec, 0x5e, 0x9f, 0xbb, 0xc6, 0xc2, 0xb0, 0xa8, 0xdb, 0xdb, 0x33, 0xe8, 0xdd, 0x8e, 0xc7, 0x68,
	0xa7, 0xd0, 0x96, 0x26, 0x5b, 0x60, 0x77, 0xf2, 0xa8, 0x9d, 0x10, 0x06, 0xc6, 0x6d, 0xf8, 0x14,
	0xed, 0x8c, 0x1f, 0xa2, 0x55, 0x98, 0x2d, 0x29, 0x6c, 0xfd, 0x0c, 0xb6, 0x33, 0x98, 0xae, 0x81,
	0x04, 0x9e, 0x9b, 0x1a, 0x95, 0xc3, 0x94, 0x16, 0x45, 0xb8, 0xa4, 0x3f, 0xcb, 0xf4, 0x9e, 0x32,
	0x32, 0xa8, 0x73, 0x30, 0xa0, 0x68, 0xc1, 0x63, 0xf9, 0x1f, 0xca, 0x20, 0x1e, 0x31, 0x36, 0x02,
	0xaf, 0x0e, 0xc5, 0x0e, 0xda, 0xfb, 0x1e, 0x45, 0xac, 0xa4, 0x8b, 0x49, 0xc6, 0x29, 0xa3, 0x05,
	0x27, 0xdb, 0xe0, 0xd7, 0xc5, 0xf8, 0x18, 0x21, 0xaf, 0xac, 0xa0, 0xe3, 0x89, 0x4f, 0x5e, 0x43,
	0xa7, 0x1a, 0x11, 0xb3, 0xab, 0x13, 0xcc, 0xfe, 0xa6, 0x9e, 0x5d, 0x67, 0xf8, 0x08, 0x0d, 0xd4,
	0x99, 0x91, 0x03, 0xb0, 0x68, 0x81, 0xd8, 0xcc, 0x2b, 0xab, 0x51, 0x42, 0xc3, 0xec, 0x96, 0xc5,
	0x11, 0x25, 0x87, 0xf5, 0x66, 0x06, 0x94, 0x1e, 0xbf, 0xf2, 0xa4, 0x4c, 0x29, 0x21, 0xa0, 0x68,
	0x81, 0xf4, 0xf0, 0x63, 0xc6, 0xab, 0xda, 0xe3, 0x6d, 0xe3, 0xd1, 0x42, 0x7c, 0x89, 0xf6, 0xbd,
	0xb2, 0x9a, 0x86, 0x1c, 0xea, 0xf9, 0x1d, 0xd3, 0x64, 0x41, 0x86, 0x20, 0xec, 0x71, 0x91, 0x90,
	0x1a, 0x71, 0x94, 0xa7, 0x29, 0xcd, 0x38, 0x79, 0x07, 0x93, 0x77, 0xb1, 0x91, 0x80, 0x48, 0xfc,
	0xa8, 0x93, 0x80, 0xc8, 0x5b, 0xf4, 0x52, 0x56, 0x5e, 0xc8, 0x96, 0x61, 0x9c, 0x81, 0xea, 0xbd,
	0xec, 0xc5, 0xa0, 0xf2, 0xa9, 0x01, 0x4d, 0x42, 0x1e, 0xe7, 0xb5, 0xf0, 0x18, 0x84, 0x5d, 0x2c,
	0x95, 0xe3, 0x07, 0xca, 0xa2, 0xb8, 0xa8, 0x1f, 0xfc, 0xa1, 0x51, 0xea, 0x58, 0xce, 0x07, 0x5b,
	0x81, 0xcc, 0x6e, 0xe6, 0x6b, 0x18, 0xb6, 0xd1, 0xf6, 0x8c, 0x26, 0x89, 0xaa, 0xf9, 0x04, 0x36,
	0xd5, 0x91, 0x48, 0xb8, 0x39, 0x42, 0xd1, 0xa7, 0x60, 0x63, 0x42, 0xf1, 0xb6, 0x34, 0x80, 0x91,
	0x8f, 0xf5, 0xdb, 0xd2, 0x12, 0x91, 0x83, 0x38, 0x69, 0x65, 0x9f, 0x41, 0xfe, 0x1d, 0xaa, 0x7c,
	0x64, 0xdd, 0xe7, 0xa0, 0xd1, 0x88, 0xf2, 0xd1, 0x0a, 0xbf, 0x68, 0x7d, 0xb4, 0xc6, 0xaf, 0xd0,
	0x4b, 0x41, 0xcc, 0xca, 0x3f, 0x81, 0xb4, 0x7f, 0x41, 0xbc, 0x1f, 0xcd, 0xac, 0xaa, 0x74, 0x07,
	0x76, 0xe8, 0x71, 0x33, 0x0f, 0x11, 0xeb, 0xe7, 0x6e, 0x1e, 0xb2, 0x25, 0x01, 0xf4, 0xe2, 0x2f,
	0xeb, 0x96, 0x3a, 0x58, 0x3d, 0xdb, 0xa8, 0xfe, 0x0b, 0x48, 0x7b, 0x5c, 0x69, 0x8d, 0xf2, 0xaf,
	0x5a, 0xad, 0xd1, 0xbe, 0x9c, 0xb3, 0xad, 0xff, 0x6b, 0x3b, 0x67, 0x03, 0xbd, 0x6b, 0xf4, 0xc8,
	0x77, 0xc1, 0x6b, 0x3f, 0x22, 0xb7, 0xe2, 0x9f, 0xbf, 0xb8, 0xb1, 0xfe, 0x5a, 0xd6, 0xfd, 0x16,
	0x7c, 0x06, 0xbe, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xac, 0x24, 0x5a, 0xec, 0x63, 0x06, 0x00,
	0x00,
}
