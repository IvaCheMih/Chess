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

	clientIds := headers["X-Client-Id"]
	clientId, _ := strconv.ParseInt(clientIds[0], 10, 0)
	return int(clientId)
}

func GetIdAndPassword(c *fiber.Ctx) (RequestUserIdAndPassword, error) {
	body := c.Body()

	var request RequestUserIdAndPassword

	err := json.Unmarshal(body, &request)

	return request, err
}

func GetPassword(c *fiber.Ctx) (RequestPassword, error) {
	body := c.Body()

	var request RequestPassword

	err := json.Unmarshal(body, &request)

	return request, err
}
