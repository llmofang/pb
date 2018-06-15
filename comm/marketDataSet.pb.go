// Code generated by protoc-gen-go.
// source: marketDataSet.proto
// DO NOT EDIT!

/*
Package comm is a generated protocol buffer package.

It is generated from these files:
	marketDataSet.proto

It has these top-level messages:
	MarketDataSet
*/
package comm

import proto "code.google.com/p/goprotobuf/proto"
import math "math"




// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MarketDataSet struct {
	DataType         *int32                 `protobuf:"varint,1,req" json:"DataType,omitempty"`
	Tick             *MarketData      `protobuf:"bytes,2,opt" json:"Tick,omitempty"`
	Transaction      *TransactionData `protobuf:"bytes,3,opt" json:"Transaction,omitempty"`
	OrderQueue       *OrderQueueData  `protobuf:"bytes,4,opt" json:"OrderQueue,omitempty"`
	Index            *IndexData       `protobuf:"bytes,5,opt" json:"Index,omitempty"`
	Future           *FutureData      `protobuf:"bytes,6,opt,name=future" json:"future,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *MarketDataSet) Reset()         { *m = MarketDataSet{} }
func (m *MarketDataSet) String() string { return proto.CompactTextString(m) }
func (*MarketDataSet) ProtoMessage()    {}

func (m *MarketDataSet) GetDataType() int32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *MarketDataSet) GetTick() *MarketData {
	if m != nil {
		return m.Tick
	}
	return nil
}

func (m *MarketDataSet) GetTransaction() *TransactionData {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *MarketDataSet) GetOrderQueue() *OrderQueueData {
	if m != nil {
		return m.OrderQueue
	}
	return nil
}

func (m *MarketDataSet) GetIndex() *IndexData {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *MarketDataSet) GetFuture() *FutureData {
	if m != nil {
		return m.Future
	}
	return nil
}

func init() {
}
