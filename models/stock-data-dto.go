package models

import "time"

type StockDTO struct {
	ID                     int64     `json:"id,string,omitempty"`
	Symbol                 string    `json:"symbol"`
	CompanyName            string    `json:"companyName"`
	PrimaryExchange        string    `json:"primaryExchange"`
	CalculationPrice       string    `json:"calculationPrice"`
	Open                   float64   `json:"open"`
	OpenTime               int64     `json:"openTime"`
	OpenSource             string    `json:"openSource"`
	Close                  float64   `json:"close"`
	CloseTime              int64     `json:"closeTime"`
	CloseSource            string    `json:"closeSource"`
	High                   int64     `json:"high"`
	HighTime               int64     `json:"highTime"`
	HighSource             string    `json:"highSource"`
	Low                    float64   `json:"low"`
	LowTime                int64     `json:"lowTime"`
	LowSource              string    `json:"lowSource"`
	LatestPrice            float64   `json:"latestPrice"`
	LatestSource           string    `json:"latestSource"`
	LatestTime             string    `json:"latestTime"`
	LatestUpdate           int64     `json:"latestUpdate"`
	LatestVolume           int64     `json:"latestVolume"`
	IexRealtimePrice       float32   `json:"iexRealtimePrice"`
	IexRealtimeSize        int64     `json:"iexRealtimeSize"`
	IexLastUpdated         int64     `json:"iexLastUpdated"`
	DelayedPrice           float64   `json:"delayedPrice"`
	DelayedPriceTime       int64     `json:"delayedPriceTime"`
	OddLotDelayedPrice     float64   `json:"oddLotDelayedPrice"`
	OddLotDelayedPriceTime int64     `json:"oddLotDelayedPriceTime"`
	ExtendedPrice          float64   `json:"extendedPrice"`
	ExtendedChange         float64   `json:"extendedChange"`
	ExtendedChangePercent  float64   `json:"extendedChangePercent"`
	ExtendedPriceTime      int64     `json:"extendedPriceTime"`
	PreviousClose          float64   `json:"previousClose"`
	PreviousVolume         int64     `json:"previousVolume"`
	Change                 float32   `json:"change"`
	ChangePercent          float32   `json:"changePercent"`
	Volume                 int64     `json:"volume"`
	IexMarketPercent       float64   `json:"iexMarketPercent"`
	IexVolume              int64     `json:"iexVolume"`
	AvgTotalVolume         int64     `json:"avgTotalVolume"`
	IexBidPrice            int64     `json:"iexBidPrice"`
	IexBidSize             int64     `json:"iexBidSize"`
	IexAskPrice            int64     `json:"iexAskPrice"`
	IexAskSize             int64     `json:"iexAskSize"`
	IexOpen                int64     `json:"iexOpen"`
	IexOpenTime            int64     `json:"iexOpenTime"`
	IexClose               int64     `json:"iexClose"`
	IexCloseTime           int64     `json:"iexCloseTime"`
	MarketCap              int64     `json:"marketCap"`
	PeRatio                float64   `json:"peRatio"`
	Week52High             float64   `json:"week52High"`
	Week52Low              float64   `json:"week52Low"`
	YtdChange              float64   `json:"ytdChange"`
	LastTradeTime          int64     `json:"lastTradeTime"`
	IsUSMarketOpen         bool      `json:"isUSMarketOpen"`
	Stamp                  time.Time `json:"stamp"`
}
