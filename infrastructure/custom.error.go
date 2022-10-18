package infrastructure

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
	return &DbErr{message, innerMessage} // 構造体の初期化 + ポインタ化 + 関数に戻り値として返す
}
