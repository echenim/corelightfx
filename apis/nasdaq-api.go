package apis

import (
	"net/http"

	"github.com/echenim/corelightfx/mappers"
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
