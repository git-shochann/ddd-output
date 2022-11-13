// interface層

package validator

import (
	"ddd/domain/model"

	"github.com/go-playground/validator"
)

// まずとりあえず他の層で使えるようにinterfaceを定義する
type UserValidator interface {
	SignUpValidate(*model.UserSignUpValidation) (string, error)
	SignInValidate(*model.UserSignInValidation) (string, error)
}

// ここの構造体で、他の層に依存している(使いたいメソッドを持ったインターフェースはある？)
type userValidation struct {
}

func NewUserValidation() UserValidator {
	return &userValidation{}
}

func (uv *userValidation) SignUpValidate(UserSignUpValidation *model.UserSignUpValidation) (string, error) {

	validate := validator.New()
	err := validate.Struct(UserSignUpValidation)

	if err != nil {

		for _, fieldErr := range err.(validator.ValidationErrors) {

			fieldName := fieldErr.Field()

			switch fieldName {
			case "FirstName":
				return "Invalid First Name", err
			case "LastName":
				return "Invalid Last Name", err
			case "Email":
				return "Invalid Email", err
			case "Password":
				return "Invalid Password", err
			default:
				return "Unknown Error", err
			}

		}
	}
	return "", err
}

func (uv *userValidation) SignInValidate(UserSignInValidation *model.UserSignInValidation) (string, error) {

	validate := validator.New()
	err := validate.Struct(UserSignInValidation)

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {

			fieldName := fieldErr.Field()

			switch fieldName {
			case "Email":
				return "Invalid Email", err
			case "Password":
				return "Invalid Password", err
			}
		}
	}
	return "", err
}
