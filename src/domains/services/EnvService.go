package services

import "os"

var PostgresqlUrl string
var JWT_secret string

func GetFromEnv(env []string) {

	PostgresqlUrl, _ = os.LookupEnv(env[0])

	JWT_secret, _ = os.LookupEnv(env[1])
}
