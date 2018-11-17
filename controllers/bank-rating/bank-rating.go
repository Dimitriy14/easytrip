package bankRatingController

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/bank-rating"
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
	r := models.MainRequest{
		Currency: this.GetStrings("currency"),
		Option:   this.GetString("option"),
		Bank:     this.GetStrings("bank"),
	}
<<<<<<< HEAD
	if r.Currency == nil {
		this.Data["IncorrectCurrency"] = true
		this.TplName = "index.tpl"
		return
	} else if r.Bank == nil {
		this.Data["IncorrectBank"] = true
		this.TplName = "index.tpl"
		return
=======
	flash := beego.NewFlash()
	if r.Currency == nil || r.Bank == nil {
		flash.Error("You should have chosen at least one currency and bank!")
		flash.Store(&this.Controller)
		//this.Redirect("/", 302)
>>>>>>> 35954cd064d125c358f7a971cf8f89cd0441b712
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

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) ErrorDb() {
	this.TplName = "err.tpl"
}
