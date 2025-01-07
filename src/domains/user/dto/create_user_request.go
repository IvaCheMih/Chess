package dto

type CreateUserRequest struct {
	TelegramId int64  `json:"telegramId"`
	Password   string `json:"password"`
}

type CreateUserResponse struct {
	Id       int
	Password string
}
