# Go pay taxes

This small binary could calculate your FOP 5% taxes, taking incomes in currency you need.

## Installation

- check out latest [release](https://github.com/alex-popov-tech/go_pay_taxes/releases) for binaries
- `go get -u github.com/alex-popov-tech/go_pay_taxes`

## How to use

```sh
./go_pay_taxes ./incomes.json
# OR
./go_pay_taxes ./incomes.yaml
```

## File examples

```json
[
  { "date": "10.01.2023", "currency": "USD", "amount": 1000 },
  { "date": "31.03.2023", "currency": "EUR", "amount": 1000 }
]
```

```yaml
- date: "10.01.2023"
  currency: "USD"
  amount: 1000
- date: "31.03.2023"
  currency: "EUR"
  amount: 1000
```

## Demo

https://user-images.githubusercontent.com/21224705/238207208-34e63901-f71d-47bb-abff-bb729f1363c6.mp4
