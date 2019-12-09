// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDSimpleTick.proto

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

//证券简化Tick行情(MDSimpleTick)
type MDSimpleTick struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	TradingPhaseCode     string            `protobuf:"bytes,5,opt,name=TradingPhaseCode,proto3" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,6,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,7,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	NumTrades            int64             `protobuf:"varint,8,opt,name=NumTrades,proto3" json:"NumTrades,omitempty"`
	TotalVolumeTrade     int64             `protobuf:"varint,9,opt,name=TotalVolumeTrade,proto3" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade      int64             `protobuf:"varint,10,opt,name=TotalValueTrade,proto3" json:"TotalValueTrade,omitempty"`
	PreClosePx           int64             `protobuf:"varint,11,opt,name=PreClosePx,proto3" json:"PreClosePx,omitempty"`
	LastPx               int64             `protobuf:"varint,12,opt,name=LastPx,proto3" json:"LastPx,omitempty"`
	OpenPx               int64             `protobuf:"varint,13,opt,name=OpenPx,proto3" json:"OpenPx,omitempty"`
	ClosePx              int64             `protobuf:"varint,14,opt,name=ClosePx,proto3" json:"ClosePx,omitempty"`
	HighPx               int64             `protobuf:"varint,15,opt,name=HighPx,proto3" json:"HighPx,omitempty"`
	LowPx                int64             `protobuf:"varint,16,opt,name=LowPx,proto3" json:"LowPx,omitempty"`
	IOPV                 int64             `protobuf:"varint,17,opt,name=IOPV,proto3" json:"IOPV,omitempty"`
	PreIOPV              int64             `protobuf:"varint,18,opt,name=PreIOPV,proto3" json:"PreIOPV,omitempty"`
	OpenInterest         int64             `protobuf:"varint,19,opt,name=OpenInterest,proto3" json:"OpenInterest,omitempty"`
	PreOpenInterest      int64             `protobuf:"varint,20,opt,name=PreOpenInterest,proto3" json:"PreOpenInterest,omitempty"`
	SettlePrice          int64             `protobuf:"varint,21,opt,name=SettlePrice,proto3" json:"SettlePrice,omitempty"`
	PreSettlePrice       int64             `protobuf:"varint,22,opt,name=PreSettlePrice,proto3" json:"PreSettlePrice,omitempty"`
	ExchangeDate         int32             `protobuf:"varint,23,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32             `protobuf:"varint,24,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	ADIndicators         *ADIndicators     `protobuf:"bytes,40,opt,name=ADIndicators,proto3" json:"ADIndicators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MDSimpleTick) Reset()         { *m = MDSimpleTick{} }
func (m *MDSimpleTick) String() string { return proto.CompactTextString(m) }
func (*MDSimpleTick) ProtoMessage()    {}
func (*MDSimpleTick) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd1b7320875a227, []int{0}
}

func (m *MDSimpleTick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDSimpleTick.Unmarshal(m, b)
}
func (m *MDSimpleTick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDSimpleTick.Marshal(b, m, deterministic)
}
func (m *MDSimpleTick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDSimpleTick.Merge(m, src)
}
func (m *MDSimpleTick) XXX_Size() int {
	return xxx_messageInfo_MDSimpleTick.Size(m)
}
func (m *MDSimpleTick) XXX_DiscardUnknown() {
	xxx_messageInfo_MDSimpleTick.DiscardUnknown(m)
}

var xxx_messageInfo_MDSimpleTick proto.InternalMessageInfo

func (m *MDSimpleTick) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDSimpleTick) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDSimpleTick) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDSimpleTick) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDSimpleTick) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDSimpleTick) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDSimpleTick) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDSimpleTick) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func (m *MDSimpleTick) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *MDSimpleTick) GetTotalValueTrade() int64 {
	if m != nil {
		return m.TotalValueTrade
	}
	return 0
}

func (m *MDSimpleTick) GetPreClosePx() int64 {
	if m != nil {
		return m.PreClosePx
	}
	return 0
}

func (m *MDSimpleTick) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *MDSimpleTick) GetOpenPx() int64 {
	if m != nil {
		return m.OpenPx
	}
	return 0
}

func (m *MDSimpleTick) GetClosePx() int64 {
	if m != nil {
		return m.ClosePx
	}
	return 0
}

func (m *MDSimpleTick) GetHighPx() int64 {
	if m != nil {
		return m.HighPx
	}
	return 0
}

func (m *MDSimpleTick) GetLowPx() int64 {
	if m != nil {
		return m.LowPx
	}
	return 0
}

func (m *MDSimpleTick) GetIOPV() int64 {
	if m != nil {
		return m.IOPV
	}
	return 0
}

func (m *MDSimpleTick) GetPreIOPV() int64 {
	if m != nil {
		return m.PreIOPV
	}
	return 0
}

func (m *MDSimpleTick) GetOpenInterest() int64 {
	if m != nil {
		return m.OpenInterest
	}
	return 0
}

func (m *MDSimpleTick) GetPreOpenInterest() int64 {
	if m != nil {
		return m.PreOpenInterest
	}
	return 0
}

func (m *MDSimpleTick) GetSettlePrice() int64 {
	if m != nil {
		return m.SettlePrice
	}
	return 0
}

func (m *MDSimpleTick) GetPreSettlePrice() int64 {
	if m != nil {
		return m.PreSettlePrice
	}
	return 0
}

func (m *MDSimpleTick) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *MDSimpleTick) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func (m *MDSimpleTick) GetADIndicators() *ADIndicators {
	if m != nil {
		return m.ADIndicators
	}
	return nil
}

//行情指标
type ADIndicators struct {
	Ind1101              int64    `protobuf:"varint,1,opt,name=Ind1101,proto3" json:"Ind1101,omitempty"`
	Ind1102              int64    `protobuf:"varint,2,opt,name=Ind1102,proto3" json:"Ind1102,omitempty"`
	Ind1103              int64    `protobuf:"varint,3,opt,name=Ind1103,proto3" json:"Ind1103,omitempty"`
	Ind1104              int64    `protobuf:"varint,4,opt,name=Ind1104,proto3" json:"Ind1104,omitempty"`
	Ind1105              int64    `protobuf:"varint,5,opt,name=Ind1105,proto3" json:"Ind1105,omitempty"`
	Ind1106              int64    `protobuf:"varint,6,opt,name=Ind1106,proto3" json:"Ind1106,omitempty"`
	Ind1107              int64    `protobuf:"varint,7,opt,name=Ind1107,proto3" json:"Ind1107,omitempty"`
	Ind1108              int64    `protobuf:"varint,8,opt,name=Ind1108,proto3" json:"Ind1108,omitempty"`
	Ind1109              int64    `protobuf:"varint,9,opt,name=Ind1109,proto3" json:"Ind1109,omitempty"`
	Ind1110              int64    `protobuf:"varint,10,opt,name=Ind1110,proto3" json:"Ind1110,omitempty"`
	Ind1111              int64    `protobuf:"varint,11,opt,name=Ind1111,proto3" json:"Ind1111,omitempty"`
	Ind1112              int64    `protobuf:"varint,12,opt,name=Ind1112,proto3" json:"Ind1112,omitempty"`
	Ind1113              int64    `protobuf:"varint,13,opt,name=Ind1113,proto3" json:"Ind1113,omitempty"`
	Ind1114              int64    `protobuf:"varint,14,opt,name=Ind1114,proto3" json:"Ind1114,omitempty"`
	Ind1115              int64    `protobuf:"varint,15,opt,name=Ind1115,proto3" json:"Ind1115,omitempty"`
	Ind1116              int64    `protobuf:"varint,16,opt,name=Ind1116,proto3" json:"Ind1116,omitempty"`
	Ind1117              int64    `protobuf:"varint,17,opt,name=Ind1117,proto3" json:"Ind1117,omitempty"`
	Ind1118              int64    `protobuf:"varint,18,opt,name=Ind1118,proto3" json:"Ind1118,omitempty"`
	Ind1119              int64    `protobuf:"varint,19,opt,name=Ind1119,proto3" json:"Ind1119,omitempty"`
	Ind1120              int64    `protobuf:"varint,20,opt,name=Ind1120,proto3" json:"Ind1120,omitempty"`
	Ind1121              int64    `protobuf:"varint,21,opt,name=Ind1121,proto3" json:"Ind1121,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ADIndicators) Reset()         { *m = ADIndicators{} }
func (m *ADIndicators) String() string { return proto.CompactTextString(m) }
func (*ADIndicators) ProtoMessage()    {}
func (*ADIndicators) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd1b7320875a227, []int{1}
}

func (m *ADIndicators) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADIndicators.Unmarshal(m, b)
}
func (m *ADIndicators) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADIndicators.Marshal(b, m, deterministic)
}
func (m *ADIndicators) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADIndicators.Merge(m, src)
}
func (m *ADIndicators) XXX_Size() int {
	return xxx_messageInfo_ADIndicators.Size(m)
}
func (m *ADIndicators) XXX_DiscardUnknown() {
	xxx_messageInfo_ADIndicators.DiscardUnknown(m)
}

var xxx_messageInfo_ADIndicators proto.InternalMessageInfo

func (m *ADIndicators) GetInd1101() int64 {
	if m != nil {
		return m.Ind1101
	}
	return 0
}

func (m *ADIndicators) GetInd1102() int64 {
	if m != nil {
		return m.Ind1102
	}
	return 0
}

func (m *ADIndicators) GetInd1103() int64 {
	if m != nil {
		return m.Ind1103
	}
	return 0
}

func (m *ADIndicators) GetInd1104() int64 {
	if m != nil {
		return m.Ind1104
	}
	return 0
}

func (m *ADIndicators) GetInd1105() int64 {
	if m != nil {
		return m.Ind1105
	}
	return 0
}

func (m *ADIndicators) GetInd1106() int64 {
	if m != nil {
		return m.Ind1106
	}
	return 0
}

func (m *ADIndicators) GetInd1107() int64 {
	if m != nil {
		return m.Ind1107
	}
	return 0
}

func (m *ADIndicators) GetInd1108() int64 {
	if m != nil {
		return m.Ind1108
	}
	return 0
}

func (m *ADIndicators) GetInd1109() int64 {
	if m != nil {
		return m.Ind1109
	}
	return 0
}

func (m *ADIndicators) GetInd1110() int64 {
	if m != nil {
		return m.Ind1110
	}
	return 0
}

func (m *ADIndicators) GetInd1111() int64 {
	if m != nil {
		return m.Ind1111
	}
	return 0
}

func (m *ADIndicators) GetInd1112() int64 {
	if m != nil {
		return m.Ind1112
	}
	return 0
}

func (m *ADIndicators) GetInd1113() int64 {
	if m != nil {
		return m.Ind1113
	}
	return 0
}

func (m *ADIndicators) GetInd1114() int64 {
	if m != nil {
		return m.Ind1114
	}
	return 0
}

func (m *ADIndicators) GetInd1115() int64 {
	if m != nil {
		return m.Ind1115
	}
	return 0
}

func (m *ADIndicators) GetInd1116() int64 {
	if m != nil {
		return m.Ind1116
	}
	return 0
}

func (m *ADIndicators) GetInd1117() int64 {
	if m != nil {
		return m.Ind1117
	}
	return 0
}

func (m *ADIndicators) GetInd1118() int64 {
	if m != nil {
		return m.Ind1118
	}
	return 0
}

func (m *ADIndicators) GetInd1119() int64 {
	if m != nil {
		return m.Ind1119
	}
	return 0
}

func (m *ADIndicators) GetInd1120() int64 {
	if m != nil {
		return m.Ind1120
	}
	return 0
}

func (m *ADIndicators) GetInd1121() int64 {
	if m != nil {
		return m.Ind1121
	}
	return 0
}

func init() {
	proto.RegisterType((*MDSimpleTick)(nil), "com.htsc.mdc.insight.model.MDSimpleTick")
	proto.RegisterType((*ADIndicators)(nil), "com.htsc.mdc.insight.model.ADIndicators")
}

func init() { proto.RegisterFile("MDSimpleTick.proto", fileDescriptor_2cd1b7320875a227) }

var fileDescriptor_2cd1b7320875a227 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0x5b, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x15, 0x76, 0x29, 0xf3, 0xba, 0x0b, 0xde, 0xd8, 0xac, 0x09, 0xa1, 0x32, 0x01, 0x8a,
	0x78, 0x88, 0xd6, 0x76, 0xdd, 0xe5, 0x91, 0x2d, 0x93, 0x56, 0xa9, 0x63, 0x26, 0xad, 0xf6, 0x1e,
	0x12, 0xab, 0x8d, 0xc8, 0xa5, 0x8a, 0x5d, 0x91, 0x7d, 0x1b, 0x3e, 0x02, 0xcf, 0x7c, 0x3a, 0xe4,
	0x7f, 0xec, 0x38, 0x69, 0xc5, 0xde, 0x7a, 0xce, 0xef, 0xd8, 0x49, 0xfe, 0xf5, 0x91, 0x11, 0x7e,
	0x70, 0xc7, 0x51, 0x32, 0x8f, 0xd9, 0x24, 0x0a, 0x7e, 0x3a, 0xf3, 0x3c, 0x13, 0x19, 0x3e, 0x09,
	0xb2, 0xc4, 0x99, 0x09, 0x1e, 0x38, 0x49, 0x18, 0x38, 0x51, 0xca, 0xa3, 0xe9, 0x4c, 0x38, 0x49,
	0x16, 0xb2, 0xf8, 0xe4, 0xe0, 0x6e, 0xcc, 0x82, 0x45, 0x1e, 0x89, 0xe7, 0xc9, 0xf3, 0x9c, 0x95,
	0x0b, 0x4e, 0x8e, 0x2b, 0x73, 0xe8, 0x8e, 0xb3, 0x45, 0x1e, 0x28, 0x70, 0xfa, 0xb7, 0x85, 0xda,
	0xf5, 0x07, 0xe0, 0xcf, 0x68, 0xf7, 0x7e, 0x32, 0xbe, 0x35, 0x71, 0x62, 0x75, 0x2c, 0x7b, 0xcb,
	0x5b, 0x72, 0xf1, 0x11, 0xda, 0x7c, 0x70, 0x5d, 0x5f, 0x30, 0xf2, 0xaa, 0x63, 0xd9, 0x1b, 0x9e,
	0x52, 0xa5, 0x3f, 0x89, 0x12, 0x46, 0xd6, 0xb4, 0x2f, 0x15, 0xfe, 0x88, 0x76, 0x5c, 0x5f, 0xf8,
	0xf2, 0x37, 0x17, 0x7e, 0x32, 0x27, 0xeb, 0x1d, 0xcb, 0x5e, 0xf3, 0x9a, 0x26, 0xfe, 0x82, 0xf6,
	0x27, 0xb9, 0x1f, 0x46, 0xe9, 0x94, 0xce, 0x7c, 0xce, 0x6e, 0xb3, 0x90, 0x91, 0x0d, 0x78, 0xfe,
	0x8a, 0x8f, 0xbf, 0xa3, 0x7d, 0xbe, 0xf4, 0x51, 0x64, 0xb3, 0x63, 0xd9, 0xbb, 0xbd, 0x4f, 0x4e,
	0x63, 0x3e, 0x30, 0x17, 0x67, 0x65, 0x02, 0xde, 0xca, 0x72, 0x7c, 0x87, 0xda, 0xbc, 0x36, 0x3c,
	0xd2, 0x82, 0xed, 0x3e, 0xbc, 0xb8, 0x9d, 0x0c, 0x7a, 0x8d, 0x65, 0xf8, 0x1d, 0xda, 0xfa, 0xb6,
	0x48, 0xe4, 0x0b, 0x33, 0x4e, 0x5e, 0xc3, 0x77, 0x1a, 0x03, 0xbe, 0x31, 0x13, 0x7e, 0xfc, 0x94,
	0xc5, 0x8b, 0x84, 0x81, 0x49, 0xb6, 0x20, 0xb4, 0xe2, 0x63, 0x1b, 0xed, 0x95, 0x9e, 0x1f, 0x2f,
	0x54, 0x14, 0x41, 0x74, 0xd9, 0xc6, 0xef, 0x11, 0xa2, 0x39, 0xbb, 0x8d, 0x33, 0xce, 0x68, 0x41,
	0xb6, 0x21, 0x54, 0x73, 0xe4, 0xff, 0x32, 0xf2, 0xb9, 0xa0, 0x05, 0x69, 0x03, 0x53, 0x4a, 0xfa,
	0x8f, 0x73, 0x96, 0xd2, 0x82, 0xec, 0x94, 0x7e, 0xa9, 0x30, 0x41, 0x2d, 0xbd, 0xd9, 0x2e, 0x80,
	0x56, 0x6d, 0xa7, 0xfb, 0x68, 0x3a, 0xa3, 0x05, 0xd9, 0x2b, 0x57, 0x94, 0x0a, 0x1f, 0xa2, 0x8d,
	0x51, 0xf6, 0x8b, 0x16, 0x64, 0x1f, 0xec, 0x52, 0x60, 0x8c, 0xd6, 0x87, 0x8f, 0xf4, 0x89, 0xbc,
	0x01, 0x13, 0x7e, 0xcb, 0xbd, 0x69, 0xce, 0xc0, 0xc6, 0xe5, 0xde, 0x4a, 0xe2, 0x53, 0xd4, 0x96,
	0xcf, 0x1f, 0xa6, 0x82, 0xe5, 0x8c, 0x0b, 0x72, 0x00, 0xb8, 0xe1, 0xc9, 0x99, 0xd0, 0x9c, 0x35,
	0x62, 0x87, 0xe5, 0x4c, 0x96, 0x6c, 0xdc, 0x41, 0xdb, 0x63, 0x26, 0x44, 0xcc, 0x68, 0x1e, 0x05,
	0x8c, 0xbc, 0x85, 0x54, 0xdd, 0x92, 0xa7, 0x9d, 0xe6, 0xac, 0x1e, 0x3a, 0x82, 0xd0, 0x92, 0x2b,
	0xdf, 0xeb, 0xae, 0x08, 0x66, 0x7e, 0x3a, 0x65, 0x70, 0xe6, 0x8f, 0xe1, 0x6c, 0x37, 0xbc, 0x7a,
	0x06, 0xce, 0x3f, 0x69, 0x66, 0xa0, 0x05, 0x23, 0xd4, 0xfe, 0xea, 0x0e, 0xd3, 0x30, 0x0a, 0x7c,
	0x91, 0xe5, 0x9c, 0xd8, 0x1d, 0xcb, 0xde, 0xee, 0xd9, 0xce, 0xff, 0xfb, 0xec, 0xd4, 0xf3, 0x5e,
	0x63, 0xf5, 0xe9, 0x9f, 0xf5, 0xe6, 0x76, 0x72, 0xb0, 0xc3, 0x34, 0xec, 0x76, 0xcf, 0xba, 0xd0,
	0xda, 0x35, 0x4f, 0x4b, 0x43, 0x7a, 0xd0, 0xd7, 0x8a, 0xf4, 0x0c, 0xe9, 0x43, 0x63, 0x2b, 0xd2,
	0x37, 0xe4, 0x5c, 0x95, 0x55, 0x4b, 0x43, 0x06, 0xd0, 0xce, 0x8a, 0x0c, 0x0c, 0xb9, 0x80, 0x2e,
	0x56, 0xe4, 0xc2, 0x90, 0x4b, 0xa8, 0x55, 0x45, 0x2e, 0x0d, 0xb9, 0x52, 0x65, 0xd1, 0xd2, 0x90,
	0x6b, 0xd5, 0x10, 0x2d, 0x2b, 0xd2, 0x3d, 0x53, 0x85, 0xd0, 0xd2, 0x90, 0xae, 0x6a, 0x81, 0x96,
	0x86, 0xf4, 0x54, 0x07, 0xb4, 0x34, 0xa4, 0xaf, 0x5a, 0xa0, 0xa5, 0x21, 0xe7, 0xba, 0x06, 0x4a,
	0x1a, 0x32, 0x50, 0x3d, 0xd0, 0xd2, 0x90, 0x0b, 0x55, 0x05, 0x2d, 0x0d, 0xb9, 0x54, 0x7d, 0xd0,
	0xd2, 0x90, 0x2b, 0x5d, 0x09, 0x25, 0x0d, 0xb9, 0x56, 0x6d, 0xd0, 0xb2, 0x22, 0xbd, 0x33, 0x55,
	0x00, 0x2d, 0x0d, 0xe9, 0xaa, 0x43, 0xaf, 0xe5, 0xcd, 0x00, 0xbd, 0x70, 0x77, 0xdc, 0x34, 0xee,
	0x1a, 0x2a, 0x2f, 0x08, 0x7e, 0x6f, 0xfd, 0xb6, 0xac, 0x1f, 0x9b, 0x70, 0x5b, 0xf4, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x5d, 0xbd, 0x72, 0xfe, 0x8d, 0x06, 0x00, 0x00,
}