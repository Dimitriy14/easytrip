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

	b, err := this.RatesService.GetBankRates(r)
	if err != nil {
		beego.Error("Error:%v", err)
	}

	this.Data["Banks"] = b
	this.Layout = "comparision_layout.tpl"
	this.TplName = "comparision.tpl"
}