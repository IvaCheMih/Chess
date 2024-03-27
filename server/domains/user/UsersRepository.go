package user

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/user/dto"
)

type UsersRepository struct {
	db *sql.DB
}

func CreateUsersRepository(db *sql.DB) UsersRepository {
	return UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) GetClientPassword(clientId int, tx *sql.Tx) (string, error) {
	var password string

	row, err := tx.Query(`
		select password
		from users 
			where id = $1
		`,
		clientId,
	)

	if err != nil {
		return "", err
	}

	err = row.Scan(&password)

	return password, err
}

func (r *UsersRepository) CreateUser(password string, tx *sql.Tx) (dto.CreateUsersResponse, error) {
	var response dto.CreateUsersResponse

	row := tx.QueryRow(`
		insert into users (password)
		    values ($1)
		RETURNING *
		`,
		password,
	)

	err := row.Scan(&response)

	return response, err
}
