package apis

import (
	"log"
	"net/http"

	"github.com/echenim/corelightfx/mappers"
	"github.com/echenim/corelightfx/models"
	"github.com/echenim/corelightfx/services"
	"github.com/gin-gonic/gin"
)

type NasdaqAPI struct {
	nasdaqService services.NasdaqService
}

func ProviderNasdaqAPI(s services.NasdaqService) NasdaqAPI {
	return NasdaqAPI{nasdaqService: s}
}

func (s *NasdaqAPI) GetAll(ctx *gin.Context) {
	nasdaq := s.nasdaqService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"nasdaq": mappers.ToNasdaqDTOs(nasdaq)})
}

func (s *NasdaqAPI) Create(ctx *gin.Context) {
	var nasdaqDto models.NasdaqDTO
	err := ctx.BindJSON(&nasdaqDto)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	nasdaq := s.nasdaqService.Save(mappers.ToNasdaq(nasdaqDto))
	ctx.JSON(http.StatusOK, gin.H{"nasdaq": mappers.ToNasdaqDTO(nasdaq)})
}
