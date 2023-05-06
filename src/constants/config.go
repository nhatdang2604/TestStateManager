package constants

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type ConfigConstant struct {
	Host string
	Port string

	DBHost                          string
	DBPort                          string
	DBUser                          string
	DBPass                          string
	DBName                          string
	DBProtocol                      string
	DBMaxConnectionLifeTimeInMinute int
	DBMaxOpenConnection             int
	DBMaxIdleConnection             int
}

var configConstant *ConfigConstant
var once sync.Once

func GetConfigConstant() *ConfigConstant {
	once.Do(func() {
		configConstant, _ = NewConfigConstant()
	})

	return configConstant
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

	configConstant.DBHost = os.Getenv("DB_HOST")
	configConstant.DBPort = os.Getenv("DB_PORT")
	configConstant.DBUser = os.Getenv("DB_USER")
	configConstant.DBPass = os.Getenv("DB_PASS")
	configConstant.DBName = os.Getenv("DB_NAME")
	configConstant.DBProtocol = os.Getenv("DB_PROTOCOL")

	configConstant.DBMaxConnectionLifeTimeInMinute, err = strconv.Atoi(os.Getenv("DB_MAX_CONN_LIFE_IN_MINUTE"))
	configConstant.DBMaxOpenConnection, err = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	configConstant.DBMaxIdleConnection, err = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))

	if nil != err {
		log.Fatalf("Error on parsing string to int: %v", err)
		return nil, nil
	}

	return configConstant, nil
}
