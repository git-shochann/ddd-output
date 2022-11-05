// interface層 (usecase層に依存)

package handler

import (
	"ddd/domain/model"
	"ddd/infrastructure"
	"ddd/interface/util"
	"ddd/interface/validator"
	"ddd/usecase"
	"encoding/json"
	"errors"
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
	userUseCase         usecase.UserUseCase      // usecase層
	userValidation      validator.UserValidation // interface層
	jwtUtil             util.JwtUtil             // interface層
	responseUtil        util.ResponseUtil        // interface層
	encryptPasswordUtil util.EncryptPasswordUtil // interface層
}

func NewUserHandler(userUseCase usecase.UserUseCase, userValidation validator.UserValidation, jwtUtil util.JwtUtil, responseUtil util.ResponseUtil, encryptPasswordUtil util.EncryptPasswordUtil) UserHandler {
	return &userHandler{
		userUseCase:         userUseCase,
		userValidation:      userValidation,
		jwtUtil:             jwtUtil,
		responseUtil:        responseUtil,
		encryptPasswordUtil: encryptPasswordUtil,
	}
}

func (uh *userHandler) SignUpFunc(w http.ResponseWriter, r *http.Request) {

	// Jsonでくるので、まずGoで使用できるようにする
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, "failed to read json", http.StatusBadRequest)
		return // router.HandleFunc())の第二引数に関数を渡すだけなので戻り値なし
	}

	var signUpUserValidation model.UserSignUpValidation
	err = json.Unmarshal(reqBody, &signUpUserValidation)

	if err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, "failed to read json", http.StatusBadRequest)
		return
	}

	errorMessage, err := uh.userValidation.SignupValidator(&signUpUserValidation)
	if err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// DBに登録する内容の準備 -> リファクタリング後 -> ここで構造体の初期化をする
	createUser := model.User{
		FirstName: signUpUserValidation.FirstName,
		LastName:  signUpUserValidation.LastName,
		Email:     signUpUserValidation.Email,
		Password:  uh.encryptPasswordUtil.EncryptPassword(signUpUserValidation.Password),
	}

	// 保存処理
	newUser, err := uh.userUseCase.CreateUser(&createUser)

	if err != nil {
		log.Println(err)
		var DbErr *infrastructure.DbErr
		switch {
		case errors.As(err, &DbErr):
			uh.responseUtil.SendErrorResponse(w, "failed to create user", http.StatusBadRequest)
		default:
			uh.responseUtil.SendErrorResponse(w, "unknown error occured", http.StatusInternalServerError)
		}
		return
	}

	// createUser -> ポインタ型(アドレス)
	if err := uh.responseUtil.SendAuthResponse(w, newUser, http.StatusOK); err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, "unknown error occurred", http.StatusBadRequest)
		return
	}
}

func (uh *userHandler) SignInFunc(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, "failed to read json", http.StatusBadRequest)
		return // router.HandleFunc())の第二引数に関数を渡すだけなので戻り値なし
	}

	var signInUserValidation model.UserSignInValidation
	if err := json.Unmarshal(reqBody, &signInUserValidation); err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, "failed to read json", http.StatusBadRequest)
		return
	}

	errorMessage, err := uh.userValidation.SigninValidator(&signInUserValidation)
	if err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		return
	}

	// emailでユーザーを検索する -> 成功したらuserに値が入る
	user, err := uh.userUseCase.GetUserByEmail(signInUserValidation.Email)
	if err != nil {
		log.Println(err)
		var DbErr *infrastructure.DbErr

		switch {
		case errors.Is(err, infrastructure.ErrRecordNotFound):
			uh.responseUtil.SendErrorResponse(w, "record not found", http.StatusBadRequest)
		case errors.Is(err, DbErr):
			uh.responseUtil.SendErrorResponse(w, "failed to get user", http.StatusBadRequest)
		default:
			uh.responseUtil.SendErrorResponse(w, "unknown error occured", http.StatusInternalServerError)
		}
		return
	}

	// 	ここでログインユーザーを取得出来たのでuserを使ってく
	// 	bcryptでDBはハッシュ化しているので比較する処理を用意

	// 	fmt.Printf("signinUser.Password: %v\n", signInUser.Password)
	// 	fmt.Printf("user.Password: %v\n", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInUserValidation.Password))
	if err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, "password invalid", http.StatusInternalServerError)
		return
	}

	if err := uh.responseUtil.SendAuthResponse(w, user, http.StatusOK); err != nil {
		log.Println(err)
		uh.responseUtil.SendErrorResponse(w, "failed to sign in", http.StatusBadRequest)
		return
	}

}
