package mappers

import "github.com/echenim/corelightfx/models"

func ToNasdaq(s models.NasdaqDTO) models.Nasdaqs {
	return models.Nasdaqs{
		ID:          s.ID,
		Symbol:      s.Symbol,
		CompanyName: s.CompanyName,
	}
}

func ToNasdaqDTO(s models.Nasdaqs) models.NasdaqDTO {
	return models.NasdaqDTO{
		ID:          s.ID,
		Symbol:      s.Symbol,
		CompanyName: s.CompanyName,
	}
}

func ToNasdaqDTOs(s []models.Nasdaqs) []models.NasdaqDTO {
	nasdaqdtos := make([]models.NasdaqDTO, len(s))
	for i, item := range s {
		nasdaqdtos[i] = ToNasdaqDTO(item)
	}
	return nasdaqdtos
}
