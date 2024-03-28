package main

import (
	"database/sql"
	jwtware "github.com/gofiber/contrib/jwt"
)

import (
	"github.com/IvaCheMih/chess/server/domains"
	"github.com/IvaCheMih/chess/server/domains/game"
	"github.com/IvaCheMih/chess/server/domains/user"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

const connect = "postgres://user:pass@postgres:5432/test?sslmode=disable"

var db *sql.DB

var userHandlers user.UserHandlers
var gamesHandlers game.GamesHandlers

func Init(postgresqlUrl string) {
	migrationService := domains.CreateMigrationService()

	time.Sleep(60 * time.Second)

	migrationService.RunUp(postgresqlUrl, "file://migrations/postgresql")

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
}

func Shutdown() {
	db.Close()
}

func main() {
	// add swagger !!!!!!!!!!

	Init(connect)

	defer Shutdown()

	server := fiber.New()

	server.Post("/user", userHandlers.CreateUser)

	server.Post("/session", userHandlers.CreateSession)

	server.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	//server.Get("/user/:userId", userHandlers.GetUser)

	server.Post("/game", gamesHandlers.CreateGame)

	server.Get("/game/:gameId/board", gamesHandlers.GetBoard)

	//http.HandleFunc("/game", handlers.Game)

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

	log.Fatal(server.Listen(":8082"))
	//err := http.ListenAndServe(":8082", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
