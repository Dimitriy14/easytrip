package registration

import (
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/sql1"
)

type RegService interface {
	CanRegistr(data models.User) (err error)
	CanLogIN(data models.User) (user models.User, err error)
}

type RegServiceStruct struct{}

func (a *RegServiceStruct) CanRegistr(data models.User) (err error) {
	err = sql1.InsertInto(data)
	return
}

func (a *RegServiceStruct) CanLogIN(data models.User) (user models.User, err error) {
	user, err = sql1.CheckUser(data)
	return
}

func New() RegService {
	return &RegServiceStruct{}
}
