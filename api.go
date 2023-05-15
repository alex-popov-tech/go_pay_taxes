package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/shopspring/decimal"
)

func getRate(date time.Time, currency string) decimal.Decimal {
	fullUrl := fmt.Sprintf("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=%s&date=%s&json", currency, date.Format("20060102"))
	client := resty.New()
	resp, err := client.R().Get(fullUrl)
	if err != nil {
		panic(err)
	}

	var rates *[]Rate
	err = json.Unmarshal(resp.Body(), &rates)
	if err != nil {
		panic(err)
	}
	rate := (*rates)[0]
	return decimal.NewFromFloat32(rate.Rate)
}
