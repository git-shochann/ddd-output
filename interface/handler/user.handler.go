// interface層 (usecase層に依存)

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

	"golang.org/x/crypto/bcrypt"
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
	epl util.EncryptPasswordUtil // interface層
}

func NewUserHandler(uuc usecase.UserUseCase, uv validator.UserValidation, ju util.JwtUtil, ru util.ResponseUtil, epl util.EncryptPasswordUtil) UserHandler {
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
	newUser, err := uh.uuc.CreateUser(&createUser)
	if err != nil {
		uh.ru.SendErrorResponse(w, "Failed to create user", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// createUser -> ポインタ型(アドレス)
	if err := uh.ru.SendAuthResponse(w, newUser, http.StatusOK); err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, "Occurred unknown error", http.StatusBadRequest)
		return
	}
}

func (uh *userHandler) SignInFunc(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		return // router.HandleFunc())の第二引数に関数を渡すだけなので戻り値なし
	}

	var signInUserValidation model.UserSignInValidation
	if err := json.Unmarshal(reqBody, &signInUserValidation); err != nil {
		uh.ru.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	errorMessage, err := uh.uv.SigninValidator(&signInUserValidation)
	if err != nil {
		log.Println(err)
		uh.ru.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// emailでユーザーを検索する -> 成功したらuserに値が入る
	user, err := uh.uuc.GetUserByEmail(signInUserValidation.Email)
	if err != nil {
		uh.ru.SendErrorResponse(w, "Failed to get user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// 	ここでログインユーザーを取得出来たのでuserを使ってく
	// 	bcryptでDBはハッシュかしているので比較する処理を用意

	// 	fmt.Printf("signinUser.Password: %v\n", signInUser.Password)
	// 	fmt.Printf("user.Password: %v\n", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInUserValidation.Password))
	if err != nil {
		uh.ru.SendErrorResponse(w, "Password error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := uh.ru.SendAuthResponse(w, user, http.StatusOK); err != nil {
		uh.ru.SendErrorResponse(w, "Failed to sign in", http.StatusBadRequest)
		log.Println(err)
		return
	}

}
