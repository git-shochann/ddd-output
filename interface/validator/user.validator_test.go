package validator

import (
	"ddd/domain/model"
	"testing"
)

// Signup
func TestSignupValidate(t *testing.T) {

	userValidation := NewUserValidation()

	// missing first name
	userSignUpValidation := model.UserSignUpValidation{
		FirstName: "",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "taro0123456",
	}
	expectMessage := "Invalid First Name"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// missing last name
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "",
		Email:     "taro@gmail.com",
		Password:  "taro0123456",
	}
	expectMessage = "Invalid Last Name"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// missing email
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "",
		Password:  "taro0123456",
	}
	expectMessage = "Invalid Email"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	//FIXME
	// invalid email format
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "tarotarotaro",
		Password:  "taro0123456",
	}
	expectMessage = "Invalid Email"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// missing password
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "",
	}
	expectMessage = "Invalid Password"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// password 8文字以下

	// password 15文字以下

}

func ExecuteSignUpValidateTest(t *testing.T, userValidation UserValidator, userSignUpValidation *model.UserSignUpValidation, expect string) {

	result, _ := userValidation.SignupValidate(userSignUpValidation)
	if result != expect {
		t.Fatalf("\nactual: %v\nexpected: %v", result, expect)
	}
}
