package apis

import (
	"log"
	"net/http"
	"time"

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

func (s *StockAPI) FindByName(ctx *gin.Context) {
	name := ctx.Param("name")
	stocks := s.stockService.FindByName(name)
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTO(stocks)})
}

func (s *StockAPI) SearchNameAndDuration(ctx *gin.Context) {

	name := ctx.Param("name")
	duration, _ := time.Parse("2006-01-02 15:04", ctx.Param("duration"))
	stocks := s.stockService.Search(name, duration)
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTOs(stocks)})
}

func (s *StockAPI) GetAll(ctx *gin.Context) {
	stocks := s.stockService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTOs(stocks)})
}

func (s *StockAPI) Create(ctx *gin.Context) {
	var stockDto models.StockDTO
	err := ctx.BindJSON(&stockDto)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	stocks := s.stockService.Save(mappers.ToStock(stockDto))
	ctx.JSON(http.StatusOK, gin.H{"stocks": mappers.ToStockDTO(stocks)})
}
