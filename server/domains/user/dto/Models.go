package dto

type CreateUsersResponse struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type ResponseUserPassword struct {
	Password string `json:"password"`
}
