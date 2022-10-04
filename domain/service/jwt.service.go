package service

import (
	"ddd/domain/model"
	"net/http"
)

// ここの層はinterfaceを提供するのみでOK！(DDDの場合)

type JwtLogic interface {
	CreateJWTToken(u *model.User) (string, error)
	CheckJWTToken(r *http.Request) (int, error)
}
