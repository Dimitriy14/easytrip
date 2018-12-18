package login

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/models"
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

	this.Data["User"] = u

}
