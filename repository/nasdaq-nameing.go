package repository

import (
	"github.com/echenim/corelightfx/models"
	"gorm.io/gorm"
)

type NasdaqRespository struct {
	DB *gorm.DB
}

func ProviderNasdaqRespository(_DB *gorm.DB) NasdaqRespository {
	return NasdaqRespository{DB: _DB}
}

func (s *NasdaqRespository) Save(nasdaq models.NasdaqDTO) models.NasdaqDTO {
	s.DB.Save(&nasdaq)
	return nasdaq
}

func (s *NasdaqRespository) FindAll() []models.Nasdaq {
	var nasdaq []models.Nasdaq
	s.DB.Find(&nasdaq)
	return nasdaq
}
