// Code generated by protoc-gen-go. DO NOT EDIT.
// source: EMDPeriodType.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	EMDPeriodType.proto

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

// 时间周期类型
type EMDPeriodType int32

const (
	EMDPeriodType_DefaultPeriod EMDPeriodType = 0
	EMDPeriodType_Period1Min    EMDPeriodType = 1
	EMDPeriodType_Period1D      EMDPeriodType = 2
	EMDPeriodType_Period5Min    EMDPeriodType = 3
	EMDPeriodType_Period15Min   EMDPeriodType = 4
	EMDPeriodType_Period30Min   EMDPeriodType = 5
	EMDPeriodType_Period1H      EMDPeriodType = 6
	EMDPeriodType_Period1S      EMDPeriodType = 7
	EMDPeriodType_Period3S      EMDPeriodType = 8
	EMDPeriodType_Period5S      EMDPeriodType = 9
	EMDPeriodType_Period15S     EMDPeriodType = 10
	EMDPeriodType_Period30S     EMDPeriodType = 11
	EMDPeriodType_Period2H      EMDPeriodType = 12
	EMDPeriodType_Period1W      EMDPeriodType = 13
	EMDPeriodType_Period15D     EMDPeriodType = 14
	EMDPeriodType_Period1Month  EMDPeriodType = 15
	EMDPeriodType_Period3Month  EMDPeriodType = 16
	EMDPeriodType_Period4Month  EMDPeriodType = 17
	EMDPeriodType_Period6Month  EMDPeriodType = 18
	EMDPeriodType_Period1Y      EMDPeriodType = 19
)

var EMDPeriodType_name = map[int32]string{
	0:  "DefaultPeriod",
	1:  "Period1Min",
	2:  "Period1D",
	3:  "Period5Min",
	4:  "Period15Min",
	5:  "Period30Min",
	6:  "Period1H",
	7:  "Period1S",
	8:  "Period3S",
	9:  "Period5S",
	10: "Period15S",
	11: "Period30S",
	12: "Period2H",
	13: "Period1W",
	14: "Period15D",
	15: "Period1Month",
	16: "Period3Month",
	17: "Period4Month",
	18: "Period6Month",
	19: "Period1Y",
}
var EMDPeriodType_value = map[string]int32{
	"DefaultPeriod": 0,
	"Period1Min":    1,
	"Period1D":      2,
	"Period5Min":    3,
	"Period15Min":   4,
	"Period30Min":   5,
	"Period1H":      6,
	"Period1S":      7,
	"Period3S":      8,
	"Period5S":      9,
	"Period15S":     10,
	"Period30S":     11,
	"Period2H":      12,
	"Period1W":      13,
	"Period15D":     14,
	"Period1Month":  15,
	"Period3Month":  16,
	"Period4Month":  17,
	"Period6Month":  18,
	"Period1Y":      19,
}

func (x EMDPeriodType) String() string {
	return proto.EnumName(EMDPeriodType_name, int32(x))
}

func init() {
	proto.RegisterEnum("com.htsc.mdc.model.EMDPeriodType", EMDPeriodType_name, EMDPeriodType_value)
}
