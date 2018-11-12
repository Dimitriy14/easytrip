package bestBankService

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

//BestBankServiceInterface ..
type BestBankServiceInterface interface {
	GetBestBanksSale(r models.MainRequest) (banks []models.CurrencyBank, err error)
	GetBestBanksBuy(r models.MainRequest) (banks []models.CurrencyBank, err error)
}

//BestBankService ..
type BestBankService struct {
	Client clients.BankUAClient
}

//New ..
func New(newClient clients.BankUAClient) BestBankServiceInterface {
	return &BestBankService{Client: newClient}
}

func (b BestBankService) GetBestBanksSale(data models.MainRequest) (banks []models.CurrencyBank, err error) {
	if data.Option == "buy" {
		return banks, nil
	}
	jsn, err := b.Client.Get()
	if err != nil {
		return banks, fmt.Errorf("Method Get in Client BankUACient: %v", err)
	}
	err = json.Unmarshal(jsn, &banks)
	if err != nil {
		return banks, fmt.Errorf("json.Unmarshal %v:", err)
	}
	return BestSale(data, FilterCurrency(data, FilterBank(data, banks))), nil
}

func (b BestBankService) GetBestBanksBuy(data models.MainRequest) (banks []models.CurrencyBank, err error) {
	if data.Option == "sale" {
		return banks, nil
	}
	jsn, err := b.Client.Get()
	if err != nil {
		return banks, fmt.Errorf("Method Get in Client BankUACient: %v", err)
	}
	err = json.Unmarshal(jsn, &banks)
	if err != nil {
		return banks, fmt.Errorf("Json.Unmarshal %v:", err)
	}
	return BestBuy(data, FilterCurrency(data, FilterBank(data, banks))), nil
}

func FilterBank(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	if len(data.Bank) == 0 {
		return inpBanks
	}
	s := strings.Join(data.Bank, "")
	if strings.Contains(s, "privat") {
		for _, value := range inpBanks {
			if value.BankName == "\u041f\u0440\u0438\u0432\u0430\u0442\u0411\u0430\u043d\u043a" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}
	if strings.Contains(s, "pireus") {
		for _, value := range inpBanks {
			if value.BankName == "\u041f\u0456\u0440\u0435\u0443\u0441 \u0411\u0430\u043d\u043a" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}
	if strings.Contains(s, "otp") {
		for _, value := range inpBanks {
			if value.BankName == "\u041e\u0422\u041f \u0411\u0430\u043d\u043a" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}
	return OutpBanks
}
func FilterCurrency(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	s := strings.Join(data.Currency, "")

	if strings.Contains(s, "usd") || len(data.Currency) == 0 {
		for _, value := range inpBanks {
			if value.CodeAlpha == "USD" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}

	if strings.Contains(s, "eur") || len(data.Currency) == 0 {
		for _, value := range inpBanks {
			if value.CodeAlpha == "EUR" {
				OutpBanks = append(OutpBanks, value)
			}
		}
	}

	return
}

//min
func BestSale(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	eur := 999999.0
	usd := 999999.0
	for _, value := range inpBanks {
		if value.CodeAlpha == "EUR" {
			if eur > value.RateSale {
				eur = value.RateSale
			}
		}
	}
	for _, value := range inpBanks {
		if value.CodeAlpha == "USD" {
			if usd > value.RateSale {
				usd = value.RateSale
			}
		}
	}
	for _, value := range inpBanks {
		if (value.CodeAlpha == "EUR" && value.RateSale == eur) || (value.CodeAlpha == "USD" && value.RateSale == usd) {
			OutpBanks = append(OutpBanks, value)
		}
	}
	return OutpBanks
}

//max
func BestBuy(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	eur := 0.0
	usd := 0.0
	for _, value := range inpBanks {
		if value.CodeAlpha == "EUR" {
			if eur < value.RateBuy {
				eur = value.RateBuy
			}
		}
	}
	for _, value := range inpBanks {
		if value.CodeAlpha == "USD" {
			if usd < value.RateBuy {
				usd = value.RateBuy
			}
		}
	}

	for _, value := range inpBanks {
		if (value.CodeAlpha == "EUR" && value.RateBuy == eur) || (value.CodeAlpha == "USD" && value.RateBuy == usd) {
			OutpBanks = append(OutpBanks, value)
		}
	}
	return OutpBanks
}
