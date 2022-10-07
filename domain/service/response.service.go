// domain

package service

import (
	"ddd/domain/model"
	"net/http"
)

// ここの層はinterfaceを提供するのみでOK！(DDDの場合)

type ResponseLogic interface {
	SendResponseLogic(w http.ResponseWriter, response []byte, code int) error
	SendErrorResponseLogic(w http.ResponseWriter, errorMessage string, code int) error
	SendAuthResponseLogic(w http.ResponseWriter, user *model.User, code int) error
}
