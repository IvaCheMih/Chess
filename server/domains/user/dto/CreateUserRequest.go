package dto

type CreateUserRequest struct {
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}
