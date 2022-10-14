// infrastructure層 (domain層に依存)

package persistence

import (
	"ddd/domain/model"
	"ddd/domain/repository"
	"errors"

	"github.com/jinzhu/gorm"
)

// infrastructure層では、domain層に依存するように書く。
// domain層でインターフェースとして定義したメソッドの具体的な技術的処理を書いていく。

// 以下の全処理はDBの処理なのでDBのメソッドとして定義する

type habitPersistence struct {
	Conn *gorm.DB
}

// ここは infrastructure層だけど、domain層のインターフェースを返す
func NewHabitPersistence(conn *gorm.DB) repository.HabitRepository {
	return &habitPersistence{Conn: conn}
}

// habitPersistence構造体のメソッドとして定義する
func (h *habitPersistence) CreateHabitPersistence(habit *model.Habit) error {

	// DB接続の用意
	db := h.Conn

	if err := db.Create(habit).Error; err != nil {
		return err
	}
	// fmt.Printf("h: %v\n", h) // h: {{2 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 <nil>} hello false 1}
	return nil

}

func (h *habitPersistence) UpdateHabitPersistence(habit *model.Habit) error {

	db := h.Conn

	result := db.Model(h).Where("id = ? AND user_id = ?", habit.Model.ID, habit.UserID).Update("content", habit.Content)

	if err := result.Error; err != nil {
		return err
	}

	// 実際にレコードが存在し、更新されたかどうかの判定は以下で行う
	if result.RowsAffected < 1 {
		err := errors.New("not found record") // 当たり前のように論理削除していたら更新は不可
		return err
	}

	return nil
}

func (h *habitPersistence) DeleteHabitPersistence(habitID, userID int, habit *model.Habit) error {

	// DB接続の用意
	db := h.Conn

	// &habitでもhabitでも問題がないのは内部でリフレクションが行われているため？
	result := db.Where("id = ? AND user_id = ?", habitID, userID).Delete(habit)

	if err := result.Error; err != nil {
		return err
	}

	// 実際にレコードが存在し、削除されたかどうかの判定は以下で行う
	if result.RowsAffected < 1 {
		err := errors.New("not found record")
		return err
	}

	return nil
}

//実体を受け取って、実体を書き換えるので、戻り値に指定する必要はない。
// 旧: 値渡し, 新: ポインタを受け取る
func (h *habitPersistence) GetAllHabitByUserIDPersistence(user *model.User, habit *[]model.Habit) error {

	// habitテーブル内の外部キーであるuseridで全てを取得する
	// fmt.Printf("u.ID: %v\n", u.ID)     // 1
	// fmt.Printf("[]habit: %v\n", habit) // 空の構造体

	db := h.Conn // 構造体のフィールドにアクセス

	// 全て取得したい
	if err := db.Where("user_id = ?", user.ID).Find(habit).Error; err != nil {
		// ここの戻り値
		return err
	}
	// fmt.Printf("habit: %v\n", habit) // habit: [{{2 2022-09-07 04:47:29 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is test false 1} {{3 2022-09-07 04:49:30 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} aaa false 1} {{4 2022-09-07 04:49:31 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is test false 1} {{5 2022-09-07 04:50:22 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is testbbb false 1} {{6 2022-09-07 04:55:55 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} aaadsadsa false 1}]
	return nil
}
