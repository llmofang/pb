// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ADIndicatorsRanking.proto

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

//行情指标排行榜
type ADIndicatorsRanking struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,5,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,6,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	RankingType          int32             `protobuf:"varint,7,opt,name=RankingType,proto3" json:"RankingType,omitempty"`
	TypicalValue         *MDSimpleTick     `protobuf:"bytes,8,opt,name=TypicalValue,proto3" json:"TypicalValue,omitempty"`
	RankingList          []*MDSimpleTick   `protobuf:"bytes,9,rep,name=RankingList,proto3" json:"RankingList,omitempty"`
	ExchangeDate         int32             `protobuf:"varint,10,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime         int32             `protobuf:"varint,11,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ADIndicatorsRanking) Reset()         { *m = ADIndicatorsRanking{} }
func (m *ADIndicatorsRanking) String() string { return proto.CompactTextString(m) }
func (*ADIndicatorsRanking) ProtoMessage()    {}
func (*ADIndicatorsRanking) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b98b5ef9c4c72a8, []int{0}
}

func (m *ADIndicatorsRanking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADIndicatorsRanking.Unmarshal(m, b)
}
func (m *ADIndicatorsRanking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADIndicatorsRanking.Marshal(b, m, deterministic)
}
func (m *ADIndicatorsRanking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADIndicatorsRanking.Merge(m, src)
}
func (m *ADIndicatorsRanking) XXX_Size() int {
	return xxx_messageInfo_ADIndicatorsRanking.Size(m)
}
func (m *ADIndicatorsRanking) XXX_DiscardUnknown() {
	xxx_messageInfo_ADIndicatorsRanking.DiscardUnknown(m)
}

var xxx_messageInfo_ADIndicatorsRanking proto.InternalMessageInfo

func (m *ADIndicatorsRanking) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *ADIndicatorsRanking) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *ADIndicatorsRanking) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *ADIndicatorsRanking) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *ADIndicatorsRanking) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *ADIndicatorsRanking) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *ADIndicatorsRanking) GetRankingType() int32 {
	if m != nil {
		return m.RankingType
	}
	return 0
}

func (m *ADIndicatorsRanking) GetTypicalValue() *MDSimpleTick {
	if m != nil {
		return m.TypicalValue
	}
	return nil
}

func (m *ADIndicatorsRanking) GetRankingList() []*MDSimpleTick {
	if m != nil {
		return m.RankingList
	}
	return nil
}

func (m *ADIndicatorsRanking) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *ADIndicatorsRanking) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ADIndicatorsRanking)(nil), "com.htsc.mdc.insight.model.ADIndicatorsRanking")
}

func init() { proto.RegisterFile("ADIndicatorsRanking.proto", fileDescriptor_6b98b5ef9c4c72a8) }

var fileDescriptor_6b98b5ef9c4c72a8 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcb, 0x4e, 0x02, 0x31,
	0x14, 0x86, 0x53, 0xb9, 0x28, 0x05, 0x89, 0x29, 0x89, 0x56, 0x56, 0x95, 0xa8, 0x99, 0x55, 0x17,
	0xb8, 0x75, 0x23, 0x0e, 0x09, 0x18, 0x48, 0xb4, 0x33, 0x71, 0x5f, 0x4b, 0x03, 0x0d, 0x73, 0xcb,
	0xb4, 0x24, 0xf2, 0x36, 0xbe, 0x82, 0x6f, 0x68, 0xa6, 0x5c, 0x9c, 0x01, 0x24, 0x71, 0x37, 0xfd,
	0xe6, 0xfc, 0x7f, 0x7b, 0xce, 0xf9, 0xe1, 0xf5, 0x93, 0x3b, 0x8c, 0x26, 0x4a, 0x70, 0x13, 0xa7,
	0x9a, 0xf1, 0x68, 0xae, 0xa2, 0x29, 0x4d, 0xd2, 0xd8, 0xc4, 0xa8, 0x2d, 0xe2, 0x90, 0xce, 0x8c,
	0x16, 0x34, 0x9c, 0x08, 0xaa, 0x22, 0xad, 0xa6, 0x33, 0x43, 0xc3, 0x78, 0x22, 0x83, 0x76, 0xab,
	0xef, 0x49, 0xb1, 0x48, 0x95, 0x59, 0xfa, 0xcb, 0x44, 0xae, 0x04, 0xed, 0xab, 0x2d, 0x1c, 0xba,
	0x5e, 0xbc, 0x48, 0xc5, 0xe6, 0x07, 0x1a, 0xbb, 0x9e, 0x0a, 0x93, 0x40, 0xfa, 0x4a, 0xcc, 0x57,
	0xac, 0xf3, 0x5d, 0x86, 0xad, 0x03, 0x77, 0xa3, 0x7b, 0xd8, 0x1c, 0xf8, 0xde, 0xf3, 0xaf, 0x13,
	0x06, 0x04, 0x38, 0x35, 0xb6, 0x43, 0xd1, 0x25, 0xac, 0x8e, 0x5d, 0x97, 0x1b, 0x89, 0x4f, 0x08,
	0x70, 0x2a, 0x6c, 0x7d, 0x5a, 0x71, 0x5f, 0x85, 0x12, 0x97, 0x36, 0x3c, 0x3b, 0xa1, 0x5b, 0x78,
	0xee, 0x72, 0xc3, 0xb3, 0x6f, 0x6d, 0x78, 0x98, 0xe0, 0x32, 0x01, 0x4e, 0x89, 0x15, 0x21, 0x7a,
	0x83, 0x17, 0x7a, 0xa7, 0x07, 0x5c, 0x21, 0xc0, 0x69, 0x76, 0xef, 0x68, 0x61, 0x1c, 0x76, 0x0c,
	0x74, 0xaf, 0x61, 0xb6, 0x27, 0x47, 0x7d, 0xd8, 0xd0, 0xb9, 0x59, 0xe1, 0xaa, 0xb5, 0xbb, 0x39,
	0x6a, 0x97, 0x15, 0xb2, 0x82, 0x0c, 0x11, 0x58, 0x5f, 0x8f, 0xc8, 0xba, 0x9c, 0xda, 0xe6, 0xf2,
	0x08, 0x8d, 0x60, 0xc3, 0x5f, 0x26, 0x4a, 0xf0, 0xe0, 0x9d, 0x07, 0x0b, 0x89, 0xcf, 0x08, 0x70,
	0xea, 0x5d, 0x87, 0xfe, 0xbd, 0x46, 0x9a, 0xdf, 0x0b, 0x2b, 0xa8, 0xd1, 0xcb, 0xf6, 0xbe, 0x91,
	0xd2, 0x06, 0xd7, 0x48, 0xe9, 0x5f, 0x66, 0x79, 0x31, 0xea, 0xc0, 0x46, 0xff, 0x53, 0xcc, 0x78,
	0x34, 0x95, 0x76, 0x63, 0xd0, 0x3e, 0xbe, 0xc0, 0xf2, 0x35, 0x76, 0x7b, 0xf5, 0x62, 0x4d, 0xc6,
	0x7a, 0x8f, 0xf0, 0x48, 0x26, 0x7b, 0x87, 0xa2, 0xfc, 0x9a, 0x65, 0x4d, 0x0f, 0xc0, 0x17, 0x00,
	0x1f, 0x55, 0x1b, 0xbc, 0x87, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xe3, 0x4b, 0x27, 0xf3,
	0x02, 0x00, 0x00,
}
