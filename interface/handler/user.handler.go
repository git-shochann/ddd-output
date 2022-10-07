// interface (usecaseに依存)

package handler

// 1. まず何のメソッドを定義して、直接interface層に依存するところ(main関数)で使いたいか？
// 2. ここの構造体はどうやって書く？ 依存する層のインターフェースがあればフィールドとして定義する。
// 3. NewUserHandler()はmain関数から呼び出すので、実際に構造体のフィールドで必要な型を受け取るように設定する
// 4. 完了!

// type UserHandler interface {
// 	SignUpFunc(http.ResponseWriter, *http.Request)
// 	SignInFunc(http.ResponseWriter, *http.Request)
// }

// type userHandler struct {
// 	UserUseCase usecase.UserUseCase
// }

// func NewUserHandler() UserHandler {
// 	return &userHandler{}
// }

// func (uh *userHandler) SignUpFunc(w http.ResponseWriter, r *http.Request) {

// 	// Jsonでくるので、まずGoで使用できるようにする
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
// 		return
// 	}

// 	var signUpUser models.UserSignUpValidation
// 	err = json.Unmarshal(reqBody, &signUpUser)

// 	if err != nil {
// 		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
// 		log.Println(err)
// 		return
// 	}

// 	errorMessage, err := signUpUser.SignupValidator()
// 	if err != nil {
// 		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
// 		log.Println(err)
// 		return
// 	}

// 	// ユーザーを登録する準備 -> リファクタリング後
// 	createUser := models.User{
// 		FirstName: signUpUser.FirstName,
// 		LastName:  signUpUser.LastName,
// 		Email:     signUpUser.Email,
// 		Password:  models.EncryptPassword(signUpUser.Password),
// 	}

// 	// 実際にDBに登録する
// 	if err := createUser.CreateUser(); err != nil {
// 		models.SendErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
// 		log.Println(err)
// 		return
// 	}

// 	// createUser -> ポインタ型(アドレス)
// 	if err := models.SendAuthResponse(w, &createUser, http.StatusOK); err != nil {
// 		models.SendErrorResponse(w, "Unknown error occurred", http.StatusBadRequest)
// 		log.Println(err)
// 		return
// 	}
// }

// func (uh *userHandler) SignInFunc(w http.ResponseWriter, r *http.Request) {

// }
