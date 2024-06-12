package game

import (
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/game/dto"
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
// @Param game body dto.CreateGameBody true "request"
// @Success 200 {object} dto.CreateGameResponse
// @Router /game/ [post]
func (h *GamesHandlers) CreateGame(c *fiber.Ctx) error {
	request, err := dto.GetRequestNewGame(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId := c.Context().Value("userId")

	responseCreateGame, err := h.gameService.CreateGame(userId.(int), request.IsWhite)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(responseCreateGame)

}

// GetBoard godoc
// @Summary get board.
// @Description get board.
// @Tags board
// @Accept json
// @Produce json
// @Security       JWT
// @Param gameId path dto.GetGameIdParam true "gameId"
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
		fmt.Println(err)
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
// @Param gameId path dto.GetGameIdParam true "gameId"
// @Success 200 {object}  dto.GetHistoryResponse
// @Router /game/{gameId}/history [get]
func (h *GamesHandlers) GetHistory(c *fiber.Ctx) error {
	request, err := dto.GetGameId(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId := c.Context().Value("userId")

	responseGetHistory, err := h.gameService.GetHistory(request.GameId, userId)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(responseGetHistory)
}

// Move godoc
// @Summary do move.
// @Description do move.
// @Tags move
// @Accept json
// @Produce json
// @Security       JWT
// @Param gameId path dto.GetGameIdParam true "gameId"
// @Param move body dto.DoMoveBody true "move"
// @Success 200 {object}  models.Move
// @Router /game/{gameId}/move [post]
func (h *GamesHandlers) Move(c *fiber.Ctx) error {
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

	responseMove, err := h.gameService.Move(request.GameId, userId, requestDoMove)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(responseMove)
}

// GiveUp godoc
// @Summary do give-up.
// @Description do give-up.
// @Tags give-up
// @Accept json
// @Produce json
// @Security       JWT
// @Param gameId path dto.GetGameIdParam true "gameId"
// @Success 200 {object}  models.Game
// @Router /game/{gameId}/give-up [post]
func (h *GamesHandlers) GiveUp(c *fiber.Ctx) error {
	request, err := dto.GetGameId(c)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	//userId := c.Context().Value("userId")

	responseMove, err := h.gameService.GiveUp(request.GameId)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(responseMove)
}
