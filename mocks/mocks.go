package mocks

import (
	_ "github.com/oreuta/easytrip/controllers/bankRating"
)

type Req Request

type ComparisionService interface {
	Comparision(r *Request) (bank []CurrencyBank)
}

func Comparision(r *Request) (banks []CurrencyBank) {
	var bank CurrencyBank
	bank.BankName = "Privat"
	bank.CodeAlpha = "USD"
	bank.RateBuy = 28.15
	bank.RateSale = 28.05
	banks = append(banks, bank)
	return
}
