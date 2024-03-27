package postgresql

import (
	"database/sql"
	"fmt"
)

func MaxId(connect string, table string) int {
	db, err := sql.Open("postrges", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var rows *sql.Rows
	var er error

	switch table {
	case "users":
		rows, er = db.Query("SELECT * FROM users WHERE ID = (SELECT MAX(ID) FROM users)")
	case "game":
		rows, er = db.Query("SELECT * FROM game WHERE ID = (SELECT MAX(ID) FROM game)")
	case "figure":
		rows, er = db.Query("SELECT * FROM figure WHERE ID = (SELECT MAX(ID) FROM figure)")
	case "move":
		rows, er = db.Query("SELECT * FROM move WHERE ID = (SELECT MAX(ID) FROM move)")
	case "board":
		rows, er = db.Query("SELECT * FROM board WHERE ID = (SELECT MAX(ID) FROM board)")
	}
	if er != nil {
		panic(err)
	}

	defer rows.Close()

	var field int
	err3 := rows.Scan(&field)
	if err3 != nil {
		fmt.Println(err3)
	}
	return field

}
