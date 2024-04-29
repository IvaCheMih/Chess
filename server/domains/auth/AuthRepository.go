package auth

import (
	"errors"
	"github.com/IvaCheMih/chess/server/domains/user/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func CreateAuthRepository(db *gorm.DB) AuthRepository {
	return AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) GetUserById(clientId any) (models.User, error) {
	var user models.User

	a.db.Take(&user, clientId)

	if user.Password == "" {
		return models.User{}, errors.New("gorm error, password is empty")
	}

	return user, nil
}
