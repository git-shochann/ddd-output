package persistence

import (
	"ddd/domain/model"
	"ddd/domain/repository"

	"gorm.io/gorm"
)

// エラーハンドリング
// https://gorm.io/ja_JP/docs/error_handling.html

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	// WIP: ここなんでポインタ型で返す？
	return &userPersistence{Conn: conn}
}

// ポインタ渡し -> 元の実体を書き換えるので
func (u userPersistence) CreateUser(user *model.User) error {

	// メソッドとして定義しているのでフィールドにアクセスして実行する
	db := u.Conn

	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil

}

// Emailを元に重複していないか検索をする
// User構造体の値はなぜ必要？ -> 結果を格納するため(out) -> First(out interface{}, where ...interface{}) *gorm.DB
func (u userPersistence) GetUserByEmail(email string) (*model.User, error) {

	// DB接続の設定
	db := u.Conn

	var user model.User
	// 検索する
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		// *Userと戻り値を設定しないと、nilで返せない。Userとするとゼロ値が埋められた構造体が返ってしまうので。
		return nil, err
	}

	return &user, nil
}
