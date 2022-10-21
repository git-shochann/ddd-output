// interface層

package validator

import (
	"ddd/domain/model"
	"fmt"

	"github.com/go-playground/validator"
)

type HabitValidation interface {
	CreateHabitValidator(*model.CreateHabitValidation) (string, error)
}

type habitValidation struct{}

func NewHabitValidation() HabitValidation {
	return &habitValidation{}
}

func (hv habitValidation) CreateHabitValidator(CreateHabitValidation *model.CreateHabitValidation) (string, error) {

	fmt.Println("Debug")

	validate := validator.New()
	err := validate.Struct(&CreateHabitValidation)

	var errorMessage string

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			fieldName := fieldErr.Field()

			switch fieldName {
			case "Content":
				errorMessage = "Invalid Content"

			}
		}
		return errorMessage, err
	}
	return "", err
}
