package main

import (
	"ddd/config"
	"ddd/infrastructure/logic"
	"ddd/infrastructure/persistence"
	"ddd/interface/handler"
	"ddd/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ここでは依存関係(どこの層がどこを利用するか)とルーティングを定義する
// 各層の依存関係を定義することで、利用可能な状態にする

func main() {

	// 依存関係を順番につなげていく
	habitPersistence := persistence.NewHabitPersistence(config.ConnectDB()) // infrastructure層の設定
	habitUseCase := usecase.NewHabitUseCase(habitPersistence)               // usecase -> domain
	habitHandler := handler.NewHabitHandler(habitUseCase)                   // interface -> usecase

	// 環境変数の読み込み
	envLogic := logic.NewEnvLogic()
	envLogic.LoadEnv()

	// ログの設定
	loggingLogic := logic.NewLoggingLogic()
	loggingLogic.LoggingSetting()

	// JWTの設定
	jwtLogic := logic.NewJwtLogic()

	// レスポンスの設定
	responseLogic := logic.NewResponseLogic()

	// ルーティングの設定
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/api/v1/signup", userHandler.SignupFunc).Methods("POST")
	// router.HandleFunc("/api/v1/signin", userHandler.SigninFunc).Methods("POST")
	router.HandleFunc("/", habitHandler.IndexFunc).Methods("GET") // 引数に関数
	// router.HandleFunc("/api/v1/get", habitHandler.GetAllHabitFunc).Methods("GET")
	router.HandleFunc("/api/v1/create", habitHandler.CreateFunc).Methods("POST")
	// router.HandleFunc("/api/v1/update/{id}", habitHandler.UpdateHabitFunc).Methods("PATCH")
	// router.HandleFunc("/api/v1/delete/{id}", habitHandler.DeteteHabitFunc).Methods("DELETE")
	fmt.Println("Start Server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
