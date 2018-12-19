package login

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/registration"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	this.TplName = "login.tpl"

	u := models.User{
		Login:    this.GetString("login"),
		Password: this.GetString("password"),
	}

	reg := registration.New()
	userName, ok := reg.CanLogIN(u)
	if !ok {
		this.Data["Errors"] = "Error"
		return
	}

	this.Data["Username"] = userName

}
