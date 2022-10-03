package logic

import (
	"ddd/domain/model"
	"net/http"
)

type ResponseLogic interface {
	SendResponse(w http.ResponseWriter, response []byte, code int) error
	SendErrorResponse(w http.ResponseWriter, errorMessage string, code int) error
	SendAuthResponse(w http.ResponseWriter, user *model.User, code int) error
}
