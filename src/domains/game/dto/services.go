package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func GetRequestNewGame(c *fiber.Ctx) (CreateGameBody, error) {
	body := c.Body()

	var request CreateGameBody

	err := json.Unmarshal(body, &request)
	if err != nil {
		return CreateGameBody{}, err
	}

	return request, nil
}

func GetRequestGetGame(c *fiber.Ctx) (GetGameRequest, error) {
	body := c.Body()

	var request GetGameRequest

	err := json.Unmarshal(body, &request)
	if err != nil {
		return GetGameRequest{}, err
	}

	return request, nil
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

func GetEndGameRequest(c *fiber.Ctx) (EndGameRequest, error) {
	body := c.Body()

	var request EndGameRequest

	err := json.Unmarshal(body, &request)
	if err != nil {
		return EndGameRequest{}, err
	}

	return request, nil
}

func GetRequestDoMoveFromBody(c *fiber.Ctx) (DoMoveBody, error) {
	body := c.Body()

	var request DoMoveBody

	err := json.Unmarshal(body, &request)
	if err != nil {
		return DoMoveBody{}, err
	}

	return request, nil
}
