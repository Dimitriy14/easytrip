package registration

import (
	"github.com/oreuta/easytrip/models"
)

type RegService interface {
	CanRegistr(models.User) bool
	CanLogIN(models.User) bool
}
