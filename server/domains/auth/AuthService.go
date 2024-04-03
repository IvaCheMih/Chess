package auth

import "fmt"

type AuthService struct {
	AuthRepo *AuthRepository
}

func CreateAuthService(authRepo *AuthRepository) AuthService {
	return AuthService{
		AuthRepo: authRepo,
	}
}

func (a *AuthService) CheckUserId(userId any) error {
	tx, err := a.AuthRepo.db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer tx.Rollback()

	err = a.AuthRepo.FindUserByUserId(userId, tx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tx.Commit()

	return err
}
