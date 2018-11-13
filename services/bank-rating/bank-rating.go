package bankRatingService

import (
	"sort"

	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

//RatesServiceInterface represents a common service to interact with BankUAClient
type RatesServiceInterface interface {
	GetBankRates(r models.MainRequest) (banks []models.CurrencyBank, err error)
}

//GetBankRates returns list of Banks response
func (obj *BankRatingService) GetBankRates(r models.MainRequest) (banks []models.CurrencyBank, err error) {
	unpack, err := obj.Client.GetCurrBank()
	banks = getOption(r, getBanks(r, getCurrency(r, unpack)))
	return
}

//BankRatingService implements RatesServiceInterface interface
type BankRatingService struct {
	Client clients.BankUAClient //Сlient interface
}

//New creates a new RatesService instance
func New(newClient clients.BankUAClient) RatesServiceInterface {
	return &BankRatingService{
		Client: newClient,
	}
}

//valueInSlice checks if the slice consists needed string
func valueInSlice(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

//getCurrency сuts currency field object of remote Bank Service according to site request
func getCurrency(r models.MainRequest, unpacked []models.CurrencyBank) (banks []models.CurrencyBank) {
	for _, v := range r.Currency {
		if v == "usd" {
			for i := range unpacked {
				if unpacked[i].CodeAlpha == "USD" {
					banks = append(banks, unpacked[i])
				}
			}
		}

		if v == "eur" {
			for i := range unpacked {
				if unpacked[i].CodeAlpha == "EUR" {
					banks = append(banks, unpacked[i])
				}
			}
		}

	}
	return
}

//getBanks сuts bank field in object of remote Bank Service according to site request
func getBanks(r models.MainRequest, unpacked []models.CurrencyBank) []models.CurrencyBank {

	var banks []models.CurrencyBank
	if valueInSlice("privat", r.Bank) {
		for i := range unpacked {
			if unpacked[i].BankName == "ПриватБанк" {
				banks = append(banks, unpacked[i])
			}
		}
	}

	if valueInSlice("otp", r.Bank) {
		for i := range unpacked {
			if unpacked[i].BankName == "ОТП Банк" {
				banks = append(banks, unpacked[i])

			}
		}
	}

	if valueInSlice("pireus", r.Bank) {
		for i := range unpacked {
			if unpacked[i].BankName == "Піреус Банк" {
				banks = append(banks, unpacked[i])

			}
		}
	}

	return banks
}

//getOption identify type of sort according to site request
func getOption(r models.MainRequest, unpacked []models.CurrencyBank) []models.CurrencyBank {

	switch r.Option {
	case "buy":
		sort.Sort(sort.Reverse(buy(unpacked)))
		return unpacked
	case "sale":
		sort.Sort(sale(unpacked))
		return unpacked
	}
	return unpacked
}

type buy []models.CurrencyBank

func (a buy) Len() int           { return len(a) }
func (a buy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a buy) Less(i, j int) bool { return a[i].RateBuy < a[j].RateBuy }

type sale []models.CurrencyBank

func (a sale) Len() int           { return len(a) }
func (a sale) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sale) Less(i, j int) bool { return a[i].RateSale < a[j].RateSale }