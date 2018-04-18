// Code generated by protoc-gen-go.
// source: transactionData.proto
// DO NOT EDIT!

/*
Package comm is a generated protocol buffer package.

It is generated from these files:
	transactionData.proto

It has these top-level messages:
	TransactionData
*/
package comm

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TransactionData struct {
	WindCode         *string `protobuf:"bytes,1,req" json:"WindCode,omitempty"`
	Code             *string `protobuf:"bytes,2,req" json:"Code,omitempty"`
	ActionDay        *int32  `protobuf:"varint,3,req" json:"ActionDay,omitempty"`
	Time             *int32  `protobuf:"varint,4,req" json:"Time,omitempty"`
	Index            *int32  `protobuf:"varint,5,req" json:"Index,omitempty"`
	Price            *int32  `protobuf:"varint,6,req" json:"Price,omitempty"`
	Volume           *int32  `protobuf:"varint,7,req" json:"Volume,omitempty"`
	Turnover         *int32  `protobuf:"varint,8,req" json:"Turnover,omitempty"`
	BSFlag           *int32  `protobuf:"varint,9,req" json:"BSFlag,omitempty"`
	OrderKind        *string `protobuf:"bytes,10,req" json:"OrderKind,omitempty"`
	FunctionCode     *string `protobuf:"bytes,11,req" json:"FunctionCode,omitempty"`
	AskOrder         *int32  `protobuf:"varint,12,req" json:"AskOrder,omitempty"`
	BidOrder         *int32  `protobuf:"varint,13,req" json:"BidOrder,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TransactionData) Reset()         { *m = TransactionData{} }
func (m *TransactionData) String() string { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()    {}

func (m *TransactionData) GetWindCode() string {
	if m != nil && m.WindCode != nil {
		return *m.WindCode
	}
	return ""
}

func (m *TransactionData) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *TransactionData) GetActionDay() int32 {
	if m != nil && m.ActionDay != nil {
		return *m.ActionDay
	}
	return 0
}

func (m *TransactionData) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *TransactionData) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *TransactionData) GetPrice() int32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *TransactionData) GetVolume() int32 {
	if m != nil && m.Volume != nil {
		return *m.Volume
	}
	return 0
}

func (m *TransactionData) GetTurnover() int32 {
	if m != nil && m.Turnover != nil {
		return *m.Turnover
	}
	return 0
}

func (m *TransactionData) GetBSFlag() int32 {
	if m != nil && m.BSFlag != nil {
		return *m.BSFlag
	}
	return 0
}

func (m *TransactionData) GetOrderKind() string {
	if m != nil && m.OrderKind != nil {
		return *m.OrderKind
	}
	return ""
}

func (m *TransactionData) GetFunctionCode() string {
	if m != nil && m.FunctionCode != nil {
		return *m.FunctionCode
	}
	return ""
}

func (m *TransactionData) GetAskOrder() int32 {
	if m != nil && m.AskOrder != nil {
		return *m.AskOrder
	}
	return 0
}

func (m *TransactionData) GetBidOrder() int32 {
	if m != nil && m.BidOrder != nil {
		return *m.BidOrder
	}
	return 0
}

func init() {
}
