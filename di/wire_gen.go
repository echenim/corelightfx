// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"github.com/echenim/corelightfx/apis"
	"github.com/echenim/corelightfx/repository"
	"github.com/echenim/corelightfx/services"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitStockAPI(db *gorm.DB) apis.StockAPI {
	stockStoreRespository := repository.ProviderRespository(db)
	stockService := services.ProviderStockService(stockStoreRespository)
	stockAPI := apis.ProviderStockAPI(stockService)
	return stockAPI
}

func InitNasdaqAPI(db *gorm.DB) apis.NasdaqAPI {
	nasdaqRespository := repository.ProviderNasdaqRespository(db)
	nasdaqService := services.ProviderNasdaqService(nasdaqRespository)
	nasdaqAPI := apis.ProviderNasdaqAPI(nasdaqService)
	return nasdaqAPI
}
