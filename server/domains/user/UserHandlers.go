package user

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/user/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandlers struct {
	usersService *UsersService
}

func CreateUserHandlers(usersService *UsersService) UserHandlers {
	return UserHandlers{
		usersService: usersService,
	}
}

func (h *UserHandlers) CreateSession(c *fiber.Ctx) error {
	clientId := dto.GetClientId(c)

	fmt.Println(clientId)

	clientPassword, err := dto.GetClientPassword(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !h.usersService.CreateSession(clientId, clientPassword) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	clientPassword, err := dto.GetClientPassword(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userData, err := h.usersService.CreateUser(clientPassword)
	if err != nil {
		return c.SendStatus(fiber.StatusNonAuthoritativeInformation)
	}

	return c.JSON(fiber.Map{"userData": userData})
}
