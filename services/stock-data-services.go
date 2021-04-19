package services

import (
	"github.com/echenim/corelightfx/models"
	"github.com/echenim/corelightfx/repository"
)

type StockService struct {
	stockStoreRespository repository.StockStoreRespository
}

func ProviderStockService(s repository.StockStoreRespository) StockService {
	return StockService{stockStoreRespository: s}
}

func (s *StockService) FindAll() []models.Stock {
	return s.stockStoreRespository.FindAll()
}

func (s *StockService) FindByName(name string) models.Stock {
	return s.stockStoreRespository.FindByName(name)
}

func (s *StockService) Search(paramenter []string) []models.Stock {
	return s.stockStoreRespository.Search(paramenter)
}

func (s *StockService) Save(stock models.Stock) models.Stock {
	s.stockStoreRespository.Save(stock)
	return stock
}
