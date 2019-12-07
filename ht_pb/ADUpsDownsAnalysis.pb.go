// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ADUpsDownsAnalysis.proto

package ht_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//涨跌分析
type ADUpsDownsAnalysis struct {
	HTSCSecurityID          string                       `protobuf:"bytes,1,opt,name=HTSCSecurityID,proto3" json:"HTSCSecurityID,omitempty"`
	MDDate                  int32                        `protobuf:"varint,2,opt,name=MDDate,proto3" json:"MDDate,omitempty"`
	MDTime                  int32                        `protobuf:"varint,3,opt,name=MDTime,proto3" json:"MDTime,omitempty"`
	DataTimestamp           int64                        `protobuf:"varint,4,opt,name=DataTimestamp,proto3" json:"DataTimestamp,omitempty"`
	SecurityIDSource        ESecurityIDSource            `protobuf:"varint,5,opt,name=securityIDSource,proto3,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType            ESecurityType                `protobuf:"varint,6,opt,name=securityType,proto3,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	UpsDownsCount           *ADUpsDownsCount             `protobuf:"bytes,7,opt,name=UpsDownsCount,proto3" json:"UpsDownsCount,omitempty"`
	UpsDownsLimitCount      *ADUpsDownsLimitCount        `protobuf:"bytes,8,opt,name=UpsDownsLimitCount,proto3" json:"UpsDownsLimitCount,omitempty"`
	UpsDownsPartitionDetail []*ADUpsDownsPartitionDetail `protobuf:"bytes,9,rep,name=UpsDownsPartitionDetail,proto3" json:"UpsDownsPartitionDetail,omitempty"`
	ExchangeDate            int32                        `protobuf:"varint,10,opt,name=ExchangeDate,proto3" json:"ExchangeDate,omitempty"`
	ExchangeTime            int32                        `protobuf:"varint,11,opt,name=ExchangeTime,proto3" json:"ExchangeTime,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                     `json:"-"`
	XXX_unrecognized        []byte                       `json:"-"`
	XXX_sizecache           int32                        `json:"-"`
}

func (m *ADUpsDownsAnalysis) Reset()         { *m = ADUpsDownsAnalysis{} }
func (m *ADUpsDownsAnalysis) String() string { return proto.CompactTextString(m) }
func (*ADUpsDownsAnalysis) ProtoMessage()    {}
func (*ADUpsDownsAnalysis) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee0c61378c34e23, []int{0}
}

func (m *ADUpsDownsAnalysis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADUpsDownsAnalysis.Unmarshal(m, b)
}
func (m *ADUpsDownsAnalysis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADUpsDownsAnalysis.Marshal(b, m, deterministic)
}
func (m *ADUpsDownsAnalysis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADUpsDownsAnalysis.Merge(m, src)
}
func (m *ADUpsDownsAnalysis) XXX_Size() int {
	return xxx_messageInfo_ADUpsDownsAnalysis.Size(m)
}
func (m *ADUpsDownsAnalysis) XXX_DiscardUnknown() {
	xxx_messageInfo_ADUpsDownsAnalysis.DiscardUnknown(m)
}

var xxx_messageInfo_ADUpsDownsAnalysis proto.InternalMessageInfo

func (m *ADUpsDownsAnalysis) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *ADUpsDownsAnalysis) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *ADUpsDownsAnalysis) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *ADUpsDownsAnalysis) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *ADUpsDownsAnalysis) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *ADUpsDownsAnalysis) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *ADUpsDownsAnalysis) GetUpsDownsCount() *ADUpsDownsCount {
	if m != nil {
		return m.UpsDownsCount
	}
	return nil
}

func (m *ADUpsDownsAnalysis) GetUpsDownsLimitCount() *ADUpsDownsLimitCount {
	if m != nil {
		return m.UpsDownsLimitCount
	}
	return nil
}

func (m *ADUpsDownsAnalysis) GetUpsDownsPartitionDetail() []*ADUpsDownsPartitionDetail {
	if m != nil {
		return m.UpsDownsPartitionDetail
	}
	return nil
}

func (m *ADUpsDownsAnalysis) GetExchangeDate() int32 {
	if m != nil {
		return m.ExchangeDate
	}
	return 0
}

func (m *ADUpsDownsAnalysis) GetExchangeTime() int32 {
	if m != nil {
		return m.ExchangeTime
	}
	return 0
}

//今日涨跌数据
type ADUpsDownsCount struct {
	Ups                  int32    `protobuf:"varint,1,opt,name=Ups,proto3" json:"Ups,omitempty"`
	Downs                int32    `protobuf:"varint,2,opt,name=Downs,proto3" json:"Downs,omitempty"`
	Equals               int32    `protobuf:"varint,3,opt,name=Equals,proto3" json:"Equals,omitempty"`
	PreUps               int32    `protobuf:"varint,4,opt,name=PreUps,proto3" json:"PreUps,omitempty"`
	PreDowns             int32    `protobuf:"varint,5,opt,name=PreDowns,proto3" json:"PreDowns,omitempty"`
	PreEquals            int32    `protobuf:"varint,6,opt,name=PreEquals,proto3" json:"PreEquals,omitempty"`
	UpsPercent           float64  `protobuf:"fixed64,7,opt,name=UpsPercent,proto3" json:"UpsPercent,omitempty"`
	PreUpsPercent        float64  `protobuf:"fixed64,8,opt,name=PreUpsPercent,proto3" json:"PreUpsPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ADUpsDownsCount) Reset()         { *m = ADUpsDownsCount{} }
func (m *ADUpsDownsCount) String() string { return proto.CompactTextString(m) }
func (*ADUpsDownsCount) ProtoMessage()    {}
func (*ADUpsDownsCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee0c61378c34e23, []int{1}
}

func (m *ADUpsDownsCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADUpsDownsCount.Unmarshal(m, b)
}
func (m *ADUpsDownsCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADUpsDownsCount.Marshal(b, m, deterministic)
}
func (m *ADUpsDownsCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADUpsDownsCount.Merge(m, src)
}
func (m *ADUpsDownsCount) XXX_Size() int {
	return xxx_messageInfo_ADUpsDownsCount.Size(m)
}
func (m *ADUpsDownsCount) XXX_DiscardUnknown() {
	xxx_messageInfo_ADUpsDownsCount.DiscardUnknown(m)
}

var xxx_messageInfo_ADUpsDownsCount proto.InternalMessageInfo

func (m *ADUpsDownsCount) GetUps() int32 {
	if m != nil {
		return m.Ups
	}
	return 0
}

func (m *ADUpsDownsCount) GetDowns() int32 {
	if m != nil {
		return m.Downs
	}
	return 0
}

func (m *ADUpsDownsCount) GetEquals() int32 {
	if m != nil {
		return m.Equals
	}
	return 0
}

func (m *ADUpsDownsCount) GetPreUps() int32 {
	if m != nil {
		return m.PreUps
	}
	return 0
}

func (m *ADUpsDownsCount) GetPreDowns() int32 {
	if m != nil {
		return m.PreDowns
	}
	return 0
}

func (m *ADUpsDownsCount) GetPreEquals() int32 {
	if m != nil {
		return m.PreEquals
	}
	return 0
}

func (m *ADUpsDownsCount) GetUpsPercent() float64 {
	if m != nil {
		return m.UpsPercent
	}
	return 0
}

func (m *ADUpsDownsCount) GetPreUpsPercent() float64 {
	if m != nil {
		return m.PreUpsPercent
	}
	return 0
}

//涨跌停数据
type ADUpsDownsLimitCount struct {
	NoReachedLimitPx                int32    `protobuf:"varint,1,opt,name=NoReachedLimitPx,proto3" json:"NoReachedLimitPx,omitempty"`
	UpLimits                        int32    `protobuf:"varint,2,opt,name=UpLimits,proto3" json:"UpLimits,omitempty"`
	DownLimits                      int32    `protobuf:"varint,3,opt,name=DownLimits,proto3" json:"DownLimits,omitempty"`
	PreNoReachedLimitPx             int32    `protobuf:"varint,4,opt,name=PreNoReachedLimitPx,proto3" json:"PreNoReachedLimitPx,omitempty"`
	PreUpLimits                     int32    `protobuf:"varint,5,opt,name=PreUpLimits,proto3" json:"PreUpLimits,omitempty"`
	PreDownLimits                   int32    `protobuf:"varint,6,opt,name=PreDownLimits,proto3" json:"PreDownLimits,omitempty"`
	PreUpLimitsAverageChangePercent float64  `protobuf:"fixed64,7,opt,name=PreUpLimitsAverageChangePercent,proto3" json:"PreUpLimitsAverageChangePercent,omitempty"`
	UpLimitsPercent                 float64  `protobuf:"fixed64,8,opt,name=UpLimitsPercent,proto3" json:"UpLimitsPercent,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *ADUpsDownsLimitCount) Reset()         { *m = ADUpsDownsLimitCount{} }
func (m *ADUpsDownsLimitCount) String() string { return proto.CompactTextString(m) }
func (*ADUpsDownsLimitCount) ProtoMessage()    {}
func (*ADUpsDownsLimitCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee0c61378c34e23, []int{2}
}

func (m *ADUpsDownsLimitCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADUpsDownsLimitCount.Unmarshal(m, b)
}
func (m *ADUpsDownsLimitCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADUpsDownsLimitCount.Marshal(b, m, deterministic)
}
func (m *ADUpsDownsLimitCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADUpsDownsLimitCount.Merge(m, src)
}
func (m *ADUpsDownsLimitCount) XXX_Size() int {
	return xxx_messageInfo_ADUpsDownsLimitCount.Size(m)
}
func (m *ADUpsDownsLimitCount) XXX_DiscardUnknown() {
	xxx_messageInfo_ADUpsDownsLimitCount.DiscardUnknown(m)
}

var xxx_messageInfo_ADUpsDownsLimitCount proto.InternalMessageInfo

func (m *ADUpsDownsLimitCount) GetNoReachedLimitPx() int32 {
	if m != nil {
		return m.NoReachedLimitPx
	}
	return 0
}

func (m *ADUpsDownsLimitCount) GetUpLimits() int32 {
	if m != nil {
		return m.UpLimits
	}
	return 0
}

func (m *ADUpsDownsLimitCount) GetDownLimits() int32 {
	if m != nil {
		return m.DownLimits
	}
	return 0
}

func (m *ADUpsDownsLimitCount) GetPreNoReachedLimitPx() int32 {
	if m != nil {
		return m.PreNoReachedLimitPx
	}
	return 0
}

func (m *ADUpsDownsLimitCount) GetPreUpLimits() int32 {
	if m != nil {
		return m.PreUpLimits
	}
	return 0
}

func (m *ADUpsDownsLimitCount) GetPreDownLimits() int32 {
	if m != nil {
		return m.PreDownLimits
	}
	return 0
}

func (m *ADUpsDownsLimitCount) GetPreUpLimitsAverageChangePercent() float64 {
	if m != nil {
		return m.PreUpLimitsAverageChangePercent
	}
	return 0
}

func (m *ADUpsDownsLimitCount) GetUpLimitsPercent() float64 {
	if m != nil {
		return m.UpLimitsPercent
	}
	return 0
}

//涨跌分布明细
type ADUpsDownsPartitionDetail struct {
	Numbers                int32    `protobuf:"varint,1,opt,name=Numbers,proto3" json:"Numbers,omitempty"`
	PartitionChangePercent int32    `protobuf:"varint,2,opt,name=PartitionChangePercent,proto3" json:"PartitionChangePercent,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ADUpsDownsPartitionDetail) Reset()         { *m = ADUpsDownsPartitionDetail{} }
func (m *ADUpsDownsPartitionDetail) String() string { return proto.CompactTextString(m) }
func (*ADUpsDownsPartitionDetail) ProtoMessage()    {}
func (*ADUpsDownsPartitionDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee0c61378c34e23, []int{3}
}

func (m *ADUpsDownsPartitionDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADUpsDownsPartitionDetail.Unmarshal(m, b)
}
func (m *ADUpsDownsPartitionDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADUpsDownsPartitionDetail.Marshal(b, m, deterministic)
}
func (m *ADUpsDownsPartitionDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADUpsDownsPartitionDetail.Merge(m, src)
}
func (m *ADUpsDownsPartitionDetail) XXX_Size() int {
	return xxx_messageInfo_ADUpsDownsPartitionDetail.Size(m)
}
func (m *ADUpsDownsPartitionDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ADUpsDownsPartitionDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ADUpsDownsPartitionDetail proto.InternalMessageInfo

func (m *ADUpsDownsPartitionDetail) GetNumbers() int32 {
	if m != nil {
		return m.Numbers
	}
	return 0
}

func (m *ADUpsDownsPartitionDetail) GetPartitionChangePercent() int32 {
	if m != nil {
		return m.PartitionChangePercent
	}
	return 0
}

func init() {
	proto.RegisterType((*ADUpsDownsAnalysis)(nil), "com.htsc.mdc.insight.model.ADUpsDownsAnalysis")
	proto.RegisterType((*ADUpsDownsCount)(nil), "com.htsc.mdc.insight.model.ADUpsDownsCount")
	proto.RegisterType((*ADUpsDownsLimitCount)(nil), "com.htsc.mdc.insight.model.ADUpsDownsLimitCount")
	proto.RegisterType((*ADUpsDownsPartitionDetail)(nil), "com.htsc.mdc.insight.model.ADUpsDownsPartitionDetail")
}

func init() { proto.RegisterFile("ADUpsDownsAnalysis.proto", fileDescriptor_bee0c61378c34e23) }

var fileDescriptor_bee0c61378c34e23 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x95, 0x9b, 0x3a, 0x4d, 0x26, 0xfd, 0x89, 0xa6, 0x55, 0xeb, 0x2f, 0xfa, 0x04, 0x26, 0x02,
	0x64, 0x81, 0x64, 0x55, 0x41, 0x70, 0xc3, 0x55, 0x1a, 0x47, 0x0a, 0x12, 0x54, 0xae, 0x93, 0xdc,
	0xe3, 0x3a, 0xab, 0xc4, 0x52, 0xfc, 0x83, 0x77, 0x03, 0xcd, 0xab, 0x70, 0xc5, 0xeb, 0x21, 0xf1,
	0x10, 0x68, 0x77, 0xed, 0xc4, 0x9b, 0x9f, 0xd2, 0x3b, 0x9f, 0x33, 0x33, 0x67, 0x77, 0x8e, 0x8f,
	0x16, 0x8c, 0xae, 0x33, 0x4e, 0xa9, 0x93, 0xfc, 0x88, 0x69, 0x37, 0xf6, 0xe7, 0x4b, 0x1a, 0x52,
	0x3b, 0xcd, 0x12, 0x96, 0x60, 0x2b, 0x48, 0x22, 0x7b, 0xc6, 0x68, 0x60, 0x47, 0x93, 0xc0, 0x0e,
	0x63, 0x1a, 0x4e, 0x67, 0xcc, 0x8e, 0x92, 0x09, 0x99, 0xb7, 0xce, 0xfb, 0x43, 0x12, 0x2c, 0xb2,
	0x90, 0x2d, 0x47, 0xcb, 0x94, 0xc8, 0x81, 0xd6, 0xd5, 0x8a, 0xfc, 0xe4, 0x0c, 0x93, 0x45, 0x16,
	0xe4, 0x85, 0xf6, 0x4f, 0x1d, 0x70, 0xfb, 0x18, 0x7c, 0x0d, 0xa7, 0x83, 0xd1, 0xb0, 0xb7, 0x1e,
	0x32, 0x34, 0x53, 0xb3, 0xea, 0xde, 0x06, 0x8b, 0x97, 0x50, 0xfd, 0xe2, 0x38, 0x3e, 0x23, 0xc6,
	0x81, 0xa9, 0x59, 0xba, 0x97, 0x23, 0xc9, 0x8f, 0xc2, 0x88, 0x18, 0x95, 0x82, 0xe7, 0x08, 0x5f,
	0xc2, 0x89, 0xe3, 0x33, 0x9f, 0x7f, 0x53, 0xe6, 0x47, 0xa9, 0x71, 0x68, 0x6a, 0x56, 0xc5, 0x53,
	0x49, 0xbc, 0x83, 0x26, 0xdd, 0xb8, 0xae, 0xa1, 0x9b, 0x9a, 0x75, 0xda, 0x79, 0x65, 0x2b, 0x9b,
	0x8b, 0x8d, 0xed, 0xad, 0xdd, 0xbc, 0xad, 0x71, 0xec, 0xc3, 0x31, 0x2d, 0xd9, 0x62, 0x54, 0x85,
	0xdc, 0x8b, 0x47, 0xe5, 0x78, 0xa3, 0xa7, 0x8c, 0xe1, 0x1d, 0x9c, 0x14, 0x5e, 0xf5, 0x92, 0x45,
	0xcc, 0x8c, 0x23, 0x53, 0xb3, 0x1a, 0x9d, 0xb7, 0xf6, 0xfe, 0x1f, 0x62, 0xaf, 0xed, 0x15, 0x23,
	0x9e, 0xaa, 0x80, 0x5f, 0x01, 0x0b, 0xe2, 0x73, 0x18, 0x85, 0x4c, 0xea, 0xd6, 0x84, 0xee, 0xf5,
	0xd3, 0x74, 0xd7, 0x73, 0xde, 0x0e, 0x2d, 0x4c, 0xe0, 0xaa, 0x60, 0x5d, 0x3f, 0x63, 0x21, 0x0b,
	0x93, 0xd8, 0x21, 0xcc, 0x0f, 0xe7, 0x46, 0xdd, 0xac, 0x58, 0x8d, 0xce, 0xfb, 0xa7, 0x1d, 0xb3,
	0x31, 0xec, 0xed, 0x53, 0xc5, 0x36, 0x1c, 0xf7, 0x1f, 0x82, 0x99, 0x1f, 0x4f, 0x89, 0xc8, 0x06,
	0x88, 0x0c, 0x28, 0x5c, 0xb9, 0x47, 0xe4, 0xa4, 0xa1, 0xf6, 0x70, 0xae, 0xfd, 0x5b, 0x83, 0xb3,
	0x0d, 0xf7, 0xb0, 0x09, 0x95, 0x71, 0x4a, 0x45, 0x1c, 0x75, 0x8f, 0x7f, 0xe2, 0x05, 0xe8, 0xa2,
	0x9e, 0x47, 0x50, 0x02, 0x9e, 0xc0, 0xfe, 0xb7, 0x85, 0x3f, 0xa7, 0x45, 0x02, 0x25, 0xe2, 0xbc,
	0x9b, 0x11, 0x2e, 0x71, 0x28, 0x79, 0x89, 0xb0, 0x05, 0x35, 0x37, 0x23, 0x52, 0x48, 0x17, 0x95,
	0x15, 0xc6, 0xff, 0xa1, 0xee, 0x66, 0x24, 0x97, 0xab, 0x8a, 0xe2, 0x9a, 0xc0, 0x67, 0x00, 0xe3,
	0x94, 0xba, 0x24, 0x0b, 0x48, 0x1e, 0x08, 0xcd, 0x2b, 0x31, 0x3c, 0xf3, 0xf2, 0x8c, 0xa2, 0xa5,
	0x26, 0x5a, 0x54, 0xb2, 0xfd, 0xe7, 0x00, 0x2e, 0x76, 0xfd, 0x51, 0x7c, 0x03, 0xcd, 0xdb, 0xc4,
	0x23, 0x7e, 0x30, 0x23, 0x13, 0x41, 0xbb, 0x0f, 0xf9, 0xf6, 0x5b, 0x3c, 0x5f, 0x62, 0x9c, 0x0a,
	0x50, 0xb8, 0xb1, 0xc2, 0xfc, 0x9a, 0x5c, 0x3a, 0xaf, 0x4a, 0x53, 0x4a, 0x0c, 0x5e, 0xc3, 0xb9,
	0x9b, 0x91, 0xad, 0xa3, 0xa4, 0x4b, 0xbb, 0x4a, 0x68, 0x42, 0x43, 0xec, 0x90, 0x4b, 0x4a, 0xd7,
	0xca, 0x54, 0xbe, 0x7a, 0xe9, 0x58, 0x69, 0x9e, 0x4a, 0xe2, 0x00, 0x9e, 0x97, 0x86, 0xba, 0xdf,
	0x49, 0xe6, 0x4f, 0x49, 0x4f, 0xc4, 0x40, 0x75, 0xf5, 0x5f, 0x6d, 0x68, 0xc1, 0x59, 0x51, 0x57,
	0xcd, 0xde, 0xa4, 0xdb, 0x11, 0xfc, 0xb7, 0x37, 0xd8, 0x68, 0xc0, 0xd1, 0xed, 0x22, 0xba, 0x27,
	0x59, 0x91, 0xb3, 0x02, 0xe2, 0x07, 0xb8, 0x5c, 0x35, 0xab, 0x37, 0x94, 0x76, 0xef, 0xa9, 0xde,
	0x7c, 0x84, 0x47, 0x9e, 0xec, 0x9b, 0x1d, 0x0f, 0xbd, 0xcb, 0x5f, 0x67, 0x3a, 0xd0, 0x7e, 0x69,
	0xda, 0x7d, 0x55, 0x3c, 0xd5, 0xef, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xff, 0x1c, 0x2d, 0xf2,
	0x10, 0x06, 0x00, 0x00,
}
