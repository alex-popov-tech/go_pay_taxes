package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func validate(incomes []Income) {
	v := validator.New()
	// Register a custom validation function for the "date" tag
	for _, i := range incomes {
		err := v.Struct(i)
		if err != nil {
			fmt.Printf("\nDate:%s, currency:%s, amount:%s\n, error:%s", i.Date, i.Currency, i.Amount.String(), err)
			panic(err)
		}
	}
}
