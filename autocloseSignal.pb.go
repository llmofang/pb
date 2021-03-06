// Code generated by protoc-gen-go.
// source: autocloseSignal.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	autocloseSignal.proto

It has these top-level messages:
	AutoCloseSignal
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

type AutoCloseSignal struct {
	Sym              *string  `protobuf:"bytes,1,req,name=Sym" json:"Sym,omitempty"`
	Time             *int32   `protobuf:"varint,2,req,name=Time" json:"Time,omitempty"`
	Name             *string  `protobuf:"bytes,3,req,name=Name" json:"Name,omitempty"`
	Signaltype       *string  `protobuf:"bytes,4,req,name=Signaltype" json:"Signaltype,omitempty"`
	Signal           *float32 `protobuf:"fixed32,5,req,name=Signal" json:"Signal,omitempty"`
	Strength         *float32 `protobuf:"fixed32,6,req,name=Strength" json:"Strength,omitempty"`
	Note             *string  `protobuf:"bytes,7,req,name=Note" json:"Note,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AutoCloseSignal) Reset()                    { *m = AutoCloseSignal{} }
func (m *AutoCloseSignal) String() string            { return proto.CompactTextString(m) }
func (*AutoCloseSignal) ProtoMessage()               {}
func (*AutoCloseSignal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AutoCloseSignal) GetSym() string {
	if m != nil && m.Sym != nil {
		return *m.Sym
	}
	return ""
}

func (m *AutoCloseSignal) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *AutoCloseSignal) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AutoCloseSignal) GetSignaltype() string {
	if m != nil && m.Signaltype != nil {
		return *m.Signaltype
	}
	return ""
}

func (m *AutoCloseSignal) GetSignal() float32 {
	if m != nil && m.Signal != nil {
		return *m.Signal
	}
	return 0
}

func (m *AutoCloseSignal) GetStrength() float32 {
	if m != nil && m.Strength != nil {
		return *m.Strength
	}
	return 0
}

func (m *AutoCloseSignal) GetNote() string {
	if m != nil && m.Note != nil {
		return *m.Note
	}
	return ""
}

func init() {
	proto.RegisterType((*AutoCloseSignal)(nil), "pb.AutoCloseSignal")
}

