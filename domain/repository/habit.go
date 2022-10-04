package repository

import "ddd/domain/model"

// ここの層はinterfaceを提供するのみでOK！

type HabitRepository interface {
	CreateHabit(h *model.Habit) error
	DeleteHabit(habitID, userID int, habit *model.Habit) error
	UpdateHabit(habit *model.Habit) error
	GetAllHabitByUserID(user model.User, habit *[]model.Habit) error
}
