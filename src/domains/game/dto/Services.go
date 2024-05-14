package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func GetRequestNewGame(c *fiber.Ctx) (CreateGameBody, error) {
	body := c.Body()

	var request CreateGameBody

	err := json.Unmarshal(body, &request)

	return request, err
}

func GetGameId(c *fiber.Ctx) (GetGameIdParam, error) {
	gameId, err := c.ParamsInt("gameId")
	if err != nil {
		return GetGameIdParam{}, err
	}

	var request = GetGameIdParam{
		GameId: gameId,
	}

	return request, nil
}

func GetRequestDoMoveFromBody(c *fiber.Ctx) (DoMoveBody, error) {
	body := c.Body()

	var request DoMoveBody

	err := json.Unmarshal(body, &request)

	return request, err
}
