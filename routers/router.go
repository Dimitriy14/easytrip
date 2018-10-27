package routers

import (
	"github.com/Dimitriy14/easytrip/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
