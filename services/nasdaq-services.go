package services

import (
	"github.com/echenim/corelightfx/models"
	"github.com/echenim/corelightfx/repository"
)

type NasdaqService struct {
	nasdaqRespository repository.NasdaqRespository
}

func ProviderNasdaqService(s repository.NasdaqRespository) NasdaqService {
	return NasdaqService{nasdaqRespository: s}
}

func (s *NasdaqService) FindAll() []models.Nasdaq {
	return s.nasdaqRespository.FindAll()
}

func (s *NasdaqService) Save(nasdaq models.Nasdaq) models.Nasdaq {
	s.nasdaqRespository.Save(nasdaq)
	return nasdaq
}
