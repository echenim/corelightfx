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

// GetAll godoc
// @Summary list stocks
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} models.Stocks
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /stocks [get]
func (s *StockAPI) GetAll(ctx *gin.Context) {

	stocks := s.stockService.FindAll()
	//ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTOs(stocks)})
	ctx.JSON(http.StatusOK, mappers.ToStockDTOs(stocks))
}
