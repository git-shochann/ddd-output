package service

import (
	"ddd/domain/model"
	"net/http"
)

type JwtLogic interface {
	CreateJWTToken(u *model.User) (string, error)
	CheckJWTToken(r *http.Request) (int, error)
}
