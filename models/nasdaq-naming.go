package models

type Nasdaq struct {
	ID          int64  `gorm:"primaryKey"`
	CompanyName string `gorm:"unique; not null"`
	Symbol      string `gorm:"unique; not null"`
}
