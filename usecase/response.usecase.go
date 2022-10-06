// usecase

package usecase

import (
	"ddd/infrastructure/logic"
	"net/http"
)

// 1. 以下のメソッドを使いたい層のためにインターフェースをまずは書く
type ResponseUseCase interface {
	SendResponseUseCase(w http.ResponseWriter, response []byte, code int)
	SendErrorResponseUseCase()
	SendAuthResponseUseCase()
}

// 2. ここの層から見た直接依存する先のインターフェース型をフィールドとして設定する -> ここではdomain層
//    こちらの構造体のメソッドとして定義し、そのメソッド内でインターフェースのメソッドをさらに使用することが可能なため。
type responseUseCase struct {
	rl logic.ResponseLogic
}

// 3. １で作成したインターフェースを戻り値として設定して提供できるようにする
//    引数には、ここの層から見た直接依存する先のインターフェース型をフィールドとして設定する
//    main関数で依存関係を繋げる
func NewResponseUseCase(rl logic.ResponseLogic) ResponseUseCase {
	return &responseUseCase{
		rl: rl,
	}
}

// 4. 実際にメソッドを書いていく
func (ruc *responseUseCase) SendResponseUseCase(w http.ResponseWriter, response []byte, code int) {
	err := ruc.rl.SendResponseLogic(w, response, code)
	if err != nil {
		ruc.rl.SendErrorResponseLogic()
	}
}

func (ruc *responseUseCase) SendErrorResponseUseCase() {}

func (ruc *responseUseCase) SendAuthResponseUseCase() {}
