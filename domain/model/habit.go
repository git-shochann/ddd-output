package model

import "github.com/jinzhu/gorm"

// システムが扱う業務領域に関するコードを置くところ。
// 他のレイヤーに依存しない。

type Habit struct {
	gorm.Model
	Content  string `gorm:"not null"`
	Finished bool   `gorm:"not null"`
	UserID   int    `gorm:"not null"`
}

// バリデーション関連もここにまとめた
type CreateHabitValidation struct {
	Content string `json:"content" validate:"required"`
}
