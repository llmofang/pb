// Code generated by protoc-gen-go.
// source: trade_service.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	trade_service.proto

It has these top-level messages:
	TradeServiceMessage
	TradeServiceNetMsg
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TradeServiceMessage struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Ip               *string `protobuf:"bytes,2,req,name=ip" json:"ip,omitempty"`
	Port             *string `protobuf:"bytes,3,req,name=port" json:"port,omitempty"`
	Fundaccount      *string `protobuf:"bytes,4,req,name=fundaccount" json:"fundaccount,omitempty"`
	Stat             *string `protobuf:"bytes,5,req,name=stat" json:"stat,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TradeServiceMessage) Reset()         { *m = TradeServiceMessage{} }
func (m *TradeServiceMessage) String() string { return proto.CompactTextString(m) }
func (*TradeServiceMessage) ProtoMessage()    {}

func (m *TradeServiceMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TradeServiceMessage) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *TradeServiceMessage) GetPort() string {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return ""
}

func (m *TradeServiceMessage) GetFundaccount() string {
	if m != nil && m.Fundaccount != nil {
		return *m.Fundaccount
	}
	return ""
}

func (m *TradeServiceMessage) GetStat() string {
	if m != nil && m.Stat != nil {
		return *m.Stat
	}
	return ""
}

type TradeServiceNetMsg struct {
	Qid              *string `protobuf:"bytes,1,req,name=qid" json:"qid,omitempty"`
	Type             *string `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	Data             []byte  `protobuf:"bytes,3,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TradeServiceNetMsg) Reset()         { *m = TradeServiceNetMsg{} }
func (m *TradeServiceNetMsg) String() string { return proto.CompactTextString(m) }
func (*TradeServiceNetMsg) ProtoMessage()    {}

func (m *TradeServiceNetMsg) GetQid() string {
	if m != nil && m.Qid != nil {
		return *m.Qid
	}
	return ""
}

func (m *TradeServiceNetMsg) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *TradeServiceNetMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
}
