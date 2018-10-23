package routers

import (
	"github.com/oreuta/easytrip/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
