// usecase層 (domain層に依存)

package usecase

import (
	"ddd/domain/model"
	"ddd/domain/repository"
	"ddd/interface/validator"
	"log"
)

// habitの取得や登録などでDBにアクセスする時に、domain層のrepository(インターフェースとして設定した部分)を介してアクセスすることによって、infrastructure層にアクセスするのではなく、
// domain層のみに直接依存するだけで完結出来る！ 単一方向であるので。infrastructure層を触れたりすることはしない。

// ここのusecase層がすることは、図の上のinterface層から情報を受け取り、下のdomain層のインターフェースで定義してあるメソッドを用いてビジネスロジックを実行すること

// インターフェース -> 窓口である
type HabitUseCase interface {
	CreateHabit(habit *model.Habit) (*model.Habit, error)
	// DeleteHabit(habitID, userID int, habit *model.Habit) error
	// UpdateHabit(habit *model.Habit) error
	// GetAllHabitByUserID(user model.User, habit *[]model.Habit) error
	// ここにJWTのロジックを使用する関数を追加してあげる
	// CheckJWTToken
}

type habitUseCase struct {
	hr repository.HabitRepository //以下全てdomain層のインターフェース。 この構造体に紐づいているメソッドでそのメソッドを使用したいので！
	hv validator.HabitValidation
}

func NewHabitUseCase(hr repository.HabitRepository, hv validator.HabitValidation) HabitUseCase {
	return &habitUseCase{
		hr: hr,
		hv: hv,
	}
}

// WIP: ここで引数にhttp.ResponseWriterが来ることはない

// domainのインターフェースを使って、実際に処理を行う
func (hu *habitUseCase) CreateHabit(habit *model.Habit) (*model.Habit, error) {

	err := hu.hr.CreateHabitPersistence(habit)
	if err != nil {
		// hu.rl.SendErrorResponseLogic(w, "Failed to create habit", http.StatusInternalServerError)-> ここではこれは行わない -> 次回のリファクタリングの段階で、エラーハンドリングを終わらせる！
		log.Println(err)
		return nil, err
	}

	// 書き変わったhabitを返す
	return habit, nil

}

// 参考 //

// ということはこのCreateTodoを読んでいるところはどこ？

// func (ts *todoService) CreateTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error) {
// 	// ioutil: ioに特化したパッケージ
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var mutationTodoRequest models.MutationTodoRequest
// 	if err := json.Unmarshal(reqBody, &mutationTodoRequest); err != nil {
// 		log.Fatal(err)
// 		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
// 		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
// 		return models.BaseTodoResponse{}, err
// 	}
// 	// バリデーション
// 	if err := ts.tv.MutationTodoValidate(mutationTodoRequest); err != nil {
// 		// バリデーションエラーのレスポンスを送信
// 		ts.rl.SendResponse(w, ts.rl.CreateErrorResponse(err), http.StatusBadRequest)
// 		return models.BaseTodoResponse{}, err
// 	}

// 	var todo models.Todo
// 	todo.Title = mutationTodoRequest.Title
// 	todo.Comment = mutationTodoRequest.Comment
// 	todo.UserId = userId

// 	// todoデータ新規登録処理
// 	if err := ts.tr.CreateTodo(&todo); err != nil {
// 		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ取得に失敗しました。"), http.StatusInternalServerError)
// 		return models.BaseTodoResponse{}, err
// 	}

// 	// 登録したtodoデータ取得処理
// 	if err := ts.tr.GetTodoLastByUserId(&todo, userId); err != nil {
// 		var errMessage string
// 		var statusCode int
// 		// https://gorm.io/ja_JP/docs/error_handling.html
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			statusCode = http.StatusBadRequest
// 			errMessage = "該当データは存在しません。"
// 		} else {
// 			statusCode = http.StatusInternalServerError
// 			errMessage = "データ取得に失敗しました。"
// 		}
// 		// エラーレスポンス送信
// 		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse(errMessage), statusCode)
// 		return models.BaseTodoResponse{}, err
// 	}

// 	// レスポンス用の構造体に変換
// 	responseTodos := ts.tl.CreateTodoResponse(&todo)

// 	return responseTodos, nil
// }
