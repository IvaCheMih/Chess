package user

import (
	"errors"
	"github.com/IvaCheMih/chess/src/domains/user/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func CreateUsersRepository(dbGorm *gorm.DB) UsersRepository {
	return UsersRepository{
		db: dbGorm,
	}
}

func (r *UsersRepository) GetUserById(accountId int) (models.User, error) {
	var user models.User

	err := r.db.Table(`users`).
		Where("id=?", accountId).
		Take(&user).
		Error
	if err != nil {
		return models.User{}, err
	}

	if user.Password == "" {
		return models.User{}, errors.New("gorm error, password is empty")
	}

	return user, nil
}

func (r *UsersRepository) GetUserByTelegramId(telegramId int64) (models.User, error) {
	var user models.User

	err := r.db.Table(`users`).
		Where("telegram_id=?", telegramId).
		Take(&user).
		Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UsersRepository) Create(telegramId int64, password string) (models.User, error) {
	var user = models.User{
		TelegramId: telegramId,
		Password:   password,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
