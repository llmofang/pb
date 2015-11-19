// Code generated by protoc-gen-go.
// source: userinfo.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	userinfo.proto

It has these top-level messages:
	FundInfo
	UserInfo
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FundInfo struct {
	FundName         *string `protobuf:"bytes,1,req" json:"FundName,omitempty"`
	StockCode        *string `protobuf:"bytes,2,req" json:"StockCode,omitempty"`
	StockName        *string `protobuf:"bytes,3,req" json:"StockName,omitempty"`
	StockNum         *string `protobuf:"bytes,4,req" json:"StockNum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FundInfo) Reset()         { *m = FundInfo{} }
func (m *FundInfo) String() string { return proto.CompactTextString(m) }
func (*FundInfo) ProtoMessage()    {}

func (m *FundInfo) GetFundName() string {
	if m != nil && m.FundName != nil {
		return *m.FundName
	}
	return ""
}

func (m *FundInfo) GetStockCode() string {
	if m != nil && m.StockCode != nil {
		return *m.StockCode
	}
	return ""
}

func (m *FundInfo) GetStockName() string {
	if m != nil && m.StockName != nil {
		return *m.StockName
	}
	return ""
}

func (m *FundInfo) GetStockNum() string {
	if m != nil && m.StockNum != nil {
		return *m.StockNum
	}
	return ""
}

type UserInfo struct {
	UserName         *string     `protobuf:"bytes,1,req" json:"UserName,omitempty"`
	Funds            []*FundInfo `protobuf:"bytes,2,rep" json:"Funds,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}

func (m *UserInfo) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *UserInfo) GetFunds() []*FundInfo {
	if m != nil {
		return m.Funds
	}
	return nil
}

func init() {
}
