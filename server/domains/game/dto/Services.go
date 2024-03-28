package dto

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) (int, error) {
	headers := c.GetReqHeaders()

	clientIds := headers["X-Client-Id"]
	clientId, err := strconv.ParseInt(clientIds[0], 10, 0)

	return int(clientId), err
}

func GetRequestedColor(c *fiber.Ctx) (RequestedColor, error) {
	body := c.Body()

	var requestedColor RequestedColor

	err := json.Unmarshal(body, &requestedColor)

	return requestedColor, err
}

func GetGameId(c *fiber.Ctx) (int, error) {
	reqHeaders, err := c.ParamsInt("gameId")

	fmt.Println(reqHeaders)

	//if len(reqHeaders["gameId"]) != 1 {
	//	if gameId, err := strconv.Atoi(reqHeaders["gameId"][0]); gameId != 0 && err == nil {
	//		return gameId, err
	//	}
	//}

	return reqHeaders, err
}
