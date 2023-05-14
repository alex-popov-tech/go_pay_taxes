package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/shopspring/decimal"
)

func getRates(date string) ([]Rate, error) {
	resp, err := resty.New().R().
		SetQueryString(fmt.Sprintf("date=%s&period=daily", date)).
		SetHeader("x-requested-with", "XMLHttpRequest").
		Get("https://bank.gov.ua/ua/tables/exchangerates")
	if err != nil {
		return nil, err
	}

	var rates *RatesResponse
	err = json.Unmarshal(resp.Body(), &rates)
	if err != nil {
		return nil, err
	}

	return rates.Data, nil
}

func getRate(date, currency string) decimal.Decimal {
	rates, err := getRates(date)
	if err != nil {
		panic(err)
	}
	for _, rate := range rates {
		if rate.Cc == currency {
			dec, err := decimal.NewFromString(strings.ReplaceAll(rate.Rate, ",", "."))
			if err != nil {
				panic(err)
			}
			return dec
		}
	}
	panic(errors.New(fmt.Sprintf("Can't find rate '%s'", currency)))
}
