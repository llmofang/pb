// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDPlayback.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	MDPlayback.proto

It has these top-level messages:
	PlaybackRequest
	PlaybackResponse
	PlaybackControlRequest
	PlaybackControlResponse
	PlaybackStatusRequest
	PlaybackStatus
	PlaybackPayload
*/
package com_htsc_mdc_insight_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_htsc_mdc_insight_model1 "."
import com_htsc_mdc_insight_model14 "."
import com_htsc_mdc_insight_model15 "."
import com_htsc_mdc_insight_model16 "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 复权类型
type EPlaybackExrightsType int32

const (
	EPlaybackExrightsType_DEFAULT_EXRIGHTS_TYPE EPlaybackExrightsType = 0
	EPlaybackExrightsType_NO_EXRIGHTS           EPlaybackExrightsType = 10
	EPlaybackExrightsType_FORWARD_EXRIGHTS      EPlaybackExrightsType = 11
	EPlaybackExrightsType_BACKWARD_EXRIGHTS     EPlaybackExrightsType = 12
)

var EPlaybackExrightsType_name = map[int32]string{
	0:  "DEFAULT_EXRIGHTS_TYPE",
	10: "NO_EXRIGHTS",
	11: "FORWARD_EXRIGHTS",
	12: "BACKWARD_EXRIGHTS",
}
var EPlaybackExrightsType_value = map[string]int32{
	"DEFAULT_EXRIGHTS_TYPE": 0,
	"NO_EXRIGHTS":           10,
	"FORWARD_EXRIGHTS":      11,
	"BACKWARD_EXRIGHTS":     12,
}

func (x EPlaybackExrightsType) String() string {
	return proto.EnumName(EPlaybackExrightsType_name, int32(x))
}
func (EPlaybackExrightsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 任务控制类型
type EPlaybackTaskControlType int32

const (
	EPlaybackTaskControlType_DEFAULT_CONTROL_TYPE EPlaybackTaskControlType = 0
	EPlaybackTaskControlType_CANCEL_TASK          EPlaybackTaskControlType = 1
	EPlaybackTaskControlType_SET_PLAYBACK_RATE    EPlaybackTaskControlType = 2
)

var EPlaybackTaskControlType_name = map[int32]string{
	0: "DEFAULT_CONTROL_TYPE",
	1: "CANCEL_TASK",
	2: "SET_PLAYBACK_RATE",
}
var EPlaybackTaskControlType_value = map[string]int32{
	"DEFAULT_CONTROL_TYPE": 0,
	"CANCEL_TASK":          1,
	"SET_PLAYBACK_RATE":    2,
}

func (x EPlaybackTaskControlType) String() string {
	return proto.EnumName(EPlaybackTaskControlType_name, int32(x))
}
func (EPlaybackTaskControlType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 任务状态指示
type EPlaybackTaskStatus int32

const (
	EPlaybackTaskStatus_DEFAULT_STATUS EPlaybackTaskStatus = 0
	EPlaybackTaskStatus_INITIALIZING   EPlaybackTaskStatus = 11
	EPlaybackTaskStatus_PREPARING      EPlaybackTaskStatus = 12
	EPlaybackTaskStatus_PREPARED       EPlaybackTaskStatus = 13
	EPlaybackTaskStatus_RUNNING        EPlaybackTaskStatus = 14
	EPlaybackTaskStatus_APPENDING      EPlaybackTaskStatus = 15
	EPlaybackTaskStatus_CANCELED       EPlaybackTaskStatus = 16
	EPlaybackTaskStatus_COMPLETED      EPlaybackTaskStatus = 17
	EPlaybackTaskStatus_FAILED         EPlaybackTaskStatus = 18
)

var EPlaybackTaskStatus_name = map[int32]string{
	0:  "DEFAULT_STATUS",
	11: "INITIALIZING",
	12: "PREPARING",
	13: "PREPARED",
	14: "RUNNING",
	15: "APPENDING",
	16: "CANCELED",
	17: "COMPLETED",
	18: "FAILED",
}
var EPlaybackTaskStatus_value = map[string]int32{
	"DEFAULT_STATUS": 0,
	"INITIALIZING":   11,
	"PREPARING":      12,
	"PREPARED":       13,
	"RUNNING":        14,
	"APPENDING":      15,
	"CANCELED":       16,
	"COMPLETED":      17,
	"FAILED":         18,
}

func (x EPlaybackTaskStatus) String() string {
	return proto.EnumName(EPlaybackTaskStatus_name, int32(x))
}
func (EPlaybackTaskStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 行情回放请求
type PlaybackRequest struct {
	TaskId               string                                             `protobuf:"bytes,1,opt,name=taskId" json:"taskId,omitempty"`
	HtscSecurityIDs      []string                                           `protobuf:"bytes,2,rep,name=htscSecurityIDs" json:"htscSecurityIDs,omitempty"`
	SecuritySourceType   []*com_htsc_mdc_insight_model16.SecuritySourceType `protobuf:"bytes,3,rep,name=securitySourceType" json:"securitySourceType,omitempty"`
	StartTime            string                                             `protobuf:"bytes,4,opt,name=startTime" json:"startTime,omitempty"`
	StopTime             string                                             `protobuf:"bytes,5,opt,name=stopTime" json:"stopTime,omitempty"`
	ReplayDataType       com_htsc_mdc_insight_model1.EMarketDataType        `protobuf:"varint,6,opt,name=replayDataType,enum=com.htsc.mdc.insight.model.EMarketDataType" json:"replayDataType,omitempty"`
	ReplayRate           int32                                              `protobuf:"varint,7,opt,name=replayRate" json:"replayRate,omitempty"`
	ExrightsType         EPlaybackExrightsType                              `protobuf:"varint,8,opt,name=exrightsType,enum=com.htsc.mdc.insight.model.EPlaybackExrightsType" json:"exrightsType,omitempty"`
	IsNeedInitialData    bool                                               `protobuf:"varint,9,opt,name=isNeedInitialData" json:"isNeedInitialData,omitempty"`
	InitialDataStartTime string                                             `protobuf:"bytes,10,opt,name=initialDataStartTime" json:"initialDataStartTime,omitempty"`
}

func (m *PlaybackRequest) Reset()                    { *m = PlaybackRequest{} }
func (m *PlaybackRequest) String() string            { return proto.CompactTextString(m) }
func (*PlaybackRequest) ProtoMessage()               {}
func (*PlaybackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PlaybackRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *PlaybackRequest) GetHtscSecurityIDs() []string {
	if m != nil {
		return m.HtscSecurityIDs
	}
	return nil
}

func (m *PlaybackRequest) GetSecuritySourceType() []*com_htsc_mdc_insight_model16.SecuritySourceType {
	if m != nil {
		return m.SecuritySourceType
	}
	return nil
}

func (m *PlaybackRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *PlaybackRequest) GetStopTime() string {
	if m != nil {
		return m.StopTime
	}
	return ""
}

func (m *PlaybackRequest) GetReplayDataType() com_htsc_mdc_insight_model1.EMarketDataType {
	if m != nil {
		return m.ReplayDataType
	}
	return com_htsc_mdc_insight_model1.EMarketDataType_UNKNOWN_DATA_TYPE
}

func (m *PlaybackRequest) GetReplayRate() int32 {
	if m != nil {
		return m.ReplayRate
	}
	return 0
}

func (m *PlaybackRequest) GetExrightsType() EPlaybackExrightsType {
	if m != nil {
		return m.ExrightsType
	}
	return EPlaybackExrightsType_DEFAULT_EXRIGHTS_TYPE
}

func (m *PlaybackRequest) GetIsNeedInitialData() bool {
	if m != nil {
		return m.IsNeedInitialData
	}
	return false
}

func (m *PlaybackRequest) GetInitialDataStartTime() string {
	if m != nil {
		return m.InitialDataStartTime
	}
	return ""
}

// 回放应答
type PlaybackResponse struct {
	TaskId       string                                            `protobuf:"bytes,1,opt,name=taskId" json:"taskId,omitempty"`
	IsSuccess    bool                                              `protobuf:"varint,2,opt,name=isSuccess" json:"isSuccess,omitempty"`
	ErrorContext *com_htsc_mdc_insight_model15.InsightErrorContext `protobuf:"bytes,3,opt,name=errorContext" json:"errorContext,omitempty"`
}

func (m *PlaybackResponse) Reset()                    { *m = PlaybackResponse{} }
func (m *PlaybackResponse) String() string            { return proto.CompactTextString(m) }
func (*PlaybackResponse) ProtoMessage()               {}
func (*PlaybackResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PlaybackResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *PlaybackResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *PlaybackResponse) GetErrorContext() *com_htsc_mdc_insight_model15.InsightErrorContext {
	if m != nil {
		return m.ErrorContext
	}
	return nil
}

// 任务控制请求
type PlaybackControlRequest struct {
	TaskId      string                   `protobuf:"bytes,1,opt,name=taskId" json:"taskId,omitempty"`
	ControlType EPlaybackTaskControlType `protobuf:"varint,2,opt,name=controlType,enum=com.htsc.mdc.insight.model.EPlaybackTaskControlType" json:"controlType,omitempty"`
	ReplayRate  int32                    `protobuf:"varint,3,opt,name=replayRate" json:"replayRate,omitempty"`
}

func (m *PlaybackControlRequest) Reset()                    { *m = PlaybackControlRequest{} }
func (m *PlaybackControlRequest) String() string            { return proto.CompactTextString(m) }
func (*PlaybackControlRequest) ProtoMessage()               {}
func (*PlaybackControlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PlaybackControlRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *PlaybackControlRequest) GetControlType() EPlaybackTaskControlType {
	if m != nil {
		return m.ControlType
	}
	return EPlaybackTaskControlType_DEFAULT_CONTROL_TYPE
}

func (m *PlaybackControlRequest) GetReplayRate() int32 {
	if m != nil {
		return m.ReplayRate
	}
	return 0
}

// 回放任务控制应答
type PlaybackControlResponse struct {
	TaskId            string                                            `protobuf:"bytes,1,opt,name=taskId" json:"taskId,omitempty"`
	IsSuccess         bool                                              `protobuf:"varint,2,opt,name=isSuccess" json:"isSuccess,omitempty"`
	ErrorContext      *com_htsc_mdc_insight_model15.InsightErrorContext `protobuf:"bytes,3,opt,name=errorContext" json:"errorContext,omitempty"`
	CurrentReplayRate int32                                             `protobuf:"varint,4,opt,name=currentReplayRate" json:"currentReplayRate,omitempty"`
}

func (m *PlaybackControlResponse) Reset()                    { *m = PlaybackControlResponse{} }
func (m *PlaybackControlResponse) String() string            { return proto.CompactTextString(m) }
func (*PlaybackControlResponse) ProtoMessage()               {}
func (*PlaybackControlResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlaybackControlResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *PlaybackControlResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *PlaybackControlResponse) GetErrorContext() *com_htsc_mdc_insight_model15.InsightErrorContext {
	if m != nil {
		return m.ErrorContext
	}
	return nil
}

func (m *PlaybackControlResponse) GetCurrentReplayRate() int32 {
	if m != nil {
		return m.CurrentReplayRate
	}
	return 0
}

// 回放任务状态请求
type PlaybackStatusRequest struct {
	TaskId string `protobuf:"bytes,1,opt,name=taskId" json:"taskId,omitempty"`
}

func (m *PlaybackStatusRequest) Reset()                    { *m = PlaybackStatusRequest{} }
func (m *PlaybackStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*PlaybackStatusRequest) ProtoMessage()               {}
func (*PlaybackStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PlaybackStatusRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

// 回放任务状态
type PlaybackStatus struct {
	TaskId            string              `protobuf:"bytes,1,opt,name=taskId" json:"taskId,omitempty"`
	TaskStatus        EPlaybackTaskStatus `protobuf:"varint,2,opt,name=taskStatus,enum=com.htsc.mdc.insight.model.EPlaybackTaskStatus" json:"taskStatus,omitempty"`
	ReplayPercent     int32               `protobuf:"varint,3,opt,name=replayPercent" json:"replayPercent,omitempty"`
	CurrentReplayRate int32               `protobuf:"varint,4,opt,name=currentReplayRate" json:"currentReplayRate,omitempty"`
}

func (m *PlaybackStatus) Reset()                    { *m = PlaybackStatus{} }
func (m *PlaybackStatus) String() string            { return proto.CompactTextString(m) }
func (*PlaybackStatus) ProtoMessage()               {}
func (*PlaybackStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PlaybackStatus) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *PlaybackStatus) GetTaskStatus() EPlaybackTaskStatus {
	if m != nil {
		return m.TaskStatus
	}
	return EPlaybackTaskStatus_DEFAULT_STATUS
}

func (m *PlaybackStatus) GetReplayPercent() int32 {
	if m != nil {
		return m.ReplayPercent
	}
	return 0
}

func (m *PlaybackStatus) GetCurrentReplayRate() int32 {
	if m != nil {
		return m.CurrentReplayRate
	}
	return 0
}

// 回放数据
type PlaybackPayload struct {
	TaskId           string                                         `protobuf:"bytes,1,opt,name=taskId" json:"taskId,omitempty"`
	MarketDataStream *com_htsc_mdc_insight_model14.MarketDataStream `protobuf:"bytes,2,opt,name=marketDataStream" json:"marketDataStream,omitempty"`
}

func (m *PlaybackPayload) Reset()                    { *m = PlaybackPayload{} }
func (m *PlaybackPayload) String() string            { return proto.CompactTextString(m) }
func (*PlaybackPayload) ProtoMessage()               {}
func (*PlaybackPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PlaybackPayload) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *PlaybackPayload) GetMarketDataStream() *com_htsc_mdc_insight_model14.MarketDataStream {
	if m != nil {
		return m.MarketDataStream
	}
	return nil
}

func init() {
	proto.RegisterType((*PlaybackRequest)(nil), "com.htsc.mdc.insight.model.PlaybackRequest")
	proto.RegisterType((*PlaybackResponse)(nil), "com.htsc.mdc.insight.model.PlaybackResponse")
	proto.RegisterType((*PlaybackControlRequest)(nil), "com.htsc.mdc.insight.model.PlaybackControlRequest")
	proto.RegisterType((*PlaybackControlResponse)(nil), "com.htsc.mdc.insight.model.PlaybackControlResponse")
	proto.RegisterType((*PlaybackStatusRequest)(nil), "com.htsc.mdc.insight.model.PlaybackStatusRequest")
	proto.RegisterType((*PlaybackStatus)(nil), "com.htsc.mdc.insight.model.PlaybackStatus")
	proto.RegisterType((*PlaybackPayload)(nil), "com.htsc.mdc.insight.model.PlaybackPayload")
	proto.RegisterEnum("com.htsc.mdc.insight.model.EPlaybackExrightsType", EPlaybackExrightsType_name, EPlaybackExrightsType_value)
	proto.RegisterEnum("com.htsc.mdc.insight.model.EPlaybackTaskControlType", EPlaybackTaskControlType_name, EPlaybackTaskControlType_value)
	proto.RegisterEnum("com.htsc.mdc.insight.model.EPlaybackTaskStatus", EPlaybackTaskStatus_name, EPlaybackTaskStatus_value)
}

func init() { proto.RegisterFile("MDPlayback.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x8e, 0xda, 0x46,
	0x14, 0xce, 0x2c, 0x9b, 0x0d, 0x1c, 0xb3, 0x30, 0x3b, 0x5d, 0x52, 0x07, 0x45, 0x15, 0x42, 0xbd,
	0xb0, 0xb6, 0x91, 0x57, 0xa5, 0x7d, 0x01, 0x2f, 0xf6, 0x26, 0x56, 0xc0, 0x58, 0x63, 0xd3, 0x26,
	0x55, 0x55, 0xe4, 0x98, 0x51, 0x63, 0x2d, 0x60, 0xea, 0x19, 0xa4, 0x70, 0xdd, 0xc7, 0xa8, 0x22,
	0xf5, 0xb2, 0x2f, 0xd3, 0x9b, 0x3e, 0x51, 0x35, 0x36, 0xc6, 0xfc, 0x27, 0xbd, 0xeb, 0x1d, 0xe7,
	0x3b, 0x33, 0xdf, 0xf9, 0xce, 0xf9, 0x0e, 0x1e, 0xc0, 0x7d, 0xd3, 0x9d, 0x04, 0xcb, 0x77, 0x41,
	0xf8, 0xa0, 0xcf, 0x93, 0x58, 0xc4, 0xa4, 0x19, 0xc6, 0x53, 0xfd, 0xbd, 0xe0, 0xa1, 0x3e, 0x1d,
	0x87, 0x7a, 0x34, 0xe3, 0xd1, 0xaf, 0xef, 0x85, 0x3e, 0x8d, 0xc7, 0x6c, 0xd2, 0x6c, 0x58, 0xfd,
	0x20, 0x79, 0x60, 0xc2, 0x0c, 0x44, 0xe0, 0x2f, 0xe7, 0x2c, 0xbb, 0xd2, 0xc4, 0x05, 0xba, 0x42,
	0x9e, 0xd9, 0xd9, 0x3d, 0x2b, 0x49, 0xe2, 0xa4, 0x1b, 0xcf, 0x04, 0xfb, 0x20, 0x56, 0x29, 0xd5,
	0x63, 0xe1, 0x22, 0x89, 0xc4, 0xd2, 0x8b, 0x17, 0x49, 0xc8, 0x0a, 0x9a, 0xf6, 0x1f, 0xe7, 0x50,
	0xcf, 0xc5, 0x50, 0xf6, 0xdb, 0x82, 0x71, 0x41, 0x9e, 0xc2, 0x85, 0x08, 0xf8, 0x83, 0x3d, 0x56,
	0x51, 0x0b, 0x69, 0x15, 0xba, 0x8a, 0x88, 0x06, 0x75, 0xa9, 0x31, 0xe7, 0xb2, 0x4d, 0xae, 0x9e,
	0xb5, 0x4a, 0x5a, 0x85, 0xee, 0xc2, 0xe4, 0x17, 0x20, 0x7c, 0xaf, 0xa2, 0x5a, 0x6a, 0x95, 0x34,
	0xa5, 0xa3, 0xeb, 0xc7, 0x9b, 0xd5, 0xf7, 0x75, 0xd2, 0x03, 0x4c, 0xe4, 0x39, 0x54, 0xb8, 0x08,
	0x12, 0xe1, 0x47, 0x53, 0xa6, 0x9e, 0xa7, 0x22, 0x0b, 0x80, 0x34, 0xa1, 0xcc, 0x45, 0x3c, 0x4f,
	0x93, 0x8f, 0xd3, 0xe4, 0x3a, 0x26, 0x1e, 0xd4, 0x12, 0x36, 0x9f, 0x04, 0xcb, 0x7c, 0x9c, 0xea,
	0x45, 0x0b, 0x69, 0xb5, 0xce, 0x37, 0xa7, 0x54, 0xed, 0x38, 0x40, 0x77, 0x28, 0xc8, 0x57, 0x00,
	0x19, 0x42, 0x03, 0xc1, 0xd4, 0x27, 0x2d, 0xa4, 0x3d, 0xa6, 0x1b, 0x08, 0x19, 0x42, 0x95, 0x7d,
	0x48, 0x24, 0x21, 0x4f, 0x4b, 0x96, 0xd3, 0x92, 0xdf, 0x9e, 0x2c, 0x99, 0x9b, 0x62, 0x6d, 0x5c,
	0xa4, 0x5b, 0x34, 0xe4, 0x05, 0x5c, 0x45, 0xdc, 0x61, 0x6c, 0x6c, 0xcf, 0x22, 0x11, 0x05, 0x13,
	0xa9, 0x47, 0xad, 0xb4, 0x90, 0x56, 0xa6, 0xfb, 0x09, 0xd2, 0x81, 0xeb, 0xa8, 0x08, 0xbd, 0xf5,
	0xf8, 0x20, 0x9d, 0xd0, 0xc1, 0x5c, 0xfb, 0x23, 0x02, 0x5c, 0x6c, 0x07, 0x9f, 0xc7, 0x33, 0xce,
	0x8e, 0xae, 0xc7, 0x73, 0xa8, 0x44, 0xdc, 0x5b, 0x84, 0x21, 0xe3, 0x72, 0x31, 0xa4, 0x8c, 0x02,
	0x20, 0x1e, 0x54, 0xd9, 0xc6, 0x62, 0xaa, 0xa5, 0x16, 0xd2, 0x94, 0xce, 0xed, 0xa9, 0x19, 0x1c,
	0xd8, 0x67, 0xba, 0x45, 0xd2, 0xfe, 0x0b, 0xc1, 0xd3, 0x5c, 0x9f, 0xc4, 0x92, 0x78, 0xf2, 0xa9,
	0x25, 0xfe, 0x01, 0x94, 0x30, 0x3b, 0x99, 0x5a, 0x71, 0x96, 0x5a, 0xf1, 0xfd, 0x67, 0x59, 0xe1,
	0x07, 0x3c, 0xaf, 0x92, 0xba, 0xb1, 0x49, 0xb4, 0xb3, 0x03, 0xa5, 0xdd, 0x1d, 0x68, 0xff, 0x83,
	0xe0, 0xcb, 0x3d, 0xa9, 0xff, 0xbb, 0x89, 0xca, 0x9d, 0x0a, 0x17, 0x49, 0xc2, 0x66, 0x82, 0x16,
	0xdd, 0x9c, 0xa7, 0xdd, 0xec, 0x27, 0xda, 0xb7, 0xd0, 0xc8, 0x7b, 0xf2, 0x44, 0x20, 0x16, 0xfc,
	0x13, 0xd3, 0x6f, 0xff, 0x8d, 0xa0, 0xb6, 0x7d, 0xe3, 0x68, 0xf3, 0x03, 0x00, 0xf9, 0x2b, 0x3b,
	0xb5, 0xf2, 0xe9, 0xf6, 0xb3, 0x7d, 0x5a, 0xc9, 0xd9, 0xa0, 0x20, 0x5f, 0xc3, 0x65, 0xe6, 0x87,
	0xcb, 0x92, 0x90, 0xcd, 0xc4, 0xca, 0xa4, 0x6d, 0xf0, 0x3f, 0x0e, 0xe0, 0x77, 0x54, 0x7c, 0x3e,
	0xdd, 0x60, 0x39, 0x89, 0x83, 0xf1, 0xd1, 0x86, 0xde, 0x00, 0x9e, 0xae, 0xbf, 0x23, 0x9e, 0x48,
	0x58, 0x30, 0x4d, 0xdb, 0x52, 0x3a, 0x2f, 0x4e, 0xb5, 0xd5, 0xdf, 0xb9, 0x43, 0xf7, 0x58, 0x6e,
	0xe6, 0xd0, 0x38, 0xf8, 0xbd, 0x20, 0xcf, 0xa0, 0x61, 0x5a, 0xf7, 0xc6, 0xb0, 0xe7, 0x8f, 0xac,
	0x37, 0xd4, 0x7e, 0xf9, 0xca, 0xf7, 0x46, 0xfe, 0x5b, 0xd7, 0xc2, 0x8f, 0x48, 0x1d, 0x14, 0x67,
	0xb0, 0x46, 0x31, 0x90, 0x6b, 0xc0, 0xf7, 0x03, 0xfa, 0xa3, 0x41, 0xcd, 0x02, 0x55, 0x48, 0x03,
	0xae, 0xee, 0x8c, 0xee, 0xeb, 0x6d, 0xb8, 0x7a, 0xf3, 0x33, 0xa8, 0xc7, 0xfe, 0x16, 0x44, 0x85,
	0xeb, 0xbc, 0x68, 0x77, 0xe0, 0xf8, 0x74, 0xd0, 0xdb, 0xa8, 0xd9, 0x35, 0x9c, 0xae, 0xd5, 0x1b,
	0xf9, 0x86, 0xf7, 0x1a, 0x23, 0xc9, 0xee, 0x59, 0xfe, 0xc8, 0xed, 0x19, 0x6f, 0x65, 0x95, 0x11,
	0x35, 0x7c, 0x0b, 0x9f, 0xdd, 0x7c, 0x44, 0xf0, 0xc5, 0x01, 0x37, 0x09, 0x81, 0x5a, 0xce, 0xec,
	0xf9, 0x86, 0x3f, 0xf4, 0xf0, 0x23, 0x82, 0xa1, 0x6a, 0x3b, 0xb6, 0x6f, 0x1b, 0x3d, 0xfb, 0x27,
	0xdb, 0x79, 0x89, 0x15, 0x72, 0x09, 0x15, 0x97, 0x5a, 0xae, 0x41, 0x65, 0x58, 0x25, 0x55, 0x28,
	0x67, 0xa1, 0x65, 0xe2, 0x4b, 0xa2, 0xc0, 0x13, 0x3a, 0x74, 0x1c, 0x99, 0xaa, 0xc9, 0x93, 0x86,
	0xeb, 0x5a, 0x8e, 0x29, 0xc3, 0xba, 0x3c, 0x99, 0xc9, 0xb3, 0x4c, 0x8c, 0x65, 0xb2, 0x3b, 0xe8,
	0xbb, 0x3d, 0xcb, 0xb7, 0x4c, 0x7c, 0x45, 0x00, 0x2e, 0xee, 0x0d, 0x5b, 0xa6, 0xc8, 0x5d, 0x07,
	0x4e, 0x3c, 0xd8, 0x77, 0xf5, 0xe2, 0x79, 0x77, 0xe5, 0x1b, 0xfb, 0x0a, 0xfd, 0x89, 0xd0, 0xbb,
	0x8b, 0xf4, 0xbd, 0xfd, 0xee, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x5b, 0x01, 0xfd, 0xfd,
	0x07, 0x00, 0x00,
}
