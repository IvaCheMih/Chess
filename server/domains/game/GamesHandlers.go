package game

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/gofiber/fiber/v2"
)

type GamesHandlers struct {
	gameService *GamesService
}

func CreateGamesHandlers(gameService *GamesService) GamesHandlers {
	return GamesHandlers{
		gameService: gameService,
	}
}

func (h *GamesHandlers) CreateGame(c *fiber.Ctx) error {
	userId, err1 := dto.GetClientId(c)

	userRequestedColor, err2 := dto.GetRequestedColor(c)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	requestCreateGame, err := h.gameService.CreateGame(userId, userRequestedColor)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"gameId": requestCreateGame.GameId})

}

func (h *GamesHandlers) GetBoard(c *fiber.Ctx) error {
	userId, err1 := dto.GetClientId(c)

	gameId, err2 := dto.GetGameId(c)
	if err1 != nil || err2 != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Println(-1)

	responseGetBoard, err := h.gameService.GetBoard(gameId, userId)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"board": responseGetBoard})
}
