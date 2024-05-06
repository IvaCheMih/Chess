package user

import (
	"errors"
	"github.com/IvaCheMih/chess/server/domains/user/models"
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

func (r *UsersRepository) GetUserById(clientId int) (models.User, error) {
	var user models.User

	r.db.Take(&user, clientId)

	if user.Password == "" {
		return models.User{}, errors.New("gorm error, password is empty")
	}

	return user, nil
}

func (r *UsersRepository) Create(password string) (models.User, error) {
	var response models.User

	var user = models.User{
		Password: password,
	}

	result := r.db.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	err := result.Row().Scan(&response.Id, &response.Password)

	return response, err
}
