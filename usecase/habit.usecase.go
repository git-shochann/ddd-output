// usecase層 (domain層に依存)

package usecase

import (
	"ddd/domain"
	"ddd/domain/model"
	"log"
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
	hr domain.HabitRepository //以下全てdomain層のインターフェース。 この構造体に紐づいているメソッドでそのメソッドを使用したいので！
}

func NewHabitUseCase(hr domain.HabitRepository) HabitUseCase {
	return &habitUseCase{
		hr: hr,
	}
}

// WIP: ここで引数にhttp.ResponseWriterが来ることはない

// domainのインターフェースを使って、実際に処理を行う
func (huc *habitUseCase) CreateHabit(habit *model.Habit) (*model.Habit, error) {

	err := huc.hr.CreateHabitPersistence(habit)
	if err != nil {
		// hu.rl.SendErrorResponseLogic(w, "Failed to create habit", http.StatusInternalServerError)-> ここではこれは行わない -> 次回のリファクタリングの段階で、エラーハンドリングを終わらせる！
		log.Println(err)
		return nil, err
	}

	// 書き変わったhabitを返す
	return habit, nil

}

func (huc *habitUseCase) UpdateHabit(habit *model.Habit) (*model.Habit, error) {

	err := huc.hr.UpdateHabitPersistence(habit)
	if err != nil {
		// hu.rl.SendErrorResponseLogic(w, "Failed to create habit", http.StatusInternalServerError)-> ここではこれは行わない -> 次回のリファクタリングの段階で、エラーハンドリングを終わらせる！
		log.Println(err)
		return nil, err
	}

	// 書き変わったhabitを返す
	return habit, nil
}

func (huc *habitUseCase) DeleteHabit(habitID, userID int, habit *model.Habit) error {

	err := huc.hr.DeleteHabitPersistence(habitID, userID, habit)
	if err != nil {
		// hu.rl.SendErrorResponseLogic(w, "Failed to create habit", http.StatusInternalServerError)-> ここではこれは行わない -> 次回のリファクタリングの段階で、エラーハンドリングを終わらせる！
		log.Println(err)
		return err
	}

	// 書き変わったhabitは返さない
	return nil
}

func (huc *habitUseCase) GetAllHabitByUserID(user *model.User, habit *[]model.Habit) (*[]model.Habit, error) {

	err := huc.hr.GetAllHabitByUserIDPersistence(user, habit)
	if err != nil {
		// hu.rl.SendErrorResponseLogic(w, "Failed to create habit", http.StatusInternalServerError)-> ここではこれは行わない -> 次回のリファクタリングの段階で、エラーハンドリングを終わらせる！
		log.Println(err)
		return nil, err
	}

	// 書き変わったhabitを返す
	return habit, nil
}
