// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketData.proto

package comm

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

type MarketData struct {
	HTSCSecurityID             *string  `protobuf:"bytes,1,req,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	MDDate                     *int32   `protobuf:"varint,2,req,name=MDDate" json:"MDDate,omitempty"`
	MDTime                     *int32   `protobuf:"varint,3,req,name=MDTime" json:"MDTime,omitempty"`
	TradingPhaseCode           *string  `protobuf:"bytes,4,opt,name=TradingPhaseCode" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource           *string  `protobuf:"bytes,5,opt,name=securityIDSource" json:"securityIDSource,omitempty"`
	SecurityType               *string  `protobuf:"bytes,6,opt,name=securityType" json:"securityType,omitempty"`
	MaxPx                      *int64   `protobuf:"varint,7,opt,name=MaxPx" json:"MaxPx,omitempty"`
	MinPx                      *int64   `protobuf:"varint,8,opt,name=MinPx" json:"MinPx,omitempty"`
	PreClosePx                 *int64   `protobuf:"varint,9,opt,name=PreClosePx" json:"PreClosePx,omitempty"`
	NumTrades                  *int64   `protobuf:"varint,10,opt,name=NumTrades" json:"NumTrades,omitempty"`
	TotalVolumeTrade           *int64   `protobuf:"varint,11,opt,name=TotalVolumeTrade" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade            *int64   `protobuf:"varint,12,opt,name=TotalValueTrade" json:"TotalValueTrade,omitempty"`
	LastPx                     *int64   `protobuf:"varint,13,opt,name=LastPx" json:"LastPx,omitempty"`
	OpenPx                     *int64   `protobuf:"varint,14,opt,name=OpenPx" json:"OpenPx,omitempty"`
	ClosePx                    *int64   `protobuf:"varint,15,opt,name=ClosePx" json:"ClosePx,omitempty"`
	HighPx                     *int64   `protobuf:"varint,16,opt,name=HighPx" json:"HighPx,omitempty"`
	LowPx                      *int64   `protobuf:"varint,17,opt,name=LowPx" json:"LowPx,omitempty"`
	DiffPx1                    *int64   `protobuf:"varint,18,opt,name=DiffPx1" json:"DiffPx1,omitempty"`
	DiffPx2                    *int64   `protobuf:"varint,19,opt,name=DiffPx2" json:"DiffPx2,omitempty"`
	TotalBuyQty                *int64   `protobuf:"varint,20,opt,name=TotalBuyQty" json:"TotalBuyQty,omitempty"`
	TotalSellQty               *int64   `protobuf:"varint,21,opt,name=TotalSellQty" json:"TotalSellQty,omitempty"`
	WeightedAvgBuyPx           *int64   `protobuf:"varint,22,opt,name=WeightedAvgBuyPx" json:"WeightedAvgBuyPx,omitempty"`
	WeightedAvgSellPx          *int64   `protobuf:"varint,23,opt,name=WeightedAvgSellPx" json:"WeightedAvgSellPx,omitempty"`
	WithdrawBuyNumber          *int64   `protobuf:"varint,24,opt,name=WithdrawBuyNumber" json:"WithdrawBuyNumber,omitempty"`
	WithdrawBuyAmount          *int64   `protobuf:"varint,25,opt,name=WithdrawBuyAmount" json:"WithdrawBuyAmount,omitempty"`
	WithdrawBuyMoney           *int64   `protobuf:"varint,26,opt,name=WithdrawBuyMoney" json:"WithdrawBuyMoney,omitempty"`
	WithdrawSellNumber         *int64   `protobuf:"varint,27,opt,name=WithdrawSellNumber" json:"WithdrawSellNumber,omitempty"`
	WithdrawSellAmount         *int64   `protobuf:"varint,28,opt,name=WithdrawSellAmount" json:"WithdrawSellAmount,omitempty"`
	WithdrawSellMoney          *int64   `protobuf:"varint,29,opt,name=WithdrawSellMoney" json:"WithdrawSellMoney,omitempty"`
	TotalBuyNumber             *int64   `protobuf:"varint,30,opt,name=TotalBuyNumber" json:"TotalBuyNumber,omitempty"`
	TotalSellNumber            *int64   `protobuf:"varint,31,opt,name=TotalSellNumber" json:"TotalSellNumber,omitempty"`
	BuyTradeMaxDuration        *int64   `protobuf:"varint,32,opt,name=BuyTradeMaxDuration" json:"BuyTradeMaxDuration,omitempty"`
	SellTradeMaxDuration       *int64   `protobuf:"varint,33,opt,name=SellTradeMaxDuration" json:"SellTradeMaxDuration,omitempty"`
	NumBuyOrders               *int32   `protobuf:"varint,34,opt,name=NumBuyOrders" json:"NumBuyOrders,omitempty"`
	NumSellOrders              *int32   `protobuf:"varint,35,opt,name=NumSellOrders" json:"NumSellOrders,omitempty"`
	NorminalPx                 *int64   `protobuf:"varint,36,opt,name=NorminalPx" json:"NorminalPx,omitempty"`
	ShortSellSharesTraded      *int64   `protobuf:"varint,37,opt,name=ShortSellSharesTraded" json:"ShortSellSharesTraded,omitempty"`
	ShortSellTurnover          *int64   `protobuf:"varint,38,opt,name=ShortSellTurnover" json:"ShortSellTurnover,omitempty"`
	ReferencePx                *int64   `protobuf:"varint,39,opt,name=ReferencePx" json:"ReferencePx,omitempty"`
	ComplexEventStartTime      *int64   `protobuf:"varint,40,opt,name=ComplexEventStartTime" json:"ComplexEventStartTime,omitempty"`
	ComplexEventEndTime        *int64   `protobuf:"varint,41,opt,name=ComplexEventEndTime" json:"ComplexEventEndTime,omitempty"`
	ExchangeDate               *int32   `protobuf:"varint,42,opt,name=ExchangeDate" json:"ExchangeDate,omitempty"`
	ExchangeTime               *int32   `protobuf:"varint,43,opt,name=ExchangeTime" json:"ExchangeTime,omitempty"`
	AfterHoursNumTrades        *int64   `protobuf:"varint,44,opt,name=AfterHoursNumTrades" json:"AfterHoursNumTrades,omitempty"`
	AfterHoursTotalVolumeTrade *int64   `protobuf:"varint,45,opt,name=AfterHoursTotalVolumeTrade" json:"AfterHoursTotalVolumeTrade,omitempty"`
	AfterHoursTotalValueTrade  *int64   `protobuf:"varint,46,opt,name=AfterHoursTotalValueTrade" json:"AfterHoursTotalValueTrade,omitempty"`
	ChannelNo                  *int32   `protobuf:"varint,47,opt,name=ChannelNo" json:"ChannelNo,omitempty"`
	BuyPriceQueue              []int64  `protobuf:"varint,48,rep,name=BuyPriceQueue" json:"BuyPriceQueue,omitempty"`
	BuyOrderQtyQueue           []int64  `protobuf:"varint,49,rep,name=BuyOrderQtyQueue" json:"BuyOrderQtyQueue,omitempty"`
	SellPriceQueue             []int64  `protobuf:"varint,50,rep,name=SellPriceQueue" json:"SellPriceQueue,omitempty"`
	SellOrderQtyQueue          []int64  `protobuf:"varint,51,rep,name=SellOrderQtyQueue" json:"SellOrderQtyQueue,omitempty"`
	BuyOrderQueue              []int64  `protobuf:"varint,52,rep,name=BuyOrderQueue" json:"BuyOrderQueue,omitempty"`
	SellOrderQueue             []int64  `protobuf:"varint,53,rep,name=SellOrderQueue" json:"SellOrderQueue,omitempty"`
	BuyNumOrdersQueue          []int64  `protobuf:"varint,54,rep,name=BuyNumOrdersQueue" json:"BuyNumOrdersQueue,omitempty"`
	SellNumOrdersQueue         []int64  `protobuf:"varint,55,rep,name=SellNumOrdersQueue" json:"SellNumOrdersQueue,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *MarketData) Reset()         { *m = MarketData{} }
func (m *MarketData) String() string { return proto.CompactTextString(m) }
func (*MarketData) ProtoMessage()    {}
func (*MarketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cc15ad5492c5deb, []int{0}
}

func (m *MarketData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketData.Unmarshal(m, b)
}
func (m *MarketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketData.Marshal(b, m, deterministic)
}
func (m *MarketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketData.Merge(m, src)
}
func (m *MarketData) XXX_Size() int {
	return xxx_messageInfo_MarketData.Size(m)
}
func (m *MarketData) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketData.DiscardUnknown(m)
}

var xxx_messageInfo_MarketData proto.InternalMessageInfo

func (m *MarketData) GetHTSCSecurityID() string {
	if m != nil && m.HTSCSecurityID != nil {
		return *m.HTSCSecurityID
	}
	return ""
}

func (m *MarketData) GetMDDate() int32 {
	if m != nil && m.MDDate != nil {
		return *m.MDDate
	}
	return 0
}

func (m *MarketData) GetMDTime() int32 {
	if m != nil && m.MDTime != nil {
		return *m.MDTime
	}
	return 0
}

func (m *MarketData) GetTradingPhaseCode() string {
	if m != nil && m.TradingPhaseCode != nil {
		return *m.TradingPhaseCode
	}
	return ""
}

func (m *MarketData) GetSecurityIDSource() string {
	if m != nil && m.SecurityIDSource != nil {
		return *m.SecurityIDSource
	}
	return ""
}

func (m *MarketData) GetSecurityType() string {
	if m != nil && m.SecurityType != nil {
		return *m.SecurityType
	}
	return ""
}

func (m *MarketData) GetMaxPx() int64 {
	if m != nil && m.MaxPx != nil {
		return *m.MaxPx
	}
	return 0
}

func (m *MarketData) GetMinPx() int64 {
	if m != nil && m.MinPx != nil {
		return *m.MinPx
	}
	return 0
}

func (m *MarketData) GetPreClosePx() int64 {
	if m != nil && m.PreClosePx != nil {
		return *m.PreClosePx
	}
	return 0
}

func (m *MarketData) GetNumTrades() int64 {
	if m != nil && m.NumTrades != nil {
		return *m.NumTrades
	}
	return 0
}

func (m *MarketData) GetTotalVolumeTrade() int64 {
	if m != nil && m.TotalVolumeTrade != nil {
		return *m.TotalVolumeTrade
	}
	return 0
}

func (m *MarketData) GetTotalValueTrade() int64 {
	if m != nil && m.TotalValueTrade != nil {
		return *m.TotalValueTrade
	}
	return 0
}

func (m *MarketData) GetLastPx() int64 {
	if m != nil && m.LastPx != nil {
		return *m.LastPx
	}
	return 0
}

func (m *MarketData) GetOpenPx() int64 {
	if m != nil && m.OpenPx != nil {
		return *m.OpenPx
	}
	return 0
}

func (m *MarketData) GetClosePx() int64 {
	if m != nil && m.ClosePx != nil {
		return *m.ClosePx
	}
	return 0
}

func (m *MarketData) GetHighPx() int64 {
	if m != nil && m.HighPx != nil {
		return *m.HighPx
	}
	return 0
}

func (m *MarketData) GetLowPx() int64 {
	if m != nil && m.LowPx != nil {
		return *m.LowPx
	}
	return 0
}

func (m *MarketData) GetDiffPx1() int64 {
	if m != nil && m.DiffPx1 != nil {
		return *m.DiffPx1
	}
	return 0
}

func (m *MarketData) GetDiffPx2() int64 {
	if m != nil && m.DiffPx2 != nil {
		return *m.DiffPx2
	}
	return 0
}

func (m *MarketData) GetTotalBuyQty() int64 {
	if m != nil && m.TotalBuyQty != nil {
		return *m.TotalBuyQty
	}
	return 0
}

func (m *MarketData) GetTotalSellQty() int64 {
	if m != nil && m.TotalSellQty != nil {
		return *m.TotalSellQty
	}
	return 0
}

func (m *MarketData) GetWeightedAvgBuyPx() int64 {
	if m != nil && m.WeightedAvgBuyPx != nil {
		return *m.WeightedAvgBuyPx
	}
	return 0
}

func (m *MarketData) GetWeightedAvgSellPx() int64 {
	if m != nil && m.WeightedAvgSellPx != nil {
		return *m.WeightedAvgSellPx
	}
	return 0
}

func (m *MarketData) GetWithdrawBuyNumber() int64 {
	if m != nil && m.WithdrawBuyNumber != nil {
		return *m.WithdrawBuyNumber
	}
	return 0
}

func (m *MarketData) GetWithdrawBuyAmount() int64 {
	if m != nil && m.WithdrawBuyAmount != nil {
		return *m.WithdrawBuyAmount
	}
	return 0
}

func (m *MarketData) GetWithdrawBuyMoney() int64 {
	if m != nil && m.WithdrawBuyMoney != nil {
		return *m.WithdrawBuyMoney
	}
	return 0
}

func (m *MarketData) GetWithdrawSellNumber() int64 {
	if m != nil && m.WithdrawSellNumber != nil {
		return *m.WithdrawSellNumber
	}
	return 0
}

func (m *MarketData) GetWithdrawSellAmount() int64 {
	if m != nil && m.WithdrawSellAmount != nil {
		return *m.WithdrawSellAmount
	}
	return 0
}

func (m *MarketData) GetWithdrawSellMoney() int64 {
	if m != nil && m.WithdrawSellMoney != nil {
		return *m.WithdrawSellMoney
	}
	return 0
}

func (m *MarketData) GetTotalBuyNumber() int64 {
	if m != nil && m.TotalBuyNumber != nil {
		return *m.TotalBuyNumber
	}
	return 0
}

func (m *MarketData) GetTotalSellNumber() int64 {
	if m != nil && m.TotalSellNumber != nil {
		return *m.TotalSellNumber
	}
	return 0
}

func (m *MarketData) GetBuyTradeMaxDuration() int64 {
	if m != nil && m.BuyTradeMaxDuration != nil {
		return *m.BuyTradeMaxDuration
	}
	return 0
}

func (m *MarketData) GetSellTradeMaxDuration() int64 {
	if m != nil && m.SellTradeMaxDuration != nil {
		return *m.SellTradeMaxDuration
	}
	return 0
}

func (m *MarketData) GetNumBuyOrders() int32 {
	if m != nil && m.NumBuyOrders != nil {
		return *m.NumBuyOrders
	}
	return 0
}

func (m *MarketData) GetNumSellOrders() int32 {
	if m != nil && m.NumSellOrders != nil {
		return *m.NumSellOrders
	}
	return 0
}

func (m *MarketData) GetNorminalPx() int64 {
	if m != nil && m.NorminalPx != nil {
		return *m.NorminalPx
	}
	return 0
}

func (m *MarketData) GetShortSellSharesTraded() int64 {
	if m != nil && m.ShortSellSharesTraded != nil {
		return *m.ShortSellSharesTraded
	}
	return 0
}

func (m *MarketData) GetShortSellTurnover() int64 {
	if m != nil && m.ShortSellTurnover != nil {
		return *m.ShortSellTurnover
	}
	return 0
}

func (m *MarketData) GetReferencePx() int64 {
	if m != nil && m.ReferencePx != nil {
		return *m.ReferencePx
	}
	return 0
}

func (m *MarketData) GetComplexEventStartTime() int64 {
	if m != nil && m.ComplexEventStartTime != nil {
		return *m.ComplexEventStartTime
	}
	return 0
}

func (m *MarketData) GetComplexEventEndTime() int64 {
	if m != nil && m.ComplexEventEndTime != nil {
		return *m.ComplexEventEndTime
	}
	return 0
}

func (m *MarketData) GetExchangeDate() int32 {
	if m != nil && m.ExchangeDate != nil {
		return *m.ExchangeDate
	}
	return 0
}

func (m *MarketData) GetExchangeTime() int32 {
	if m != nil && m.ExchangeTime != nil {
		return *m.ExchangeTime
	}
	return 0
}

func (m *MarketData) GetAfterHoursNumTrades() int64 {
	if m != nil && m.AfterHoursNumTrades != nil {
		return *m.AfterHoursNumTrades
	}
	return 0
}

func (m *MarketData) GetAfterHoursTotalVolumeTrade() int64 {
	if m != nil && m.AfterHoursTotalVolumeTrade != nil {
		return *m.AfterHoursTotalVolumeTrade
	}
	return 0
}

func (m *MarketData) GetAfterHoursTotalValueTrade() int64 {
	if m != nil && m.AfterHoursTotalValueTrade != nil {
		return *m.AfterHoursTotalValueTrade
	}
	return 0
}

func (m *MarketData) GetChannelNo() int32 {
	if m != nil && m.ChannelNo != nil {
		return *m.ChannelNo
	}
	return 0
}

func (m *MarketData) GetBuyPriceQueue() []int64 {
	if m != nil {
		return m.BuyPriceQueue
	}
	return nil
}

func (m *MarketData) GetBuyOrderQtyQueue() []int64 {
	if m != nil {
		return m.BuyOrderQtyQueue
	}
	return nil
}

func (m *MarketData) GetSellPriceQueue() []int64 {
	if m != nil {
		return m.SellPriceQueue
	}
	return nil
}

func (m *MarketData) GetSellOrderQtyQueue() []int64 {
	if m != nil {
		return m.SellOrderQtyQueue
	}
	return nil
}

func (m *MarketData) GetBuyOrderQueue() []int64 {
	if m != nil {
		return m.BuyOrderQueue
	}
	return nil
}

func (m *MarketData) GetSellOrderQueue() []int64 {
	if m != nil {
		return m.SellOrderQueue
	}
	return nil
}

func (m *MarketData) GetBuyNumOrdersQueue() []int64 {
	if m != nil {
		return m.BuyNumOrdersQueue
	}
	return nil
}

func (m *MarketData) GetSellNumOrdersQueue() []int64 {
	if m != nil {
		return m.SellNumOrdersQueue
	}
	return nil
}

func init() {
	proto.RegisterType((*MarketData)(nil), "comm.MarketData")
}

func init() { proto.RegisterFile("marketData.proto", fileDescriptor_4cc15ad5492c5deb) }

var fileDescriptor_4cc15ad5492c5deb = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x73, 0xdb, 0x44,
	0x14, 0xc5, 0x27, 0x49, 0xd3, 0x92, 0xed, 0x57, 0xb2, 0x4d, 0xcb, 0xb6, 0x94, 0x22, 0x42, 0x29,
	0xa2, 0x84, 0xd0, 0x86, 0x02, 0x2f, 0x0c, 0x33, 0xb1, 0x9d, 0x99, 0x30, 0x53, 0xbb, 0x8a, 0xed,
	0x81, 0xe7, 0xc5, 0xbe, 0xb1, 0x34, 0x48, 0xda, 0xcc, 0x6a, 0x37, 0x95, 0xfe, 0x77, 0x1e, 0x98,
	0x7b, 0x57, 0xb2, 0x64, 0x49, 0xf0, 0x78, 0x7f, 0xe7, 0x5c, 0xed, 0xd1, 0x7e, 0x5c, 0xb6, 0x9f,
	0x48, 0xfd, 0x37, 0x98, 0x91, 0x34, 0xf2, 0xe4, 0x5a, 0x2b, 0xa3, 0xf8, 0xad, 0x85, 0x4a, 0x92,
	0xa3, 0x7f, 0xf6, 0x19, 0x1b, 0xaf, 0x25, 0xfe, 0x8a, 0x3d, 0xb8, 0x98, 0xcf, 0x86, 0x33, 0x58,
	0x58, 0x1d, 0x99, 0xe2, 0xf7, 0x91, 0xd8, 0xf2, 0xb6, 0xfd, 0xbd, 0x69, 0x8b, 0xf2, 0x27, 0xec,
	0xf6, 0x78, 0x34, 0x92, 0x06, 0xc4, 0xb6, 0xb7, 0xed, 0xef, 0x4e, 0xcb, 0xca, 0xf1, 0x79, 0x94,
	0x80, 0xd8, 0xa9, 0x38, 0x56, 0xfc, 0x35, 0xdb, 0x9f, 0x6b, 0xb9, 0x8c, 0xd2, 0x55, 0x10, 0xca,
	0x0c, 0x86, 0x6a, 0x09, 0xe2, 0x96, 0xb7, 0xe5, 0xef, 0x4d, 0x3b, 0x1c, 0xbd, 0xd9, 0x7a, 0xa5,
	0x99, 0xb2, 0x7a, 0x01, 0x62, 0xd7, 0x79, 0xdb, 0x9c, 0x1f, 0xb1, 0x7b, 0x15, 0x9b, 0x17, 0xd7,
	0x20, 0x6e, 0x93, 0x6f, 0x83, 0xf1, 0x43, 0xb6, 0x3b, 0x96, 0x79, 0x90, 0x8b, 0x3b, 0xde, 0x96,
	0xbf, 0x33, 0x75, 0x05, 0xd1, 0x28, 0x0d, 0x72, 0xf1, 0x49, 0x49, 0xb1, 0xe0, 0x2f, 0x18, 0x0b,
	0x34, 0x0c, 0x63, 0x95, 0x41, 0x90, 0x8b, 0x3d, 0x92, 0x1a, 0x84, 0x3f, 0x67, 0x7b, 0x13, 0x9b,
	0x60, 0x64, 0xc8, 0x04, 0x23, 0xb9, 0x06, 0xf4, 0x97, 0xca, 0xc8, 0xf8, 0x0f, 0x15, 0xdb, 0x04,
	0x08, 0x8a, 0xbb, 0x64, 0xea, 0x70, 0xee, 0xb3, 0x87, 0x8e, 0xc9, 0xd8, 0x96, 0xd6, 0x7b, 0x64,
	0x6d, 0x63, 0xdc, 0xd3, 0xf7, 0x32, 0x33, 0x41, 0x2e, 0xee, 0x93, 0xa1, 0xac, 0x90, 0x7f, 0xb8,
	0x06, 0xfc, 0x85, 0x07, 0x8e, 0xbb, 0x8a, 0x0b, 0x76, 0xa7, 0xfa, 0x81, 0x87, 0x24, 0x54, 0x25,
	0x76, 0x5c, 0x44, 0xab, 0x30, 0xc8, 0xc5, 0xbe, 0xeb, 0x70, 0x15, 0xee, 0xc5, 0x7b, 0xf5, 0x31,
	0xc8, 0xc5, 0x81, 0xdb, 0x0b, 0x2a, 0xf0, 0x3b, 0xa3, 0xe8, 0xea, 0x2a, 0xc8, 0xdf, 0x0a, 0xee,
	0xbe, 0x53, 0x96, 0xb5, 0x72, 0x2a, 0x1e, 0x35, 0x95, 0x53, 0xee, 0xb1, 0xbb, 0x14, 0x7f, 0x60,
	0x8b, 0x4b, 0x53, 0x88, 0x43, 0x52, 0x9b, 0x08, 0x4f, 0x8c, 0xca, 0x19, 0xc4, 0x31, 0x5a, 0x1e,
	0x93, 0x65, 0x83, 0xe1, 0x3e, 0xfe, 0x09, 0xd1, 0x2a, 0x34, 0xb0, 0x3c, 0xbb, 0x59, 0x0d, 0x6c,
	0x11, 0xe4, 0xe2, 0x89, 0xdb, 0xc7, 0x36, 0xe7, 0xc7, 0xec, 0xa0, 0xc1, 0xf0, 0x0b, 0x41, 0x2e,
	0x3e, 0x25, 0x73, 0x57, 0x20, 0x77, 0x64, 0xc2, 0xa5, 0x96, 0x1f, 0x07, 0xb6, 0x98, 0xd8, 0xe4,
	0x2f, 0xd0, 0x42, 0x94, 0xee, 0xb6, 0xd0, 0x72, 0x9f, 0x25, 0xca, 0xa6, 0x46, 0x3c, 0xed, 0xb8,
	0x9d, 0x40, 0xa9, 0x6b, 0x38, 0x56, 0x29, 0x14, 0xe2, 0x59, 0x99, 0xba, 0xc5, 0xf9, 0x09, 0xe3,
	0x15, 0xc3, 0x64, 0x65, 0x90, 0xcf, 0xc8, 0xdd, 0xa3, 0xb4, 0xfd, 0x65, 0x94, 0xe7, 0x5d, 0x7f,
	0x99, 0xa5, 0x91, 0x1c, 0xa9, 0x0b, 0xf3, 0xf9, 0x66, 0xf2, 0xb5, 0x80, 0xaf, 0xbe, 0x3a, 0xa2,
	0x32, 0xc9, 0x0b, 0xb2, 0xb6, 0xe8, 0xfa, 0xce, 0x36, 0x22, 0x7f, 0xd1, 0xb8, 0xb3, 0x8d, 0xbc,
	0x6f, 0xd8, 0xa3, 0x81, 0x2d, 0xe8, 0xfe, 0x8e, 0x65, 0x3e, 0xb2, 0x5a, 0x9a, 0x48, 0xa5, 0xc2,
	0x23, 0x77, 0x9f, 0xc4, 0x4f, 0xd9, 0x21, 0xf6, 0x77, 0x5a, 0xbe, 0xa4, 0x96, 0x5e, 0x0d, 0xef,
	0xd2, 0xc4, 0x26, 0x03, 0x5b, 0x7c, 0xd0, 0x4b, 0xd0, 0x99, 0x38, 0xf2, 0xb6, 0xfc, 0xdd, 0xe9,
	0x06, 0xe3, 0x2f, 0xd9, 0xfd, 0x89, 0x4d, 0xb0, 0xbd, 0x34, 0x7d, 0x45, 0xa6, 0x4d, 0x88, 0xef,
	0x7e, 0xa2, 0x74, 0x12, 0xa5, 0x12, 0xaf, 0xcf, 0x4b, 0xf7, 0xee, 0x6b, 0xc2, 0xdf, 0xb1, 0xc7,
	0xb3, 0x50, 0x69, 0x83, 0x2d, 0xb3, 0x50, 0x6a, 0xc8, 0x28, 0xcc, 0x52, 0x7c, 0x4d, 0xd6, 0x7e,
	0x11, 0x4f, 0x61, 0x2d, 0xcc, 0xad, 0x4e, 0xd5, 0x0d, 0x68, 0xf1, 0xca, 0x9d, 0x42, 0x47, 0xc0,
	0xb7, 0x33, 0x85, 0x2b, 0xd0, 0x90, 0x2e, 0xf0, 0xed, 0x7e, 0xe3, 0xde, 0x4e, 0x03, 0x61, 0x8a,
	0xa1, 0x4a, 0xae, 0x63, 0xc8, 0xcf, 0x6f, 0x20, 0x35, 0x33, 0x23, 0xb5, 0xa1, 0x61, 0xeb, 0xbb,
	0x14, 0xbd, 0x22, 0x9e, 0x45, 0x53, 0x38, 0x4f, 0x97, 0xd4, 0xf3, 0xad, 0x3b, 0x8b, 0x1e, 0x09,
	0xf7, 0xf5, 0x3c, 0x5f, 0x84, 0x32, 0x5d, 0x01, 0xcd, 0xf8, 0xd7, 0x6e, 0x5f, 0x9b, 0xac, 0xe9,
	0xa1, 0xcf, 0x7d, 0xb7, 0xe9, 0xa9, 0x56, 0x3e, 0xbb, 0x32, 0xa0, 0x2f, 0x94, 0xd5, 0x59, 0x3d,
	0x37, 0x8f, 0xdd, 0xca, 0x3d, 0x12, 0xff, 0x8d, 0x3d, 0xab, 0x71, 0x67, 0x96, 0x7e, 0x4f, 0x8d,
	0xff, 0xe3, 0xe0, 0xbf, 0xb2, 0xa7, 0x6d, 0xb5, 0x9e, 0xaf, 0x27, 0xd4, 0xfe, 0xdf, 0x06, 0x9c,
	0xee, 0xc3, 0x50, 0xa6, 0x29, 0xc4, 0x13, 0x25, 0x7e, 0xa0, 0x1f, 0xaa, 0x01, 0xde, 0x24, 0x1c,
	0x39, 0x3a, 0x5a, 0xc0, 0xa5, 0x05, 0x0b, 0xe2, 0x8d, 0xb7, 0xe3, 0xef, 0x4c, 0x37, 0x21, 0x4e,
	0x81, 0xea, 0xf2, 0x5d, 0x9a, 0xc2, 0x19, 0xdf, 0x92, 0xb1, 0xc3, 0xf1, 0xdd, 0xd1, 0x5c, 0xaa,
	0x3f, 0x79, 0x4a, 0xce, 0x16, 0xa5, 0x7b, 0x54, 0xdd, 0xd5, 0xf5, 0x47, 0x7f, 0x24, 0x6b, 0x57,
	0x28, 0x73, 0x3a, 0x46, 0xce, 0x77, 0xeb, 0x9c, 0x35, 0xac, 0xd6, 0x6e, 0xd8, 0x7e, 0xaa, 0xd7,
	0x6e, 0xf8, 0x8e, 0xd9, 0x81, 0x1b, 0x00, 0xee, 0xa5, 0x38, 0xeb, 0xcf, 0x6e, 0xed, 0x8e, 0x80,
	0x73, 0xaa, 0x9c, 0x02, 0x4d, 0xfb, 0x2f, 0x64, 0xef, 0x51, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x5b, 0x0f, 0x1a, 0x97, 0x08, 0x00, 0x00,
}
