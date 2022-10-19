package util

import "errors"

// JWT関連

// *** 独自エラーの作成 ここから *** //

type jwtErr struct {
	message      string
	innerMessage error
}

func (e *jwtErr) Error() string {
	return e.message
}

// *** 独自エラーの作成 ここまで *** //

// 戻り値で設定するのは`型`ね
func NewJwtErr(message string, innerMessage error) *jwtErr {
	return &jwtErr{message, innerMessage} // これは`値`ね
}

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrInvalidSignature = errors.New("invalid signature method")
	ErrAssertType       = errors.New("faild to assert type")
)
