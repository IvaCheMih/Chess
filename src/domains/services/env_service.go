package services

import (
	"github.com/spf13/viper"
	"log"
)

type EnvService struct {
	PostgresqlUrl string
	JWTSecret     string
	AppURL        string
	MODE          string
	Migrations    string
}

func NewEnvService() *EnvService {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	var postgresqlUrl string
	mode := viper.GetString("MODE")

	if mode == "LOCAL" {
		postgresqlUrl = viper.GetString("POSTGRES_URL_LOCAL")
	} else {
		postgresqlUrl = viper.GetString("POSTGRES_URL_REMOTE")
	}

	return &EnvService{
		PostgresqlUrl: postgresqlUrl,
		JWTSecret:     viper.GetString("JWT_SECRET"),
		AppURL:        viper.GetString("APP_URL"),
		MODE:          mode,
		Migrations:    viper.GetString("MIGRATIONS"),
	}
}
