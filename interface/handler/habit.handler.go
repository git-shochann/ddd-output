// interface層 (domain層に依存)

package handler

import (
	"ddd/domain/model"
	"ddd/interface/util"
	"ddd/interface/validator"
	"ddd/usecase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ここの層に依存する箇所で使用する メソッドの窓口を用意してあげる
type HabitHandler interface {
	IndexFunc(w http.ResponseWriter, r *http.Request)
	CreateFunc(w http.ResponseWriter, r *http.Request)
	// UpdateFunc(http.ResponseWriter, *http.Request)
	// DeleteFunc(http.ResponseWriter, *http.Request)
	// GetAllHabitFunc(http.ResponseWriter, *http.Request)
}

type habitHandler struct {
	huc usecase.HabitUseCase      // usecase層
	hv  validator.HabitValidation // interface層
	ju  util.JwtUtil              // interface層
	ru  util.ResponseUtil         // interface層
}

// main関数で依存関係同士で繋ぐために必要
// ここの構造体のフィールドに書くのは、依存先のインターフェースを書けばOK
func NewHabitHandler(huc usecase.HabitUseCase, hv validator.HabitValidation, ju util.JwtUtil, ru util.ResponseUtil) HabitHandler {
	return &habitHandler{
		huc: huc,
		hv:  hv,
		ju:  ju,
		ru:  ru,
	}
}

// 第一引数にはHTTPサーバーからのレスポンスを出力することが出来るメソッドを持っている(該当のメソッドを実装している)構造体の値が来る
func (hh *habitHandler) IndexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body: %v\n", r.Body)
	fmt.Printf("%T\n", w)                   // *http.response構造体
	fmt.Fprintf(w, "This is Go's Rest API") // メソッド内でw.Write()をするため
}

// ** interface層では具体的なロジックを書くのは発生しない ** //

// main（）のrouter.HandleFunc()の第二引数として以下の関数を渡すだけ
func (hh *habitHandler) CreateFunc(w http.ResponseWriter, r *http.Request) {

	// JWTの検証
	userID, err := hh.ju.CheckJWTToken(r)
	if err != nil {
		log.Println(err)
		hh.ru.SendErrorResponse(w, "Failed to authenticate", http.StatusBadRequest)
		return
	}

	// Bodyの読み込み
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		hh.ru.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		return // router.HandleFunc())の第二引数に関数を渡すだけなので戻り値なし
	}

	// バリデーション
	var habitValidation model.CreateHabitValidation
	err = json.Unmarshal(reqBody, &habitValidation)
	if err != nil {
		hh.ru.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	errorMessage, err := hh.hv.CreateHabitValidator(&habitValidation)
	if err != nil {
		hh.ru.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		log.Println(err)
		return
	}

	// DBに登録する内容の準備
	habit := model.Habit{
		Content:  habitValidation.Content,
		Finished: false,
		UserID:   userID,
	}

	// 保存処理
	newHabit, err := hh.huc.CreateHabit(&habit) // -> usecase層に依存
	if err != nil {
		hh.ru.SendErrorResponse(w, "failed to create habit", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// 登録が完了したhabitを上書きしてレスポンスとして返すためにjson形式にする([]byte)
	response, err := json.Marshal(newHabit)
	if err != nil {
		hh.ru.SendErrorResponse(w, "failed to encode json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	hh.ru.SendResponse(w, response, http.StatusOK)

}
