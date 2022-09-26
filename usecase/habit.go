package usecase

import (
	"ddd/domain/model"
	"ddd/domain/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// habitの取得や登録などでDBにアクセスする時に、domain層のrepository(インターフェースとして設定した部分)を介してアクセスすることによって、infrastructure層にアクセスするのではなく、
// domain層のみに直接依存するだけで完結出来る！ 単一方向であるので。

// ここのusecase層がすることは、図の上のinterface層から情報を受け取り、下のdomain層のインターフェースで定義してあるメソッドを用いてビジネスロジックを実行すること

// インターフェース -> 窓口である
// ここをどこで使う？
type HabitUseCase interface {
	CreateHabit(h *model.Habit) error
	DeleteHabit(habitID, userID int, habit *model.Habit) error
	UpdateHabit(habit *model.Habit) error
	GetAllHabitByUserID(user model.User, habit *[]model.Habit) error
}

// これはなに？ -> ここの層の責務を構造体で表現する。
// !!!!! またusecase層の関数をメソッドとして定義し、以下構造体のフィールド内にインターフェース型として設定すれば、インターフェースはメソッドを使用することの出来る窓口であるので、
// いつでもそのメソッドを使うことが出来る その使うことが出来るメソッドは domain層のインターフェースで設定したインターフェース。

// どの方向に依存しているかで考えると分かりやすい。
type habitUseCase struct {
	habitRepository repository.HabitRepository
}

// インターフェースを引数にとってインターフェースを返す？
func NewHabitUseCase(hr repository.HabitRepository) HabitUseCase {
	// なぜポインタ型？
	return &habitUseCase{
		habitRepository: hr,
	}
}

// domainのインターフェースを使って、実際に処理を行う
func (hu habitUseCase) CreateHabit(h *model.Habit) error {
	// 実際のDBの処理であるhu.CreateHabit() としてアクセスをすることが可能

	// Bodyを検証
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// バリデーションの実施
	var habitValidation models.CreateHabitValidation
	err = json.Unmarshal(reqBody, &habitValidation)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// 	errorMessage, err := habitValidation.CreateHabitValidator()

	// 	if err != nil {
	// 		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
	// 		log.Println(err)
	// 		return
	// 	}

	// 	if err != nil {
	// 		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
	// 		log.Println(err)
	// 		return
	// 	}

	// 	// JWTの検証
	// 	userID, err := models.CheckJWTToken(r)
	// 	if err != nil {
	// 		models.SendErrorResponse(w, "Authentication error", http.StatusBadRequest)
	// 		log.Println(err)
	// 		return
	// 	}

	// 	// JWTにIDが乗っているので、IDをもとに保存処理をする

	// 	habit := models.Habit{
	// 		Content: habitValidation.Content,
	// 		UserID:  userID,
	// 	}

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

}
