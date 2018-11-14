package services

import (
	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

//BestBankServiceInterface represents a common service to interact with BankUAClient.
type BestBankServiceInterface interface {
	GetBestBankInfo(r models.MainRequest) (banks []models.CurrencyBank, err error)
}

//BestBankService implements BestBankServiceInterface interface.
type BestBankService struct {
	Client clients.BankUAClient
}

//New creates a new BestBankService instance.
func New(newClient clients.BankUAClient) BestBankServiceInterface {
	return &BestBankService{
		Client: newClient,
	}
}

/*func BestBankSale(r models.MainRequest, banks []models.CurrencyBank) (banks []models.CurrencyBank) {
Math.max(...array.map(func(obj){return obj.rateSale, obj.codeAlpha}))
 }*/

//BestBankBuy returns the best buy in banks.
func BestBankBuy(r models.MainRequest, banks []models.CurrencyBank) (banks []models.CurrencyBank) {
	for _, minValue := range banks {
		if minValue.CodeAlpha == "USD" {
			if usd < minValue.RateBuy {
				usd = minValue.RateBuy
			}
			banks = append(banks, minValue)
		}
	}
	for _, minValue := range banks {
		if minValue.CodeAlpha == "EUR" {
			if eur < minValue.RateBuy {
				eur = minValue.RateBuy
			}
			banks = append(banks, minValue)
		}
	}
	return banks
}

//BestBankSale returns the best sale in banks.
func BestBankSale(r models.MainRequest, banks []models.CurrencyBank) (banks []models.CurrencyBank) {
	for _, maxValue := range banks {
		if maxValue.CodeAlpha == "USD" {
			if usd > maxValue.RateSale {
				usd = maxValue.RateSale
			}
			banks = append(banks, maxValue)
		}
	}
	for _, maxValue := range banks {
		if maxValue.CodeAlpha == "EUR" {
			if eur > maxValue.RateSale {
				eur = maxValue.RateSale
			}
			banks = append(banks, maxValue)
		}
	}
	return banks
}
