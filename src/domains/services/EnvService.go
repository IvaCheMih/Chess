package services

import "os"

var PostgresqlUrl string
var JWT_secret string

func GetFromEnv() {
	var exist bool

	PostgresqlUrl, exist = os.LookupEnv("POSTGRES_URL")
	if !exist {
		panic("PostgresqlUrl is not found")
	}

	JWT_secret, exist = os.LookupEnv("JWT_SECRET")
	if !exist {
		panic("JWT_secret is not found")
	}
}
