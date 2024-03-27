package postgresql

import (
	"database/sql"
	"fmt"
)

func ExecMove(connect string, data Data) {
	db, err := sql.Open("postrges", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, er := db.Exec("insert into move (gameId,moveNumber,fromID,toID, figureId, killedFigureId, newFigureId, isCheckWhite,isCheckBlack) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)",
		data.Move.GameId, data.Move.MoveNumber, data.Move.From, data.Move.To, data.Move.FigureId, data.Move.KilledFigureId, data.Move.NewFigureId, data.Move.IsCheckWhite, data.Move.IsCheckBlack)
	if er != nil {
		panic(err)
	}
}

func DeleteMove(connect string, id int) {

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	q := "delete from move where id = $1"

	result, err := db.Exec(q, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}

func GetMoveById(connect string, id int) (Move, error) {
	db, err1 := sql.Open("postgres", connect)
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	rows, err := db.Query("select * from move where ID in($1)", id)

	if err != nil {
		return Move{}, err
	}
	defer rows.Close()

	move := Move{}

	for rows.Next() {
		//err1 := rows.Scan(&move.GameId, &move.MoveNumber, &move.From, &move.To,&move.FigureId,&move.KilledFigureId,&move.NewFigureId,&move.IsCheckWhite,&move.IsCheckBlack)
		err := rows.Scan(&move)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	defer rows.Close()

	return move, nil
}

func UpdateMove[T comparable](connect string, column string, data T, id int) {

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result sql.Result

	switch column {
	case "gameId":
		result, err = db.Exec("update move set gameId = $1 where id = $2", data, id)
	case "moveNumber":
		result, err = db.Exec("update move set moveNumber = $1 where id = $2", data, id)
	case "from":

		result, err = db.Exec("update move set fromID = $1 where id = $2", data, id)
	case "to":
		result, err = db.Exec("update move set toID = $1 where id = $2", data, id)
	case "figureId":
		result, err = db.Exec("update move set figureId = $1 where id = $2", data, id)
	case "killedFigureId":
		result, err = db.Exec("update move set killedFigureId = $1 where id = $2", data, id)
	case "newFigureId":

		result, err = db.Exec("update move set newFigureId = $1 where id = $2", data, id)
	case "isCheckWhite":
		result, err = db.Exec("update move set isCheckWhite = $1 where id = $2", data, id)
	case "isCheckBlack":
		result, err = db.Exec("update move set isCheckBlack = $1 where id = $2", data, id)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
