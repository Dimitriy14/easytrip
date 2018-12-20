package historyController

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/models"
)

type HistoryController struct {
	beego.Controller
}

func (this *HistoryController) Get() {

	session := this.GetSession("session")
	if session == nil {
		return
	}
	usermap := session.(map[string]interface{})
	var user models.User
	user.Name = usermap["name"].(string)
	user.Login = usermap["login"].(string)
	user.Password = usermap["password"].(string)
	r := []models.MainRequest{
		{
			Currency: []string{"eur", "usd"},
			Option:   "both",
			Bank:     []string{"pireus", "otp", "privat"},
		},
		{
			Currency: []string{"eur", "usd"},
			Option:   "buy",
			Bank:     []string{"pireus", "otp", "privat"},
		},
		{
			Currency: []string{"eur", "usd"},
			Option:   "sale",
			Bank:     []string{"pireus", "otp", "privat"},
		},
	}
	_ = r

}
