package regController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/oreuta/easytrip/models"
)

type RegController struct {
	beego.Controller
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

	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		beego.Error("ValidationError: %v", err)
	}
	if !b {
		this.Data["Errors"] = valid.ErrorsMap
		return
	}
	// if !registration.CanRegistr(u){

	// }

	this.Redirect("/login", 303)

}
