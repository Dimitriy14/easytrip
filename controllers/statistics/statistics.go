package statistics

import (
	_ "github.com/oreuta/easytrip/controllers"
	_ "github.com/oreuta/easytrip/controllers/bank-rating"
	_ "github.com/oreuta/easytrip/controllers/best-bank"
<<<<<<< HEAD
=======
	"github.com/oreuta/easytrip/translate"
>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type StatisticController struct {
	beego.Controller
}

func (this *StatisticController) Get() {
<<<<<<< HEAD
=======

	translate := translate.New()
	lang := this.GetString("lang")
	if lang != "" {
		translate.Lang = lang
		this.Ctx.SetCookie("lang", translate.Lang)
	} else {
		translate.Lang = this.Ctx.GetCookie("lang")
		if translate.Lang == "" {
			translate.Lang = "en-US"
		}
	}
	translate.Path = "conf/locale_" + translate.Lang + ".ini"
	this.Data["i18n"] = translate.Tr

>>>>>>> 6e627ef8d2ec6bdf8aa5a102215e4b86d314f2a9
	statistic := toolbox.StatisticsMap.GetMap()
	this.Data["Stat"] = statistic["Data"]
	this.TplName = "statistic.tpl"
}
