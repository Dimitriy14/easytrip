package bankRatingService

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

//RatesServiceInterface represents a common service to interact with BankUAClient
type RatesServiceInterface interface {
	GetBankRates(r models.MainRequest) (banks []models.CurrencyBank, err error)
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

//GetCurrency сuts currency field object of remote Bank Service according to site request
func GetCurrency(r models.MainRequest, unpacked []models.CurrencyBank) (banks []models.CurrencyBank) {
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

//GetBanks сuts bank field in object of remote Bank Service according to site request
func GetBanks(r models.MainRequest, unpacked []models.CurrencyBank) []models.CurrencyBank {

	var banks []models.CurrencyBank
	if valueInSlice("privat", r.Bank) == true {
		for i := range unpacked {
			if unpacked[i].BankName == "ПриватБанк" {
				banks = append(banks, unpacked[i])
			}
		}
	}

	if valueInSlice("otp", r.Bank) == true {
		for i := range unpacked {
			if unpacked[i].BankName == "ОТП Банк" {
				banks = append(banks, unpacked[i])

			}
		}
	}

	if valueInSlice("pireus", r.Bank) == true {
		for i := range unpacked {
			if unpacked[i].BankName == "Піреус Банк" {
				banks = append(banks, unpacked[i])

			}
		}
	}

	return banks
}

//bigger compare two floats, finds bigger one
func bigger(a, b float64) (k bool) {
	if a > b {
		k = true
	}
	return k
}

//less compare two floats, finds lower one
func less(a, b float64) (k bool) {
	if a < b {
		k = true
	}
	return k
}

//SortOnPlace sorts banks array according to site request
func SortOnPlace(unpacked []models.CurrencyBank, operation func(a, b float64) bool, str string) []models.CurrencyBank {
	if str == "buy" {

		k := len(unpacked)
		for i := 0; i < k; i++ {
			indMin := i
			for j := i; j < k; j++ {
				if operation(unpacked[j].RateBuy, unpacked[indMin].RateBuy) {
					indMin = j
				}
			}
			unpacked[i], unpacked[indMin] = unpacked[indMin], unpacked[i]
		}
		return unpacked
	}

	k := len(unpacked)
	for i := 0; i < k; i++ {
		indMin := i
		for j := i; j < k; j++ {
			if operation(unpacked[j].RateSale, unpacked[indMin].RateSale) {
				indMin = j
			}
		}
		unpacked[i], unpacked[indMin] = unpacked[indMin], unpacked[i]
	}
	return unpacked

}

//GetOption identify type of sort according to site request
func GetOption(r models.MainRequest, unpacked []models.CurrencyBank) []models.CurrencyBank {
	if r.Option == "buy" {
		return SortOnPlace(unpacked, bigger, "buy")
	}

	if r.Option == "sale" {
		return SortOnPlace(unpacked, less, "sale")
	}
	return unpacked
}

//GetBankRates returns list of Banks response
func (obj *BankRatingService) GetBankRates(r models.MainRequest) (banks []models.CurrencyBank, err error) {
	var unpack []models.CurrencyBank
	body, err := obj.Client.Get()
	if err != nil {
		return nil, fmt.Errorf("Error GetBankRates:%v", err)
	}
	err = json.Unmarshal(body, &unpack)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
		return nil, fmt.Errorf("Error GetBankRates(unmarshal):%v", err)
	}
	banks = GetOption(r, GetBanks(r, GetCurrency(r, unpack)))
	return
}
