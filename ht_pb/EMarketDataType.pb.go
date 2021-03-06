// Code generated by protoc-gen-go. DO NOT EDIT.
// source: EMarketDataType.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	EMarketDataType.proto

It has these top-level messages:
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

// 行情数据类型
type EMarketDataType int32

const (
	EMarketDataType_UNKNOWN_DATA_TYPE EMarketDataType = 0
	EMarketDataType_MD_TICK           EMarketDataType = 1
	EMarketDataType_MD_TRANSACTION    EMarketDataType = 2
	EMarketDataType_MD_ORDER          EMarketDataType = 3
	EMarketDataType_MD_CONSTANT       EMarketDataType = 4
	EMarketDataType_MD_KLINE_1MIN     EMarketDataType = 20
	EMarketDataType_MD_KLINE_5MIN     EMarketDataType = 21
	EMarketDataType_MD_KLINE_15MIN    EMarketDataType = 22
	EMarketDataType_MD_KLINE_30MIN    EMarketDataType = 23
	EMarketDataType_MD_KLINE_60MIN    EMarketDataType = 24
	EMarketDataType_MD_KLINE_1D       EMarketDataType = 25
	EMarketDataType_MD_TWAP_1MIN      EMarketDataType = 30
	EMarketDataType_MD_VWAP_1MIN      EMarketDataType = 40
	EMarketDataType_MD_VWAP_1S        EMarketDataType = 41
)

var EMarketDataType_name = map[int32]string{
	0:  "UNKNOWN_DATA_TYPE",
	1:  "MD_TICK",
	2:  "MD_TRANSACTION",
	3:  "MD_ORDER",
	4:  "MD_CONSTANT",
	20: "MD_KLINE_1MIN",
	21: "MD_KLINE_5MIN",
	22: "MD_KLINE_15MIN",
	23: "MD_KLINE_30MIN",
	24: "MD_KLINE_60MIN",
	25: "MD_KLINE_1D",
	30: "MD_TWAP_1MIN",
	40: "MD_VWAP_1MIN",
	41: "MD_VWAP_1S",
}
var EMarketDataType_value = map[string]int32{
	"UNKNOWN_DATA_TYPE": 0,
	"MD_TICK":           1,
	"MD_TRANSACTION":    2,
	"MD_ORDER":          3,
	"MD_CONSTANT":       4,
	"MD_KLINE_1MIN":     20,
	"MD_KLINE_5MIN":     21,
	"MD_KLINE_15MIN":    22,
	"MD_KLINE_30MIN":    23,
	"MD_KLINE_60MIN":    24,
	"MD_KLINE_1D":       25,
	"MD_TWAP_1MIN":      30,
	"MD_VWAP_1MIN":      40,
	"MD_VWAP_1S":        41,
}

func (x EMarketDataType) String() string {
	return proto.EnumName(EMarketDataType_name, int32(x))
}

func init() {
	proto.RegisterEnum("com.htsc.mdc.insight.model.EMarketDataType", EMarketDataType_name, EMarketDataType_value)
}
