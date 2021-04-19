package models

type Nasdaqs struct {
	ID          int64  `gorm:"primaryKey"`
	CompanyName string `gorm:"unique; not null"`
	Symbol      string `gorm:"unique; not null"`
}
