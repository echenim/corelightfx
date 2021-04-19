package repository

import (
	"reflect"
	"strings"

	"github.com/echenim/corelightfx/models"
	"gorm.io/gorm"
)

type NasdaqRespository struct {
	DB *gorm.DB
}

func ProviderNasdaqRespository(_DB *gorm.DB) NasdaqRespository {
	return NasdaqRespository{DB: _DB}
}

func (s *NasdaqRespository) Save(nasdaq models.Nasdaqs) models.Nasdaqs {
	s.migrate()
	s.DB.Save(&nasdaq)
	return nasdaq
}

func (s *NasdaqRespository) FindAll() []models.Nasdaqs {
	s.migrate()
	var nasdaq []models.Nasdaqs
	s.DB.Find(&nasdaq)
	return nasdaq
}

func (s *NasdaqRespository) migrate() {
	if !s.DB.Migrator().HasTable(&models.Nasdaqs{}) {
		e := s.DB.Migrator().CreateTable(&models.Nasdaqs{})
		if e == nil {
			tbname := strings.ToLower(reflect.TypeOf(&models.Nasdaqs{}).Elem().Name())
			s.DB.Exec("INSERT INTO " + tbname + "(symbol,company_name) VALUES('FB','Facebook, Inc. Common Stock');")
			s.DB.Exec("INSERT INTO " + tbname + "(symbol,company_name) VALUES('NFLX','Netflix Inc');")
			s.DB.Exec("INSERT INTO " + tbname + "(symbol,company_name) VALUES('AMZN','Amazon.com, Inc');")
			s.DB.Exec("INSERT INTO " + tbname + "(symbol,company_name) VALUES('GOOGL','Alphabet Inc Class A');")
			s.DB.Exec("INSERT INTO " + tbname + "(symbol,company_name) VALUES('AAPL','Apple In');")
		}
	}
}
