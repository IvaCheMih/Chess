package user

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/user/dto"
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
	tx, err := u.usersRepo.db.Begin()
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer tx.Rollback()

	passwordFromBase, err := u.usersRepo.GetClientPassword(clientId, tx)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = tx.Commit()
	if err != nil || password != passwordFromBase {
		fmt.Println(err)
		return false
	}

	return true
}

func (u *UsersService) CreateUser(password string) (dto.CreateUserResponse, error) {
	tx, err := u.usersRepo.db.Begin()
	if err != nil {
		fmt.Println(err)
		return dto.CreateUserResponse{}, err
	}

	defer tx.Rollback()

	query, err := u.usersRepo.CreateUser(password, tx)
	if err != nil {
		fmt.Println(err)
		return dto.CreateUserResponse{}, err
	}

	err = tx.Commit()
	if err != nil || query.Password != password {
		fmt.Println(err)
		return dto.CreateUserResponse{}, err
	}

	return query, err
}
