// Code generated by protoc-gen-go.
// source: indexData.proto
// DO NOT EDIT!

/*
Package comm is a generated protocol buffer package.

It is generated from these files:
	indexData.proto

It has these top-level messages:
	IndexData
*/
package comm

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type IndexData struct {
	WindCode         *string `protobuf:"bytes,1,req" json:"WindCode,omitempty"`
	Code             *string `protobuf:"bytes,2,req" json:"Code,omitempty"`
	ActionDay        *int32  `protobuf:"varint,3,req" json:"ActionDay,omitempty"`
	TradingDay       *int32  `protobuf:"varint,4,req" json:"TradingDay,omitempty"`
	Time             *int32  `protobuf:"varint,5,req" json:"Time,omitempty"`
	OpenIndex        *int32  `protobuf:"varint,6,req" json:"OpenIndex,omitempty"`
	HighIndex        *int32  `protobuf:"varint,7,req" json:"HighIndex,omitempty"`
	LowIndex         *int32  `protobuf:"varint,8,req" json:"LowIndex,omitempty"`
	LastIndex        *int32  `protobuf:"varint,9,req" json:"LastIndex,omitempty"`
	TotalVolume      *int64  `protobuf:"varint,10,req" json:"TotalVolume,omitempty"`
	Turnover         *int64  `protobuf:"varint,11,req" json:"Turnover,omitempty"`
	PreCloseIndex    *int32  `protobuf:"varint,12,req" json:"PreCloseIndex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IndexData) Reset()         { *m = IndexData{} }
func (m *IndexData) String() string { return proto.CompactTextString(m) }
func (*IndexData) ProtoMessage()    {}

func (m *IndexData) GetWindCode() string {
	if m != nil && m.WindCode != nil {
		return *m.WindCode
	}
	return ""
}

func (m *IndexData) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *IndexData) GetActionDay() int32 {
	if m != nil && m.ActionDay != nil {
		return *m.ActionDay
	}
	return 0
}

func (m *IndexData) GetTradingDay() int32 {
	if m != nil && m.TradingDay != nil {
		return *m.TradingDay
	}
	return 0
}

func (m *IndexData) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *IndexData) GetOpenIndex() int32 {
	if m != nil && m.OpenIndex != nil {
		return *m.OpenIndex
	}
	return 0
}

func (m *IndexData) GetHighIndex() int32 {
	if m != nil && m.HighIndex != nil {
		return *m.HighIndex
	}
	return 0
}

func (m *IndexData) GetLowIndex() int32 {
	if m != nil && m.LowIndex != nil {
		return *m.LowIndex
	}
	return 0
}

func (m *IndexData) GetLastIndex() int32 {
	if m != nil && m.LastIndex != nil {
		return *m.LastIndex
	}
	return 0
}

func (m *IndexData) GetTotalVolume() int64 {
	if m != nil && m.TotalVolume != nil {
		return *m.TotalVolume
	}
	return 0
}

func (m *IndexData) GetTurnover() int64 {
	if m != nil && m.Turnover != nil {
		return *m.Turnover
	}
	return 0
}

func (m *IndexData) GetPreCloseIndex() int32 {
	if m != nil && m.PreCloseIndex != nil {
		return *m.PreCloseIndex
	}
	return 0
}

func init() {
}
