package validator

import (
	"ddd/domain/model"
	"testing"
)

// Signup
func TestSignupValidate(t *testing.T) {

	userValidation := NewUserValidation()

	// first name 空白
	userSignUpValidation := model.UserSignUpValidation{
		FirstName: "",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "taro0123456",
	}
	expectMessage := "Invalid First Name"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// last name 空白
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "",
		Email:     "taro@gmail.com",
		Password:  "taro0123456",
	}
	expectMessage = "Invalid Last Name"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// email 空白
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "",
		Password:  "taro0123456",
	}
	expectMessage = "Invalid Email"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// emailフォーマットではない
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "tarotarotaro",
		Password:  "taro0123456",
	}
	expectMessage = "Invalid Email"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// password 空白
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "",
	}
	expectMessage = "Invalid Password"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// password 8文字以下
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "tarotar", // 7
	}
	expectMessage = "Invalid Password"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// password 15文字以上
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "tarotarotarotaro", // 16
	}
	expectMessage = "Invalid Password"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// password 英語のみ
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "tarotaro",
	}
	expectMessage = "Invalid Password"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

	// password 数字のみ
	userSignUpValidation = model.UserSignUpValidation{
		FirstName: "Taro",
		LastName:  "Taro",
		Email:     "taro@gmail.com",
		Password:  "123456789",
	}
	expectMessage = "Invalid Password"
	ExecuteSignUpValidateTest(t, userValidation, &userSignUpValidation, expectMessage)

}

func ExecuteSignUpValidateTest(t *testing.T, userValidation UserValidator, userSignUpValidation *model.UserSignUpValidation, expect string) {

	result, _ := userValidation.SignupValidate(userSignUpValidation)

	if result == expect {
		t.Skip() // エラーメッセージが同じであればOK
	} else {
		t.Fatalf("\nactual: %v\nexpected: %v", result, expect) //
	}

}
