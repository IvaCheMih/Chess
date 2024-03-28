package user

import (
	"database/sql"
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/user/dto"
	_ "github.com/lib/pq"
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

	rows, err := tx.Query(`
		select password
		from users 
			where id = $1
		`,
		clientId,
	)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for rows.Next() {
		err = rows.Scan(&password)
	}

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

	err := row.Scan(&response.Id, &response.Password)

	return response, err
}
