package model

import "gorm.io/gorm"

type User struct {
	gorm.Model         // ID, CreatedAt, UpdatedAt, DeletedAt を作成
	FirstName  string  `gorm:"not null"`
	LastName   string  `gorm:"not null"`
	Email      string  `gorm:"not null;unique"`
	Password   string  `gorm:"not null"`
	Habits     []Habit // User has many Habit
}
