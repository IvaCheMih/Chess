package postgresql

import (
	"database/sql"
	"fmt"
)

func ExecUsers(connect string, data Data) {
	db, err := sql.Open("postrges", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, er := db.Exec("insert into users (password) values ($1)", data.User.Password)
	if er != nil {
		panic(err)
	}
}

func DeleteUsers(connect string, id int) {

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from users where id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}

func GetUserById(connect string, id int) (User, error) {
	db, err1 := sql.Open("postgres", connect)
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	rows, err := db.Query("select * from users where ID in($1)", id)

	if err != nil {
		return User{}, nil
	}
	defer rows.Close()

	user := User{}

	for rows.Next() {
		err := rows.Scan(&user)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	defer rows.Close()

	return user, nil
}

func UpdateUsers[T comparable](connect string, data T, id int) {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result sql.Result

	result, err = db.Exec("update users set password = $1 where id = $2", data, id)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк
}
