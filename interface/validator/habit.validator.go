// interface層

package validator

import (
	"ddd/domain/model"

	"github.com/go-playground/validator"
)

// バリデーターを公開してあげる
type HabitValidator interface {
	HabitValidate(*model.CreateHabitValidation) (string, error)
}

type HabitValidation struct{}

func NewHabitValidation() HabitValidator {
	return &HabitValidation{}
}

func (chv *HabitValidation) HabitValidate(createHabitValidation *model.CreateHabitValidation) (string, error) {

	validate := validator.New()
	err := validate.Struct(createHabitValidation)

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			fieldName := fieldErr.Field()

			switch fieldName {
			case "Content":
				return "Invalid Content", err
			default:
				return "Unknown Error", err
			}
		}
	}
	return "", err
}
