package dto

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) (int, error) {
	headers := c.GetReqHeaders()

	clientIds := headers["X-Client-Id"]
	clientId, err := strconv.ParseInt(clientIds[0], 10, 0)

	return int(clientId), err
}

func GetRequestNewGame(c *fiber.Ctx) (RequestedCreateGame, error) {
	body := c.Body()

	var request RequestedCreateGame

	err := json.Unmarshal(body, &request)

	return request, err
}

func GetRequestGetBoard(c *fiber.Ctx) (RequestGetBoard, error) {
	headers := c.GetReqHeaders()

	gameIdString := headers["Gameid"]
	gameId, err1 := strconv.ParseInt(gameIdString[0], 10, 0)

	userIdString := headers["Userid"]
	userId, err2 := strconv.ParseInt(userIdString[0], 10, 0)

	if err1 != nil || err2 != nil {
		return RequestGetBoard{}, errors.New("headers error")
	}

	var request = RequestGetBoard{
		GameId: int(gameId),
		UserId: int(userId),
	}

	return request, nil
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
