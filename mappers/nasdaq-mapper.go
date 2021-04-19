package mappers

import "github.com/echenim/corelightfx/models"

func ToNasdaq(s models.NasdaqDTO) models.Nasdaq {
	return models.Nasdaq{
		ID:          s.ID,
		Symbol:      s.Symbol,
		CompanyName: s.CompanyName,
	}
}

func ToNasdaqDTO(s models.Nasdaq) models.NasdaqDTO {
	return models.NasdaqDTO{
		ID:          s.ID,
		Symbol:      s.Symbol,
		CompanyName: s.CompanyName,
	}
}

func ToNasdaqDTOs(s []models.Nasdaq) []models.NasdaqDTO {
	nasdaqdtos := make([]models.NasdaqDTO, len(s))
	for i, item := range s {
		nasdaqdtos[i] = ToNasdaqDTO(item)
	}
	return nasdaqdtos
}
