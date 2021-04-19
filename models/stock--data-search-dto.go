package models

import "time"

type StockDataSearchDTO struct {
	Symbol       string    `json:"symbol"`
	LatestSource string    `json:"latestSource"`
	LatestTime   string    `json:"latestTime"`
	Stamp        time.Time `json:"stamp"`
}
