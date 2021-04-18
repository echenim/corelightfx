package apis

import (
	"github.com/echenim/corelightfx/services"
	"github.com/gin-gonic/gin"
)

type StockAPI struct {
	stockService services.StockService
}

func ProviderStockAPI(s services.StockService) StockAPI {
	return StockAPI{stockService: s}
}

func (s *StockAPI) FindAll(ctx *gin.Context) {

}
