package bestbank

import (
	"github.com/astaxie/beego"
)

type Request struct {
	Currency []string
	Option   string
	Bank     []string
}

type CurrencyBank struct {
	BankName  string
	CodeAlpha string
	RateBuy   float64 `json:",string"`
	RateSale  float64 `json:",string"`
}

type BestController struct {
	beego.Controller
}

func (this *BestController) Get() {
	/*
		request := Request{
			Currency: this.GetStrings("currency"),
			Option:   this.GetString("option"),
			Bank:     this.GetStrings("bank"),
		}
	*/
	data := []CurrencyBank{}

	tmpBank := CurrencyBank{BankName: "BankName", CodeAlpha: "CodeAlpha", RateBuy: 28.01, RateSale: 29.01}
	data = append(data, tmpBank, tmpBank)
	//data=lilia(...)

	this.Data["Data"] = data
	this.TplName = "best.tpl"
}
