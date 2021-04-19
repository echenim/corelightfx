package models

import "time"

type StockDataSearch struct {
	Symbol       string
	LatestSource string
	LatestTime   string
	Stamp        time.Time
}
