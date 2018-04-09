// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDStock.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	MDStock.proto

It has these top-level messages:
	MDStock
*/
package com_htsc_mdc_insight_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_htsc_mdc_model "."
import com_htsc_mdc_model1 "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 股票信息
type MDStock struct {
	HTSCSecurityID        string                               `protobuf:"bytes,1,opt,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	MDDate                int32                                `protobuf:"varint,2,opt,name=MDDate" json:"MDDate,omitempty"`
	MDTime                int32                                `protobuf:"varint,3,opt,name=MDTime" json:"MDTime,omitempty"`
	DataTimestamp         int64                                `protobuf:"varint,4,opt,name=DataTimestamp" json:"DataTimestamp,omitempty"`
	TradingPhaseCode      string                               `protobuf:"bytes,5,opt,name=TradingPhaseCode" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource      com_htsc_mdc_model.ESecurityIDSource `protobuf:"varint,6,opt,name=securityIDSource,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType          com_htsc_mdc_model1.ESecurityType    `protobuf:"varint,7,opt,name=securityType,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	MaxPx                 int64                                `protobuf:"varint,8,opt,name=MaxPx" json:"MaxPx,omitempty"`
	MinPx                 int64                                `protobuf:"varint,9,opt,name=MinPx" json:"MinPx,omitempty"`
	PreClosePx            int64                                `protobuf:"varint,10,opt,name=PreClosePx" json:"PreClosePx,omitempty"`
	NumTrades             int64                                `protobuf:"varint,11,opt,name=NumTrades" json:"NumTrades,omitempty"`
	TotalVolumeTrade      int64                                `protobuf:"varint,12,opt,name=TotalVolumeTrade" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade       int64                                `protobuf:"varint,13,opt,name=TotalValueTrade" json:"TotalValueTrade,omitempty"`
	LastPx                int64                                `protobuf:"varint,14,opt,name=LastPx" json:"LastPx,omitempty"`
	OpenPx                int64                                `protobuf:"varint,15,opt,name=OpenPx" json:"OpenPx,omitempty"`
	ClosePx               int64                                `protobuf:"varint,16,opt,name=ClosePx" json:"ClosePx,omitempty"`
	HighPx                int64                                `protobuf:"varint,17,opt,name=HighPx" json:"HighPx,omitempty"`
	LowPx                 int64                                `protobuf:"varint,18,opt,name=LowPx" json:"LowPx,omitempty"`
	DiffPx1               int64                                `protobuf:"varint,19,opt,name=DiffPx1" json:"DiffPx1,omitempty"`
	DiffPx2               int64                                `protobuf:"varint,20,opt,name=DiffPx2" json:"DiffPx2,omitempty"`
	TotalBuyQty           int64                                `protobuf:"varint,21,opt,name=TotalBuyQty" json:"TotalBuyQty,omitempty"`
	TotalSellQty          int64                                `protobuf:"varint,22,opt,name=TotalSellQty" json:"TotalSellQty,omitempty"`
	WeightedAvgBuyPx      int64                                `protobuf:"varint,23,opt,name=WeightedAvgBuyPx" json:"WeightedAvgBuyPx,omitempty"`
	WeightedAvgSellPx     int64                                `protobuf:"varint,24,opt,name=WeightedAvgSellPx" json:"WeightedAvgSellPx,omitempty"`
	WithdrawBuyNumber     int64                                `protobuf:"varint,25,opt,name=WithdrawBuyNumber" json:"WithdrawBuyNumber,omitempty"`
	WithdrawBuyAmount     int64                                `protobuf:"varint,26,opt,name=WithdrawBuyAmount" json:"WithdrawBuyAmount,omitempty"`
	WithdrawBuyMoney      int64                                `protobuf:"varint,27,opt,name=WithdrawBuyMoney" json:"WithdrawBuyMoney,omitempty"`
	WithdrawSellNumber    int64                                `protobuf:"varint,28,opt,name=WithdrawSellNumber" json:"WithdrawSellNumber,omitempty"`
	WithdrawSellAmount    int64                                `protobuf:"varint,29,opt,name=WithdrawSellAmount" json:"WithdrawSellAmount,omitempty"`
	WithdrawSellMoney     int64                                `protobuf:"varint,30,opt,name=WithdrawSellMoney" json:"WithdrawSellMoney,omitempty"`
	TotalBuyNumber        int64                                `protobuf:"varint,31,opt,name=TotalBuyNumber" json:"TotalBuyNumber,omitempty"`
	TotalSellNumber       int64                                `protobuf:"varint,32,opt,name=TotalSellNumber" json:"TotalSellNumber,omitempty"`
	BuyTradeMaxDuration   int64                                `protobuf:"varint,33,opt,name=BuyTradeMaxDuration" json:"BuyTradeMaxDuration,omitempty"`
	SellTradeMaxDuration  int64                                `protobuf:"varint,34,opt,name=SellTradeMaxDuration" json:"SellTradeMaxDuration,omitempty"`
	NumBuyOrders          int32                                `protobuf:"varint,35,opt,name=NumBuyOrders" json:"NumBuyOrders,omitempty"`
	NumSellOrders         int32                                `protobuf:"varint,36,opt,name=NumSellOrders" json:"NumSellOrders,omitempty"`
	NorminalPx            int64                                `protobuf:"varint,37,opt,name=NorminalPx" json:"NorminalPx,omitempty"`
	ShortSellSharesTraded int64                                `protobuf:"varint,38,opt,name=ShortSellSharesTraded" json:"ShortSellSharesTraded,omitempty"`
	ShortSellTurnover     int64                                `protobuf:"varint,39,opt,name=ShortSellTurnover" json:"ShortSellTurnover,omitempty"`
	ReferencePx           int64                                `protobuf:"varint,40,opt,name=ReferencePx" json:"ReferencePx,omitempty"`
	ComplexEventStartTime int64                                `protobuf:"varint,41,opt,name=ComplexEventStartTime" json:"ComplexEventStartTime,omitempty"`
	ComplexEventEndTime   int64                                `protobuf:"varint,42,opt,name=ComplexEventEndTime" json:"ComplexEventEndTime,omitempty"`
	BuyOrderQueue         []int64                              `protobuf:"varint,100,rep,packed,name=BuyOrderQueue" json:"BuyOrderQueue,omitempty"`
	SellOrderQueue        []int64                              `protobuf:"varint,101,rep,packed,name=SellOrderQueue" json:"SellOrderQueue,omitempty"`
	Buy1Price             int64                                `protobuf:"varint,102,opt,name=Buy1Price" json:"Buy1Price,omitempty"`
	Buy1OrderQty          int64                                `protobuf:"varint,103,opt,name=Buy1OrderQty" json:"Buy1OrderQty,omitempty"`
	Buy1NumOrders         int64                                `protobuf:"varint,104,opt,name=Buy1NumOrders" json:"Buy1NumOrders,omitempty"`
	Sell1Price            int64                                `protobuf:"varint,105,opt,name=Sell1Price" json:"Sell1Price,omitempty"`
	Sell1OrderQty         int64                                `protobuf:"varint,106,opt,name=Sell1OrderQty" json:"Sell1OrderQty,omitempty"`
	Sell1NumOrders        int64                                `protobuf:"varint,107,opt,name=Sell1NumOrders" json:"Sell1NumOrders,omitempty"`
	Buy2Price             int64                                `protobuf:"varint,108,opt,name=Buy2Price" json:"Buy2Price,omitempty"`
	Buy2OrderQty          int64                                `protobuf:"varint,109,opt,name=Buy2OrderQty" json:"Buy2OrderQty,omitempty"`
	Buy2NumOrders         int64                                `protobuf:"varint,110,opt,name=Buy2NumOrders" json:"Buy2NumOrders,omitempty"`
	Sell2Price            int64                                `protobuf:"varint,111,opt,name=Sell2Price" json:"Sell2Price,omitempty"`
	Sell2OrderQty         int64                                `protobuf:"varint,112,opt,name=Sell2OrderQty" json:"Sell2OrderQty,omitempty"`
	Sell2NumOrders        int64                                `protobuf:"varint,113,opt,name=Sell2NumOrders" json:"Sell2NumOrders,omitempty"`
	Buy3Price             int64                                `protobuf:"varint,114,opt,name=Buy3Price" json:"Buy3Price,omitempty"`
	Buy3OrderQty          int64                                `protobuf:"varint,115,opt,name=Buy3OrderQty" json:"Buy3OrderQty,omitempty"`
	Buy3NumOrders         int64                                `protobuf:"varint,116,opt,name=Buy3NumOrders" json:"Buy3NumOrders,omitempty"`
	Sell3Price            int64                                `protobuf:"varint,117,opt,name=Sell3Price" json:"Sell3Price,omitempty"`
	Sell3OrderQty         int64                                `protobuf:"varint,118,opt,name=Sell3OrderQty" json:"Sell3OrderQty,omitempty"`
	Sell3NumOrders        int64                                `protobuf:"varint,119,opt,name=Sell3NumOrders" json:"Sell3NumOrders,omitempty"`
	Buy4Price             int64                                `protobuf:"varint,120,opt,name=Buy4Price" json:"Buy4Price,omitempty"`
	Buy4OrderQty          int64                                `protobuf:"varint,121,opt,name=Buy4OrderQty" json:"Buy4OrderQty,omitempty"`
	Buy4NumOrders         int64                                `protobuf:"varint,122,opt,name=Buy4NumOrders" json:"Buy4NumOrders,omitempty"`
	Sell4Price            int64                                `protobuf:"varint,123,opt,name=Sell4Price" json:"Sell4Price,omitempty"`
	Sell4OrderQty         int64                                `protobuf:"varint,124,opt,name=Sell4OrderQty" json:"Sell4OrderQty,omitempty"`
	Sell4NumOrders        int64                                `protobuf:"varint,125,opt,name=Sell4NumOrders" json:"Sell4NumOrders,omitempty"`
	Buy5Price             int64                                `protobuf:"varint,126,opt,name=Buy5Price" json:"Buy5Price,omitempty"`
	Buy5OrderQty          int64                                `protobuf:"varint,127,opt,name=Buy5OrderQty" json:"Buy5OrderQty,omitempty"`
	Buy5NumOrders         int64                                `protobuf:"varint,128,opt,name=Buy5NumOrders" json:"Buy5NumOrders,omitempty"`
	Sell5Price            int64                                `protobuf:"varint,129,opt,name=Sell5Price" json:"Sell5Price,omitempty"`
	Sell5OrderQty         int64                                `protobuf:"varint,130,opt,name=Sell5OrderQty" json:"Sell5OrderQty,omitempty"`
	Sell5NumOrders        int64                                `protobuf:"varint,131,opt,name=Sell5NumOrders" json:"Sell5NumOrders,omitempty"`
	Buy6Price             int64                                `protobuf:"varint,132,opt,name=Buy6Price" json:"Buy6Price,omitempty"`
	Buy6OrderQty          int64                                `protobuf:"varint,133,opt,name=Buy6OrderQty" json:"Buy6OrderQty,omitempty"`
	Buy6NumOrders         int64                                `protobuf:"varint,134,opt,name=Buy6NumOrders" json:"Buy6NumOrders,omitempty"`
	Sell6Price            int64                                `protobuf:"varint,135,opt,name=Sell6Price" json:"Sell6Price,omitempty"`
	Sell6OrderQty         int64                                `protobuf:"varint,136,opt,name=Sell6OrderQty" json:"Sell6OrderQty,omitempty"`
	Sell6NumOrders        int64                                `protobuf:"varint,137,opt,name=Sell6NumOrders" json:"Sell6NumOrders,omitempty"`
	Buy7Price             int64                                `protobuf:"varint,138,opt,name=Buy7Price" json:"Buy7Price,omitempty"`
	Buy7OrderQty          int64                                `protobuf:"varint,139,opt,name=Buy7OrderQty" json:"Buy7OrderQty,omitempty"`
	Buy7NumOrders         int64                                `protobuf:"varint,140,opt,name=Buy7NumOrders" json:"Buy7NumOrders,omitempty"`
	Sell7Price            int64                                `protobuf:"varint,141,opt,name=Sell7Price" json:"Sell7Price,omitempty"`
	Sell7OrderQty         int64                                `protobuf:"varint,142,opt,name=Sell7OrderQty" json:"Sell7OrderQty,omitempty"`
	Sell7NumOrders        int64                                `protobuf:"varint,143,opt,name=Sell7NumOrders" json:"Sell7NumOrders,omitempty"`
	Buy8Price             int64                                `protobuf:"varint,144,opt,name=Buy8Price" json:"Buy8Price,omitempty"`
	Buy8OrderQty          int64                                `protobuf:"varint,145,opt,name=Buy8OrderQty" json:"Buy8OrderQty,omitempty"`
	Buy8NumOrders         int64                                `protobuf:"varint,146,opt,name=Buy8NumOrders" json:"Buy8NumOrders,omitempty"`
	Sell8Price            int64                                `protobuf:"varint,147,opt,name=Sell8Price" json:"Sell8Price,omitempty"`
	Sell8OrderQty         int64                                `protobuf:"varint,148,opt,name=Sell8OrderQty" json:"Sell8OrderQty,omitempty"`
	Sell8NumOrders        int64                                `protobuf:"varint,149,opt,name=Sell8NumOrders" json:"Sell8NumOrders,omitempty"`
	Buy9Price             int64                                `protobuf:"varint,150,opt,name=Buy9Price" json:"Buy9Price,omitempty"`
	Buy9OrderQty          int64                                `protobuf:"varint,151,opt,name=Buy9OrderQty" json:"Buy9OrderQty,omitempty"`
	Buy9NumOrders         int64                                `protobuf:"varint,152,opt,name=Buy9NumOrders" json:"Buy9NumOrders,omitempty"`
	Sell9Price            int64                                `protobuf:"varint,153,opt,name=Sell9Price" json:"Sell9Price,omitempty"`
	Sell9OrderQty         int64                                `protobuf:"varint,154,opt,name=Sell9OrderQty" json:"Sell9OrderQty,omitempty"`
	Sell9NumOrders        int64                                `protobuf:"varint,155,opt,name=Sell9NumOrders" json:"Sell9NumOrders,omitempty"`
	Buy10Price            int64                                `protobuf:"varint,156,opt,name=Buy10Price" json:"Buy10Price,omitempty"`
	Buy10OrderQty         int64                                `protobuf:"varint,157,opt,name=Buy10OrderQty" json:"Buy10OrderQty,omitempty"`
	Buy10NumOrders        int64                                `protobuf:"varint,158,opt,name=Buy10NumOrders" json:"Buy10NumOrders,omitempty"`
	Sell10Price           int64                                `protobuf:"varint,159,opt,name=Sell10Price" json:"Sell10Price,omitempty"`
	Sell10OrderQty        int64                                `protobuf:"varint,160,opt,name=Sell10OrderQty" json:"Sell10OrderQty,omitempty"`
	Sell10NumOrders       int64                                `protobuf:"varint,163,opt,name=Sell10NumOrders" json:"Sell10NumOrders,omitempty"`
}

func (m *MDStock) Reset()                    { *m = MDStock{} }
func (m *MDStock) String() string            { return proto.CompactTextString(m) }
func (*MDStock) ProtoMessage()               {}
func (*MDStock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MDStock) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDStock) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDStock) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDStock) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDStock) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDStock) GetSecurityIDSource() com_htsc_mdc_model.ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return com_htsc_mdc_model.ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDStock) GetSecurityType() com_htsc_mdc_model1.ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return com_htsc_mdc_model1.ESecurityType_DefaultSecurityType
}

func (m *MDStock) GetMaxPx() int64 {
	if m != nil {
		return m.MaxPx
	}
	return 0
}

func (m *MDStock) GetMinPx() int64 {
	if m != nil {
		return m.MinPx
	}
	return 0
}

func (m *MDStock) GetPreClosePx() int64 {
	if m != nil {
		return m.PreClosePx
	}
	return 0
}

func (m *MDStock) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func (m *MDStock) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *MDStock) GetTotalValueTrade() int64 {
	if m != nil {
		return m.TotalValueTrade
	}
	return 0
}

func (m *MDStock) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *MDStock) GetOpenPx() int64 {
	if m != nil {
		return m.OpenPx
	}
	return 0
}

func (m *MDStock) GetClosePx() int64 {
	if m != nil {
		return m.ClosePx
	}
	return 0
}

func (m *MDStock) GetHighPx() int64 {
	if m != nil {
		return m.HighPx
	}
	return 0
}

func (m *MDStock) GetLowPx() int64 {
	if m != nil {
		return m.LowPx
	}
	return 0
}

func (m *MDStock) GetDiffPx1() int64 {
	if m != nil {
		return m.DiffPx1
	}
	return 0
}

func (m *MDStock) GetDiffPx2() int64 {
	if m != nil {
		return m.DiffPx2
	}
	return 0
}

func (m *MDStock) GetTotalBuyQty() int64 {
	if m != nil {
		return m.TotalBuyQty
	}
	return 0
}

func (m *MDStock) GetTotalSellQty() int64 {
	if m != nil {
		return m.TotalSellQty
	}
	return 0
}

func (m *MDStock) GetWeightedAvgBuyPx() int64 {
	if m != nil {
		return m.WeightedAvgBuyPx
	}
	return 0
}

func (m *MDStock) GetWeightedAvgSellPx() int64 {
	if m != nil {
		return m.WeightedAvgSellPx
	}
	return 0
}

func (m *MDStock) GetWithdrawBuyNumber() int64 {
	if m != nil {
		return m.WithdrawBuyNumber
	}
	return 0
}

func (m *MDStock) GetWithdrawBuyAmount() int64 {
	if m != nil {
		return m.WithdrawBuyAmount
	}
	return 0
}

func (m *MDStock) GetWithdrawBuyMoney() int64 {
	if m != nil {
		return m.WithdrawBuyMoney
	}
	return 0
}

func (m *MDStock) GetWithdrawSellNumber() int64 {
	if m != nil {
		return m.WithdrawSellNumber
	}
	return 0
}

func (m *MDStock) GetWithdrawSellAmount() int64 {
	if m != nil {
		return m.WithdrawSellAmount
	}
	return 0
}

func (m *MDStock) GetWithdrawSellMoney() int64 {
	if m != nil {
		return m.WithdrawSellMoney
	}
	return 0
}

func (m *MDStock) GetTotalBuyNumber() int64 {
	if m != nil {
		return m.TotalBuyNumber
	}
	return 0
}

func (m *MDStock) GetTotalSellNumber() int64 {
	if m != nil {
		return m.TotalSellNumber
	}
	return 0
}

func (m *MDStock) GetBuyTradeMaxDuration() int64 {
	if m != nil {
		return m.BuyTradeMaxDuration
	}
	return 0
}

func (m *MDStock) GetSellTradeMaxDuration() int64 {
	if m != nil {
		return m.SellTradeMaxDuration
	}
	return 0
}

func (m *MDStock) GetNumBuyOrders() int32 {
	if m != nil {
		return m.NumBuyOrders
	}
	return 0
}

func (m *MDStock) GetNumSellOrders() int32 {
	if m != nil {
		return m.NumSellOrders
	}
	return 0
}

func (m *MDStock) GetNorminalPx() int64 {
	if m != nil {
		return m.NorminalPx
	}
	return 0
}

func (m *MDStock) GetShortSellSharesTraded() int64 {
	if m != nil {
		return m.ShortSellSharesTraded
	}
	return 0
}

func (m *MDStock) GetShortSellTurnover() int64 {
	if m != nil {
		return m.ShortSellTurnover
	}
	return 0
}

func (m *MDStock) GetReferencePx() int64 {
	if m != nil {
		return m.ReferencePx
	}
	return 0
}

func (m *MDStock) GetComplexEventStartTime() int64 {
	if m != nil {
		return m.ComplexEventStartTime
	}
	return 0
}

func (m *MDStock) GetComplexEventEndTime() int64 {
	if m != nil {
		return m.ComplexEventEndTime
	}
	return 0
}

func (m *MDStock) GetBuyOrderQueue() []int64 {
	if m != nil {
		return m.BuyOrderQueue
	}
	return nil
}

func (m *MDStock) GetSellOrderQueue() []int64 {
	if m != nil {
		return m.SellOrderQueue
	}
	return nil
}

func (m *MDStock) GetBuy1Price() int64 {
	if m != nil {
		return m.Buy1Price
	}
	return 0
}

func (m *MDStock) GetBuy1OrderQty() int64 {
	if m != nil {
		return m.Buy1OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy1NumOrders() int64 {
	if m != nil {
		return m.Buy1NumOrders
	}
	return 0
}

func (m *MDStock) GetSell1Price() int64 {
	if m != nil {
		return m.Sell1Price
	}
	return 0
}

func (m *MDStock) GetSell1OrderQty() int64 {
	if m != nil {
		return m.Sell1OrderQty
	}
	return 0
}

func (m *MDStock) GetSell1NumOrders() int64 {
	if m != nil {
		return m.Sell1NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy2Price() int64 {
	if m != nil {
		return m.Buy2Price
	}
	return 0
}

func (m *MDStock) GetBuy2OrderQty() int64 {
	if m != nil {
		return m.Buy2OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy2NumOrders() int64 {
	if m != nil {
		return m.Buy2NumOrders
	}
	return 0
}

func (m *MDStock) GetSell2Price() int64 {
	if m != nil {
		return m.Sell2Price
	}
	return 0
}

func (m *MDStock) GetSell2OrderQty() int64 {
	if m != nil {
		return m.Sell2OrderQty
	}
	return 0
}

func (m *MDStock) GetSell2NumOrders() int64 {
	if m != nil {
		return m.Sell2NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy3Price() int64 {
	if m != nil {
		return m.Buy3Price
	}
	return 0
}

func (m *MDStock) GetBuy3OrderQty() int64 {
	if m != nil {
		return m.Buy3OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy3NumOrders() int64 {
	if m != nil {
		return m.Buy3NumOrders
	}
	return 0
}

func (m *MDStock) GetSell3Price() int64 {
	if m != nil {
		return m.Sell3Price
	}
	return 0
}

func (m *MDStock) GetSell3OrderQty() int64 {
	if m != nil {
		return m.Sell3OrderQty
	}
	return 0
}

func (m *MDStock) GetSell3NumOrders() int64 {
	if m != nil {
		return m.Sell3NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy4Price() int64 {
	if m != nil {
		return m.Buy4Price
	}
	return 0
}

func (m *MDStock) GetBuy4OrderQty() int64 {
	if m != nil {
		return m.Buy4OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy4NumOrders() int64 {
	if m != nil {
		return m.Buy4NumOrders
	}
	return 0
}

func (m *MDStock) GetSell4Price() int64 {
	if m != nil {
		return m.Sell4Price
	}
	return 0
}

func (m *MDStock) GetSell4OrderQty() int64 {
	if m != nil {
		return m.Sell4OrderQty
	}
	return 0
}

func (m *MDStock) GetSell4NumOrders() int64 {
	if m != nil {
		return m.Sell4NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy5Price() int64 {
	if m != nil {
		return m.Buy5Price
	}
	return 0
}

func (m *MDStock) GetBuy5OrderQty() int64 {
	if m != nil {
		return m.Buy5OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy5NumOrders() int64 {
	if m != nil {
		return m.Buy5NumOrders
	}
	return 0
}

func (m *MDStock) GetSell5Price() int64 {
	if m != nil {
		return m.Sell5Price
	}
	return 0
}

func (m *MDStock) GetSell5OrderQty() int64 {
	if m != nil {
		return m.Sell5OrderQty
	}
	return 0
}

func (m *MDStock) GetSell5NumOrders() int64 {
	if m != nil {
		return m.Sell5NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy6Price() int64 {
	if m != nil {
		return m.Buy6Price
	}
	return 0
}

func (m *MDStock) GetBuy6OrderQty() int64 {
	if m != nil {
		return m.Buy6OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy6NumOrders() int64 {
	if m != nil {
		return m.Buy6NumOrders
	}
	return 0
}

func (m *MDStock) GetSell6Price() int64 {
	if m != nil {
		return m.Sell6Price
	}
	return 0
}

func (m *MDStock) GetSell6OrderQty() int64 {
	if m != nil {
		return m.Sell6OrderQty
	}
	return 0
}

func (m *MDStock) GetSell6NumOrders() int64 {
	if m != nil {
		return m.Sell6NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy7Price() int64 {
	if m != nil {
		return m.Buy7Price
	}
	return 0
}

func (m *MDStock) GetBuy7OrderQty() int64 {
	if m != nil {
		return m.Buy7OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy7NumOrders() int64 {
	if m != nil {
		return m.Buy7NumOrders
	}
	return 0
}

func (m *MDStock) GetSell7Price() int64 {
	if m != nil {
		return m.Sell7Price
	}
	return 0
}

func (m *MDStock) GetSell7OrderQty() int64 {
	if m != nil {
		return m.Sell7OrderQty
	}
	return 0
}

func (m *MDStock) GetSell7NumOrders() int64 {
	if m != nil {
		return m.Sell7NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy8Price() int64 {
	if m != nil {
		return m.Buy8Price
	}
	return 0
}

func (m *MDStock) GetBuy8OrderQty() int64 {
	if m != nil {
		return m.Buy8OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy8NumOrders() int64 {
	if m != nil {
		return m.Buy8NumOrders
	}
	return 0
}

func (m *MDStock) GetSell8Price() int64 {
	if m != nil {
		return m.Sell8Price
	}
	return 0
}

func (m *MDStock) GetSell8OrderQty() int64 {
	if m != nil {
		return m.Sell8OrderQty
	}
	return 0
}

func (m *MDStock) GetSell8NumOrders() int64 {
	if m != nil {
		return m.Sell8NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy9Price() int64 {
	if m != nil {
		return m.Buy9Price
	}
	return 0
}

func (m *MDStock) GetBuy9OrderQty() int64 {
	if m != nil {
		return m.Buy9OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy9NumOrders() int64 {
	if m != nil {
		return m.Buy9NumOrders
	}
	return 0
}

func (m *MDStock) GetSell9Price() int64 {
	if m != nil {
		return m.Sell9Price
	}
	return 0
}

func (m *MDStock) GetSell9OrderQty() int64 {
	if m != nil {
		return m.Sell9OrderQty
	}
	return 0
}

func (m *MDStock) GetSell9NumOrders() int64 {
	if m != nil {
		return m.Sell9NumOrders
	}
	return 0
}

func (m *MDStock) GetBuy10Price() int64 {
	if m != nil {
		return m.Buy10Price
	}
	return 0
}

func (m *MDStock) GetBuy10OrderQty() int64 {
	if m != nil {
		return m.Buy10OrderQty
	}
	return 0
}

func (m *MDStock) GetBuy10NumOrders() int64 {
	if m != nil {
		return m.Buy10NumOrders
	}
	return 0
}

func (m *MDStock) GetSell10Price() int64 {
	if m != nil {
		return m.Sell10Price
	}
	return 0
}

func (m *MDStock) GetSell10OrderQty() int64 {
	if m != nil {
		return m.Sell10OrderQty
	}
	return 0
}

func (m *MDStock) GetSell10NumOrders() int64 {
	if m != nil {
		return m.Sell10NumOrders
	}
	return 0
}

func init() {
	proto.RegisterType((*MDStock)(nil), "com.htsc.mdc.insight.model.MDStock")
}

func init() { proto.RegisterFile("MDStock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x97, 0x67, 0x6f, 0xdb, 0x46,
	0x18, 0xc7, 0xc1, 0xa6, 0x49, 0x9a, 0xcb, 0xbe, 0xac, 0xab, 0x9b, 0xa1, 0x0c, 0x27, 0x4a, 0x50,
	0x08, 0x36, 0x25, 0x6b, 0xbc, 0x8c, 0x2c, 0x03, 0x2e, 0x10, 0x2b, 0xb4, 0x64, 0xb4, 0xaf, 0x19,
	0xe9, 0x6c, 0xb1, 0x11, 0x49, 0x95, 0xc3, 0x26, 0xbb, 0xf7, 0xde, 0x7b, 0x37, 0x40, 0x3f, 0x54,
	0xbf, 0x4e, 0x71, 0x43, 0xf7, 0x1c, 0x8f, 0x42, 0xde, 0xf9, 0xfe, 0xcf, 0x73, 0xfc, 0x3d, 0xbf,
	0xe3, 0x30, 0x84, 0x4e, 0x6e, 0xf5, 0x86, 0x49, 0x38, 0x7a, 0x5c, 0x9b, 0x45, 0x61, 0x12, 0xe2,
	0xa5, 0x51, 0xe8, 0xd7, 0x26, 0x49, 0x3c, 0xaa, 0xf9, 0xe3, 0x51, 0xcd, 0x0b, 0x62, 0x6f, 0x6f,
	0x92, 0xd4, 0xfc, 0x70, 0x4c, 0xa7, 0x4b, 0x97, 0x36, 0x86, 0x74, 0x94, 0x46, 0x5e, 0x92, 0xbf,
	0xd4, 0x1b, 0x86, 0x69, 0x34, 0xa2, 0x62, 0xd3, 0xd2, 0x39, 0x55, 0xd8, 0xc9, 0x67, 0x32, 0xbc,
	0xf1, 0xdf, 0x55, 0x74, 0x54, 0x5e, 0x1b, 0xdf, 0x46, 0xa7, 0x36, 0x77, 0x86, 0xeb, 0xb0, 0x9d,
	0x58, 0x15, 0xab, 0x7a, 0x6c, 0x60, 0xa4, 0xf8, 0x22, 0x3a, 0xb2, 0xd5, 0xeb, 0xb9, 0x09, 0x25,
	0xcf, 0x54, 0xac, 0xea, 0xe1, 0x81, 0x5c, 0x89, 0x7c, 0xc7, 0xf3, 0x29, 0x39, 0x34, 0xcf, 0xd9,
	0x0a, 0xdf, 0x42, 0x27, 0x7b, 0x6e, 0xe2, 0xb2, 0xbf, 0xe3, 0xc4, 0xf5, 0x67, 0xe4, 0xd9, 0x8a,
	0x55, 0x3d, 0x34, 0x28, 0x86, 0xf8, 0x1e, 0x3a, 0xb3, 0x13, 0xb9, 0x63, 0x2f, 0xd8, 0x73, 0x26,
	0x6e, 0x4c, 0xd7, 0xc3, 0x31, 0x25, 0x87, 0x39, 0xbf, 0x94, 0xe3, 0x6d, 0x74, 0x26, 0x36, 0x24,
	0xc9, 0x91, 0x8a, 0x55, 0x3d, 0x65, 0x2f, 0xd7, 0x0a, 0x47, 0xc3, 0x8f, 0xa4, 0x56, 0x3a, 0x91,
	0x41, 0x69, 0x3b, 0xde, 0x40, 0x27, 0x62, 0xed, 0x78, 0xc8, 0x51, 0x7e, 0xb9, 0xeb, 0x4f, 0xbd,
	0x1c, 0x6b, 0x1c, 0x14, 0xb6, 0xe1, 0xf3, 0xe8, 0xf0, 0x96, 0x9b, 0x39, 0x19, 0x79, 0x8e, 0x3b,
	0x8a, 0x05, 0x4f, 0xbd, 0xc0, 0xc9, 0xc8, 0x31, 0x99, 0xb2, 0x05, 0xbe, 0x8a, 0x90, 0x13, 0xd1,
	0xf5, 0x69, 0x18, 0x53, 0x27, 0x23, 0x88, 0x97, 0xb4, 0x04, 0x5f, 0x46, 0xc7, 0xfa, 0xa9, 0xcf,
	0xe4, 0x69, 0x4c, 0x8e, 0xf3, 0x32, 0x04, 0xfc, 0xbc, 0xc2, 0xc4, 0x9d, 0xbe, 0x1c, 0x4e, 0x53,
	0x9f, 0xf2, 0x90, 0x9c, 0xe0, 0x4d, 0xa5, 0x1c, 0x57, 0xd1, 0x69, 0x91, 0xb9, 0xd3, 0x54, 0xb6,
	0x9e, 0xe4, 0xad, 0x66, 0xcc, 0xee, 0xe1, 0x03, 0x37, 0x4e, 0x9c, 0x8c, 0x9c, 0xe2, 0x0d, 0x72,
	0xc5, 0xf2, 0x87, 0x33, 0xca, 0x14, 0x4e, 0x8b, 0x5c, 0xac, 0x30, 0x41, 0x47, 0xe7, 0x02, 0x67,
	0x78, 0x61, 0xbe, 0x64, 0x3b, 0x36, 0xbd, 0xbd, 0x89, 0x93, 0x91, 0xb3, 0x62, 0x87, 0x58, 0xb1,
	0xb3, 0x78, 0x10, 0x1e, 0x38, 0x19, 0xc1, 0xe2, 0x2c, 0xf8, 0x82, 0x5d, 0xa7, 0xe7, 0xed, 0xee,
	0x3a, 0xd9, 0x2a, 0x39, 0x27, 0xae, 0x23, 0x97, 0x50, 0xb1, 0xc9, 0x79, 0xbd, 0x62, 0xe3, 0x0a,
	0x3a, 0xce, 0xc7, 0xef, 0xa6, 0xf9, 0x76, 0x92, 0x93, 0x0b, 0xbc, 0xaa, 0x47, 0xf8, 0x06, 0x3a,
	0xc1, 0x97, 0x43, 0x3a, 0x9d, 0xb2, 0x96, 0x8b, 0xbc, 0xa5, 0x90, 0xb1, 0x73, 0x7c, 0x85, 0xb2,
	0xf7, 0x87, 0x8e, 0xef, 0xef, 0xef, 0x75, 0xd3, 0xdc, 0xc9, 0xc8, 0x25, 0x71, 0x8e, 0x66, 0x8e,
	0x5f, 0x44, 0x67, 0xb5, 0x8c, 0x5d, 0xc1, 0xc9, 0x08, 0xe1, 0xcd, 0xe5, 0x02, 0xef, 0xf6, 0x92,
	0xc9, 0x38, 0x72, 0x0f, 0xba, 0x69, 0xde, 0x4f, 0xfd, 0x47, 0x34, 0x22, 0xcf, 0xcb, 0x6e, 0xb3,
	0x60, 0x74, 0xdf, 0xf7, 0xc3, 0x34, 0x48, 0xc8, 0x52, 0xa9, 0x5b, 0x14, 0xf8, 0xd4, 0x10, 0x6e,
	0x85, 0x01, 0xcd, 0xc9, 0x0b, 0x72, 0x6a, 0x23, 0xc7, 0x35, 0x84, 0xe7, 0x19, 0x9b, 0x4c, 0x0e,
	0x72, 0x99, 0x77, 0x2f, 0xa8, 0x98, 0xfd, 0x72, 0x94, 0x2b, 0xe5, 0x7e, 0x39, 0x8b, 0x36, 0x39,
	0x4b, 0xc5, 0x30, 0x57, 0x8b, 0x93, 0xab, 0x02, 0xfb, 0xca, 0xcc, 0x6f, 0x91, 0x9c, 0xe4, 0x1a,
	0x6f, 0x35, 0x52, 0xf5, 0xcc, 0x6a, 0x23, 0x57, 0xb4, 0x67, 0x56, 0x9b, 0x77, 0x05, 0x9d, 0xeb,
	0xa6, 0x39, 0x7f, 0x7e, 0xb7, 0xdc, 0xac, 0x97, 0x46, 0x6e, 0xe2, 0x85, 0x01, 0xb9, 0xce, 0xbb,
	0x17, 0x95, 0xb0, 0x8d, 0xce, 0xb3, 0xfd, 0xa5, 0x2d, 0x37, 0xf8, 0x96, 0x85, 0x35, 0xf6, 0x2c,
	0xf5, 0x53, 0xbf, 0x9b, 0xe6, 0x0f, 0xa3, 0x31, 0x8d, 0x62, 0x72, 0x93, 0x7f, 0xe3, 0x0a, 0x19,
	0xfb, 0xd2, 0xf5, 0x53, 0x9f, 0x6d, 0x97, 0x4d, 0xb7, 0x78, 0x53, 0x31, 0x64, 0xef, 0x7d, 0x3f,
	0x8c, 0x7c, 0x2f, 0x70, 0xd9, 0xe3, 0xb3, 0x2c, 0xde, 0x7b, 0x48, 0x70, 0x03, 0x5d, 0x18, 0x4e,
	0xc2, 0x28, 0x61, 0x5b, 0x86, 0x13, 0x37, 0xa2, 0x31, 0x1f, 0x66, 0x4c, 0x6e, 0xf3, 0xd6, 0xc5,
	0x45, 0x76, 0x17, 0x54, 0x61, 0x27, 0x8d, 0x82, 0x70, 0x9f, 0x46, 0xe4, 0x8e, 0xb8, 0x0b, 0xa5,
	0x02, 0x7b, 0x77, 0x06, 0x74, 0x97, 0x46, 0x34, 0x18, 0xb1, 0x77, 0xb7, 0x2a, 0xde, 0x1d, 0x2d,
	0x62, 0x53, 0xac, 0x87, 0xfe, 0x6c, 0x4a, 0xb3, 0x8d, 0x7d, 0x1a, 0x24, 0xc3, 0xc4, 0x8d, 0x12,
	0xfe, 0x71, 0xbf, 0x2b, 0xa6, 0x58, 0x58, 0x64, 0xf7, 0x42, 0x2f, 0x6c, 0x04, 0x63, 0xbe, 0xe7,
	0x9e, 0xb8, 0x17, 0x0b, 0x4a, 0xec, 0xcc, 0xe6, 0x07, 0xb8, 0x9d, 0xd2, 0x94, 0x92, 0x71, 0xe5,
	0x10, 0xfb, 0xef, 0x50, 0x08, 0xd9, 0x53, 0xa3, 0x4e, 0x50, 0xb4, 0x51, 0xde, 0x66, 0xa4, 0xec,
	0x9b, 0xd9, 0x4d, 0xf3, 0x55, 0x27, 0xf2, 0x46, 0x94, 0xec, 0x8a, 0x6f, 0xa6, 0x0a, 0xd8, 0x3d,
	0x64, 0x0b, 0xd1, 0x9f, 0xe4, 0x64, 0x4f, 0x7c, 0x0f, 0xf4, 0x4c, 0xce, 0xb3, 0xda, 0x4f, 0x7d,
	0x79, 0x0f, 0x27, 0xe2, 0xbf, 0x55, 0x21, 0x64, 0xf7, 0x90, 0x91, 0x25, 0xc8, 0x13, 0xf7, 0x10,
	0x12, 0x76, 0x15, 0xbe, 0x52, 0xa8, 0x57, 0xc5, 0x55, 0x0a, 0xe1, 0xdc, 0x4a, 0x83, 0x3d, 0x16,
	0xef, 0x42, 0x31, 0x95, 0x56, 0xb6, 0x80, 0x4d, 0x95, 0x95, 0xad, 0x5b, 0xd9, 0x0a, 0xe5, 0x2b,
	0x2b, 0xdb, 0xb0, 0xb2, 0x01, 0x14, 0x28, 0x2b, 0xbb, 0x64, 0x25, 0x41, 0x21, 0x58, 0xd9, 0x05,
	0x2b, 0x40, 0xcd, 0xc0, 0xca, 0x36, 0xad, 0x34, 0xd8, 0x6b, 0x60, 0x65, 0x9b, 0x56, 0x75, 0x01,
	0x8b, 0x94, 0x55, 0x5d, 0xb7, 0xaa, 0x2b, 0x54, 0xac, 0xac, 0xea, 0x86, 0x55, 0x1d, 0x40, 0x89,
	0xb2, 0xaa, 0x97, 0xac, 0x24, 0x28, 0x05, 0xab, 0x7a, 0xc1, 0x0a, 0x50, 0xfb, 0x60, 0x55, 0x37,
	0xad, 0x34, 0xd8, 0x01, 0x58, 0xd5, 0x4d, 0xab, 0x86, 0x80, 0x65, 0xca, 0xaa, 0xa1, 0x5b, 0x35,
	0x14, 0x2a, 0x57, 0x56, 0x0d, 0xc3, 0xaa, 0x01, 0xa0, 0xd7, 0x95, 0x55, 0xa3, 0x64, 0x25, 0x41,
	0x6f, 0x80, 0x55, 0xa3, 0x60, 0x05, 0xa8, 0x37, 0xc1, 0xaa, 0x61, 0x5a, 0x69, 0xb0, 0xb7, 0xc0,
	0xaa, 0x61, 0x5a, 0xad, 0x09, 0xd8, 0xdb, 0xca, 0x6a, 0x4d, 0xb7, 0x5a, 0x53, 0xa8, 0x77, 0x94,
	0x95, 0xca, 0xf0, 0x32, 0xb7, 0x5a, 0x03, 0xd0, 0xbb, 0x96, 0xd2, 0x82, 0x14, 0x5f, 0x13, 0x5a,
	0x92, 0xf4, 0x9e, 0x05, 0x5e, 0x92, 0xb5, 0x2c, 0xbc, 0x00, 0xf6, 0xbe, 0x05, 0x62, 0x80, 0xbb,
	0x23, 0xc4, 0x34, 0xde, 0x07, 0x16, 0x98, 0x69, 0xc0, 0x2b, 0xdc, 0xac, 0x29, 0x78, 0x1f, 0x5a,
	0x4a, 0x4d, 0x24, 0xf8, 0x26, 0x57, 0x6b, 0x2a, 0xda, 0x47, 0x96, 0x72, 0x6b, 0x1a, 0x6e, 0x4d,
	0x60, 0x7d, 0x0c, 0x6e, 0xcd, 0x92, 0x9b, 0x64, 0x7d, 0xa2, 0xb9, 0x35, 0x0b, 0x6e, 0x40, 0xfb,
	0x54, 0x73, 0x6b, 0x9a, 0x6e, 0x1a, 0xef, 0x33, 0xcd, 0xad, 0x69, 0xba, 0xb5, 0x04, 0xef, 0x73,
	0x70, 0x6b, 0xe9, 0x6e, 0x2d, 0x45, 0xfb, 0x02, 0xdc, 0x5a, 0x86, 0x5b, 0x0b, 0x58, 0x5f, 0x82,
	0x5b, 0xab, 0xe4, 0x26, 0x59, 0x5f, 0x69, 0x6e, 0xad, 0x82, 0x1b, 0xd0, 0xbe, 0xd6, 0xdc, 0x5a,
	0xa6, 0x9b, 0xc6, 0xfb, 0x46, 0x73, 0x6b, 0x99, 0x6e, 0x6d, 0xc1, 0xfb, 0x16, 0xdc, 0xda, 0xba,
	0x5b, 0x5b, 0xd1, 0xbe, 0x03, 0xb7, 0xb6, 0xe1, 0xd6, 0x06, 0xd6, 0xf7, 0xe0, 0xd6, 0x2e, 0xb9,
	0x49, 0xd6, 0x0f, 0x9a, 0x5b, 0xbb, 0xe0, 0x06, 0xb4, 0x1f, 0x35, 0xb7, 0xb6, 0xe9, 0xa6, 0xf1,
	0x7e, 0xd2, 0xdc, 0xda, 0xa6, 0x5b, 0x47, 0xf0, 0x7e, 0x06, 0xb7, 0x8e, 0xee, 0xd6, 0x51, 0xb4,
	0x5f, 0xc0, 0xad, 0x63, 0xb8, 0x75, 0x80, 0xf5, 0x2b, 0xb8, 0x75, 0x4a, 0x6e, 0x92, 0xf5, 0x9b,
	0xe6, 0xd6, 0x29, 0xb8, 0x01, 0xed, 0x77, 0xcd, 0xad, 0x63, 0xba, 0x69, 0xbc, 0x3f, 0x34, 0xb7,
	0x22, 0x90, 0xfd, 0x2b, 0x5d, 0x11, 0xc0, 0x3f, 0x25, 0x10, 0x22, 0x39, 0xf8, 0xea, 0x8a, 0x02,
	0xfe, 0x05, 0x83, 0x43, 0xca, 0x80, 0x3c, 0x00, 0xe0, 0xdf, 0x12, 0x58, 0x8c, 0xf1, 0x75, 0x74,
	0x9c, 0xff, 0x3b, 0x95, 0xc4, 0x7f, 0x44, 0x97, 0x9e, 0xcd, 0x87, 0xd7, 0x98, 0x4f, 0xb4, 0xe1,
	0x35, 0xe8, 0x5d, 0x74, 0x5a, 0x24, 0x40, 0xfd, 0x57, 0x74, 0x9a, 0x79, 0x77, 0x05, 0x3d, 0xe5,
	0x57, 0x7a, 0x77, 0xfe, 0x83, 0xde, 0x61, 0xbf, 0xc2, 0xe3, 0x4d, 0xeb, 0x89, 0x65, 0x3d, 0x3a,
	0xc2, 0x7f, 0x92, 0xd7, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x01, 0xa2, 0x51, 0xf9, 0xed, 0x0f,
	0x00, 0x00,
}
