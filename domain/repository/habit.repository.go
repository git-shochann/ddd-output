// domain

package repository

import "ddd/domain/model"

// ここの層はinterfaceを提供するのみでOK！ (DDDの場合)

type HabitRepository interface {
	CreateHabitPersistence(h *model.Habit) error
	DeleteHabitPersistence(habitID, userID int, habit *model.Habit) error
	UpdateHabitPersistence(habit *model.Habit) error
	GetAllHabitByUserIDPersistence(user model.User, habit *[]model.Habit) error
}
