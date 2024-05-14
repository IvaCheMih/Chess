package user

import (
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/user/dto"
	_ "github.com/lib/pq"
)

type UsersService struct {
	usersRepo *UsersRepository
}

func CreateUsersService(usersRepo *UsersRepository) UsersService {
	return UsersService{
		usersRepo: usersRepo,
	}
}

func (u *UsersService) CreateSession(clientId int, password string) bool {

	query, err := u.usersRepo.GetUserById(clientId)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if err != nil || password != query.Password {
		fmt.Println(err)
		return false
	}

	return true
}

func (u *UsersService) CreateUser(password string) (dto.CreateUserResponse, error) {

	query, err := u.usersRepo.Create(password)

	response := dto.CreateUserResponse{
		Id:       query.Id,
		Password: query.Password,
	}

	return response, err
}
