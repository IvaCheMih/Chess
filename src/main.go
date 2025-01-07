package main

import (
	_ "github.com/IvaCheMih/chess/src/docs"
	"github.com/IvaCheMih/chess/src/domains/auth"
	"github.com/IvaCheMih/chess/src/domains/game"
	"github.com/IvaCheMih/chess/src/domains/services/env"
	"github.com/IvaCheMih/chess/src/domains/services/migrations"
	telegram "github.com/IvaCheMih/chess/src/domains/services/telegrem"
	"github.com/IvaCheMih/chess/src/domains/user"
	swagger "github.com/arsmn/fiber-swagger/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

var userHandlers user.UserHandlers
var gamesHandlers game.GamesHandlers
var authHandlers auth.AuthHandlers

func Init(envs *env.EnvService) {
	migrationService := migrations.CreateMigrationService()

	err := migrationService.RunUp(envs.PostgresqlUrl, envs.Migrations)
	if err != nil {
		log.Fatalln(err)
	}

	db, err = gorm.Open(postgres.Open(envs.PostgresqlUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	usersRepository := user.CreateUsersRepository(db)
	boardCellsRepository := game.CreateBoardCellsRepository(db)
	movesRepository := game.CreateMovesRepository(db)
	gamesRepository := game.CreateGamesRepository(db)

	usersServices := user.CreateUsersService(&usersRepository)
	userHandlers = user.CreateUserHandlers(&usersServices, envs.JWTSecret)

	gamesService := game.CreateGamesService(&boardCellsRepository, &gamesRepository, &movesRepository)
	gamesHandlers = game.CreateGamesHandlers(&gamesService)

	authHandlers = auth.CreateAuthHandlers()
}

func Shutdown() {
	sqlDB, _ := db.DB()

	sqlDB.Close()
}

// @title 						Fiber Swagger Example API
// @version 					2.0
// @description 				This is a sample server.
// @termsOfService 				http://swagger.io/terms/

// @contact.name				API Support
// @contact.url 				http://www.swagger.io/support
// @contact.email				support@swagger.io

// @license.name 				Apache 2.0
// @license.url 				http://www.apache.org/licenses/LICENSE-2.0.html

// @host 						127.0.0.1:8080
// @BasePath 					/
// @schemes 					http
//
//	@securityDefinitions.apiKey JWT
//	@in                         header
//	@name                       Authorization
//	@description                JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func main() {
	envs := env.NewEnvService()

	Init(envs)

	if envs.TelegramBot != "" {
		telegramBot, err := telegram.NewTelegramBot(envs.TelegramBot, envs.AppURL)
		if err != nil {
			panic(err)
		}

		go telegramBot.StartBot()
	}

	defer Shutdown()

	server := fiber.New()

	server.Get("/swagger/*", swagger.HandlerDefault) // default

	server.Post("/user", userHandlers.CreateUser)

	server.Post("/user/sign-in/telegram/", userHandlers.TelegramSignIn)

	server.Post("/session", userHandlers.CreateSession)

	server.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(envs.JWTSecret)},
	}))

	server.Post("/game", authHandlers.Auth, gamesHandlers.CreateGame)

	server.Get("/game/:gameId/board", authHandlers.Auth, gamesHandlers.GetBoard)

	server.Get("/game/:gameId/history", authHandlers.Auth, gamesHandlers.GetHistory)

	server.Post("/game/:gameId/move", authHandlers.Auth, gamesHandlers.Move)

	server.Post("/game/:gameId/give-up", authHandlers.Auth, gamesHandlers.GiveUp)

	if err := server.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
