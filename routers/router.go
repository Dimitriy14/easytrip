package routers

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/controllers"
	"github.com/oreuta/easytrip/controllers/bank-rating"
	"github.com/oreuta/easytrip/services/bank-rating"
)

func init() {
	ratesclient := clients.New()
	ratesService := bankRatingService.New(ratesclient)
	ratesController := bankRatingController.New(ratesService)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/comparision", ratesController)
	beego.Router("/err", &controllers.ErrorController{}, "get:ErrorDb")
}
