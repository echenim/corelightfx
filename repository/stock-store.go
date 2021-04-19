package repository

import (
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
	s.DB.Find(&stock)
	return stock
}

func (s *StockStoreRespository) FindByName(name string) models.Stock {
	var stock models.Stock
	s.DB.Where("Symbol =? ", name).Find(&stock)
	return stock
}

func (s *StockStoreRespository) Search(symbol string, t time.Time) []models.Stock {
	var stock []models.Stock
	s.DB.Where("Symbol = ? OR LatestSource = ? OR Stamp >= ?", symbol, t).Find(&stock)
	return stock
}

func (s *StockStoreRespository) Save(stock models.Stock) models.Stock {
	s.DB.Save(&stock)
	return stock
}
