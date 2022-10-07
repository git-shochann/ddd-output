package service

// ここの層はinterfaceを提供するのみでOK！(DDDの場合)

type EncryptPasswordLogic interface {
	EncryptPasswordLogic(password string) string
}
