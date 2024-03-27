package postgresql

import (
	"database/sql"
	"fmt"
)

func ExecFigure(connect string, data Data) {
	db, err := sql.Open("postrges", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, er := db.Exec("insert into figure (type, isWhite) values ($1,$2)",
		data.Figure.Type, data.Figure.IsWhite)
	if er != nil {
		panic(err)
	}
}

func DeleteFigure(connect string, id int) {

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from figure where id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}

func QueryFigure[T comparable](connect string, id int, column string) T {
	db, err1 := sql.Open("postgres", connect)
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	var rows *sql.Rows
	var err2 error

	switch column {
	case "type":
		rows, err2 = db.Query("select type from figure where ID in($1)", id)
	case "isWhite":
		rows, err2 = db.Query("select isWhite from figure where ID in($1)", id)
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

func UpdateFigure[T comparable](connect string, column string, data T, id int) {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result sql.Result

	switch column {
	case "type":
		result, err = db.Exec("update figure set type = $1 where id = $2", data, id)
	case "isWhite":
		result, err = db.Exec("update figure set isWhite = $1 where id = $2", data, id)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
