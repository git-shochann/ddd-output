package main

import (
	"ddd/config"
	"ddd/infrastructure"
	"ddd/interface/handler"
	"ddd/interface/util"
	"ddd/interface/validator"
	"fmt"
	"log"
	"net/http"

	"ddd/usecase"

	"github.com/gorilla/mux"
)

// ここでは依存関係(どこの層がどこを利用するか)とルーティングを定義する
// 各層の依存関係を定義することで、利用可能な状態にする

func main() {

	db := config.ConnectDB()

	/*** 1. infrastructure層からdomain層をインターフェースを提供する ***/

	habitInfrastructure := infrastructure.NewHabitInfrastructure(db)
	userInfrastructure := infrastructure.NewUserInfrastructure(db)

	/*** 2. usecase層でdomain層を使うためにdomain層のインターフェースを渡す usecase -> domain ***/

	habitUseCase := usecase.NewHabitUseCase(habitInfrastructure)

	userUseCase := usecase.NewUserUseCase(userInfrastructure)

	/*** 3. usecase層で使用するためのinterface層を準備する ***/
	habitValidation := validator.NewHabitValidation()
	userValidation := validator.NewUserValidation()
	jwtUtil := util.NewJwtUtil()
	responseUtil := util.NewResponseUtil(jwtUtil)
	encryptPasswordUtil := util.NewEncryptPasswordUtil()

	/*** 4. interface層でusecase層を使うためにusecase層のインターフェースを渡す interface -> usecase ***/

	habitHandler := handler.NewHabitHandler(habitUseCase, habitValidation, jwtUtil, responseUtil)
	userHandler := handler.NewUserHandler(userUseCase, userValidation, jwtUtil, responseUtil, encryptPasswordUtil)

	/*** 5. ルーティングの設定 ***/

	router := mux.NewRouter().StrictSlash(true)

	// HandleFunc() 第二引数 -> 引数に関数
	router.HandleFunc("/", habitHandler.IndexFunc).Methods("GET")

	// for user
	router.HandleFunc("/api/v1/signup", userHandler.SignUpFunc).Methods("POST")
	router.HandleFunc("/api/v1/signin", userHandler.SignInFunc).Methods("POST")

	// for habit
	router.HandleFunc("/api/v1/create", habitHandler.CreateFunc).Methods("POST")
	router.HandleFunc("/api/v1/update/{id}", habitHandler.UpdateFunc).Methods("PATCH")
	router.HandleFunc("/api/v1/delete/{id}", habitHandler.DeleteFunc).Methods("DELETE")
	router.HandleFunc("/api/v1/get", habitHandler.GetAllHabitFunc).Methods("GET")

	fmt.Println("Start Server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
