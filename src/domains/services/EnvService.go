package services

import (
	"errors"
	"os"
)

var PostgresqlUrl string
var JWT_secret string
var APP_URL string

//var PostgresqlUrl = "postgres://user:pass@localhost:8090/test?sslmode=disable"
//var JWT_secret = "secret"

func GetFromEnv() error {
	var exist bool

	PostgresqlUrl, exist = os.LookupEnv("POSTGRES_URL")
	if !exist {
		return errors.New("PostgresqlUrl is not found")
	}

	JWT_secret, exist = os.LookupEnv("JWT_SECRET")
	if !exist {
		return errors.New("JWT_secret is not found")
	}

	APP_URL, exist = os.LookupEnv("APP_URL")
	if !exist {
		return errors.New("APP_URL is not found")
	}

	return nil
}
