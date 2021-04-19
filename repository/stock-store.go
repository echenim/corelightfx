package repository

import (
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

func (s *StockStoreRespository) Search(stoc models.StockDataSearch) []models.Stock {
	var stock []models.Stock
	s.DB.Where("Symbol = ? OR LatestSource = ? OR Stamp >= ?", stoc.Symbol, stoc.LatestSource, stoc.Stamp).Find(&stock)
	return stock
}

func (s *StockStoreRespository) Save(stock models.StockDTO) models.StockDTO {
	s.DB.Save(&stock)
	return stock
}
