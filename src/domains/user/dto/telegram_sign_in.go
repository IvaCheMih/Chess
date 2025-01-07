package dto

type TelegramSignInRequest struct {
	TelegramId int64
	ChatId     int64
}

type TelegramSignInResponse struct {
	Token string
}
