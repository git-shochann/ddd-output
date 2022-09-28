package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// 新規登録が成功したらトークンを発行してレスポンスに含める。
// Userと紐づいているのでメソッドでOK。
func (u *User) CreateJWTToken() (string, error) {

	// クレームの作成
	claim := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダー部分とペイロードの作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 署名をして完全なjwtを生成する
	// 引数にtoken.SignedString(os.Getenv("JWTSIGNKEY")) だとエラー
	jwtToken, err := token.SignedString([]byte(os.Getenv("JWTSIGNKEY")))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
