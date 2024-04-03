package auth

import (
	"database/sql"
	"fmt"
)

type AuthRepository struct {
	db *sql.DB
}

func CreateAuthRepository(db *sql.DB) AuthRepository {
	return AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) FindUserByUserId(userId any, tx *sql.Tx) error {
	id := 0
	err := tx.QueryRow(`
		SELECT 1
		FROM users
			WHERE id = $1
		`,
		userId,
	).Scan(&id)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println(err)
		}
		fmt.Println(err)
		return err
	}

	return err
}
