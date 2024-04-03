package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandlers struct {
	AuthService *AuthService
}

func CreateAuthHandlers(authService *AuthService) AuthHandlers {
	return AuthHandlers{
		AuthService: authService,
	}
}

func (a *AuthHandlers) CheckAuth(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId64 := claims["userId"].(float64)

	userId := int(userId64)

	//err := a.AuthService.CheckUserId(userId)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}

	c.Context().SetUserValue("userId", userId)

	return c.Next()
}
