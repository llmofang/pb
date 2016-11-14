// Code generated by protoc-gen-go.
// source: klineData.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	klineData.proto

It has these top-level messages:
	KlineData
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type KlineData struct {
	Stock            *string `protobuf:"bytes,1,req,name=stock" json:"stock,omitempty"`
	Time             *int64  `protobuf:"varint,2,req,name=time" json:"time,omitempty"`
	Open             *string `protobuf:"bytes,3,req,name=open" json:"open,omitempty"`
	High             *string `protobuf:"bytes,4,req,name=high" json:"high,omitempty"`
	Low              *string `protobuf:"bytes,5,req,name=low" json:"low,omitempty"`
	Close            *string `protobuf:"bytes,6,req,name=close" json:"close,omitempty"`
	Size             *string `protobuf:"bytes,7,req,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KlineData) Reset()         { *m = KlineData{} }
func (m *KlineData) String() string { return proto.CompactTextString(m) }
func (*KlineData) ProtoMessage()    {}

func (m *KlineData) GetStock() string {
	if m != nil && m.Stock != nil {
		return *m.Stock
	}
	return ""
}

func (m *KlineData) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *KlineData) GetOpen() string {
	if m != nil && m.Open != nil {
		return *m.Open
	}
	return ""
}

func (m *KlineData) GetHigh() string {
	if m != nil && m.High != nil {
		return *m.High
	}
	return ""
}

func (m *KlineData) GetLow() string {
	if m != nil && m.Low != nil {
		return *m.Low
	}
	return ""
}

func (m *KlineData) GetClose() string {
	if m != nil && m.Close != nil {
		return *m.Close
	}
	return ""
}

func (m *KlineData) GetSize() string {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return ""
}

func init() {
}