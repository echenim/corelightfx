package di

import (
	"github.com/echenim/corelightfx/apis"
	"github.com/echenim/corelightfx/repository"
	"github.com/echenim/corelightfx/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func initStockAPI(db *gorm.DB) apis.StockAPI {
	wire.Build(repository.ProviderRespository, services.ProviderStockService, apis.ProviderStockAPI)
	return apis.StockAPI{}
}
