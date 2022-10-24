// usecase (domainに依存)

package usecase

import (
	"ddd/domain"
	"ddd/domain/model"
	"fmt"
)

// メソッドを提供する
type UserUseCase interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

// 依存先のインターフェースがあれば書く
type userUseCase struct {
	ur domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) UserUseCase {
	return &userUseCase{
		ur: ur,
	}
}

func (uuc *userUseCase) CreateUser(user *model.User) (*model.User, error) {

	if err := uuc.ur.CreateUserInfrastructure(user); err != nil {
		err := fmt.Errorf("userUseCase: failed to create user %w", err)
		return nil, err
	}

	// 書き変わったuserを返す
	return user, nil
}

func (uuc *userUseCase) GetUserByEmail(email string) (*model.User, error) {

	user, err := uuc.ur.GetUserByEmailInfrastructure(email)
	if err != nil {
		err := fmt.Errorf("userUseCase: failed to get user %w", err)
		return nil, err
	}
	// 書き変わったuserを返す
	return user, nil
}
