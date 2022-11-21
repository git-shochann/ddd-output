// usecase層 (domain層に依存)

package usecase

import (
	"ddd/domain"
	"ddd/domain/model"
	"fmt"
)

// habitの取得や登録などでDBにアクセスする時に、domain層のrepository(インターフェースとして設定した部分)を介してアクセスすることによって、infrastructure層にアクセスするのではなく、
// domain層のみに直接依存するだけで完結出来る！ 単一方向であるので。infrastructure層を触れたりすることはしない。

// ここのusecase層がすることは、図の上のinterface層から情報を受け取り、下のdomain層のインターフェースで定義してあるメソッドを用いてビジネスロジックを実行すること

// インターフェース -> 窓口である
type HabitUseCase interface {
	CreateHabit(habit *model.Habit) (*model.Habit, error)
	UpdateHabit(habit *model.Habit) (*model.Habit, error)
	DeleteHabit(habitID, userID int, habit *model.Habit) error
	GetAllHabitByUserID(user *model.User, habit *[]model.Habit) (*[]model.Habit, error)
}

type habitUseCase struct {
	HabitRepository domain.HabitRepository //以下全てdomain層のインターフェース。 この構造体に紐づいているメソッドでそのメソッドを使用したいので！
}

func NewHabitUseCase(habitRepository domain.HabitRepository) HabitUseCase {
	return &habitUseCase{
		HabitRepository: habitRepository,
	}
}

// ここの層で引数にhttp.ResponseWriterが来ることはない

// domainのインターフェースを使って、実際に処理を行う
func (huc *habitUseCase) CreateHabit(habit *model.Habit) (*model.Habit, error) {

	err := huc.HabitRepository.CreateHabitInfrastructure(habit)
	if err != nil {
		err := fmt.Errorf("habit usecase: failed to create habit %w", err)
		return nil, err
	}

	// 書き変わったhabitを返す
	return habit, nil

}

func (huc *habitUseCase) UpdateHabit(habit *model.Habit) (*model.Habit, error) {

	err := huc.HabitRepository.UpdateHabitInfrastructure(habit)
	if err != nil {
		err := fmt.Errorf("habit usecase: failed to update habit %w", err)
		return nil, err
	}

	// 書き変わったhabitを返す
	return habit, nil
}

func (huc *habitUseCase) DeleteHabit(habitID, userID int, habit *model.Habit) error {

	err := huc.HabitRepository.DeleteHabitInfrastructure(habitID, userID, habit)
	if err != nil {
		err := fmt.Errorf("habit usecase: failed to delete habit %w", err)
		return err
	}

	// 書き変わったhabitは返さない
	return nil
}

func (huc *habitUseCase) GetAllHabitByUserID(user *model.User, habit *[]model.Habit) (*[]model.Habit, error) {

	err := huc.HabitRepository.GetAllHabitByUserIDInfrastructure(user, habit)
	if err != nil {
		err := fmt.Errorf("habit usecase: failed to get all habit %w", err)
		return nil, err
	}

	// 書き変わったhabitを返す
	return habit, nil
}
