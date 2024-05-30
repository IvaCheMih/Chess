package services

import (
	"errors"
	"os"
)

var PostgresqlUrl string
var JWT_secret string

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

	return nil
}
