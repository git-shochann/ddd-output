package usecase

import (
	"ddd/domain/model"
	"ddd/domain/repository"
)

// habitの取得や登録などでDBにアクセスする時に、domain層のrepository(インターフェースとして設定した部分)を介してアクセスすることによって、infrastructure層にアクセスするのではなく、
// domain層のみに直接依存するだけで完結出来る！ 単一方向であるので。infrastructure層を触れたりすることはしない。

// ここのusecase層がすることは、図の上のinterface層から情報を受け取り、下のdomain層のインターフェースで定義してあるメソッドを用いてビジネスロジックを実行すること

// インターフェース -> 窓口である
type HabitUseCase interface {
	CreateHabit(h *model.Habit) error
	// DeleteHabit(habitID, userID int, habit *model.Habit) error
	// UpdateHabit(habit *model.Habit) error
	// GetAllHabitByUserID(user model.User, habit *[]model.Habit) error
}

// これはなに？ -> ここの層でやることを構造体で表現する。
// usecase層の関数をメソッドとして定義し、
// 構造体のフィールド内にインターフェース型として設定すれば、インターフェースはメソッドを使用することの出来る窓口であるので、
// いつでもそのメソッドを使うことが出来る
// その使うことが出来るメソッドは domain層のインターフェースで設定したインターフェース。

// どの方向に依存しているかで考えると分かりやすい。
type habitUseCase struct {
	HabitRepository repository.HabitRepository // domain層のインターフェース
}

// インターフェースを引数にとってインターフェースを返す？ -> この引数はどこでそもそも呼び出す？
func NewHabitUseCase(hr repository.HabitRepository) HabitUseCase {
	// なぜポインタ型？
	return &habitUseCase{
		HabitRepository: hr,
	}
}

// domainのインターフェースを使って、実際に処理を行う
func (hu habitUseCase) CreateHabit(h *model.Habit) error {
	// 実際のDBの処理であるhu.CreateHabit() としてアクセスをすることが可能

	// 	err = habit.CreateHabit()
	// 	if err != nil {
	// 		models.SendErrorResponse(w, "Failed to create habit", http.StatusInternalServerError)
	// 		log.Println(err)
	// 		return
	// 	}

	// 	response, err := json.Marshal(habit)
	// 	if err != nil {
	// 		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
	// 		log.Println(err)
	// 		return
	// 	}

	// 	models.SendResponse(w, response, http.StatusOK)

	return nil

}
