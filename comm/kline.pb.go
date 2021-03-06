// Code generated by protoc-gen-go.
// source: kline.proto
// DO NOT EDIT!

/*
Package comm is a generated protocol buffer package.

It is generated from these files:
	kline.proto

It has these top-level messages:
	KLineData
	KLineDataList
*/
package comm

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type KLineData struct {
	Code             *string `protobuf:"bytes,1,req" json:"Code,omitempty"`
	Date             *int32  `protobuf:"varint,2,req" json:"Date,omitempty"`
	Time             *int32  `protobuf:"varint,3,req" json:"Time,omitempty"`
	Timestamp        *int64  `protobuf:"varint,4,req" json:"Timestamp,omitempty"`
	Period           *int32  `protobuf:"varint,5,opt" json:"Period,omitempty"`
	Open             *int64  `protobuf:"varint,6,req" json:"Open,omitempty"`
	Close            *int64  `protobuf:"varint,7,req" json:"Close,omitempty"`
	High             *int64  `protobuf:"varint,8,req" json:"High,omitempty"`
	Low              *int64  `protobuf:"varint,9,req" json:"Low,omitempty"`
	NumTrades        *int64  `protobuf:"varint,10,req" json:"NumTrades,omitempty"`
	TotalVolumeTrade *int64  `protobuf:"varint,11,req" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade  *int64  `protobuf:"varint,12,req" json:"TotalValueTrade,omitempty"`
	IOPV             *int64  `protobuf:"varint,13,opt" json:"IOPV,omitempty"`
	OpenInterest     *int64  `protobuf:"varint,14,opt" json:"OpenInterest,omitempty"`
	SettlePrice      *int64  `protobuf:"varint,15,opt" json:"SettlePrice,omitempty"`
	PreClose         *int64  `protobuf:"varint,16,opt" json:"PreClose,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KLineData) Reset()         { *m = KLineData{} }
func (m *KLineData) String() string { return proto.CompactTextString(m) }
func (*KLineData) ProtoMessage()    {}

func (m *KLineData) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *KLineData) GetDate() int32 {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return 0
}

func (m *KLineData) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *KLineData) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *KLineData) GetPeriod() int32 {
	if m != nil && m.Period != nil {
		return *m.Period
	}
	return 0
}

func (m *KLineData) GetOpen() int64 {
	if m != nil && m.Open != nil {
		return *m.Open
	}
	return 0
}

func (m *KLineData) GetClose() int64 {
	if m != nil && m.Close != nil {
		return *m.Close
	}
	return 0
}

func (m *KLineData) GetHigh() int64 {
	if m != nil && m.High != nil {
		return *m.High
	}
	return 0
}

func (m *KLineData) GetLow() int64 {
	if m != nil && m.Low != nil {
		return *m.Low
	}
	return 0
}

func (m *KLineData) GetNumTrades() int64 {
	if m != nil && m.NumTrades != nil {
		return *m.NumTrades
	}
	return 0
}

func (m *KLineData) GetTotalVolumeTrade() int64 {
	if m != nil && m.TotalVolumeTrade != nil {
		return *m.TotalVolumeTrade
	}
	return 0
}

func (m *KLineData) GetTotalValueTrade() int64 {
	if m != nil && m.TotalValueTrade != nil {
		return *m.TotalValueTrade
	}
	return 0
}

func (m *KLineData) GetIOPV() int64 {
	if m != nil && m.IOPV != nil {
		return *m.IOPV
	}
	return 0
}

func (m *KLineData) GetOpenInterest() int64 {
	if m != nil && m.OpenInterest != nil {
		return *m.OpenInterest
	}
	return 0
}

func (m *KLineData) GetSettlePrice() int64 {
	if m != nil && m.SettlePrice != nil {
		return *m.SettlePrice
	}
	return 0
}

func (m *KLineData) GetPreClose() int64 {
	if m != nil && m.PreClose != nil {
		return *m.PreClose
	}
	return 0
}

type KLineDataList struct {
	Klinedata        []*KLineData `protobuf:"bytes,1,rep,name=klinedata" json:"klinedata,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *KLineDataList) Reset()         { *m = KLineDataList{} }
func (m *KLineDataList) String() string { return proto.CompactTextString(m) }
func (*KLineDataList) ProtoMessage()    {}

func (m *KLineDataList) GetKlinedata() []*KLineData {
	if m != nil {
		return m.Klinedata
	}
	return nil
}

func init() {
}
