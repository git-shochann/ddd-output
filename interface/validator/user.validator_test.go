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
	expect := "Invalid First Name"
	result, _ := userValidation.SignupValidate(&userSignUpValidation)
	if result != expect {
		t.Fatalf("\n実際: %v\n理想: %v", result, expect)
	}

}

// Signin
