package dto

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) int {
	headers := c.GetReqHeaders()
	fmt.Println(headers)

	clientIds := headers["X-Client-GameId"]
	clientId, _ := strconv.ParseInt(clientIds[0], 10, 0)
	return int(clientId)
}

func GetIdAndPassword(c *fiber.Ctx) (CreateSessionRequest, error) {
	body := c.Body()

	var request CreateSessionRequest

	err := json.Unmarshal(body, &request)

	return request, err
}

func GetPassword(c *fiber.Ctx) (CreateUserRequest, error) {
	body := c.Body()

	var request CreateUserRequest

	err := json.Unmarshal(body, &request)

	return request, err
}
