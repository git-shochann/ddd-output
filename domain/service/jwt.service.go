package service

import (
	"ddd/domain/model"
	"net/http"
)

// ここの層はinterfaceを提供するのみでOK！(DDDの場合)

type JwtLogic interface {
	CreateJWTTokenLogic(u *model.User) (string, error)
	CheckJWTTokenLogic(r *http.Request) (int, error)
}
