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

func GetClientPassword(c *fiber.Ctx) (string, error) {
	body := c.Body()

	var password ResponseUserPassword

	err := json.Unmarshal(body, &password)

	fmt.Println(password)

	return password.Password, err
}
