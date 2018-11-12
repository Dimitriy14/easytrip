package bestBankService

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
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
	jsn, err := b.Client.Get()
	if err != nil {
		beego.Error("Method Get in Client BankUACient: %v", err)
		return
	}
	err = json.Unmarshal(jsn, &banks)
	if err != nil {
		beego.Error("json.Unmarshal %v:", err)
		return
	}

	return BestSale(FilterCurrency(data, FilterBank(data, banks))), err
}

func (b BestBankService) GetBestBanksBuy(data models.MainRequest) (banks []models.CurrencyBank, err error) {
	jsn, err := b.Client.Get()
	if err != nil {
		beego.Error("Method Get in Client BankUACient: %v", err)
		return
	}
	err = json.Unmarshal(jsn, &banks)
	if err != nil {
		beego.Error("json.Unmarshal %v:", err)
		return
	}
	return BestBuy(FilterCurrency(data, FilterBank(data, banks))), err
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
func BestSale(inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	var eur, usd = 999999.0, 999999.0
	for _, value := range inpBanks {
		if value.CodeAlpha == "EUR" {
			if value.RateSale < eur {
				eur = value.RateSale
			}
		}
		if value.CodeAlpha == "USD" {
			if value.RateSale < usd {
				usd = value.RateSale
			}
		}
	}
	for _, value := range inpBanks {
		if value.CodeAlpha == "EUR" && value.RateSale == eur {
			OutpBanks = append(OutpBanks, value)
		}
	}
	for _, value := range inpBanks {
		if value.CodeAlpha == "USD" && value.RateSale == usd {
			OutpBanks = append(OutpBanks, value)
		}
	}
	return OutpBanks
}

//max
func BestBuy(inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {

	var eur, usd = -1.0, -1.0
	for _, value := range inpBanks {
		if value.CodeAlpha == "EUR" {
			if value.RateSale > eur {
				eur = value.RateSale
			}
		}
		if value.CodeAlpha == "USD" {
			if value.RateSale > usd {
				usd = value.RateSale
			}
		}
	}
	for _, value := range inpBanks {
		if value.CodeAlpha == "EUR" && value.RateSale == eur {
			OutpBanks = append(OutpBanks, value)
		}
	}
	for _, value := range inpBanks {
		if value.CodeAlpha == "USD" && value.RateSale == usd {
			OutpBanks = append(OutpBanks, value)
		}
	}
	return OutpBanks

}
