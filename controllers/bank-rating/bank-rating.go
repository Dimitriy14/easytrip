package bankRatingController

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/bank-rating"
	"github.com/oreuta/easytrip/translate"
)

//RatesController is a controller for comparing page
type RatesController struct {
	beego.Controller
	RatesService bankRatingService.RatesServiceInterface
}

//New create a new RatesController
func New(service bankRatingService.RatesServiceInterface) *RatesController {
	return &RatesController{
		RatesService: service,
	}
}

//Get function gets request gives and output data on display
func (this *RatesController) Get() {
<<<<<<< HEAD
=======

	translate := translate.New()
	lang := this.GetString("lang")
	if lang != "" {
		translate.Lang = lang
		this.Ctx.SetCookie("lang", translate.Lang)
	} else {
		translate.Lang = this.Ctx.GetCookie("lang")
		if translate.Lang == "" {
			translate.Lang = "en-US"
		}
	}
	translate.Path = "conf/locale_" + translate.Lang + ".ini"
	this.Data["i18n"] = translate.Tr

>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
	toolbox.StatisticsMap.AddStatistics("GET", "/comparision", "&controllers.bankRatingController.RatesController", time.Duration(13000))
	r := models.MainRequest{
		Currency: this.GetStrings("currency"),
		Option:   this.GetString("option"),
		Bank:     this.GetStrings("bank"),
	}
<<<<<<< HEAD
	if r.Currency == nil || r.Bank == nil {
		this.Data["IncorrectCurrency"] = true
		this.Data["IncorrectBank"] = true
=======

	if r.Currency == nil {
		this.Data["warningCurrency"] = "*Select Currency"
		this.Data["isWarnCurr"] = true
	} else {
		this.Data["isWarnCurr"] = false
	}
	if r.Bank == nil {
		this.Data["warningBank"] = "*Select Bank"
		this.Data["isWarnBank"] = true
	} else {
		this.Data["isWarnBank"] = false
	}
	if r.Currency == nil || r.Bank == nil {
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
		this.TplName = "index.tpl"
		return
	}

	b, err := this.RatesService.GetBankRates(r)
	if err != nil {
		beego.Error("Error:%v", err)
		return
	}

	this.Data["Banks"] = b

	this.Layout = "comparision_layout.tpl"
	this.TplName = "comparision.tpl"
}
