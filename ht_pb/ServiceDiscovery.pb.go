// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ServiceDiscovery.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	ServiceDiscovery.proto

It has these top-level messages:
	ServiceDiscoveryRequest
	ServiceDiscoveryResponse
	ServerInfo
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

// 服务发现请求,在大部分场景下，服务发现网关会在用户登录成功后直接回送ServiceDiscoveryResponse，因此一般用户无需发送此请求
type ServiceDiscoveryRequest struct {
	AppType    int32  `protobuf:"varint,1,opt,name=appType" json:"appType,omitempty"`
	AppVersion string `protobuf:"bytes,2,opt,name=appVersion" json:"appVersion,omitempty"`
	UserName   string `protobuf:"bytes,3,opt,name=userName" json:"userName,omitempty"`
	DeviceId   string `protobuf:"bytes,4,opt,name=deviceId" json:"deviceId,omitempty"`
}

func (m *ServiceDiscoveryRequest) Reset()         { *m = ServiceDiscoveryRequest{} }
func (m *ServiceDiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceDiscoveryRequest) ProtoMessage()    {}

func (m *ServiceDiscoveryRequest) GetAppType() int32 {
	if m != nil {
		return m.AppType
	}
	return 0
}

func (m *ServiceDiscoveryRequest) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *ServiceDiscoveryRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ServiceDiscoveryRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

// 服务发现回复
type ServiceDiscoveryResponse struct {
	IsSuccess    bool                 `protobuf:"varint,1,opt,name=isSuccess" json:"isSuccess,omitempty"`
	ErrorContext *InsightErrorContext `protobuf:"bytes,2,opt,name=errorContext" json:"errorContext,omitempty"`
	Servers      []*ServerInfo        `protobuf:"bytes,3,rep,name=servers" json:"servers,omitempty"`
}

func (m *ServiceDiscoveryResponse) Reset()         { *m = ServiceDiscoveryResponse{} }
func (m *ServiceDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceDiscoveryResponse) ProtoMessage()    {}

func (m *ServiceDiscoveryResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *ServiceDiscoveryResponse) GetErrorContext() *InsightErrorContext {
	if m != nil {
		return m.ErrorContext
	}
	return nil
}

func (m *ServiceDiscoveryResponse) GetServers() []*ServerInfo {
	if m != nil {
		return m.Servers
	}
	return nil
}

// 服务器信息
type ServerInfo struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}

func (m *ServerInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ServerInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*ServiceDiscoveryRequest)(nil), "com.htsc.mdc.insight.model.ServiceDiscoveryRequest")
	proto.RegisterType((*ServiceDiscoveryResponse)(nil), "com.htsc.mdc.insight.model.ServiceDiscoveryResponse")
	proto.RegisterType((*ServerInfo)(nil), "com.htsc.mdc.insight.model.ServerInfo")
}
