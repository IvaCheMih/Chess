package dto

type CreateSessionRequest struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type CreateSessionResponse struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}
