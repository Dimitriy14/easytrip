package statistics

import (
	"time"

	_ "github.com/oreuta/easytrip/controllers"
	_ "github.com/oreuta/easytrip/controllers/bank-rating"
	_ "github.com/oreuta/easytrip/controllers/best-bank"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type StatisticController struct {
	beego.Controller
}

func (this *StatisticController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/", "&controllers.MainController", time.Duration(13000))
	toolbox.StatisticsMap.AddStatistics("GET", "/comparision", "&controllers.bankRatingController.RatesController", time.Duration(14000))
	toolbox.StatisticsMap.AddStatistics("GET", "/best", "&controllers.bestBankController.bestBankController", time.Duration(15000))

	statistic := toolbox.StatisticsMap.GetMap()
	this.Data["Stat"] = statistic["Data"]
	this.TplName = "statistic.tpl"
}
