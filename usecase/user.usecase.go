// usecase (domainに依存)

package usecase

import (
	"ddd/domain/model"
	"ddd/domain/repository"
	"log"
)

// メソッドを提供する
type UserUseCase interface {
	CreateUser(user *model.User) (*model.User, error)
}

// 依存先のインターフェースがあれば書く
type userUseCase struct {
	ur repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		ur: ur,
	}
}

func (uuc *userUseCase) CreateUser(user *model.User) (*model.User, error) {

	if err := uuc.ur.CreateUser(user); err != nil {
		log.Println(err)
		return nil, err
	}

	// 書き変わったuserを返す
	return user, nil
}
