package dto

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) (int, error) {
	clientIdByte := c.Context().FormValue("clientId")

	var clientId int

	err := json.Unmarshal(clientIdByte, &clientId)

	return clientId, err
}

func GetRequestedColor(c *fiber.Ctx) (RequestedColor, error) {
	body := c.Body()

	var requestedColor RequestedColor

	err := json.Unmarshal(body, &requestedColor)

	return requestedColor, err
}

func GetGameId(c *fiber.Ctx) (int, error) {
	reqHeaders := c.GetReqHeaders()

	if len(reqHeaders["gameId"]) != 1 {
		if gameId, err := strconv.Atoi(reqHeaders["gameId"][0]); gameId != 0 && err != nil {
			return gameId, err
		}
	}

	return 0, errors.New("bad request")
}
