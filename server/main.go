package main

import (
	"database/sql"
	_ "github.com/IvaCheMih/chess/server/docs"
	"github.com/IvaCheMih/chess/server/domains"
	"github.com/IvaCheMih/chess/server/domains/auth"
	"github.com/IvaCheMih/chess/server/domains/game"
	"github.com/IvaCheMih/chess/server/domains/game/move_service"
	"github.com/IvaCheMih/chess/server/domains/user"
	swagger "github.com/arsmn/fiber-swagger/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"log"
)

const connect = "postgres://user:pass@localhost:8090/test?sslmode=disable"

var db *sql.DB

var userHandlers user.UserHandlers
var gamesHandlers game.GamesHandlers
var authHandlers auth.AuthHandlers

func Init(postgresqlUrl string) {
	migrationService := domains.CreateMigrationService()

	//time.Sleep(5 * time.Second)

	migrationService.RunUp(postgresqlUrl, "file://migrations/postgresql")

	move_service.FigureRepo = move_service.CreateFigureRepo()

	var err error

	db, err = sql.Open("postgres", postgresqlUrl)
	if err != nil {
		panic(err)
	}

	usersRepository := user.CreateUsersRepository(db)
	boardCellsRepository := game.CreateBoardCellsRepository(db)
	figuresRepository := game.CreateFiguresRepository(db)
	movesRepository := game.CreateMovesRepository(db)
	gamesRepository := game.CreateGamesRepository(db)

	usersServices := user.CreateUsersService(&usersRepository)
	userHandlers = user.CreateUserHandlers(&usersServices)

	gamesService := game.CreateGamesService(&boardCellsRepository, &figuresRepository, &gamesRepository, &movesRepository)
	gamesHandlers = game.CreateGamesHandlers(&gamesService)

	authRepository := auth.CreateAuthRepository(db)
	authService := auth.CreateAuthService(&authRepository)
	authHandlers = auth.CreateAuthHandlers(&authService)
}

func Shutdown() {
	db.Close()
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

// @host 						localhost:8082
// @BasePath 					/
// @schemes 					http
//
//	@securityDefinitions.apiKey JWT
//	@in                         header
//	@name                       Authorization
//	@description                JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func main() {
	// add swagger !!!!!!!!!!

	Init(connect)

	defer Shutdown()

	server := fiber.New()

	server.Get("/swagger/*", swagger.HandlerDefault) // default

	server.Post("/user", userHandlers.CreateUser)

	server.Post("/session", userHandlers.CreateSession)

	server.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	//server.Get("/user/:userId", userHandlers.GetUser)

	server.Post("/game", authHandlers.CheckAuth, gamesHandlers.CreateGame)

	server.Get("/game/:gameId/board", authHandlers.CheckAuth, gamesHandlers.GetBoard)

	server.Get("/game/:gameId/history", authHandlers.CheckAuth, gamesHandlers.GetHistory)

	server.Post("/game/:gameId/move", authHandlers.CheckAuth, gamesHandlers.DoMove)

	//server.Get("/game/:gameId/board", func(c *fiber.Ctx) error {
	//	clientId := GetClientId(c)
	//	gameId, _ := c.ParamsInt("gameId")
	//
	//	if !game.IsGameMember(clientId) {
	//		server.SendMessage(clientId, "get-board", "Вы не имеете доступа игре")
	//		return
	//	}
	//
	//	board := game.GetBoard()
	//
	//	server.SendMessage(clientId, "get-board", board)
	//})
	//
	//server.Get("/game/:gameId/side", func(c *fiber.Ctx) error {
	//	if !game.IsGameMember(clientId) {
	//		server.SendMessage(clientId, "get-move-side", "Вы не имеете доступа игре")
	//		return
	//	}
	//
	//	server.SendMessage(clientId, "get-move-side", game.GetMoveSide())
	//})
	//
	//server.Get("/game/:gameId/history", func(c *fiber.Ctx) error {
	//	if !game.IsGameMember(clientId) {
	//		server.SendMessage(clientId, "get-history", "Вы не имеете доступа игре")
	//		return
	//	}
	//
	//	server.SendMessage(clientId, "get-history", game.GetHistoryString())
	//})
	//
	//server.Post("/game/:gameId/give-up", func(c *fiber.Ctx) error {
	//	if !game.IsGameMember(clientId) {
	//		server.SendMessage(clientId, "give-up", "Вы не имеете доступа игре")
	//		return
	//	}
	//
	//	game.IsEnded = true
	//
	//	if game.WhiteClientId != nil {
	//		server.SendMessage(*game.WhiteClientId, "give-up", "GAME OVER!")
	//	}
	//	if game.BlackClientId != nil {
	//		server.SendMessage(*game.BlackClientId, "give-up", "GAME OVER!")
	//	}
	//})
	//
	//server.Post("/game/:gameId/move", func(c *fiber.Ctx) error {
	//	if !game.IsGameMember(clientId) {
	//		server.SendMessage(clientId, "do-move", "Вы не имеете доступа игре")
	//		return
	//	}
	//
	//	if game.GetMoveSide() == "white" && clientId != *game.WhiteClientId {
	//		server.SendMessage(clientId, "do-move", "Сейчас ход противника")
	//		return
	//	}
	//
	//	if game.GetMoveSide() == "black" && clientId != *game.BlackClientId {
	//		server.SendMessage(clientId, "do-move", "Сейчас ход противника")
	//		return
	//	}
	//
	//	if !game.CheckCorrectRequest(message) {
	//		server.SendMessage(clientId, "do-move", "Невозможный ход! Введите корректный ход:")
	//		return
	//	}
	//
	//	if !game.DoStep(message) {
	//		server.SendMessage(clientId, "do-move", "Невозможный ход! Введите корректный ход:")
	//		return
	//	}
	//
	//	server.SendMessage(*game.WhiteClientId, "do-move", game.GetBoard())
	//	server.SendMessage(*game.BlackClientId, "do-move", game.GetBoard())
	//})

	if err := server.Listen(":8082"); err != nil {
		log.Fatal(err)
	}

}
