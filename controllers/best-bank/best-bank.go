package bestBankController

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/best-bank"
)

type RatesController struct {
	beego.Controller
	BestService bestBankService.BestBankServiceInterface
}

func New(s bestBankService.BestBankServiceInterface) *RatesController {
	return &RatesController{BestService: s}
}

func (r *RatesController) Get() {
	inpData := models.MainRequest{
		Currency: r.GetStrings("currency"),
		Option:   r.GetString("option"),
		Bank:     r.GetStrings("bank"),
	}
	Sale, err := r.BestService.GetBestBanksSale(inpData)
	if err != nil {
		beego.Error("GetBestBanks func in BestService: %v", err)
	}
	Buy, err := r.BestService.GetBestBanksBuy(inpData)
	if err != nil {
		beego.Error("GetBestBanks func in BestService: %v", err)
	}

	r.Data["Error"] = err
	r.Data["Sale"] = Sale
	r.Data["Buy"] = Buy
	r.TplName = "best.tpl"
}
