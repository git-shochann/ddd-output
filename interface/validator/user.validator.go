// interface層

package validator

import (
	"ddd/domain/model"

	"github.com/go-playground/validator"
)

// まずとりあえず他の層で使えるようにinterfaceを定義する
type UserValidation interface {
	SignupValidator(*model.UserSignUpValidation) (string, error)
	SigninValidator(*model.UserSignInValidation) (string, error)
}

// ここの構造体で、他の層に依存している(使いたいメソッドを持ったインターフェースはある？)
type userValidation struct {
}

func NewUserValidation() UserValidation {
	return &userValidation{}
}

func (uv userValidation) SignupValidator(UserSignUpValidation *model.UserSignUpValidation) (string, error) {

	validate := validator.New()
	err := validate.Struct(UserSignUpValidation)

	var errorMessage string

	if err != nil {

		for _, fieldErr := range err.(validator.ValidationErrors) {

			fieldName := fieldErr.Field()

			switch fieldName {
			case "FirstName":
				errorMessage = "Invalid First Name"
			case "LastName":
				errorMessage = "Invalid Last Name"
			case "Email":
				errorMessage = "Invalid Email"
			case "Password":
				errorMessage = "Invalid Password"
			}
		}
		return errorMessage, err
	}
	return "", err
}

func (uv userValidation) SigninValidator(UserSignInValidation *model.UserSignInValidation) (string, error) {

	validate := validator.New()
	err := validate.Struct(UserSignInValidation)

	var errorMessage string

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {

			fieldName := fieldErr.Field()

			switch fieldName {
			case "Email":
				errorMessage = "Invalid Email"
			case "Password":
				errorMessage = "Invalid Password"
			}
		}
		return errorMessage, err
	}
	return "", err
}
