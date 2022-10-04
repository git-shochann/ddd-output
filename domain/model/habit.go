package model

import "github.com/jinzhu/gorm"

// 業務領域(このソフトウェアが解決、扱うもの)に関するコードをこのdomain層に置く
// 他のレイヤーに依存することはない

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
