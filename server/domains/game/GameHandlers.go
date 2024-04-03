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

	userId := c.Context().Value("userId")

	fmt.Println(userId)

	responseCreateGame, err := h.gameService.CreateGame(userId, request.IsWhite)
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
// @Param gameId path int true "gameId"
// @Success 200 {object} dto.GetBoardResponse
// @Router /game/{gameId}/board [get]
func (h *GamesHandlers) GetBoard(c *fiber.Ctx) error {
	request, err := dto.GetGameId(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId := c.Context().Value("userId")

	getBoardResponse, err := h.gameService.GetBoard(request.GameId, userId)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(getBoardResponse)
}

// GetHistory godoc
// @Summary get history.
// @Description get history.
// @Tags history
// @Accept json
// @Produce json
// @Security       JWT
// @Param gameId header string true "gameId"
// @Param userId header string true "userId"
// @Success 200 {object} map[string]interface{}
// @Router /game/:gameId/history [get]
func (h *GamesHandlers) GetHistory(c *fiber.Ctx) error {
	request, err := dto.GetGameId(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId := c.Context().Value("userId")

	responseGetHistory, err := h.gameService.GetHistory(request.GameId, userId)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"history": responseGetHistory})
}

// DoMove godoc
// @Summary do move.
// @Description do move.
// @Tags move
// @Accept json
// @Produce json
// @Security       JWT
// @Param gameId header string true "gameId"
// @Param userId header string true "userId"
// @Param move body dto.RequestDoMove true "move"
// @Success 200 {object} map[string]interface{}
// @Router /game/:gameId/move [post]
func (h *GamesHandlers) DoMove(c *fiber.Ctx) error {
	request, err := dto.GetGameId(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	requestDoMove, err := dto.GetRequestDoMoveFromBody(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId := c.Context().Value("userId")

	responseDoMove, err := h.gameService.DoMove(request.GameId, userId, requestDoMove)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"board": responseDoMove})
}
