package main

import (
	"ddd/config"
	"ddd/infrastructure/logic"
	"ddd/infrastructure/persistence"
	"ddd/infrastructure/validator"
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

	db := config.ConnectDB()

	/* infrastructure層 */

	// 暗号化パスワードの設定
	encryptPasswordLogic := logic.NewEncryptPasswordLogic()

	// 環境変数の設定
	envLogic := logic.NewEnvLogic()

	// ログの設定
	loggingLogic := logic.NewLoggingLogic()

	// JWTの設定
	jwtLogic := logic.NewJwtLogic()

	// レスポンスの設定
	responseLogic := logic.NewResponseLogic(jwtLogic)

	// バリデーションの設定
	habitValidation := validator.NewHabitValidation()
	// userValidation := validator.NewUserValidation()

	// infrastructure層に初期設定を渡す(技術的関心事の処理はinfrastructure層で実装する)
	habitPersistence := persistence.NewHabitPersistence(db)
	// userPersistence := persistence.NewUserPersistence(db)

	// usecase -> domain

	// usecase層にinfrastructure層を渡す？ -> usecase層内でinfrastructure層のメソッドにアクセスできるように
	habitUseCase := usecase.NewHabitUseCase(habitPersistence, habitValidation, encryptPasswordLogic, envLogic, jwtLogic, loggingLogic, responseLogic)
	// userUseCase := usecase.NewUserUseCase(userPersistence, userValidation, encryptPasswordLogic, envLogic, jwtLogic, loggingLogic, responseLogic)

	// interface層にusecase層を渡す -> メソッドにアクセスできるように
	habitHandler := handler.NewHabitHandler(habitUseCase) // interface -> usecase
	// userHandler := handler.NewUserHandler(userUseCase)    // interface -> usecase

	// ルーティングの設定
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/api/v1/signup", userHandler.SignUpFunc).Methods("POST")
	// router.HandleFunc("/api/v1/signin", userHandler.SignInFunc).Methods("POST")

	router.HandleFunc("/", habitHandler.IndexFunc).Methods("GET") // 引数に関数
	router.HandleFunc("/api/v1/create", habitHandler.CreateFunc).Methods("POST")
	// router.HandleFunc("/api/v1/update/{id}", habitHandler.UpdateFunc).Methods("PATCH")
	// router.HandleFunc("/api/v1/delete/{id}", habitHandler.DeleteFunc).Methods("DELETE")
	// router.HandleFunc("/api/v1/get", habitHandler.GetAllHabitFunc).Methods("GET")

	fmt.Println("Start Server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
