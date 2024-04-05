package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
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

func (m *MovesRepository) Find(gameId int, tx *sql.Tx) ([]dto.Move, error) {
	rows, err := tx.Query(`
		SELECT * FROM moves
		    where gameId = $1 ORDER BY moveNumber
		`,
		gameId,
	)

	if err != nil {
		return []dto.Move{}, err
	}

	var moves []dto.Move

	err = FromRowsToMove(rows, &moves)

	return moves, nil
}

func (m *MovesRepository) AddMove(gameId, from, to int, board models.Board, isCheckWhite, isCheckBlack move_service.IsCheck, tx *sql.Tx) (dto.Move, error) {

	killedFigureId := 0
	if board.Cells[to] != nil {
		killedFigureId = board.Cells[to].FigureId
	}

	rows, err := tx.Query(`
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

	defer rows.Close()

	if err != nil {
		return dto.Move{}, err
	}

	var moves []dto.Move

	err = FromRowsToMove(rows, &moves)

	move := moves[0]

	return move, err
}

func FromRowsToMove(rows *sql.Rows, movesOut *[]dto.Move) error {
	var moves []dto.Move

	for rows.Next() {
		var move dto.Move
		err := rows.Scan(&move.Id, &move.GameId, &move.MoveNumber, &move.From, &move.To, &move.FigureId, &move.KilledFigureId, &move.NewFigureId, &move.IsCheckWhite, &move.WhiteKingCell, &move.IsCheckBlack, &move.BlackKingCell)
		if err != nil {
			return err
		}
		moves = append(moves, move)
	}

	movesOut = &moves

	return nil
}
