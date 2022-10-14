// interface層 (usecase層に依存)

package util

import (
	"log"

	"github.com/joho/godotenv"
)

type EnvUtil interface {
	LoadEnv()
}

type envUtil struct{}

func NewEnvUtil() EnvUtil {
	return &envUtil{}
}

func (ev envUtil) LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load env file: %v", err)
	}
}
