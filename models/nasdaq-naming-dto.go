package models

type NasdaqDTO struct {
	ID          int64  `json:"id,string,omitempty"`
	CompanyName string `json:"companyName"`
	Symbol      string `json:"symbol"`
}
