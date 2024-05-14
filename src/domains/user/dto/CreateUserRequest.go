package dto

type CreateUserRequest struct {
	Password string
}

type CreateUserResponse struct {
	Id       int
	Password string
}
