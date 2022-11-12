package validator

import (
	"ddd/domain/model"
	"testing"
)

func TestCreateHabitValidator(t *testing.T) {

	habitValidation := NewHabitValidation()

	// 正常時のパターン
	createHabitValidation := model.CreateHabitValidation{
		Content: "testtest",
	}
	message, _ := habitValidation.HabitValidate(&createHabitValidation)
	if message != "" {
		t.Fatal("failed test")
	}

	// contentが空白の場合
	createHabitValidation = model.CreateHabitValidation{
		Content: "",
	}
	message, _ = habitValidation.HabitValidate(&createHabitValidation)
	if message != "Invalid Content" {
		t.Fatal("failed test")
	}

}
