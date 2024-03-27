package postgresql

import (
	"database/sql"
	"fmt"
)

func ExecBoard(connect string, data Data) {
	db, err := sql.Open("postrges", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, er := db.Exec("insert into board (gameId, indexCell,figureId) values ($1,$2,$3)",
		data.Board.GameId, data.Board.IndexCell, data.Board.FigureId)
	if er != nil {
		panic(err)
	}
}

func DeleteBoard(connect string, id int) {

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from board where id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}

func QueryBoard[T comparable](connect string, id int, column string) T {
	db, err1 := sql.Open("postgres", connect)
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	var rows *sql.Rows
	var err2 error

	switch column {
	case "gameId":
		rows, err2 = db.Query("select gameId from board where ID in($1)", id)
	case "index":
		rows, err2 = db.Query("select index from board where ID in($1)", id)
	case "figureId":
		rows, err2 = db.Query("select figureId from board where ID in($1)", id)
	}

	if err2 != nil {
		panic(err2)
	}
	defer rows.Close()

	var field T
	err3 := rows.Scan(&field)
	if err3 != nil {
		fmt.Println(err3)
	}
	return field
}

func UpdateBoard[T comparable](connect string, column string, data T, id int) {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result sql.Result

	switch column {
	case "gameId":
		result, err = db.Exec("update board set type = $1 where id = $2", data, id)
	case "index":
		result, err = db.Exec("update board set index = $1 where id = $2", data, id)
	case "figureId":
		result, err = db.Exec("update board set figureId = $1 where id = $2", data, id)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
