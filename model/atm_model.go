package model

import "time"

type Account struct {
	PIN         string
	Transaction []Transaction
	Balance     float64
}

type Transaction struct {
	Type      string
	Amount    float64
	Timestamp time.Time
}
