package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func read(filePath string) []byte {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return bytes
}

func main() {
	fileName := os.Args[1]
	extension := filepath.Ext(fileName)

	fmt.Printf("Reading, parsing, validating file %s...", fileName)
	bytes := read(fileName)
	incomes := parse(extension, bytes)
	validate(incomes)
	fmt.Printf("done.\n")

	fmt.Println("=====INCOMES=====")
	totals := []string{}
	sum := decimal.Zero
	for _, i := range incomes {
		fmt.Printf("[%s] %s %s * ", i.Date, i.Amount.String(), i.Currency)
		rate := getRate(i.Date.Time, i.Currency)
		total := rate.Mul(i.Amount)
		totals = append(totals, total.String())
		sum = sum.Add(total)
		fmt.Printf("%s = %s\n", rate.String(), total.String())
	}
	fmt.Println("=====INCOMES=====")

	sumExpression := strings.Join(totals, " + ") + " = " + sum.String()
	fmt.Println(sumExpression)

	const PERCENT_TAX = 5
	f, _ := sum.Float64()
	tax := (PERCENT_TAX * f) / 100
	expression := fmt.Sprintf("%s / %d%% = %s", sum.String(), PERCENT_TAX, strconv.FormatFloat(tax, 'f', 2, 64))
	fmt.Println("To pay: " + expression)
}
