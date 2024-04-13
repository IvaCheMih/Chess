package dto

type CreateSessionRequest struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type CreateSessionResponse struct {
	Token string `json:"token"`
}
