// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactionData.proto

package comm

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

type TransactionData struct {
	HTSCSecurityID       *string  `protobuf:"bytes,1,req,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	MDDate               *int32   `protobuf:"varint,2,req,name=MDDate" json:"MDDate,omitempty"`
	MDTime               *int32   `protobuf:"varint,3,req,name=MDTime" json:"MDTime,omitempty"`
	SecurityIDSource     *int64   `protobuf:"varint,4,opt,name=securityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         *int64   `protobuf:"varint,5,opt,name=securityType" json:"securityType,omitempty"`
	TradeIndex           *int64   `protobuf:"varint,6,opt,name=TradeIndex" json:"TradeIndex,omitempty"`
	TradeBuyNo           *int64   `protobuf:"varint,7,opt,name=TradeBuyNo" json:"TradeBuyNo,omitempty"`
	TradeSellNo          *int64   `protobuf:"varint,8,opt,name=TradeSellNo" json:"TradeSellNo,omitempty"`
	TradeType            *int32   `protobuf:"varint,9,opt,name=TradeType" json:"TradeType,omitempty"`
	TradeBSFlag          *int32   `protobuf:"varint,10,opt,name=TradeBSFlag" json:"TradeBSFlag,omitempty"`
	TradePrice           *int64   `protobuf:"varint,11,opt,name=TradePrice" json:"TradePrice,omitempty"`
	TradeQty             *int64   `protobuf:"varint,12,opt,name=TradeQty" json:"TradeQty,omitempty"`
	TradeMoney           *int64   `protobuf:"varint,13,opt,name=TradeMoney" json:"TradeMoney,omitempty"`
	ChannelNo            *int32   `protobuf:"varint,14,opt,name=ChannelNo" json:"ChannelNo,omitempty"`
	ExchangeDate         *int32   `protobuf:"varint,15,opt,name=ExchangeDate" json:"ExchangeDate,omitempty"`
	ExchangeTime         *int32   `protobuf:"varint,16,opt,name=ExchangeTime" json:"ExchangeTime,omitempty"`
	TradeCleanPrice      *int64   `protobuf:"varint,17,opt,name=TradeCleanPrice" json:"TradeCleanPrice,omitempty"`
	AccruedInterestAmt   *int64   `protobuf:"varint,18,opt,name=AccruedInterestAmt" json:"AccruedInterestAmt,omitempty"`
	TradeDirtyPrice      *int64   `protobuf:"varint,19,opt,name=TradeDirtyPrice" json:"TradeDirtyPrice,omitempty"`
	MaturityYield        *int64   `protobuf:"varint,20,opt,name=MaturityYield" json:"MaturityYield,omitempty"`
	FITradingMethod      *string  `protobuf:"bytes,21,opt,name=FITradingMethod" json:"FITradingMethod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionData) Reset()         { *m = TransactionData{} }
func (m *TransactionData) String() string { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()    {}
func (*TransactionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d1d2ae531cdc1c, []int{0}
}

func (m *TransactionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionData.Unmarshal(m, b)
}
func (m *TransactionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionData.Marshal(b, m, deterministic)
}
func (m *TransactionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionData.Merge(m, src)
}
func (m *TransactionData) XXX_Size() int {
	return xxx_messageInfo_TransactionData.Size(m)
}
func (m *TransactionData) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionData.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionData proto.InternalMessageInfo

func (m *TransactionData) GetHTSCSecurityID() string {
	if m != nil && m.HTSCSecurityID != nil {
		return *m.HTSCSecurityID
	}
	return ""
}

func (m *TransactionData) GetMDDate() int32 {
	if m != nil && m.MDDate != nil {
		return *m.MDDate
	}
	return 0
}

func (m *TransactionData) GetMDTime() int32 {
	if m != nil && m.MDTime != nil {
		return *m.MDTime
	}
	return 0
}

func (m *TransactionData) GetSecurityIDSource() int64 {
	if m != nil && m.SecurityIDSource != nil {
		return *m.SecurityIDSource
	}
	return 0
}

func (m *TransactionData) GetSecurityType() int64 {
	if m != nil && m.SecurityType != nil {
		return *m.SecurityType
	}
	return 0
}

func (m *TransactionData) GetTradeIndex() int64 {
	if m != nil && m.TradeIndex != nil {
		return *m.TradeIndex
	}
	return 0
}

func (m *TransactionData) GetTradeBuyNo() int64 {
	if m != nil && m.TradeBuyNo != nil {
		return *m.TradeBuyNo
	}
	return 0
}

func (m *TransactionData) GetTradeSellNo() int64 {
	if m != nil && m.TradeSellNo != nil {
		return *m.TradeSellNo
	}
	return 0
}

func (m *TransactionData) GetTradeType() int32 {
	if m != nil && m.TradeType != nil {
		return *m.TradeType
	}
	return 0
}

func (m *TransactionData) GetTradeBSFlag() int32 {
	if m != nil && m.TradeBSFlag != nil {
		return *m.TradeBSFlag
	}
	return 0
}

func (m *TransactionData) GetTradePrice() int64 {
	if m != nil && m.TradePrice != nil {
		return *m.TradePrice
	}
	return 0
}

func (m *TransactionData) GetTradeQty() int64 {
	if m != nil && m.TradeQty != nil {
		return *m.TradeQty
	}
	return 0
}

func (m *TransactionData) GetTradeMoney() int64 {
	if m != nil && m.TradeMoney != nil {
		return *m.TradeMoney
	}
	return 0
}

func (m *TransactionData) GetChannelNo() int32 {
	if m != nil && m.ChannelNo != nil {
		return *m.ChannelNo
	}
	return 0
}

func (m *TransactionData) GetExchangeDate() int32 {
	if m != nil && m.ExchangeDate != nil {
		return *m.ExchangeDate
	}
	return 0
}

func (m *TransactionData) GetExchangeTime() int32 {
	if m != nil && m.ExchangeTime != nil {
		return *m.ExchangeTime
	}
	return 0
}

func (m *TransactionData) GetTradeCleanPrice() int64 {
	if m != nil && m.TradeCleanPrice != nil {
		return *m.TradeCleanPrice
	}
	return 0
}

func (m *TransactionData) GetAccruedInterestAmt() int64 {
	if m != nil && m.AccruedInterestAmt != nil {
		return *m.AccruedInterestAmt
	}
	return 0
}

func (m *TransactionData) GetTradeDirtyPrice() int64 {
	if m != nil && m.TradeDirtyPrice != nil {
		return *m.TradeDirtyPrice
	}
	return 0
}

func (m *TransactionData) GetMaturityYield() int64 {
	if m != nil && m.MaturityYield != nil {
		return *m.MaturityYield
	}
	return 0
}

func (m *TransactionData) GetFITradingMethod() string {
	if m != nil && m.FITradingMethod != nil {
		return *m.FITradingMethod
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionData)(nil), "comm.TransactionData")
}

func init() { proto.RegisterFile("transactionData.proto", fileDescriptor_41d1d2ae531cdc1c) }

var fileDescriptor_41d1d2ae531cdc1c = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5f, 0x8b, 0x9b, 0x40,
	0x14, 0xc5, 0x31, 0xbb, 0x6e, 0xd7, 0xd9, 0x7f, 0xe9, 0xb4, 0x29, 0x43, 0x29, 0x45, 0x42, 0x29,
	0xd2, 0x87, 0x7c, 0x87, 0x24, 0x36, 0xd4, 0x07, 0x43, 0xab, 0xbe, 0xf4, 0x71, 0x18, 0x2f, 0x89,
	0xa0, 0x33, 0x61, 0x32, 0x42, 0xfc, 0xd8, 0xfd, 0x06, 0xc5, 0x6b, 0x33, 0x19, 0xd3, 0x7d, 0x3c,
	0xbf, 0x73, 0xee, 0xe5, 0xe8, 0x1d, 0x32, 0x33, 0x9a, 0xcb, 0x23, 0x17, 0xa6, 0x52, 0x32, 0xe6,
	0x86, 0x2f, 0x0e, 0x5a, 0x19, 0x45, 0x6f, 0x85, 0x6a, 0x9a, 0xf9, 0x1f, 0x9f, 0xbc, 0x14, 0x63,
	0x9f, 0x7e, 0x25, 0xcf, 0x3f, 0x8a, 0x7c, 0x9d, 0x83, 0x68, 0x75, 0x65, 0xba, 0x24, 0x66, 0x5e,
	0x38, 0x89, 0x82, 0xec, 0x8a, 0xd2, 0x0f, 0xe4, 0x2e, 0x8d, 0x63, 0x6e, 0x80, 0x4d, 0xc2, 0x49,
	0xe4, 0x67, 0xff, 0xd4, 0xc0, 0x8b, 0xaa, 0x01, 0x76, 0x73, 0xe6, 0xbd, 0xa2, 0xdf, 0xc8, 0xf4,
	0x68, 0xa7, 0x73, 0xd5, 0x6a, 0x01, 0xec, 0x36, 0xf4, 0xa2, 0x9b, 0xec, 0x3f, 0x4e, 0xe7, 0xe4,
	0xf1, 0xcc, 0x8a, 0xee, 0x00, 0xcc, 0xc7, 0xdc, 0x88, 0xd1, 0xcf, 0x84, 0x14, 0x9a, 0x97, 0x90,
	0xc8, 0x12, 0x4e, 0xec, 0x0e, 0x13, 0x0e, 0xb1, 0xfe, 0xaa, 0xed, 0xb6, 0x8a, 0xbd, 0x71, 0x7c,
	0x24, 0x34, 0x24, 0x0f, 0xa8, 0x72, 0xa8, 0xeb, 0xad, 0x62, 0xf7, 0x18, 0x70, 0x11, 0xfd, 0x44,
	0x02, 0x94, 0x58, 0x21, 0x08, 0xbd, 0xc8, 0xcf, 0x2e, 0xc0, 0xce, 0xaf, 0xf2, 0x4d, 0xcd, 0x77,
	0x8c, 0xa0, 0xef, 0x22, 0xdb, 0xe0, 0xa7, 0xae, 0x04, 0xb0, 0x07, 0xa7, 0x01, 0x12, 0xfa, 0x91,
	0xdc, 0xa3, 0xfa, 0x65, 0x3a, 0xf6, 0x88, 0xae, 0xd5, 0x76, 0x36, 0x55, 0x12, 0x3a, 0xf6, 0xe4,
	0xcc, 0x22, 0xe9, 0xbb, 0xad, 0xf7, 0x5c, 0x4a, 0xe8, 0xbb, 0x3f, 0x0f, 0xdd, 0x2c, 0xe8, 0xff,
	0xdf, 0xf7, 0x93, 0xd8, 0x73, 0xb9, 0x03, 0xbc, 0xd0, 0x0b, 0x06, 0x46, 0xcc, 0xcd, 0xe0, 0xb5,
	0xa6, 0xe3, 0x0c, 0xde, 0x2c, 0xc2, 0xe7, 0x51, 0xc2, 0xba, 0x06, 0x2e, 0x87, 0xcf, 0x78, 0x8b,
	0x55, 0xae, 0x31, 0x5d, 0x10, 0xba, 0x14, 0x42, 0xb7, 0x50, 0x26, 0xd2, 0x80, 0x86, 0xa3, 0x59,
	0x36, 0x86, 0x51, 0x0c, 0xbf, 0xe2, 0xd8, 0xcd, 0x71, 0xa5, 0x4d, 0x37, 0x6c, 0x7e, 0xe7, 0x6c,
	0xbe, 0x60, 0xfa, 0x85, 0x3c, 0xa5, 0xdc, 0xe0, 0xdd, 0x7f, 0x57, 0x50, 0x97, 0xec, 0x3d, 0xe6,
	0xc6, 0xb0, 0xdf, 0xb7, 0x49, 0xfa, 0xd1, 0x4a, 0xee, 0x52, 0x30, 0x7b, 0x55, 0xb2, 0x59, 0xe8,
	0x45, 0x41, 0x76, 0x8d, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xef, 0xcb, 0xd2, 0xfa, 0x11, 0x03,
	0x00, 0x00,
}
