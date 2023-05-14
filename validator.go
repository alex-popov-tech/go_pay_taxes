package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

func validate(incomes []Income) {
	v := validator.New()
	// Register a custom validation function for the "date" tag
	v.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		date := fl.Field().String()
		_, err := time.Parse("02.01.2006", date)
		return err == nil
	})
	v.RegisterValidation("floatformat", func(fl validator.FieldLevel) bool {
		number := fl.Field().String()
		_, err := strconv.ParseFloat(number, 64)
		return err == nil
	})
	for _, i := range incomes {
		err := v.Struct(i)
		if err != nil {
			fmt.Printf("\nDate:%s, currency:%s, amount:%s\n, error:%s", i.Date, i.Currency, i.Amount.String(), err)
			panic(err)
		}
	}
}
