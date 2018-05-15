// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDBasicInfo.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	MDBasicInfo.proto

It has these top-level messages:
	MDBasicInfo
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

// 证券基础信息
type MDBasicInfo struct {
	HTSCSecurityID           string            `protobuf:"bytes,1,opt,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	SecurityID               string            `protobuf:"bytes,2,opt,name=SecurityID" json:"SecurityID,omitempty"`
	Symbol                   string            `protobuf:"bytes,3,opt,name=Symbol" json:"Symbol,omitempty"`
	ChiSpelling              string            `protobuf:"bytes,4,opt,name=ChiSpelling" json:"ChiSpelling,omitempty"`
	EnglishName              string            `protobuf:"bytes,5,opt,name=EnglishName" json:"EnglishName,omitempty"`
	SecurityIDSource         ESecurityIDSource `protobuf:"varint,6,opt,name=securityIDSource,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType             ESecurityType     `protobuf:"varint,7,opt,name=securityType,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	SecuritySubType          string            `protobuf:"bytes,8,opt,name=SecuritySubType" json:"SecuritySubType,omitempty"`
	ListDate                 string            `protobuf:"bytes,9,opt,name=ListDate" json:"ListDate,omitempty"`
	Currency                 string            `protobuf:"bytes,10,opt,name=Currency" json:"Currency,omitempty"`
	OutstandingShare         int64             `protobuf:"varint,11,opt,name=OutstandingShare" json:"OutstandingShare,omitempty"`
	PublicFloatShareQuantity int64             `protobuf:"varint,12,opt,name=PublicFloatShareQuantity" json:"PublicFloatShareQuantity,omitempty"`
	MDDate                   int32             `protobuf:"varint,13,opt,name=MDDate" json:"MDDate,omitempty"`
	TradingPhaseCode         string            `protobuf:"bytes,14,opt,name=TradingPhaseCode" json:"TradingPhaseCode,omitempty"`
	PreClosePx               int64             `protobuf:"varint,15,opt,name=PreClosePx" json:"PreClosePx,omitempty"`
	MaxPx                    int64             `protobuf:"varint,16,opt,name=MaxPx" json:"MaxPx,omitempty"`
	MinPx                    int64             `protobuf:"varint,17,opt,name=MinPx" json:"MinPx,omitempty"`
}

func (m *MDBasicInfo) Reset()         { *m = MDBasicInfo{} }
func (m *MDBasicInfo) String() string { return proto.CompactTextString(m) }
func (*MDBasicInfo) ProtoMessage()    {}

func (m *MDBasicInfo) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDBasicInfo) GetSecurityID() string {
	if m != nil {
		return m.SecurityID
	}
	return ""
}

func (m *MDBasicInfo) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MDBasicInfo) GetChiSpelling() string {
	if m != nil {
		return m.ChiSpelling
	}
	return ""
}

func (m *MDBasicInfo) GetEnglishName() string {
	if m != nil {
		return m.EnglishName
	}
	return ""
}

func (m *MDBasicInfo) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDBasicInfo) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDBasicInfo) GetSecuritySubType() string {
	if m != nil {
		return m.SecuritySubType
	}
	return ""
}

func (m *MDBasicInfo) GetListDate() string {
	if m != nil {
		return m.ListDate
	}
	return ""
}

func (m *MDBasicInfo) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *MDBasicInfo) GetOutstandingShare() int64 {
	if m != nil {
		return m.OutstandingShare
	}
	return 0
}

func (m *MDBasicInfo) GetPublicFloatShareQuantity() int64 {
	if m != nil {
		return m.PublicFloatShareQuantity
	}
	return 0
}

func (m *MDBasicInfo) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDBasicInfo) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDBasicInfo) GetPreClosePx() int64 {
	if m != nil {
		return m.PreClosePx
	}
	return 0
}

func (m *MDBasicInfo) GetMaxPx() int64 {
	if m != nil {
		return m.MaxPx
	}
	return 0
}

func (m *MDBasicInfo) GetMinPx() int64 {
	if m != nil {
		return m.MinPx
	}
	return 0
}

func init() {
	proto.RegisterType((*MDBasicInfo)(nil), "com.htsc.mdc.insight.model.MDBasicInfo")
}
