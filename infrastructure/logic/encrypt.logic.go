package logic

import "golang.org/x/crypto/bcrypt"

type EncryptPasswordLogic interface {
	EncryptPassword(password string) string
}

type encryptPasswordLogic struct{}

// main関数で使うために用意
// インターフェース型を返せば、呼び出し元でそのメソッドが使用することが出来る
func NewEncryptPasswordLogic() EncryptPasswordLogic {
	return &encryptPasswordLogic{}
}

func (epl *encryptPasswordLogic) EncryptPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
