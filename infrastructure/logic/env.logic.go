package logic

import (
	"log"

	"github.com/joho/godotenv"
)

type EnvLogic interface {
	LoadEnv()
}

type envLogic struct{}

func NewEnvLogic() EnvLogic {
	return &envLogic{}
}

func (ev envLogic) LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load env file: %v", err)
	}
}
