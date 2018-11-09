package bankRatingService

import "github.com/oreuta/easytrip/models"

type RatesServiceIterface interface {
	GetBankRates(r *models.MainRequest) (bank []models.CurrencyBank, err error)
}

type RatesService struct{}

func New() RatesServiceIterface {
	return &RatesService{}
}

func (RatesService) GetBankRates(_ *models.MainRequest) (banks []models.CurrencyBank, err error) {
	bank := models.CurrencyBank{
		BankName:  "Privat",
		CodeAlpha: "USD",
		RateBuy:   28.15,
		RateSale:  28.05,
	}
	banks = append(banks, bank)
	return
}
