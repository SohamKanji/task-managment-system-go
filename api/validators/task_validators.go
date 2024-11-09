package validators

import (
	"github.com/SohamKanji/task-management-system-go/utils"
	"github.com/go-playground/validator/v10"
)

var ValidStatus validator.Func = func(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	return utils.IsValidStatus(status)
}
