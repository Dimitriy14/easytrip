package services

import (
	"fmt"
	"net/http"
)

//var usdBank = [3]float64{bankPrivatBuy, bankPireusBuy, bankOTPBuy}
//var eurBank = [3]float64{bankPrivatSale, bankPireusSale, bankOTPSale}

func UsdBankBestBuy(bankPrivatBuy, bankPireusBuy, bankOTPBuy float64) func(float64) float64 {

	if bankPrivatBuy > bankPireusBuy || bankPrivatBuy > bankOTPBuy {
		fmt.Print(bankPrivatBuy)
	} else if bankPireusBuy > bankPrivatBuy || bankPireusBuy > bankOTPBuy {
		fmt.Print(bankPireusBuy)
	} else {
		fmt.Print(bankOTPBuy)
	}
}
func EurBankBestBuy(bankPrivatBuy, bankPireusBuy, bankOTPBuy float64) func(float64) float64 {

	if bankPrivatBuy > bankPireusBuy || bankPrivatBuy > bankOTPBuy {
		fmt.Print(bankPrivatBuy)
	} else if bankPireusBuy > bankPrivatBuy || bankPireusBuy > bankOTPBuy {
		fmt.Print(bankPireusBuy)
	} else {
		fmt.Print(bankOTPBuy)
	}
}
func UsdBankBestSale(bankPrivatSale, bankPireusSale, bankOTPSale float64) func(float64) float64 {

	if bankPrivatSale < bankPireusSale || bankPrivatSale < bankOTPSale {
		fmt.Print(bankPrivatSale)
	} else if bankPireusSale < bankPrivatSale || bankPireusSale < bankOTPSale {
		fmt.Print(bankPireusSale)
	} else {
		fmt.Print(bankOTPSale)
	}
}
func EurBankBestSale(bankPrivatSale, bankPireusSale, bankOTPSale float64) func(float64) float64 {

	if bankPrivatSale < bankPireusSale || bankPrivatSale < bankOTPSale {
		fmt.Print(bankPrivatSale)
	} else if bankPireusSale < bankPrivatSale || bankPireusSale < bankOTPSale {
		fmt.Print(bankPireusSale)
	} else {
		fmt.Print(bankOTPSale)
	}
}
func NewCurrencyBank(BankName string, CodeAlpha string, RateBuy float64, RateSale float64) RatesController {
	return &CurrencyBank{
		BankName,
		CodeAlpha,
		RateBuy,
		RateSale,
	}
}

func main() {

	http.HandleFunc("/usdbestbuy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "The best buy USD")
	})
	http.HandleFunc("/usdbestsale", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "The best sale USD")
	})
	http.HandleFunc("/eurobestbuy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "The best buy EURO")
	})
	http.HandleFunc("/eurobestsale", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "The best sale EURO")
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8080", nil)
}
