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
	result, _ := habitValidation.HabitValidate(&createHabitValidation)
	if result != "" {
		t.Fatal("failed test")
	}

	// contentが空白の場合
	createHabitValidation = model.CreateHabitValidation{
		Content: "",
	}
	result, _ = habitValidation.HabitValidate(&createHabitValidation)
	if result != "Invalid Content" {
		t.Fatal("failed test")
	}

}
