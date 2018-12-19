package registration

import (
	"github.com/astaxie/beego/logs"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/sql1"
)

type RegService interface {
	CanRegistr(data models.User) (res bool)
	CanLogIN(data models.User) (user models.User, ok bool)
}

type RegServiceStruct struct{}

func (a *RegServiceStruct) CanRegistr(data models.User) (res bool) {
	res = sql1.InsertInto(data)
	return
}

func (a *RegServiceStruct) CanLogIN(data models.User) (user models.User, ok bool) {
	user, ok = sql1.CheckUser(data)
	logs.Info(user)
	return user, true
}

func New() RegService {
	return &RegServiceStruct{}
}
