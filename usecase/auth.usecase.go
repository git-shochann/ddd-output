package usecase

import (
	"ddd/domain/logic"
	"net/http"
)

// 一旦JWTの認証を引き受けるusecaseとして実装する

type JwtUseCase interface {
	CheckJWTToken(w http.ResponseWriter, r *http.Request)
}

// 依存する層のインターフェースをフィールドとして設定する -> domain層の実際のロジックを記載
type jwtUseCase struct {
	logic.JwtLogic
}

func NewJwtLogic() JwtUseCase {
	return &jwtUseCase{}
}

func (juc *jwtUseCase) CheckJWTToken(w http.ResponseWriter, r *http.Request) {

}
