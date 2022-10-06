// interface

package handler

import (
	"ddd/usecase"
	"fmt"
	"net/http"
)

// interface層はHTTPリクエスト・レスポンスを扱う層
// usecase層と切り離すことでリクエストやレスポンスの形に変わってもinterface層の修正だけで済むようになる...？

// ここの層に依存する箇所で使用する メソッドの窓口を用意してあげる
type HabitHandler interface {
	IndexFunc(w http.ResponseWriter, r *http.Request)
	CreateFunc(w http.ResponseWriter, r *http.Request)
	// UpdateFunc(http.ResponseWriter, *http.Request)
	// DeleteFunc(http.ResponseWriter, *http.Request)
	// GetAllHabitFunc(http.ResponseWriter, *http.Request)
}

// これはこの後記載するメソッドの型として、設定するために作成する
// Privateで宣言 ここのパッケージ以外では使用しないので
type habitHandler struct {
	huc usecase.HabitUseCase // usecase層のインターフェースを設定して、該当のメソッドを使用出来るようにする
	juc usecase.JwtUseCase
	ruc usecase.ResponseUseCase
}

// main関数で依存関係同士で繋ぐために必要
// ここの構造体のフィールドに書くのは、依存先のインターフェースを書けばOK
func NewHabitHandler(huc usecase.HabitUseCase, juc usecase.JwtUseCase) HabitHandler {
	return &habitHandler{
		huc: huc,
		juc: juc,
	}
}

// 第一引数にはHTTPサーバーからのレスポンスを出力することが出来るメソッドを持っている(該当のメソッドを実装している)構造体の値が来る
func (hh *habitHandler) IndexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body: %v\n", r.Body)
	fmt.Printf("%T\n", w)                   // *http.response構造体
	fmt.Fprintf(w, "This is Go's Rest API") // メソッド内でw.Write()をするため
}

// ** ここのファイルは具体的なロジックを書くのは発生しない ** //

// main（）のrouter.HandleFunc()の第二引数として以下の関数を渡すだけ
func (hh *habitHandler) CreateFunc(w http.ResponseWriter, r *http.Request) {

	// usecase層に依存するので、UseCase層でJWTのロジックを使用するインターフェースを用意する
	// 必要がある

	// JWTの検証
	userID, err := hh.juc.CheckJWTToken(w, r)
	if err != nil {
		return // router.HandleFunc())の第二引数に関数を渡すだけなので戻り値なし
	}

	// 保存準備(JWTにIDが乗っているので、IDをもとに保存処理をする)

	// これは..？
	// habit := model.Habit{
	// 	Content: habitValidation.Content,
	// 	UserID:  userID,
	// }

	// 保存処理 -> ここでBodyの検証、バリデーションの実行

	response, err = hh.huc.CreateHabit(w, r, userID)
	if err != nil {
		return
	}

	// *** 結果以下でレスポンスを作成するのでusecase内の、処理ではレスポンスの構造体を返す ***

	// レスポンス
	hh.ruc.SendResponseUseCase(w, response, http.StatusOK)

}

// 参考 //

// func (tc *todoController) CreateTodo(w http.ResponseWriter, r *http.Request) {

// 	// トークンからuserIdを取得
// 	userId, err := tc.as.GetUserIdFromToken(w, r)
// 	if userId == 0 || err != nil {
// 		return
// 	}

// 	// todoデータ取得処理

// 以下の処理内でのResponseはどんな感じ？

// 	responseTodo, err := tc.ts.CreateTodo(w, r, userId)
// 	if err != nil {
// 		return
// 	}

// 	// レスポンス送信処理
// 	tc.ts.SendCreateTodoResponse(w, &responseTodo)
// }
