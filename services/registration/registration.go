package registration

import (
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/sql1"
)

type RegService interface {
	CanRegistr(data models.User) (res bool)
	CanLogIN(data models.User) (username string, res bool)
}

type RegServiceStruct struct{}

func (a *RegServiceStruct) CanRegistr(data models.User) (res bool) {
	res = sql1.InsertInto(data)
	return
}

func (a *RegServiceStruct) CanLogIN(data models.User) (username string, res bool) {
	res = sql1.CheckUser(data)
	if res {
		username = data.Name
	}
	return
}

func New() RegService {
	return &RegServiceStruct{}
}
