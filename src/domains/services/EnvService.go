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
