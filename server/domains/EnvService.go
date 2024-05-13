package domains

import "os"

var PostgresqlUrl string
var JWT_secret string

func GetURLsFromEnv(env []string) {

	PostgresqlUrl, _ = os.LookupEnv(env[0])

	JWT_secret, _ = os.LookupEnv(env[1])
}
