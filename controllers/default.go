package controllers

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/controllers/baseController"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "main_layout.tpl"
	c.TplName = "index.tpl"

	translate := baseController.New()

	lang := c.GetString("lang")

	if lang != "" {
		translate.Lang = lang
		c.Ctx.SetCookie("lang", translate.Lang)
	} else {
		translate.Lang = c.Ctx.GetCookie("lang")
		if translate.Lang == "" {
			translate.Lang = "en-US"
		}
	}

	translate.Path = "conf/locale_" + translate.Lang + ".ini"
	c.Data["i18n"] = translate.Trans
}
