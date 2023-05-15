package main

import (
	"time"

	"github.com/shopspring/decimal"
)

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"02.01.2006"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func (t *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var dateStr string
	if err := unmarshal(&dateStr); err != nil {
		return err
	}

	date, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		return err
	}

	t.Time = date
	return nil
}

type Income struct {
	Date     Date            `json:"date" yaml:"date" validate:"required,lt"`
	Currency string          `json:"currency" yaml:"currency" validate:"required"`
	Amount   decimal.Decimal `json:"amount" yaml:"amount"`
}

// {"r030":840,"txt":"Долар США","rate":36.5686,"cc":"USD","exchangedate":"10.01.2023"}
type Rate struct {
	ExchangeDate string  `json:"exchangedate"`
	R030         int     `json:"r030"`
	Cc           string  `json:"cc"`
	Txt          string  `json:"txt"`
	Rate         float32 `json:"rate"`
}
