package repository

import "ddd/domain/model"

// repositoryはDBとのやりとりを定義するが、
// 技術的関心ごとはinfrastructure層に書くため、ここではインターフェースとしてメソッドを定義して...？
// 実際の処理はinfrastructure層に書き、domain層に依存するように(domain層を使うように)実装する

// habitにおけるrepositoryのインターフェース
// 実際のメソッド群を登録する

type HabitRepository interface {
	CreateHabit(h *model.Habit) error
	DeleteHabit(habitID, userID int, habit *model.Habit) error
	UpdateHabit(habit *model.Habit) error
	GetAllHabitByUserID(user model.User, habit *[]model.Habit) error
}
