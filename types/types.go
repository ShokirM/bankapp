package types

import "time"

type Money int64

type Currency string

const (
	Somoni  Currency = "TJS"
	RUBLS   Currency = "RUB"
	Dollars Currency = "USD"
)

type PAN string

type Card struct {
	TrnHistory []Transaction
	CardBalance  int
	CardActivity bool
}

const (
	Active   = true
	Inactive = false
)

type Transaction struct {
	Amount Money
	Type string
	Time   time.Time
}