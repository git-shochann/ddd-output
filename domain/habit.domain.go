// domain層 (依存なし)

package domain

import "ddd/domain/model"

// ここの層はinterfaceを提供するのみでOK！ (DDDの場合)

type HabitRepository interface {
	CreateHabitInfrastructure(h *model.Habit) error
	UpdateHabitInfrastructure(habit *model.Habit) error
	DeleteHabitInfrastructure(habitID, userID int, habit *model.Habit) error
	GetAllHabitByUserIDInfrastructure(user *model.User, habit *[]model.Habit) error
}
