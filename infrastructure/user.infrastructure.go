// infrastructure層 (domain層に依存)

package infrastructure

import (
	"ddd/domain"
	"ddd/domain/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

// エラーハンドリング
// https://gorm.io/ja_JP/docs/error_handling.html

type userInfrastructure struct {
	Conn *gorm.DB
}

// ここは infrastructure層だけど、domain層のインターフェースを返す
func NewUserInfrastructure(conn *gorm.DB) domain.UserRepository {
	return &userInfrastructure{Conn: conn}
}

// ポインタ渡し -> 元の実体を書き換えるので
func (u userInfrastructure) CreateUserInfrastructure(user *model.User) error {

	// メソッドとして定義しているのでフィールドにアクセスして実行する
	db := u.Conn

	if err := db.Create(user).Error; err != nil {
		fmt.Println(err)
		err = NewDbErr("failed to create user", err)
		fmt.Println(err)
		return err
	}
	return nil

}

// Emailを元に重複していないか検索をする
// User構造体の値はなぜ必要？ -> 結果を格納するため(out) -> First(out interface{}, where ...interface{}) *gorm.DB
func (u userInfrastructure) GetUserByEmailInfrastructure(email string) (*model.User, error) {

	// DB接続の設定
	db := u.Conn

	var user model.User
	// 検索する
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		// *Userと戻り値を設定しないと、nilで返せない。Userとするとゼロ値が埋められた構造体が返ってしまうので。
		err = NewDbErr("faild to get user", err)
		return nil, err
	}

	return &user, nil
}
