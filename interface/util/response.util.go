// interface層 (usecase層に依存)

package util

import (
	"ddd/domain/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseUtil interface {
	SendResponse(w http.ResponseWriter, response []byte, code int) error
	SendErrorResponse(w http.ResponseWriter, err error, code int) error
	SendAuthResponse(w http.ResponseWriter, user *model.User, code int) error
}

type responseUtil struct {
	ju JwtUtil
}

func NewResponseUtil(ju JwtUtil) ResponseUtil {
	return &responseUtil{
		ju: ju,
	}
}

// ステータスコード200の場合のレスポンス
func (rl responseUtil) SendResponse(w http.ResponseWriter, response []byte, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

// ステータスコード200以外のレスポンスで使用
// message: err.Error() とする
func (rl responseUtil) SendErrorResponse(w http.ResponseWriter, err error, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := map[string]error{
		"message": err,
	}
	// jsonに変換する
	responseBody, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = w.Write(responseBody)
	if err != nil {
		return err
	}
	return nil
}

// 新規登録とログイン時のレスポンスとして、JWTトークンとUser構造体を返却する
func (rl responseUtil) SendAuthResponse(w http.ResponseWriter, user *model.User, code int) error {

	fmt.Println("SendAuthResponse!")

	// NEW!: 以下のように呼び出す
	jwtToken, err := rl.ju.CreateJWTToken(user)
	if err != nil {
		return err
	}

	// レスポンス
	response := model.UserAuthResponse{
		User:     *user, // デリファレンスする
		JwtToken: jwtToken,
	}

	// 構造体をjsonに変換
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	fmt.Printf("jsonResponse: %v\n", string(jsonResponse))

	if err := rl.SendResponse(w, jsonResponse, code); err != nil {
		return err
	}

	return nil

}
