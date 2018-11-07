package routers

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/controllers"
	"github.com/oreuta/easytrip/controllers/bankRating"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/comparision", &bankRating.RatesController{})
}
