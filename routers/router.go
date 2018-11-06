package routers

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
