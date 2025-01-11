package user

import (
	"errors"
	"github.com/IvaCheMih/chess/src/domains/user/dto"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"strconv"
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
	if err != nil || password != query.Password {
		return false
	}

	return true
}

func (u *UsersService) CreateUser(telegramId int64, password string) (dto.CreateUserResponse, error) {
	query, err := u.usersRepo.Create(telegramId, password)

	response := dto.CreateUserResponse{
		Id:       query.Id,
		Password: query.Password,
	}

	return response, err
}

func (u *UsersService) TelegramSignIn(telegramId int64, chatId int64) (int, bool, error) {
	res, err := u.usersRepo.GetUserByTelegramId(telegramId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, false, err
	}

	userId := res.Id

	if userId == 0 {
		user, err := u.CreateUser(telegramId, strconv.FormatInt(chatId, 10))
		if err != nil {
			return 0, false, err
		}

		userId = user.Id
	}

	return userId, u.CreateSession(userId, strconv.FormatInt(chatId, 10)), nil
}
