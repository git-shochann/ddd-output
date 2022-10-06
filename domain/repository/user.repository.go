// domain

package repository

import "ddd/domain/model"

// ここの層はinterfaceを提供するのみでOK！ (DDDの場合)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}
