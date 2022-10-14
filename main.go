package main

import (
	"ddd/config"
	"ddd/interface/handler"
	"ddd/interface/util"
	"ddd/interface/validator"

	"ddd/infrastructure/persistence"
	"ddd/usecase"
)

// ここでは依存関係(どこの層がどこを利用するか)とルーティングを定義する
// 各層の依存関係を定義することで、利用可能な状態にする

func main() {

	db := config.ConnectDB()

	/*** 1. infrastructure層からdomain層をインターフェースを提供する***/

	habitPersistence := persistence.NewHabitPersistence(db)
	userPersistence := persistence.NewUserPersistence(db)

	/*** 2. usecase層でdomain層を使うためにdomain層のインターフェースを渡す ***/

	habitUseCase := usecase.NewHabitUseCase(habitPersistence)

	userUseCase := usecase.NewUserUseCase(userPersistence)

	/*** 3. usecase層でdomain層を使うためにdomain層のインターフェースを渡す ***/
	habitValidation := validator.NewHabitValidation()
	userValidation := validator.NewUserValidation()
	jwtUtil := util.NewJwtLogic()
	responseUtil := util.NewResponseLogic()

	/*** 4. interface層でusecase層を使うためにusecase層のインターフェースを渡す ***/

	habitHandler := handler.NewHabitHandler(habitUseCase) // interface -> usecase
	userHandler := handler.NewUserHandler(userUseCase)    // interface -> usecase

	/*** ルーティングの設定 ***/

	// router := mux.NewRouter().StrictSlash(true)

	// // HandleFunc() 第二引数 -> 引数に関数
	// router.HandleFunc("/api/v1/signup", userHandler.SignUpFunc).Methods("POST")
	// router.HandleFunc("/api/v1/signin", userHandler.SignInFunc).Methods("POST")

	// router.HandleFunc("/", habitHandler.IndexFunc).Methods("GET")
	// router.HandleFunc("/api/v1/create", habitHandler.CreateFunc).Methods("POST")
	// router.HandleFunc("/api/v1/update/{id}", habitHandler.UpdateFunc).Methods("PATCH")
	// router.HandleFunc("/api/v1/delete/{id}", habitHandler.DeleteFunc).Methods("DELETE")
	// router.HandleFunc("/api/v1/get", habitHandler.GetAllHabitFunc).Methods("GET")

	// fmt.Println("Start Server!")
	// log.Fatal(http.ListenAndServe(":8080", router))
}
