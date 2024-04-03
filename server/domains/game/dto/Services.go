package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) (int, error) {
	headers := c.GetReqHeaders()

	clientIds := headers["X-Client-Id"]
	clientId, err := strconv.ParseInt(clientIds[0], 10, 0)

	return int(clientId), err
}

func GetRequestNewGame(c *fiber.Ctx) (CreateGameRequest, error) {
	body := c.Body()

	var request CreateGameRequest

	err := json.Unmarshal(body, &request)

	return request, err
}

func GetGameId(c *fiber.Ctx) (RequestGetBoard, error) {
	gameId, err := c.ParamsInt("gameId")
	if err != nil {
		return RequestGetBoard{}, err
	}

	var request = RequestGetBoard{
		GameId: gameId,
	}

	return request, nil
}

func GetRequestDoMoveFromBody(c *fiber.Ctx) (DoMoveRequest, error) {
	body := c.Body()

	var request DoMoveRequest

	err := json.Unmarshal(body, &request)

	return request, err
}

//func GetRequestGetBoard(c *fiber.Ctx) (RequestGetBoard, error) {
//	body := c.Body()
//
//	var request RequestGetBoard
//
//	err := json.Unmarshal(body, &request)
//
//	return request, err
//}
