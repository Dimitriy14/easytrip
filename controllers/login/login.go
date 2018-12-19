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
	this.Data["Username"] = userName
	if !ok {
		this.Data["Errors"] = "Error"
		return
	}

	// ses := make(map[string]interface{})

	// ses["name"] = userName.Name
	// ses["login"] = userName.Login
	// ses["password"] = userName.Password

	// this.SetSession("session", ses)

	this.Data["Username"] = userName

}
