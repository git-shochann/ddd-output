package infrastructure

import (
	"errors"
	"fmt"
)

// 独自エラーの作成
// デバッグがしやすいように、実際の処理で起きたパッケージも格納する

type DbErr struct {
	message      string
	innerMessage error
}

func (e *DbErr) Error() string {
	return e.message
}

func NewDbErr(message string, innerMessage error) *DbErr {
	fmt.Printf("innerMessage: %v\n", innerMessage) // innerMessage: Error 1062: Duplicate entry 'tarotaro2@gmail.com' for key 'users.email'
	return &DbErr{message, innerMessage}           // 構造体の初期化 + ポインタ化 + 関数に戻り値として返す
}

// Errから始める これも慣習
var ErrRecordNotFound = errors.New("not found record")
