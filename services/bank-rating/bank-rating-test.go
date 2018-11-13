package bankRatingService

import (
	"testing"

	"github.com/oreuta/easytrip/models"
)

func TestGetBankRates(t *testing.T) {
	r := models.MainRequest{}
	r.Bank="privat"
	r.Currency="usd"
	r.Option="buy"

	banks := []models.CurrencyBank{}
	res:=

}
