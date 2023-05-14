package main

import "github.com/shopspring/decimal"

type Income struct {
	Date     string          `json:"date" yaml:"date" validate:"date"`
	Currency string          `json:"currency" yaml:"currency" validate:"oneof=EUR USD"`
	Amount   decimal.Decimal `json:"amount" yaml:"amount"`
}

type Rate struct {
	ExchangeDate string  `json:"exchangedate"`
	R030         int     `json:"r030"`
	Cc           string  `json:"cc"`
	Txt          string  `json:"txt"`
	Enname       string  `json:"enname"`
	Rate         string  `json:"rate"`
	Units        int     `json:"units"`
	RatePerUnit  float64 `json:"rate_per_unit"`
	Group        string  `json:"group"`
	Calcdate     string  `json:"calcdate"`
}

type RatesResponse struct {
	Data []Rate `json:"data"`
}
