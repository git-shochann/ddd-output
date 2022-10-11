// interface層 (usecase層に依存)

package util

import "golang.org/x/crypto/bcrypt"

type EncryptPassword interface {
	EncryptPassword(password string) string
}

type encryptPassword struct{}

// main関数で使うために用意
// インターフェース型を返せば、呼び出し元でそのメソッドが使用することが出来る
func NewEncryptPasswordLogic() EncryptPassword {
	return &encryptPassword{}
}

func (epl *encryptPassword) EncryptPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
