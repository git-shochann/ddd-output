// domain層 (依存なし)

package domain

import "ddd/domain/model"

// ここの層はinterfaceを提供するのみでOK！ (DDDの場合)

type UserRepository interface {
	CreateUserInfrastructure(user *model.User) error
	GetUserByEmailInfrastructure(email string) (*model.User, error)
}
