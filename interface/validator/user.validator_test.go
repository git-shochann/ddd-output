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
		t.Fatalf("\nactual: %v\nexpected: %v", result, expect)
	}

	// missing last name
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "",
		Email:     "taro@gmail.com",
		Password:  "taro0123456",
	}
	expect = "Invalid Last Name"
	result, _ = userValidation.SignupValidate(&userSignUpValidation)
	if result != expect {
		t.Fatalf("\nactual: %v\nexpected: %v", result, expect)
	}

	// missing email
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "",
		Password:  "taro0123456",
	}
	expect = "Invalid Email"
	result, _ = userValidation.SignupValidate(&userSignUpValidation)
	if result != expect {
		t.Fatalf("\nactual: %v\nexpected: %v", result, expect)
	}

	// missing password
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "",
	}
	expect = "Invalid Password"
	result, _ = userValidation.SignupValidate(&userSignUpValidation)
	if result != expect {
		t.Fatalf("\nactual: %v\nexpected: %v", result, expect)
	}

}

// Signin
