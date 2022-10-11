// interface (usecaseに依存)

package handler

import (
	"ddd/domain/model"
	"ddd/interface/util"
	"ddd/interface/validator"
	"ddd/usecase"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserHandler interface {
	SignUpFunc(http.ResponseWriter, *http.Request)
	SignInFunc(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	uuc usecase.UserUseCase      // usecase層
	uv  validator.UserValidation // interface層
	ju  util.JwtUtil             // interface層
	ru  util.ResponseUtil        // interface層
	epl util.EncryptPassword     // interface層
}

func NewUserHandler(uuc usecase.UserUseCase, uv validator.UserValidation, ju util.JwtUtil, ru util.ResponseUtil, epl util.EncryptPassword) UserHandler {
	return &userHandler{
		uuc: uuc,
		uv:  uv,
		ju:  ju,
		ru:  ru,
		epl: epl,
	}
}

func (uh *userHandler) SignUpFunc(w http.ResponseWriter, r *http.Request) {

	// Jsonでくるので、まずGoで使用できるようにする
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		return // router.HandleFunc())の第二引数に関数を渡すだけなので戻り値なし
	}

	var signUpUserValidation model.UserSignUpValidation
	err = json.Unmarshal(reqBody, &signUpUserValidation)

	if err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		return
	}

	errorMessage, err := uh.uv.SignupValidator(&signUpUserValidation)
	if err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// DBに登録する内容の準備 -> リファクタリング後 -> ここで構造体の初期化をする
	createUser := model.User{
		FirstName: signUpUserValidation.FirstName,
		LastName:  signUpUserValidation.LastName,
		Email:     signUpUserValidation.Email,
		Password:  uh.epl.EncryptPassword(signUpUserValidation.Password),
	}

	// 保存処理
	// if err := createUser.CreateUser(); err != nil {
	// 	uh.ru.SendErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
	// 	log.Println(err)
	// 	return
	// }

	// createUser -> ポインタ型(アドレス)
	if err := uh.ru.SendAuthResponse(w, &createUser, http.StatusOK); err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, "Unknown error occurred", http.StatusBadRequest)
		return
	}
}

func (uh *userHandler) SignInFunc(w http.ResponseWriter, r *http.Request) {

}
