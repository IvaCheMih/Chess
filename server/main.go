package main

import (
	_ "github.com/IvaCheMih/chess/server/docs"
	"github.com/IvaCheMih/chess/server/domains"
	"github.com/IvaCheMih/chess/server/domains/auth"
	"github.com/IvaCheMih/chess/server/domains/game"
	"github.com/IvaCheMih/chess/server/domains/game/services/move_service"
	"github.com/IvaCheMih/chess/server/domains/user"
	swagger "github.com/arsmn/fiber-swagger/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

const connect = "postgres://user:pass@localhost:8090/test?sslmode=disable"

var db *gorm.DB

var userHandlers user.UserHandlers
var gamesHandlers game.GamesHandlers
var authHandlers auth.AuthHandlers

func Init() {

	postgresqlUrl, exists := os.LookupEnv("POSTGRES_URL")

	if !exists {
		panic("postgresqlUrl is not found")
	}

	time.Sleep(5 * time.Second)

	//postgresqlUrl := connect

	migrationService := domains.CreateMigrationService()

	migrationService.RunUp(postgresqlUrl, "file://migrations/postgresql")

	move_service.FigureRepo = move_service.CreateFigureRepo()

	var err error

	//db, err = sql.Open("postgres", postgresqlUrl)
	//if err != nil {
	//	panic(err)
	//}

	db, err = gorm.Open(postgres.Open(postgresqlUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
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

// @host 						localhost:8080
// @BasePath 					/
// @schemes 					http
//
//	@securityDefinitions.apiKey JWT
//	@in                         header
//	@name                       Authorization
//	@description                JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func main() {

	Init()

	defer Shutdown()

	server := fiber.New()

	server.Get("/swagger/*", swagger.HandlerDefault) // default

	server.Post("/user", userHandlers.CreateUser)

	server.Post("/session", userHandlers.CreateSession)

	server.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	server.Post("/game", authHandlers.CheckAuth, gamesHandlers.CreateGame)

	server.Get("/game/:gameId/board", authHandlers.CheckAuth, gamesHandlers.GetBoard)

	server.Get("/game/:gameId/history", authHandlers.CheckAuth, gamesHandlers.GetHistory)

	server.Post("/game/:gameId/move", authHandlers.CheckAuth, gamesHandlers.Move)

	server.Post("/game/:gameId/give-up", authHandlers.CheckAuth, gamesHandlers.GiveUp)

	if err := server.Listen(":8080"); err != nil {
		log.Fatal(err)
	}

}
