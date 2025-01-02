package dto

type CreateSessionRequest struct {
	Id       int
	Password string
}

type CreateSessionResponse struct {
	Token string
}
