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

	// infrastructure層に初期設定を渡す(技術的関心事の処理はinfrastructure層で実装する)
	habitPersistence := persistence.NewHabitPersistence(config.ConnectDB()) // infrastructure層の設定

	// usecase層にinfrastructure層を渡す？ -> メソッドにアクセスできるように
	habitUseCase := usecase.NewHabitUseCase(habitPersistence) // usecase -> domain

	// interface層にusecase層を渡す -> メソッドにアクセスできるように
	habitHandler := handler.NewHabitHandler(habitUseCase) // interface -> usecase

	// 環境変数の読み込み
	envLogic := logic.NewEnvLogic() // infrastructure層
	// envLogic.LoadEnv()

	// ログの設定
	loggingLogic := logic.NewLoggingLogic()
	// loggingLogic.LoggingSetting()

	// JWTの設定

	jwtLogic := logic.NewJwtLogic()
	// fmt.Printf("jwtLogic: %t\n", jwtLogic)
	// fmt.Printf("jwtLogic: %v\n", jwtLogic)

	// レスポンスの設定
	responseLogic := logic.NewResponseLogic(jwtLogic)

	// responseLogicは logic.ResponseLogic型(インターフェース型) で インターフェースのメソッドを利用することが出来る型

	// で、このresponseLogicを渡したいのは、usecase層...？
	// usecase層 -> infrastructure層にアクセスしたい！ は直接は出来ないので
	// domain層 -> に、interfaceを置いて、窓口を作ってあげる

	// usecase層にinfrastrcuture層を渡す

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
