package dto

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) int {
	headers := c.GetReqHeaders()
	clientIds := headers["X-Client-Id"]
	clientId, _ := strconv.ParseInt(clientIds[0], 10, 0)
	return int(clientId)
}

func GetClientPassword(c *fiber.Ctx) string {
	headers := c.GetReqHeaders()
	password := headers["X-Client-Password"]
	return password[0]
}
