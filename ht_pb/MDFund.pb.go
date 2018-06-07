// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MDFund.proto

/*
Package com_htsc_mdc_insight_model is a generated protocol buffer package.

It is generated from these files:
	MDFund.proto

It has these top-level messages:
	MDFund
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

// 基金
type MDFund struct {
	HTSCSecurityID       string            `protobuf:"bytes,1,opt,name=HTSCSecurityID" json:"HTSCSecurityID,omitempty"`
	MDDate               int32             `protobuf:"varint,2,opt,name=MDDate" json:"MDDate,omitempty"`
	MDTime               int32             `protobuf:"varint,3,opt,name=MDTime" json:"MDTime,omitempty"`
	DataTimestamp        int64             `protobuf:"varint,4,opt,name=DataTimestamp" json:"DataTimestamp,omitempty"`
	TradingPhaseCode     string            `protobuf:"bytes,5,opt,name=TradingPhaseCode" json:"TradingPhaseCode,omitempty"`
	SecurityIDSource     ESecurityIDSource `protobuf:"varint,6,opt,name=securityIDSource,enum=com.htsc.mdc.model.ESecurityIDSource" json:"securityIDSource,omitempty"`
	SecurityType         ESecurityType     `protobuf:"varint,7,opt,name=securityType,enum=com.htsc.mdc.model.ESecurityType" json:"securityType,omitempty"`
	MaxPx                int64             `protobuf:"varint,8,opt,name=MaxPx" json:"MaxPx,omitempty"`
	MinPx                int64             `protobuf:"varint,9,opt,name=MinPx" json:"MinPx,omitempty"`
	PreClosePx           int64             `protobuf:"varint,10,opt,name=PreClosePx" json:"PreClosePx,omitempty"`
	NumTrades            int64             `protobuf:"varint,11,opt,name=NumTrades" json:"NumTrades,omitempty"`
	TotalVolumeTrade     int64             `protobuf:"varint,12,opt,name=TotalVolumeTrade" json:"TotalVolumeTrade,omitempty"`
	TotalValueTrade      int64             `protobuf:"varint,13,opt,name=TotalValueTrade" json:"TotalValueTrade,omitempty"`
	LastPx               int64             `protobuf:"varint,14,opt,name=LastPx" json:"LastPx,omitempty"`
	OpenPx               int64             `protobuf:"varint,15,opt,name=OpenPx" json:"OpenPx,omitempty"`
	ClosePx              int64             `protobuf:"varint,16,opt,name=ClosePx" json:"ClosePx,omitempty"`
	HighPx               int64             `protobuf:"varint,17,opt,name=HighPx" json:"HighPx,omitempty"`
	LowPx                int64             `protobuf:"varint,18,opt,name=LowPx" json:"LowPx,omitempty"`
	DiffPx1              int64             `protobuf:"varint,19,opt,name=DiffPx1" json:"DiffPx1,omitempty"`
	DiffPx2              int64             `protobuf:"varint,20,opt,name=DiffPx2" json:"DiffPx2,omitempty"`
	TotalBuyQty          int64             `protobuf:"varint,21,opt,name=TotalBuyQty" json:"TotalBuyQty,omitempty"`
	TotalSellQty         int64             `protobuf:"varint,22,opt,name=TotalSellQty" json:"TotalSellQty,omitempty"`
	WeightedAvgBuyPx     int64             `protobuf:"varint,23,opt,name=WeightedAvgBuyPx" json:"WeightedAvgBuyPx,omitempty"`
	WeightedAvgSellPx    int64             `protobuf:"varint,24,opt,name=WeightedAvgSellPx" json:"WeightedAvgSellPx,omitempty"`
	WithdrawBuyNumber    int64             `protobuf:"varint,25,opt,name=WithdrawBuyNumber" json:"WithdrawBuyNumber,omitempty"`
	WithdrawBuyAmount    int64             `protobuf:"varint,26,opt,name=WithdrawBuyAmount" json:"WithdrawBuyAmount,omitempty"`
	WithdrawBuyMoney     int64             `protobuf:"varint,27,opt,name=WithdrawBuyMoney" json:"WithdrawBuyMoney,omitempty"`
	WithdrawSellNumber   int64             `protobuf:"varint,28,opt,name=WithdrawSellNumber" json:"WithdrawSellNumber,omitempty"`
	WithdrawSellAmount   int64             `protobuf:"varint,29,opt,name=WithdrawSellAmount" json:"WithdrawSellAmount,omitempty"`
	WithdrawSellMoney    int64             `protobuf:"varint,30,opt,name=WithdrawSellMoney" json:"WithdrawSellMoney,omitempty"`
	TotalBuyNumber       int64             `protobuf:"varint,31,opt,name=TotalBuyNumber" json:"TotalBuyNumber,omitempty"`
	TotalSellNumber      int64             `protobuf:"varint,32,opt,name=TotalSellNumber" json:"TotalSellNumber,omitempty"`
	BuyTradeMaxDuration  int64             `protobuf:"varint,33,opt,name=BuyTradeMaxDuration" json:"BuyTradeMaxDuration,omitempty"`
	SellTradeMaxDuration int64             `protobuf:"varint,34,opt,name=SellTradeMaxDuration" json:"SellTradeMaxDuration,omitempty"`
	NumBuyOrders         int32             `protobuf:"varint,35,opt,name=NumBuyOrders" json:"NumBuyOrders,omitempty"`
	NumSellOrders        int32             `protobuf:"varint,36,opt,name=NumSellOrders" json:"NumSellOrders,omitempty"`
	IOPV                 int64             `protobuf:"varint,37,opt,name=IOPV" json:"IOPV,omitempty"`
	PreIOPV              int64             `protobuf:"varint,38,opt,name=PreIOPV" json:"PreIOPV,omitempty"`
	PurchaseNumber       int64             `protobuf:"varint,39,opt,name=PurchaseNumber" json:"PurchaseNumber,omitempty"`
	PurchaseAmount       int64             `protobuf:"varint,40,opt,name=PurchaseAmount" json:"PurchaseAmount,omitempty"`
	PurchaseMoney        int64             `protobuf:"varint,41,opt,name=PurchaseMoney" json:"PurchaseMoney,omitempty"`
	RedemptionNumber     int64             `protobuf:"varint,42,opt,name=RedemptionNumber" json:"RedemptionNumber,omitempty"`
	RedemptionAmount     int64             `protobuf:"varint,43,opt,name=RedemptionAmount" json:"RedemptionAmount,omitempty"`
	RedemptionMoney      int64             `protobuf:"varint,44,opt,name=RedemptionMoney" json:"RedemptionMoney,omitempty"`
	BuyOrderQueue        []int64           `protobuf:"varint,100,rep,packed,name=BuyOrderQueue" json:"BuyOrderQueue,omitempty"`
	SellOrderQueue       []int64           `protobuf:"varint,101,rep,packed,name=SellOrderQueue" json:"SellOrderQueue,omitempty"`
	Buy1Price            int64             `protobuf:"varint,102,opt,name=Buy1Price" json:"Buy1Price,omitempty"`
	Buy1OrderQty         int64             `protobuf:"varint,103,opt,name=Buy1OrderQty" json:"Buy1OrderQty,omitempty"`
	Buy1NumOrders        int64             `protobuf:"varint,104,opt,name=Buy1NumOrders" json:"Buy1NumOrders,omitempty"`
	Sell1Price           int64             `protobuf:"varint,105,opt,name=Sell1Price" json:"Sell1Price,omitempty"`
	Sell1OrderQty        int64             `protobuf:"varint,106,opt,name=Sell1OrderQty" json:"Sell1OrderQty,omitempty"`
	Sell1NumOrders       int64             `protobuf:"varint,107,opt,name=Sell1NumOrders" json:"Sell1NumOrders,omitempty"`
	Buy2Price            int64             `protobuf:"varint,108,opt,name=Buy2Price" json:"Buy2Price,omitempty"`
	Buy2OrderQty         int64             `protobuf:"varint,109,opt,name=Buy2OrderQty" json:"Buy2OrderQty,omitempty"`
	Buy2NumOrders        int64             `protobuf:"varint,110,opt,name=Buy2NumOrders" json:"Buy2NumOrders,omitempty"`
	Sell2Price           int64             `protobuf:"varint,111,opt,name=Sell2Price" json:"Sell2Price,omitempty"`
	Sell2OrderQty        int64             `protobuf:"varint,112,opt,name=Sell2OrderQty" json:"Sell2OrderQty,omitempty"`
	Sell2NumOrders       int64             `protobuf:"varint,113,opt,name=Sell2NumOrders" json:"Sell2NumOrders,omitempty"`
	Buy3Price            int64             `protobuf:"varint,114,opt,name=Buy3Price" json:"Buy3Price,omitempty"`
	Buy3OrderQty         int64             `protobuf:"varint,115,opt,name=Buy3OrderQty" json:"Buy3OrderQty,omitempty"`
	Buy3NumOrders        int64             `protobuf:"varint,116,opt,name=Buy3NumOrders" json:"Buy3NumOrders,omitempty"`
	Sell3Price           int64             `protobuf:"varint,117,opt,name=Sell3Price" json:"Sell3Price,omitempty"`
	Sell3OrderQty        int64             `protobuf:"varint,118,opt,name=Sell3OrderQty" json:"Sell3OrderQty,omitempty"`
	Sell3NumOrders       int64             `protobuf:"varint,119,opt,name=Sell3NumOrders" json:"Sell3NumOrders,omitempty"`
	Buy4Price            int64             `protobuf:"varint,120,opt,name=Buy4Price" json:"Buy4Price,omitempty"`
	Buy4OrderQty         int64             `protobuf:"varint,121,opt,name=Buy4OrderQty" json:"Buy4OrderQty,omitempty"`
	Buy4NumOrders        int64             `protobuf:"varint,122,opt,name=Buy4NumOrders" json:"Buy4NumOrders,omitempty"`
	Sell4Price           int64             `protobuf:"varint,123,opt,name=Sell4Price" json:"Sell4Price,omitempty"`
	Sell4OrderQty        int64             `protobuf:"varint,124,opt,name=Sell4OrderQty" json:"Sell4OrderQty,omitempty"`
	Sell4NumOrders       int64             `protobuf:"varint,125,opt,name=Sell4NumOrders" json:"Sell4NumOrders,omitempty"`
	Buy5Price            int64             `protobuf:"varint,126,opt,name=Buy5Price" json:"Buy5Price,omitempty"`
	Buy5OrderQty         int64             `protobuf:"varint,127,opt,name=Buy5OrderQty" json:"Buy5OrderQty,omitempty"`
	Buy5NumOrders        int64             `protobuf:"varint,128,opt,name=Buy5NumOrders" json:"Buy5NumOrders,omitempty"`
	Sell5Price           int64             `protobuf:"varint,129,opt,name=Sell5Price" json:"Sell5Price,omitempty"`
	Sell5OrderQty        int64             `protobuf:"varint,130,opt,name=Sell5OrderQty" json:"Sell5OrderQty,omitempty"`
	Sell5NumOrders       int64             `protobuf:"varint,131,opt,name=Sell5NumOrders" json:"Sell5NumOrders,omitempty"`
	Buy6Price            int64             `protobuf:"varint,132,opt,name=Buy6Price" json:"Buy6Price,omitempty"`
	Buy6OrderQty         int64             `protobuf:"varint,133,opt,name=Buy6OrderQty" json:"Buy6OrderQty,omitempty"`
	Buy6NumOrders        int64             `protobuf:"varint,134,opt,name=Buy6NumOrders" json:"Buy6NumOrders,omitempty"`
	Sell6Price           int64             `protobuf:"varint,135,opt,name=Sell6Price" json:"Sell6Price,omitempty"`
	Sell6OrderQty        int64             `protobuf:"varint,136,opt,name=Sell6OrderQty" json:"Sell6OrderQty,omitempty"`
	Sell6NumOrders       int64             `protobuf:"varint,137,opt,name=Sell6NumOrders" json:"Sell6NumOrders,omitempty"`
	Buy7Price            int64             `protobuf:"varint,138,opt,name=Buy7Price" json:"Buy7Price,omitempty"`
	Buy7OrderQty         int64             `protobuf:"varint,139,opt,name=Buy7OrderQty" json:"Buy7OrderQty,omitempty"`
	Buy7NumOrders        int64             `protobuf:"varint,140,opt,name=Buy7NumOrders" json:"Buy7NumOrders,omitempty"`
	Sell7Price           int64             `protobuf:"varint,141,opt,name=Sell7Price" json:"Sell7Price,omitempty"`
	Sell7OrderQty        int64             `protobuf:"varint,142,opt,name=Sell7OrderQty" json:"Sell7OrderQty,omitempty"`
	Sell7NumOrders       int64             `protobuf:"varint,143,opt,name=Sell7NumOrders" json:"Sell7NumOrders,omitempty"`
	Buy8Price            int64             `protobuf:"varint,144,opt,name=Buy8Price" json:"Buy8Price,omitempty"`
	Buy8OrderQty         int64             `protobuf:"varint,145,opt,name=Buy8OrderQty" json:"Buy8OrderQty,omitempty"`
	Buy8NumOrders        int64             `protobuf:"varint,146,opt,name=Buy8NumOrders" json:"Buy8NumOrders,omitempty"`
	Sell8Price           int64             `protobuf:"varint,147,opt,name=Sell8Price" json:"Sell8Price,omitempty"`
	Sell8OrderQty        int64             `protobuf:"varint,148,opt,name=Sell8OrderQty" json:"Sell8OrderQty,omitempty"`
	Sell8NumOrders       int64             `protobuf:"varint,149,opt,name=Sell8NumOrders" json:"Sell8NumOrders,omitempty"`
	Buy9Price            int64             `protobuf:"varint,150,opt,name=Buy9Price" json:"Buy9Price,omitempty"`
	Buy9OrderQty         int64             `protobuf:"varint,151,opt,name=Buy9OrderQty" json:"Buy9OrderQty,omitempty"`
	Buy9NumOrders        int64             `protobuf:"varint,152,opt,name=Buy9NumOrders" json:"Buy9NumOrders,omitempty"`
	Sell9Price           int64             `protobuf:"varint,153,opt,name=Sell9Price" json:"Sell9Price,omitempty"`
	Sell9OrderQty        int64             `protobuf:"varint,154,opt,name=Sell9OrderQty" json:"Sell9OrderQty,omitempty"`
	Sell9NumOrders       int64             `protobuf:"varint,155,opt,name=Sell9NumOrders" json:"Sell9NumOrders,omitempty"`
	Buy10Price           int64             `protobuf:"varint,156,opt,name=Buy10Price" json:"Buy10Price,omitempty"`
	Buy10OrderQty        int64             `protobuf:"varint,157,opt,name=Buy10OrderQty" json:"Buy10OrderQty,omitempty"`
	Buy10NumOrders       int64             `protobuf:"varint,158,opt,name=Buy10NumOrders" json:"Buy10NumOrders,omitempty"`
	Sell10Price          int64             `protobuf:"varint,159,opt,name=Sell10Price" json:"Sell10Price,omitempty"`
	Sell10OrderQty       int64             `protobuf:"varint,160,opt,name=Sell10OrderQty" json:"Sell10OrderQty,omitempty"`
	Sell10NumOrders      int64             `protobuf:"varint,163,opt,name=Sell10NumOrders" json:"Sell10NumOrders,omitempty"`
}

func (m *MDFund) Reset()         { *m = MDFund{} }
func (m *MDFund) String() string { return proto.CompactTextString(m) }
func (*MDFund) ProtoMessage()    {}

func (m *MDFund) GetHTSCSecurityID() string {
	if m != nil {
		return m.HTSCSecurityID
	}
	return ""
}

func (m *MDFund) GetMDDate() int32 {
	if m != nil {
		return m.MDDate
	}
	return 0
}

func (m *MDFund) GetMDTime() int32 {
	if m != nil {
		return m.MDTime
	}
	return 0
}

func (m *MDFund) GetDataTimestamp() int64 {
	if m != nil {
		return m.DataTimestamp
	}
	return 0
}

func (m *MDFund) GetTradingPhaseCode() string {
	if m != nil {
		return m.TradingPhaseCode
	}
	return ""
}

func (m *MDFund) GetSecurityIDSource() ESecurityIDSource {
	if m != nil {
		return m.SecurityIDSource
	}
	return ESecurityIDSource_DefaultSecurityIDSource
}

func (m *MDFund) GetSecurityType() ESecurityType {
	if m != nil {
		return m.SecurityType
	}
	return ESecurityType_DefaultSecurityType
}

func (m *MDFund) GetMaxPx() int64 {
	if m != nil {
		return m.MaxPx
	}
	return 0
}

func (m *MDFund) GetMinPx() int64 {
	if m != nil {
		return m.MinPx
	}
	return 0
}

func (m *MDFund) GetPreClosePx() int64 {
	if m != nil {
		return m.PreClosePx
	}
	return 0
}

func (m *MDFund) GetNumTrades() int64 {
	if m != nil {
		return m.NumTrades
	}
	return 0
}

func (m *MDFund) GetTotalVolumeTrade() int64 {
	if m != nil {
		return m.TotalVolumeTrade
	}
	return 0
}

func (m *MDFund) GetTotalValueTrade() int64 {
	if m != nil {
		return m.TotalValueTrade
	}
	return 0
}

func (m *MDFund) GetLastPx() int64 {
	if m != nil {
		return m.LastPx
	}
	return 0
}

func (m *MDFund) GetOpenPx() int64 {
	if m != nil {
		return m.OpenPx
	}
	return 0
}

func (m *MDFund) GetClosePx() int64 {
	if m != nil {
		return m.ClosePx
	}
	return 0
}

func (m *MDFund) GetHighPx() int64 {
	if m != nil {
		return m.HighPx
	}
	return 0
}

func (m *MDFund) GetLowPx() int64 {
	if m != nil {
		return m.LowPx
	}
	return 0
}

func (m *MDFund) GetDiffPx1() int64 {
	if m != nil {
		return m.DiffPx1
	}
	return 0
}

func (m *MDFund) GetDiffPx2() int64 {
	if m != nil {
		return m.DiffPx2
	}
	return 0
}

func (m *MDFund) GetTotalBuyQty() int64 {
	if m != nil {
		return m.TotalBuyQty
	}
	return 0
}

func (m *MDFund) GetTotalSellQty() int64 {
	if m != nil {
		return m.TotalSellQty
	}
	return 0
}

func (m *MDFund) GetWeightedAvgBuyPx() int64 {
	if m != nil {
		return m.WeightedAvgBuyPx
	}
	return 0
}

func (m *MDFund) GetWeightedAvgSellPx() int64 {
	if m != nil {
		return m.WeightedAvgSellPx
	}
	return 0
}

func (m *MDFund) GetWithdrawBuyNumber() int64 {
	if m != nil {
		return m.WithdrawBuyNumber
	}
	return 0
}

func (m *MDFund) GetWithdrawBuyAmount() int64 {
	if m != nil {
		return m.WithdrawBuyAmount
	}
	return 0
}

func (m *MDFund) GetWithdrawBuyMoney() int64 {
	if m != nil {
		return m.WithdrawBuyMoney
	}
	return 0
}

func (m *MDFund) GetWithdrawSellNumber() int64 {
	if m != nil {
		return m.WithdrawSellNumber
	}
	return 0
}

func (m *MDFund) GetWithdrawSellAmount() int64 {
	if m != nil {
		return m.WithdrawSellAmount
	}
	return 0
}

func (m *MDFund) GetWithdrawSellMoney() int64 {
	if m != nil {
		return m.WithdrawSellMoney
	}
	return 0
}

func (m *MDFund) GetTotalBuyNumber() int64 {
	if m != nil {
		return m.TotalBuyNumber
	}
	return 0
}

func (m *MDFund) GetTotalSellNumber() int64 {
	if m != nil {
		return m.TotalSellNumber
	}
	return 0
}

func (m *MDFund) GetBuyTradeMaxDuration() int64 {
	if m != nil {
		return m.BuyTradeMaxDuration
	}
	return 0
}

func (m *MDFund) GetSellTradeMaxDuration() int64 {
	if m != nil {
		return m.SellTradeMaxDuration
	}
	return 0
}

func (m *MDFund) GetNumBuyOrders() int32 {
	if m != nil {
		return m.NumBuyOrders
	}
	return 0
}

func (m *MDFund) GetNumSellOrders() int32 {
	if m != nil {
		return m.NumSellOrders
	}
	return 0
}

func (m *MDFund) GetIOPV() int64 {
	if m != nil {
		return m.IOPV
	}
	return 0
}

func (m *MDFund) GetPreIOPV() int64 {
	if m != nil {
		return m.PreIOPV
	}
	return 0
}

func (m *MDFund) GetPurchaseNumber() int64 {
	if m != nil {
		return m.PurchaseNumber
	}
	return 0
}

func (m *MDFund) GetPurchaseAmount() int64 {
	if m != nil {
		return m.PurchaseAmount
	}
	return 0
}

func (m *MDFund) GetPurchaseMoney() int64 {
	if m != nil {
		return m.PurchaseMoney
	}
	return 0
}

func (m *MDFund) GetRedemptionNumber() int64 {
	if m != nil {
		return m.RedemptionNumber
	}
	return 0
}

func (m *MDFund) GetRedemptionAmount() int64 {
	if m != nil {
		return m.RedemptionAmount
	}
	return 0
}

func (m *MDFund) GetRedemptionMoney() int64 {
	if m != nil {
		return m.RedemptionMoney
	}
	return 0
}

func (m *MDFund) GetBuyOrderQueue() []int64 {
	if m != nil {
		return m.BuyOrderQueue
	}
	return nil
}

func (m *MDFund) GetSellOrderQueue() []int64 {
	if m != nil {
		return m.SellOrderQueue
	}
	return nil
}

func (m *MDFund) GetBuy1Price() int64 {
	if m != nil {
		return m.Buy1Price
	}
	return 0
}

func (m *MDFund) GetBuy1OrderQty() int64 {
	if m != nil {
		return m.Buy1OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy1NumOrders() int64 {
	if m != nil {
		return m.Buy1NumOrders
	}
	return 0
}

func (m *MDFund) GetSell1Price() int64 {
	if m != nil {
		return m.Sell1Price
	}
	return 0
}

func (m *MDFund) GetSell1OrderQty() int64 {
	if m != nil {
		return m.Sell1OrderQty
	}
	return 0
}

func (m *MDFund) GetSell1NumOrders() int64 {
	if m != nil {
		return m.Sell1NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy2Price() int64 {
	if m != nil {
		return m.Buy2Price
	}
	return 0
}

func (m *MDFund) GetBuy2OrderQty() int64 {
	if m != nil {
		return m.Buy2OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy2NumOrders() int64 {
	if m != nil {
		return m.Buy2NumOrders
	}
	return 0
}

func (m *MDFund) GetSell2Price() int64 {
	if m != nil {
		return m.Sell2Price
	}
	return 0
}

func (m *MDFund) GetSell2OrderQty() int64 {
	if m != nil {
		return m.Sell2OrderQty
	}
	return 0
}

func (m *MDFund) GetSell2NumOrders() int64 {
	if m != nil {
		return m.Sell2NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy3Price() int64 {
	if m != nil {
		return m.Buy3Price
	}
	return 0
}

func (m *MDFund) GetBuy3OrderQty() int64 {
	if m != nil {
		return m.Buy3OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy3NumOrders() int64 {
	if m != nil {
		return m.Buy3NumOrders
	}
	return 0
}

func (m *MDFund) GetSell3Price() int64 {
	if m != nil {
		return m.Sell3Price
	}
	return 0
}

func (m *MDFund) GetSell3OrderQty() int64 {
	if m != nil {
		return m.Sell3OrderQty
	}
	return 0
}

func (m *MDFund) GetSell3NumOrders() int64 {
	if m != nil {
		return m.Sell3NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy4Price() int64 {
	if m != nil {
		return m.Buy4Price
	}
	return 0
}

func (m *MDFund) GetBuy4OrderQty() int64 {
	if m != nil {
		return m.Buy4OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy4NumOrders() int64 {
	if m != nil {
		return m.Buy4NumOrders
	}
	return 0
}

func (m *MDFund) GetSell4Price() int64 {
	if m != nil {
		return m.Sell4Price
	}
	return 0
}

func (m *MDFund) GetSell4OrderQty() int64 {
	if m != nil {
		return m.Sell4OrderQty
	}
	return 0
}

func (m *MDFund) GetSell4NumOrders() int64 {
	if m != nil {
		return m.Sell4NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy5Price() int64 {
	if m != nil {
		return m.Buy5Price
	}
	return 0
}

func (m *MDFund) GetBuy5OrderQty() int64 {
	if m != nil {
		return m.Buy5OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy5NumOrders() int64 {
	if m != nil {
		return m.Buy5NumOrders
	}
	return 0
}

func (m *MDFund) GetSell5Price() int64 {
	if m != nil {
		return m.Sell5Price
	}
	return 0
}

func (m *MDFund) GetSell5OrderQty() int64 {
	if m != nil {
		return m.Sell5OrderQty
	}
	return 0
}

func (m *MDFund) GetSell5NumOrders() int64 {
	if m != nil {
		return m.Sell5NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy6Price() int64 {
	if m != nil {
		return m.Buy6Price
	}
	return 0
}

func (m *MDFund) GetBuy6OrderQty() int64 {
	if m != nil {
		return m.Buy6OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy6NumOrders() int64 {
	if m != nil {
		return m.Buy6NumOrders
	}
	return 0
}

func (m *MDFund) GetSell6Price() int64 {
	if m != nil {
		return m.Sell6Price
	}
	return 0
}

func (m *MDFund) GetSell6OrderQty() int64 {
	if m != nil {
		return m.Sell6OrderQty
	}
	return 0
}

func (m *MDFund) GetSell6NumOrders() int64 {
	if m != nil {
		return m.Sell6NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy7Price() int64 {
	if m != nil {
		return m.Buy7Price
	}
	return 0
}

func (m *MDFund) GetBuy7OrderQty() int64 {
	if m != nil {
		return m.Buy7OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy7NumOrders() int64 {
	if m != nil {
		return m.Buy7NumOrders
	}
	return 0
}

func (m *MDFund) GetSell7Price() int64 {
	if m != nil {
		return m.Sell7Price
	}
	return 0
}

func (m *MDFund) GetSell7OrderQty() int64 {
	if m != nil {
		return m.Sell7OrderQty
	}
	return 0
}

func (m *MDFund) GetSell7NumOrders() int64 {
	if m != nil {
		return m.Sell7NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy8Price() int64 {
	if m != nil {
		return m.Buy8Price
	}
	return 0
}

func (m *MDFund) GetBuy8OrderQty() int64 {
	if m != nil {
		return m.Buy8OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy8NumOrders() int64 {
	if m != nil {
		return m.Buy8NumOrders
	}
	return 0
}

func (m *MDFund) GetSell8Price() int64 {
	if m != nil {
		return m.Sell8Price
	}
	return 0
}

func (m *MDFund) GetSell8OrderQty() int64 {
	if m != nil {
		return m.Sell8OrderQty
	}
	return 0
}

func (m *MDFund) GetSell8NumOrders() int64 {
	if m != nil {
		return m.Sell8NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy9Price() int64 {
	if m != nil {
		return m.Buy9Price
	}
	return 0
}

func (m *MDFund) GetBuy9OrderQty() int64 {
	if m != nil {
		return m.Buy9OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy9NumOrders() int64 {
	if m != nil {
		return m.Buy9NumOrders
	}
	return 0
}

func (m *MDFund) GetSell9Price() int64 {
	if m != nil {
		return m.Sell9Price
	}
	return 0
}

func (m *MDFund) GetSell9OrderQty() int64 {
	if m != nil {
		return m.Sell9OrderQty
	}
	return 0
}

func (m *MDFund) GetSell9NumOrders() int64 {
	if m != nil {
		return m.Sell9NumOrders
	}
	return 0
}

func (m *MDFund) GetBuy10Price() int64 {
	if m != nil {
		return m.Buy10Price
	}
	return 0
}

func (m *MDFund) GetBuy10OrderQty() int64 {
	if m != nil {
		return m.Buy10OrderQty
	}
	return 0
}

func (m *MDFund) GetBuy10NumOrders() int64 {
	if m != nil {
		return m.Buy10NumOrders
	}
	return 0
}

func (m *MDFund) GetSell10Price() int64 {
	if m != nil {
		return m.Sell10Price
	}
	return 0
}

func (m *MDFund) GetSell10OrderQty() int64 {
	if m != nil {
		return m.Sell10OrderQty
	}
	return 0
}

func (m *MDFund) GetSell10NumOrders() int64 {
	if m != nil {
		return m.Sell10NumOrders
	}
	return 0
}

func init() {
	proto.RegisterType((*MDFund)(nil), "com.htsc.mdc.insight.model.MDFund")
}