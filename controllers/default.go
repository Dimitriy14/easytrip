package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/oreuta/easytrip/translate"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/", "&controllers.MainController", time.Duration(13000))

	translate := translate.New()
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
	c.Data["i18n"] = translate.Tr

	c.Layout = "main_layout.tpl"
	c.TplName = "index.tpl"
}
