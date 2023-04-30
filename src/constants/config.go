package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigConstant struct {
	Host string
	Port string
}

func NewConfigConstant() (*ConfigConstant, error) {
	err := godotenv.Load()
	if nil != err {
		log.Fatalf("Error in loading .env file: %v", err)
		return nil, err
	}

	configConstant := new(ConfigConstant)
	configConstant.Host = os.Getenv("APP_HOST")
	configConstant.Port = os.Getenv("APP_PORT")

	return configConstant, nil
}
