package service

import (
	"ddd/domain/model"
	"net/http"
)

// ここの層はinterfaceを提供するのみでOK！(DDDの場合)

type ResponseLogic interface {
	SendResponse(w http.ResponseWriter, response []byte, code int) error
	SendErrorResponse(w http.ResponseWriter, errorMessage string, code int) error
	SendAuthResponse(w http.ResponseWriter, user *model.User, code int) error
}
