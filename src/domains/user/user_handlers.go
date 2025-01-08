package user

import (
	"github.com/IvaCheMih/chess/src/domains/user/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type UserHandlers struct {
	usersService *UsersService
	jwtSecret    string
}

func CreateUserHandlers(usersService *UsersService, jwtSecret string) UserHandlers {
	return UserHandlers{
		usersService: usersService,
		jwtSecret:    jwtSecret,
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

	response.Token, err = token.SignedString([]byte(h.jwtSecret))
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
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userData, err := h.usersService.CreateUser(clientPassword.TelegramId, clientPassword.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusNonAuthoritativeInformation)
	}

	return c.JSON(userData)
}

// TelegramSignIn godoc
// @Summary telegram sign-in.
// @Description telegram sign-in.
// @Tags user
// @Accept json
// @Produce json
// @Param  user body  dto.TelegramSignInRequest true "request"
// @Success 200 {object} dto.TelegramSignInResponse
// @Router /user/sign-in/telegram/ [post]
func (h *UserHandlers) TelegramSignIn(c *fiber.Ctx) error {
	req, err := dto.GetTelegramSignInRequest(c)
	if err != nil || req.TelegramId == 0 || req.ChatId == 0 {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userId, singIn, err := h.usersService.TelegramSignIn(req.TelegramId, req.ChatId)
	if err != nil || !singIn {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
		},
	)

	var response = dto.TelegramSignInResponse{
		AccountId: userId,
	}

	response.Token, err = token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(response)
}
