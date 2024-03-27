package postgresql

import (
	"database/sql"
	"fmt"
)

func ExecGame(connect string, data Data) {
	db, err := sql.Open("postrges", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, er := db.Exec("insert into game (whiteUserId,blackUserId,isStarted,isEnded) values ($1,$2,$3,$4)",
		data.Game.WhiteUserId, data.Game.BlackUserId, data.Game.IsStarted, data.Game.IsEnded)
	if er != nil {
		panic(err)
	}
}

func DeleteGame(connect string, id int) {

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from game where id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}

func QueryGame(connect string, id int) (Game, error) {
	db, err1 := sql.Open("postgres", connect)
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	rows, err := db.Query("select * from game where ID in($1)", id)

	if err != nil {
		return Game{}, err
	}
	defer rows.Close()

	game := Game{}

	for rows.Next() {
		err := rows.Scan(&game)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	defer rows.Close()

	return game, nil
}

func GetLastGame(connect string) (Game, error) {
	db, err1 := sql.Open("postgres", connect)
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	rows, err := db.Query("select * from game order by id desc limit 1")

	if err != nil {
		return Game{}, err
	}
	defer rows.Close()

	game := Game{}

	for rows.Next() {
		err := rows.Scan(&game)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	defer rows.Close()

	return game, nil
}

func UpdateGame[T comparable](connect string, column string, data T, id int) {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result sql.Result

	switch column {
	case "whiteUserId":
		result, err = db.Exec("update game set whiteUserId = $1 where id = $2", data, id)
	case "blackUserId":
		result, err = db.Exec("update game set blackUserId = $1 where id = $2", data, id)
	case "isStarted":

		result, err = db.Exec("update game set isStarted = $1 where id = $2", data, id)
	case "isEnded":
		result, err = db.Exec("update game set isEnded = $1 where id = $2", data, id)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк
}
