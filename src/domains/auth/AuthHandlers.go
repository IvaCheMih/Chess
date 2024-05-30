package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandlers struct {
}

func CreateAuthHandlers() AuthHandlers {
	return AuthHandlers{}
}

func (a *AuthHandlers) Auth(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId64 := claims["userId"].(float64)

	userId := int(userId64)

	c.Context().SetUserValue("userId", userId)

	return c.Next()
}
