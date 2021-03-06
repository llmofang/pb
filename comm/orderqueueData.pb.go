// Code generated by protoc-gen-go.
// source: orderqueueData.proto
// DO NOT EDIT!

/*
Package comm is a generated protocol buffer package.

It is generated from these files:
	orderqueueData.proto

It has these top-level messages:
	OrderQueueData
*/
package comm

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type OrderQueueData struct {
	WindCode         *string `protobuf:"bytes,1,req" json:"WindCode,omitempty"`
	Code             *string `protobuf:"bytes,2,req" json:"Code,omitempty"`
	ActionDay        *int32  `protobuf:"varint,3,req" json:"ActionDay,omitempty"`
	Time             *int32  `protobuf:"varint,4,req" json:"Time,omitempty"`
	Side             *int32  `protobuf:"varint,5,req" json:"Side,omitempty"`
	Price            *int32  `protobuf:"varint,6,req" json:"Price,omitempty"`
	Orders           *int32  `protobuf:"varint,7,req" json:"Orders,omitempty"`
	ABItems          *int32  `protobuf:"varint,8,req" json:"ABItems,omitempty"`
	ABVolume         []int32 `protobuf:"varint,9,rep" json:"ABVolume,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OrderQueueData) Reset()         { *m = OrderQueueData{} }
func (m *OrderQueueData) String() string { return proto.CompactTextString(m) }
func (*OrderQueueData) ProtoMessage()    {}

func (m *OrderQueueData) GetWindCode() string {
	if m != nil && m.WindCode != nil {
		return *m.WindCode
	}
	return ""
}

func (m *OrderQueueData) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *OrderQueueData) GetActionDay() int32 {
	if m != nil && m.ActionDay != nil {
		return *m.ActionDay
	}
	return 0
}

func (m *OrderQueueData) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *OrderQueueData) GetSide() int32 {
	if m != nil && m.Side != nil {
		return *m.Side
	}
	return 0
}

func (m *OrderQueueData) GetPrice() int32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *OrderQueueData) GetOrders() int32 {
	if m != nil && m.Orders != nil {
		return *m.Orders
	}
	return 0
}

func (m *OrderQueueData) GetABItems() int32 {
	if m != nil && m.ABItems != nil {
		return *m.ABItems
	}
	return 0
}

func (m *OrderQueueData) GetABVolume() []int32 {
	if m != nil {
		return m.ABVolume
	}
	return nil
}

func init() {
}
