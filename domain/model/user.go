package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model         // ID, CreatedAt, UpdatedAt, DeletedAt を作成
	FirstName  string  `gorm:"not null"`
	LastName   string  `gorm:"not null"`
	Email      string  `gorm:"not null;unique"`
	Password   string  `gorm:"not null"`
	Habits     []Habit // User has many Habit
}

// バリデーション関連もここにまとめた

type UserSignUpValidation struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=15,containsany=0123456789"`
}

type UserSignInValidation struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15,containsany=0123456789"`
}
