package bankRating

import (
	"fmt"
	"net/url"

	//"github.com/oreuta/easytrip/mocks"

	"github.com/astaxie/beego"
)

type Request struct {
	Currency []string
	Option   []string
	Bank     []string
}

type CurrencyBank struct {
	BankName  string
	CodeAlpha string
	RateBuy   float64 `json:",string"`
	RateSale  float64 `json:",string"`
}

//CurrencyBanks is an array of CurrencyBank
type CurrencyBanks []CurrencyBank

type RatesController struct {
	beego.Controller
}

func mockComparision(r *Request) (banks []CurrencyBank) {
	var bank CurrencyBank
	bank.BankName = "Privat"
	bank.CodeAlpha = "USD"
	bank.RateBuy = 28.15
	bank.RateSale = 28.05
	banks = append(banks, bank)
	return
}

func (c *RatesController) Get() {
	uParm, err := url.Parse(c.Ctx.Request.URL.RequestURI())
	if err != nil {
		fmt.Println("Comment GET: Error parsing URL")
		return
	}
	// Get query map
	qm, err := url.ParseQuery(uParm.RawQuery)
	if err != nil {
		fmt.Println("Comment GET: Error parsing URL")
		return
	}
	var r Request
	r.Currency = qm["currency"]
	r.Option = qm["option"]
	r.Bank = qm["bank"]

	b := mockComparision(&r)

	c.Data["Banks"] = b
	c.Data["re"] = r
	c.TplName = "comparision.tpl"
}
