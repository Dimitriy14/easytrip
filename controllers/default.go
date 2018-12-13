package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
<<<<<<< HEAD
=======
	"github.com/oreuta/easytrip/translate"
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/", "&controllers.MainController", time.Duration(13000))
<<<<<<< HEAD
=======

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

>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
	c.Layout = "main_layout.tpl"
	c.TplName = "index.tpl"
}
