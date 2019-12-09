// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ESecurityType.proto

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

//证券类型
type ESecurityType int32

const (
	ESecurityType_DefaultSecurityType ESecurityType = 0
	ESecurityType_IndexType           ESecurityType = 1
	ESecurityType_StockType           ESecurityType = 2
	ESecurityType_FundType            ESecurityType = 3
	ESecurityType_BondType            ESecurityType = 4
	ESecurityType_RepoType            ESecurityType = 5
	ESecurityType_WarrantType         ESecurityType = 6
	ESecurityType_OptionType          ESecurityType = 7
	ESecurityType_FuturesType         ESecurityType = 8
	ESecurityType_ForexType           ESecurityType = 9
	ESecurityType_RateType            ESecurityType = 10
	ESecurityType_NmetalType          ESecurityType = 11
	ESecurityType_CashBondType        ESecurityType = 12
	ESecurityType_SpotType            ESecurityType = 13
	ESecurityType_InsightType         ESecurityType = 20
	ESecurityType_OtherType           ESecurityType = 99
)

var ESecurityType_name = map[int32]string{
	0:  "DefaultSecurityType",
	1:  "IndexType",
	2:  "StockType",
	3:  "FundType",
	4:  "BondType",
	5:  "RepoType",
	6:  "WarrantType",
	7:  "OptionType",
	8:  "FuturesType",
	9:  "ForexType",
	10: "RateType",
	11: "NmetalType",
	12: "CashBondType",
	13: "SpotType",
	20: "InsightType",
	99: "OtherType",
}

var ESecurityType_value = map[string]int32{
	"DefaultSecurityType": 0,
	"IndexType":           1,
	"StockType":           2,
	"FundType":            3,
	"BondType":            4,
	"RepoType":            5,
	"WarrantType":         6,
	"OptionType":          7,
	"FuturesType":         8,
	"ForexType":           9,
	"RateType":            10,
	"NmetalType":          11,
	"CashBondType":        12,
	"SpotType":            13,
	"InsightType":         20,
	"OtherType":           99,
}

func (x ESecurityType) String() string {
	return proto.EnumName(ESecurityType_name, int32(x))
}

func (ESecurityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c7d06b47222b721, []int{0}
}

func init() {
	proto.RegisterEnum("com.htsc.mdc.model.ESecurityType", ESecurityType_name, ESecurityType_value)
}

func init() { proto.RegisterFile("ESecurityType.proto", fileDescriptor_3c7d06b47222b721) }

var fileDescriptor_3c7d06b47222b721 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xd1, 0x4a, 0xf3, 0x40,
	0x10, 0x05, 0xe0, 0x3f, 0xfd, 0xb5, 0xb6, 0xd3, 0x44, 0xcb, 0x56, 0xf0, 0x1d, 0xbc, 0xc8, 0x8d,
	0x6f, 0x10, 0x35, 0xd8, 0x1b, 0x2b, 0x46, 0xf0, 0x7a, 0xdd, 0x8c, 0x26, 0x98, 0x64, 0xc2, 0xee,
	0x2c, 0xd8, 0xa7, 0xf0, 0x15, 0x7c, 0x54, 0x99, 0x69, 0x11, 0x0b, 0x5e, 0x7e, 0x87, 0xdd, 0x39,
	0x70, 0x60, 0x75, 0x5b, 0xa1, 0x8b, 0xbe, 0xe5, 0xed, 0xd3, 0x76, 0xc4, 0x7c, 0xf4, 0xc4, 0x64,
	0x8c, 0xa3, 0x3e, 0x6f, 0x38, 0xb8, 0xbc, 0xaf, 0x5d, 0xde, 0x53, 0x8d, 0xdd, 0xe5, 0xe7, 0x04,
	0xb2, 0x83, 0xb7, 0xe6, 0x02, 0x56, 0x37, 0xf8, 0x6a, 0x63, 0xc7, 0xbf, 0xe3, 0xe5, 0x3f, 0x93,
	0xc1, 0x7c, 0x3d, 0xd4, 0xf8, 0xa1, 0x4c, 0x84, 0x15, 0x93, 0x7b, 0x57, 0x4e, 0x4c, 0x0a, 0xb3,
	0x32, 0x0e, 0xb5, 0xea, 0xbf, 0xa8, 0xa0, 0xbd, 0x8e, 0x44, 0x8f, 0x38, 0x92, 0xea, 0xd8, 0x9c,
	0xc1, 0xe2, 0xd9, 0x7a, 0x6f, 0x07, 0xd6, 0x60, 0x6a, 0x4e, 0x01, 0x36, 0x23, 0xb7, 0x34, 0xa8,
	0x4f, 0xe4, 0x41, 0x19, 0x39, 0x7a, 0x0c, 0x1a, 0xcc, 0xa4, 0xaa, 0x24, 0xbf, 0x6f, 0x9e, 0xeb,
	0x39, 0xcb, 0xa8, 0x02, 0xf9, 0x7d, 0xdf, 0x23, 0xdb, 0x4e, 0xbd, 0x30, 0x4b, 0x48, 0xaf, 0x6d,
	0x68, 0x7e, 0xea, 0x53, 0x79, 0x5f, 0x8d, 0xb4, 0x6b, 0xcb, 0xe4, 0xfa, 0x7a, 0x08, 0xed, 0x5b,
	0xb3, 0x0b, 0xce, 0xe5, 0xfa, 0x86, 0x1b, 0xf4, 0x4a, 0x57, 0xe4, 0xf0, 0xc7, 0x4e, 0xc5, 0xe1,
	0xa0, 0x0f, 0xb2, 0x67, 0xb8, 0x4b, 0xbe, 0x92, 0xe4, 0x65, 0xaa, 0xe3, 0x5e, 0x7d, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x9d, 0x83, 0x1c, 0xd4, 0x73, 0x01, 0x00, 0x00,
}
