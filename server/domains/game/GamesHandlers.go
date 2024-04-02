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

// CreateGame godoc
// @Summary create new game.
// @Description create new game.
// @Tags game
// @Accept json
// @Produce json
// @Security       JWT
// @Param Body body dto.RequestedCreateGame true "request"
// @Success 200 {object} map[string]interface{}
// @Router /game/ [post]
func (h *GamesHandlers) CreateGame(c *fiber.Ctx) error {
	request, err := dto.GetRequestNewGame(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	responseCreateGame, err := h.gameService.CreateGame(request.Id, request.IsWhite)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"gameId": responseCreateGame.GameId})

}

// GetBoard godoc
// @Summary get board.
// @Description get board.
// @Tags board
// @Accept json
// @Produce json
// @Security       JWT
// @Param gameId header string true "gameId"
// @Param userId header string true "userId"
// @Success 200 {object} map[string]interface{}
// @Router /game/:gameId/board [get]
func (h *GamesHandlers) GetBoard(c *fiber.Ctx) error {
	request, err := dto.GetRequestGetBoard(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Println(-1)

	responseGetBoard, err := h.gameService.GetBoard(request.GameId, request.UserId)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"board": responseGetBoard})
}
