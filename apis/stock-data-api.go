package apis

import (
	"net/http"

	"github.com/echenim/corelightfx/mappers"
	"github.com/echenim/corelightfx/models"
	"github.com/echenim/corelightfx/services"
	"github.com/gin-gonic/gin"
)

type StockAPI struct {
	stockService services.StockService
}

func ProviderStockAPI(s services.StockService) StockAPI {
	return StockAPI{stockService: s}
}

func (s *StockAPI) FindByName(name string, ctx *gin.Context) {
	stocks := s.stockService.FindByName(name)
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTO(stocks)})
}

func (s *StockAPI) Search(stoc models.StockDataSearch, ctx *gin.Context) {
	stocks := s.stockService.Search(stoc)
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTOs(stocks)})
}

func (s *StockAPI) FindAll(ctx *gin.Context) {
	stocks := s.stockService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTOs(stocks)})
}
