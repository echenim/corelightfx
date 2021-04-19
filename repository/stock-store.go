package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/echenim/corelightfx/models"
	"gorm.io/gorm"
)

type StockStoreRespository struct {
	DB *gorm.DB
}

func ProviderRespository(_DB *gorm.DB) StockStoreRespository {
	return StockStoreRespository{DB: _DB}
}

func (s *StockStoreRespository) FindAll() []models.Stock {
	var stock []models.Stock
	s.DB.Find(&stock).Order("stamp desc")
	return stock
}

func (s *StockStoreRespository) FindByName(name string) models.Stock {
	var stock models.Stock
	s.DB.Where("Symbol = ? ", name).Find(&stock).Order("stamp desc")
	return stock
}

func (s *StockStoreRespository) Search(paramenter []string) []models.Stock {
	var stock []models.Stock
	s.DB.Where("Symbol IN ? ", paramenter).Find(&stock)
	return stock
}

func (s *StockStoreRespository) Save(stock models.Stock) models.Stock {
	s.DB.Save(&stock)
	return stock
}

func (s *StockStoreRespository) GetStockDataFromProvider(symbol string, token string) {
	s.getDataFromProvider(symbol, token)
}

func (s *StockStoreRespository) getDataFromProvider(symbol string, token string) {
	var rs models.ProviderStockData
	url := "https://sandbox.iexapis.com//stable/stock/" + symbol + "/quote?token=" + token

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	providerData, _ := ioutil.ReadAll(resp.Body)

	var data models.ProviderStockData
	json.Unmarshal(providerData, &data)

	rs.Symbol = data.Symbol
	rs.CompanyName = data.CompanyName
	rs.PrimaryExchange = data.PrimaryExchange
	rs.CalculationPrice = data.CalculationPrice
	rs.Open = data.Open
	rs.OpenTime = data.OpenTime
	rs.OpenSource = data.OpenSource
	rs.Close = data.Close
	rs.CloseTime = data.CloseTime
	rs.CloseSource = data.CloseSource
	rs.High = data.High
	rs.HighTime = data.HighTime
	rs.HighSource = data.HighSource
	rs.Low = data.Low
	rs.LowTime = data.LowTime
	rs.LowSource = data.LowSource
	rs.LatestPrice = data.LatestPrice
	rs.LatestSource = data.LatestSource
	rs.LatestTime = data.LatestTime
	rs.LatestUpdate = data.LatestUpdate
	rs.LatestVolume = data.LatestVolume
	rs.IexRealtimePrice = data.IexRealtimePrice
	rs.IexRealtimeSize = data.IexRealtimeSize
	rs.IexLastUpdated = data.IexLastUpdated
	rs.DelayedPrice = data.DelayedPrice
	rs.DelayedPriceTime = data.DelayedPriceTime
	rs.OddLotDelayedPrice = data.OddLotDelayedPrice
	rs.OddLotDelayedPriceTime = data.OddLotDelayedPriceTime
	rs.ExtendedPrice = data.ExtendedPrice
	rs.ExtendedChange = data.ExtendedChange
	rs.ExtendedChangePercent = data.ExtendedChangePercent
	rs.ExtendedPriceTime = data.ExtendedPriceTime
	rs.PreviousClose = data.PreviousClose
	rs.PreviousVolume = data.PreviousVolume
	rs.Change = data.Change
	rs.ChangePercent = data.ChangePercent
	rs.Volume = data.Volume
	rs.IexMarketPercent = data.IexMarketPercent
	rs.IexVolume = data.IexVolume
	rs.AvgTotalVolume = data.AvgTotalVolume
	rs.IexBidPrice = data.IexBidPrice
	rs.IexBidSize = data.IexBidSize
	rs.IexAskPrice = data.IexAskPrice
	rs.IexAskSize = data.IexAskSize
	rs.IexOpen = data.IexOpen
	rs.IexOpenTime = data.IexOpenTime
	rs.IexClose = data.IexClose
	rs.IexCloseTime = data.IexCloseTime
	rs.MarketCap = data.MarketCap
	rs.PeRatio = data.PeRatio
	rs.Week52High = data.Week52High
	rs.Week52Low = data.Week52Low
	rs.YtdChange = data.YtdChange
	rs.LastTradeTime = data.LastTradeTime
	rs.IsUSMarketOpen = data.IsUSMarketOpen
	s.save(rs)
}

func (s *StockStoreRespository) save(data models.ProviderStockData) {
	stock := models.Stock{
		Symbol:                 data.Symbol,
		CompanyName:            data.CompanyName,
		PrimaryExchange:        data.PrimaryExchange,
		CalculationPrice:       data.CalculationPrice,
		Open:                   data.Open,
		OpenTime:               data.OpenTime,
		OpenSource:             data.OpenSource,
		Close:                  data.Close,
		CloseTime:              data.CloseTime,
		CloseSource:            data.CloseSource,
		High:                   data.High,
		HighTime:               data.HighTime,
		HighSource:             data.HighSource,
		Low:                    data.Low,
		LowTime:                data.LowTime,
		LowSource:              data.LowSource,
		LatestPrice:            data.LatestPrice,
		LatestSource:           data.LatestSource,
		LatestTime:             data.LatestTime,
		LatestUpdate:           data.LatestUpdate,
		LatestVolume:           data.LatestVolume,
		IexRealtimePrice:       data.IexRealtimePrice,
		IexRealtimeSize:        data.IexRealtimeSize,
		IexLastUpdated:         data.IexLastUpdated,
		DelayedPrice:           data.DelayedPrice,
		DelayedPriceTime:       data.DelayedPriceTime,
		OddLotDelayedPrice:     data.OddLotDelayedPrice,
		OddLotDelayedPriceTime: data.OddLotDelayedPriceTime,
		ExtendedPrice:          data.ExtendedPrice,
		ExtendedChange:         data.ExtendedChange,
		ExtendedChangePercent:  data.ExtendedChangePercent,
		ExtendedPriceTime:      data.ExtendedPriceTime,
		PreviousClose:          data.PreviousClose,
		PreviousVolume:         data.PreviousVolume,
		Change:                 data.Change,
		ChangePercent:          data.ChangePercent,
		Volume:                 data.Volume,
		IexMarketPercent:       data.IexMarketPercent,
		IexVolume:              data.IexVolume,
		AvgTotalVolume:         data.AvgTotalVolume,
		IexBidPrice:            data.IexBidPrice,
		IexBidSize:             data.IexBidSize,
		IexAskPrice:            data.IexAskPrice,
		IexAskSize:             data.IexAskSize,
		IexOpen:                data.IexOpen,
		IexOpenTime:            data.IexOpenTime,
		IexClose:               data.IexClose,
		IexCloseTime:           data.IexCloseTime,
		MarketCap:              data.MarketCap,
		PeRatio:                data.PeRatio,
		Week52High:             data.Week52High,
		Week52Low:              data.Week52Low,
		YtdChange:              data.YtdChange,
		LastTradeTime:          data.LastTradeTime,
		IsUSMarketOpen:         data.IsUSMarketOpen,
		Stamp:                  time.Now(),
	}
	s.DB.Save(&stock)
}
