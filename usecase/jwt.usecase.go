package usecase

import (
	"ddd/domain/service"
	"net/http"
)

// 一旦JWTの認証を引き受けるusecaseとして実装する

type JwtUseCase interface {
	CheckJWTTokenUseCase(w http.ResponseWriter, r *http.Request) (int, error)
}

// ここの層から見た依存先のインターフェースをフィールドとして設定すればOK -> domain層のインターフェースをここでは記載する
type jwtUseCase struct {
	jl service.JwtLogic
	rl service.ResponseLogic
}

// これを一番最初のmain関数で使う
// domain層のインターフェースをここでは記載する
func NewJwtUseCase(jl service.JwtLogic) JwtUseCase {
	return &jwtUseCase{
		jl: jl,
	}
}

// ポイント: 具体的な処理はドメイン層に任せる

func (juc *jwtUseCase) CheckJWTTokenUseCase(w http.ResponseWriter, r *http.Request) (int, error) {
	// ここでjuc.jl.CheckJWTTokenLogic()を呼び出すだけ
	userID, err := juc.jl.CheckJWTTokenLogic(r)
	if err != nil {
		juc.rl.SendErrorResponseLogic(w, "Authentication error", http.StatusBadRequest)
		return 0, err
	}
	return userID, nil
}
