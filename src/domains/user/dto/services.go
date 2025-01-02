package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

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
