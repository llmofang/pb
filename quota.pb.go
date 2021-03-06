// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quota.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	quota.proto

It has these top-level messages:
	Quota
*/
package pb

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

type Quota struct {
	ID               *int64 `protobuf:"varint,1,req,name=ID" json:"ID,omitempty"`
	BuyVol           *int64 `protobuf:"varint,2,req,name=BuyVol" json:"BuyVol,omitempty"`
	SellVol          *int64 `protobuf:"varint,3,req,name=SellVol" json:"SellVol,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Quota) Reset()                    { *m = Quota{} }
func (m *Quota) String() string            { return proto.CompactTextString(m) }
func (*Quota) ProtoMessage()               {}
func (*Quota) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Quota) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Quota) GetBuyVol() int64 {
	if m != nil && m.BuyVol != nil {
		return *m.BuyVol
	}
	return 0
}

func (m *Quota) GetSellVol() int64 {
	if m != nil && m.SellVol != nil {
		return *m.SellVol
	}
	return 0
}

func init() {
	proto.RegisterType((*Quota)(nil), "pb.quota")
}

func init() {  }

