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

// CreateSession godoc
// @Summary create new session.
// @Description create new session.
// @Tags session
// @Accept json
// @Produce json
// @Param Body body dto.RequestUserIdAndPassword true "request"
// @Success 200 {object} map[string]interface{}
// @Router /session/ [post]
func (h *UserHandlers) CreateSession(c *fiber.Ctx) error {

	request, err := dto.GetIdAndPassword(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !h.usersService.CreateSession(request.Id, request.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": request.Id,
		},
	)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

// CreateUser godoc
// @Summary create new user.
// @Description create new user.
// @Tags user
// @Accept json
// @Produce json
// @Param    Body body  dto.RequestPassword true "request"
// @Success 200 {object} map[string]interface{}
// @Router /user/ [post]
func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	clientPassword, err := dto.GetPassword(c)
	if err != nil || clientPassword.Password == "" {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userData, err := h.usersService.CreateUser(clientPassword.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusNonAuthoritativeInformation)
	}

	return c.JSON(fiber.Map{"userData": userData})
}
