// interface層 (usecase層に依存)

package util

import (
	"ddd/domain/model"
	"ddd/interface/custom"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtUtil interface {
	CreateJWTToken(u *model.User) (string, error)
	CheckJWTToken(r *http.Request) (int, error)
}

type jwtUtil struct{}

func NewJwtUtil() JwtUtil {
	return &jwtUtil{}
}

// 新規登録が成功したらトークンを発行してレスポンスに含める。
// Userと紐づいているのでメソッドでOK。
func (jl jwtUtil) CreateJWTToken(u *model.User) (string, error) {

	// クレームの作成
	claim := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダー部分とペイロードの作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 署名をして完全なjwtを生成する
	// 引数にtoken.SignedString(os.Getenv("JWTSIGNKEY")) だとエラー
	jwtToken, err := token.SignedString([]byte(os.Getenv("JWTSIGNKEY")))
	if err != nil {
		err = custom.NewJwtErr("faild to get signed token", err)
		return "", err
	}

	return jwtToken, nil
}

// リクエスト時のJWTTokenの検証
func (jl jwtUtil) CheckJWTToken(r *http.Request) (int, error) {

	// リクエスト構造体を渡す -> リクエストヘッダーの取得する

	tokenString := r.Header.Get("Authorization")

	// authrizationが別の種類だとpanic発生するので以下のように書き換え
	// 文字列がBearerで始まるかどうか検証
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return 0, custom.ErrInvalidToken // errorインターフェースの作成
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// ここのtoken -> 無名関数である(あくまで関数の定義) -> Parse()の内部処理で使用する -> tokenの値を使用可能 -> Parse関数の説明をしっかり読めば分かる 用意するだけ
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// 型アサーション -> algの検証を行う
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, custom.ErrInvalidSignature
		}

		// 暗号鍵を返さなくてないいけないとドキュメントに書いてある。SigningMethodHMACのキーは[]byteで返してあげる
		return []byte(os.Getenv("JWTSIGNKEY")), nil

	})

	// jwt.MapClaimsだけどここの型はインターフェース -> map[exp:1.662545596e+09 user_id:1] -> {"user_id":1}
	// 型アサーションが必要だけどなぜこうなる？
	fmt.Printf("type and value of parsedToken: %+T, %+v\n", parsedToken.Claims, parsedToken.Claims)

	// 何らかのエラー
	if err != nil {
		err = custom.NewJwtErr("failed to parse token", err)
		return 0, err
	}

	// これは？
	if !parsedToken.Valid {
		return 0, custom.ErrInvalidToken
	}

	// user_idを取り出したい
	// 型アサーション -> falseだった時の処理をやっぱり加えた方がいいか？
	assertionToken, ok := parsedToken.Claims.(jwt.MapClaims)
	fmt.Printf("value: %+v\n", assertionToken)
	if !ok {
		return 0, custom.ErrAssertType
	}

	// map[string]interface{} -> {"string":"interface{}"}
	// mapのバリューに対してのアクセス -> map名["key"]
	// fmt.Printf("value[\"user_id\"]: %v\n", value["user_id"])

	// これだとまだ以下userIDはinterface型であり, int型ではない。
	// userID := value["user_id"]

	// 型を確認 -> float64と返ってくる！
	// fmt.Printf("type: assertionToken[\"user_id\"]: %T\n", assertionToken["user_id"])

	// 再度型アサーション
	assertionUserID, ok := assertionToken["user_id"].(float64)
	if !ok {
		return 0, custom.ErrAssertType
	}

	// 一応user_idを返す いずれ必要であれば*Tokenを返してあげる
	// return parsedToken, nil

	// 型キャスト
	return int(assertionUserID), nil
}
