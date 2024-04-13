package game

import (
	"database/sql"
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"github.com/IvaCheMih/chess/server/domains/game/move_service"
)

type MovesRepository struct {
	db *sql.DB
}

func CreateMovesRepository(db *sql.DB) MovesRepository {
	return MovesRepository{
		db: db,
	}
}

func (m *MovesRepository) Find(gameId int, tx *sql.Tx) ([]models.Move, error) {
	rows, err := tx.Query(`
		SELECT * FROM moves
		    where gameId = $1 ORDER BY moveNumber
		`,
		gameId,
	)

	if err != nil {
		return []models.Move{}, err
	}

	var moves []models.Move

	err = FromRowsToMove(rows, &moves)

	return moves, nil
}

func (m *MovesRepository) AddMove(gameId, from, to int, board models.Board, isCheckWhite, isCheckBlack move_service.IsCheck, tx *sql.Tx) (models.Move, error) {

	killedFigureId := 0
	if board.Cells[to] != nil {
		killedFigureId = board.Cells[to].FigureId
	}

	row := tx.QueryRow(`
		insert into moves (gameId,moveNumber, from_id,to_id,figureId, killedFigureId, newFigureId, isCheckWhite , whiteKingCell, isCheckBlack, blackKingCell)
			values ($1,(SELECT COUNT(*) FROM moves WHERE gameId = $2)+1, $3, $4, $5, $6, $7, $8,$9,$10,$11)
		RETURNING *
		`,
		gameId,
		gameId,
		from,
		to,
		board.Cells[from].FigureId,
		killedFigureId,
		0,
		isCheckWhite.IsItCheck,
		isCheckWhite.KingGameID,
		isCheckBlack.IsItCheck,
		isCheckBlack.KingGameID,
	)

	var move models.Move

	err := FromRowToMove(row, &move)

	fmt.Println(move)

	return move, err
}

func FromRowsToMove(rows *sql.Rows, movesOut *[]models.Move) error {
	var moves []models.Move

	for rows.Next() {
		var move models.Move
		err := rows.Scan(&move.Id, &move.GameId, &move.MoveNumber, &move.From, &move.To, &move.FigureId, &move.KilledFigureId, &move.NewFigureId, &move.IsCheckWhite, &move.WhiteKingCell, &move.IsCheckBlack, &move.BlackKingCell)
		if err != nil {
			return err
		}
		moves = append(moves, move)
	}

	movesOut = &moves

	return nil
}

func FromRowToMove(row *sql.Row, move *models.Move) error {
	err := row.Scan(&move.Id, &move.GameId, &move.MoveNumber, &move.From, &move.To, &move.FigureId, &move.KilledFigureId, &move.NewFigureId, &move.IsCheckWhite, &move.WhiteKingCell, &move.IsCheckBlack, &move.BlackKingCell)
	return err
}
