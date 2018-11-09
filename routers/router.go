package routers

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/controllers"
	"github.com/oreuta/easytrip/controllers/bankRating"
	"github.com/oreuta/easytrip/controllers/bestbank"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/comparision", &bankRating.RatesController{})
	beego.Router("/best", &bestbank.BestController{})
}
