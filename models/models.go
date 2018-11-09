package models

import (
	"encoding/json"
	"fmt"

	"github.com/oreuta/easytrip/clients"
)

//MainRequest gives fields for request from main page
type MainRequest struct {
	Currency []string
	Option   string
	Bank     []string
}

//CurrencyBank gives definition of banks
type CurrencyBank struct {
	BankName  string
	CodeAlpha string
	RateBuy   float64 `json:",string"`
	RateSale  float64 `json:",string"`
}

//CurrencyBanks is an array of CurrencyBank
type CurrencyBanks []CurrencyBank

func GetBanks() (bank CurrencyBanks, err error) {
	client := clients.New()
	body, err := clients.BankUAClient.Get(client)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %v", err)
	}

	err = json.Unmarshal(body, &bank)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	return
}
