// usecase (domainに依存)

package usecase

import (
	"ddd/domain"
	"ddd/domain/model"
	"log"
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

	if err := uuc.ur.CreateUser(user); err != nil {
		log.Println(err)
		return nil, err
	}

	// 書き変わったuserを返す
	return user, nil
}

func (uuc *userUseCase) GetUserByEmail(email string) (*model.User, error) {

	user, err := uuc.ur.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// 書き変わったuserを返す
	return user, nil
}
