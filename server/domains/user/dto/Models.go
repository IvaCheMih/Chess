package dto

type CreateUsersResponse struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type RequestUserIdAndPassword struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type ResponseUserIdAndPassword struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type RequestPassword struct {
	Password string `json:"password"`
}
