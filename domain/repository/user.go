package repository

import "ddd/domain/model"

// repositoryはDBとのやりとりを定義するが、
// 技術的関心ごとはinfrastructure層に書くため、ここではインターフェースとしてメソッドを定義して...？

// habitにおけるrepositoryのインターフェース
// 実際のメソッド群を登録する

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}
