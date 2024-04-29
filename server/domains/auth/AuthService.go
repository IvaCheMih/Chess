package auth

type AuthService struct {
	AuthRepo *AuthRepository
}

func CreateAuthService(authRepo *AuthRepository) AuthService {
	return AuthService{
		AuthRepo: authRepo,
	}
}

func (a *AuthService) GetUserById(userId any) error {

	_, err := a.AuthRepo.GetUserById(userId)

	return err
}
