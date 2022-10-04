package service

// ここの層はinterfaceを提供するのみでOK！

type EncryptPasswordLogic interface {
	EncryptPassword(password string) string
}
