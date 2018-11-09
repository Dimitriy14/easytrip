package bestbank

import (
	"log"
	"net/url"

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
	urlData, err := url.ParseRequestURI(this.Ctx.Request.URL.RequestURI())
	if err != nil {
		log.Fatal(err)
	}
	urlQuerry, err := url.ParseQuery(urlData.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	request:= Request{Currency : urlQuerry["currency"],Option : urlQuerry["option"],Bank : urlQuerry["bank"]}

	data:=[]CurrencyBank{}

	data=lilia(...)

	this.Data["Data"]=data
	this.TplName = "best.tpl"
}
