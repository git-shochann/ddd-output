package handler

import (
	"ddd/domain/model"
	"ddd/usecase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// interface層はHTTPリクエスト・レスポンスを扱う層
// usecase層と切り離すことでリクエストやレスポンスの形に変わってもinterface層の修正だけで済むようになる...？

// ここの層に依存する箇所で使用する メソッドの窓口を用意してあげる
// Publicで宣言
type HabitHandler interface {
	IndexFunc(http.ResponseWriter, *http.Request)
	CreateFunc(http.ResponseWriter, *http.Request)
	// Update(http.ResponseWriter, *http.Request)
}

// これはこの後記載するメソッドの型として、設定するために作成する
// Privateで宣言 ここのパッケージ以外では使用しないので
type habitHandler struct {
	HabitUseCase usecase.HabitUseCase // usecase層のインターフェースを設定して、該当のメソッドを使用出来るようにする
}

// main関数で依存関係同士で繋ぐために必要
func NewHabitHandler(hu usecase.HabitUseCase) HabitHandler {
	return &habitHandler{
		HabitUseCase: hu,
	}
}

// 第一引数にはHTTPサーバーからのレスポンスを出力することが出来るメソッドを持っている(該当のメソッドを実装している)構造体の値が来る
func (hh habitHandler) IndexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body: %v\n", r.Body)
	fmt.Printf("%T\n", w)                   // *http.response構造体
	fmt.Fprintf(w, "This is Go's Rest API") // メソッド内でw.Write()をするため
}

func (hh habitHandler) CreateFunc(w http.ResponseWriter, r *http.Request) {

	// Bodyを検証
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// バリデーションの事前設定
	var habitValidation model.CreateHabitValidation
	err = json.Unmarshal(reqBody, &habitValidation)
	if err != nil {
		// models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// 実際にバリデーションを行う

	// errorMessage, err := habitValidation.CreateHabitValidator()

	// 	if err != nil {
	// 		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
	// 		log.Println(err)
	// 		return
	// 	}

	// UseCase層の呼び出しを行う

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
	//

}
