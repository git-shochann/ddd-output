package service

type EncryptPasswordLogic interface {
	EncryptPassword(password string) string
}
