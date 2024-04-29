package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) (int, error) {
	headers := c.GetReqHeaders()

	clientIds := headers["X-Client-GameId"]
	clientId, err := strconv.ParseInt(clientIds[0], 10, 0)

	return int(clientId), err
}

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

//func GetRequestGetBoard(c *fiber.Ctx) (GetGameIdParam, error) {
//	body := c.Body()
//
//	var request GetGameIdParam
//
//	err := json.Unmarshal(body, &request)
//
//	return request, err
//}
