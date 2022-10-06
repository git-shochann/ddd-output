// domain

package model

import "github.com/jinzhu/gorm"

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
