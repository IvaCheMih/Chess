package services

import (
	"github.com/spf13/viper"
)

var PostgresqlUrl string
var JWT_secret string
var APP_URL string
var MODE string

//var PostgresqlUrl = "postgres://user:pass@localhost:8090/test?sslmode=disable"
//var JWT_secret = "secret"

func GetFromEnv() {
	MODE = viper.GetString("MODE")

	if MODE == "LOCAL" {
		PostgresqlUrl = viper.GetString("POSTGRES_URL_LOCAL")
	} else {
		PostgresqlUrl = viper.GetString("POSTGRES_URL_REMOTE")
	}

	JWT_secret = viper.GetString("JWT_SECRET")
	APP_URL = viper.GetString("APP_URL")
}

//func GetFromEnv() error {
//	var exist bool
//
//	MODE, exist = os.LookupEnv("MODE")
//	if !exist {
//		return errors.New("MODE is not found")
//	}
//
//	if MODE == "LOCAL" {
//		PostgresqlUrl, exist = os.LookupEnv("POSTGRES_URL_LOCAL")
//		if !exist {
//			return errors.New("PostgresqlUrl is not found")
//		}
//	} else {
//		PostgresqlUrl, exist = os.LookupEnv("POSTGRES_URL_REMOTE")
//		if !exist {
//			return errors.New("PostgresqlUrl is not found")
//		}
//	}
//
//	JWT_secret, exist = os.LookupEnv("JWT_SECRET")
//	if !exist {
//		return errors.New("JWT_secret is not found")
//	}
//	APP_URL, exist = os.LookupEnv("APP_URL")
//	if !exist {
//		return errors.New("APP_URL is not found")
//	}
//
//	return nil
//}
