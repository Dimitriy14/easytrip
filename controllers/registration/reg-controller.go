package regController

import (
	"github.com/astaxie/beego"

	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/registration"
)

type RegController struct {
	beego.Controller
	regist registration.RegService
}

func (this *RegController) Get() {
	this.TplName = "registration.tpl"
}

func (this *RegController) Post() {
	this.TplName = "registration.tpl"

	u := models.User{
		Name:     this.GetString("name"),
		Login:    this.GetString("login"),
		Password: this.GetString("password"),
	}

	this.Data["User"] = u
	// if !registration.CanRegistr(u){

	// }

	//this.Redirect("/signin", 303)

}

func New(reg registration.RegService) *RegController {
	return &RegController{regist: reg}
}
