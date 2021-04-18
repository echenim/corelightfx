package apis

import (
	"net/http"

	"github.com/echenim/corelightfx/mappers"
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
	stocks := s.stockService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTOs(stocks)})
}
