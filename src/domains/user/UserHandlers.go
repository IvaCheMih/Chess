package user

import (
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/services"
	"github.com/IvaCheMih/chess/src/domains/user/dto"
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
// @Param session body dto.CreateSessionRequest true "request"
// @Success 200 {object} dto.CreateSessionResponse
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

	var response dto.CreateSessionResponse

	response.Token, err = token.SignedString([]byte(services.JWT_secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}

// CreateUser godoc
// @Summary create new user.
// @Description create new user.
// @Tags user
// @Accept json
// @Produce json
// @Param  user body  dto.CreateUserRequest true "request"
// @Success 200 {object} dto.CreateUserResponse
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

	return c.JSON(userData)
}
