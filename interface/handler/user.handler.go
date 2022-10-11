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

	var signUpUserValidationValidation model.UserSignUpValidation
	err = json.Unmarshal(reqBody, &signUpUserValidationValidation)

	if err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		return
	}

	errorMessage, err := uh.uv.SignupValidator(&signUpUserValidationValidation)
	if err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// ユーザーを登録する準備 -> リファクタリング後
	createUser := model.User{
		FirstName: signUpUserValidation.FirstName,
		LastName:  signUpUserValidation.LastName,
		Email:     signUpUserValidation.Email,
		Password:  EncryptPassword(signUpUserValidation.Password),
	}

	// 実際にDBに登録する
	if err := createUser.CreateUser(); err != nil {
		models.SendErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// createUser -> ポインタ型(アドレス)
	if err := models.SendAuthResponse(w, &createUser, http.StatusOK); err != nil {
		models.SendErrorResponse(w, "Unknown error occurred", http.StatusBadRequest)
		log.Println(err)
		return
	}
}

func (uh *userHandler) SignInFunc(w http.ResponseWriter, r *http.Request) {

}
