// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SecuritySourceType.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	SecuritySourceType.proto

It has these top-level messages:
	SecuritySourceType
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

// 市场类型及证券类型，类型和市场源不能同时为空
type SecuritySourceType struct {
	SecurityIDSource ESecurityIDSource `protobuf:"varint,1,opt,name=securityIDSource,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType     ESecurityType     `protobuf:"varint,2,opt,name=securityType,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
}

func (m *SecuritySourceType) Reset()         { *m = SecuritySourceType{} }
func (m *SecuritySourceType) String() string { return proto.CompactTextString(m) }
func (*SecuritySourceType) ProtoMessage()    {}

func (m *SecuritySourceType) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *SecuritySourceType) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func init() {
	proto.RegisterType((*SecuritySourceType)(nil), "com.htsc.mdc.insight.model.SecuritySourceType")
}
