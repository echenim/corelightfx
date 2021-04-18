package repository

import (
	"github.com/echenim/corelightfx/models"
	"github.com/jinzhu/gorm"
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

func (s *StockStoreRespository) Save(stock models.Stock) models.Stock {
	s.DB.Save(&stock)
	return stock
}
