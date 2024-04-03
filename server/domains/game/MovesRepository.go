package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
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
	resultQuery, err := tx.Query(`
		SELECT * FROM moves
		    where gameId = $1 ORDER BY moveNumber
		`,
		gameId,
	)

	if err != nil {
		return []dto.Move{}, err
	}

	var moves []dto.Move

	for resultQuery.Next() {
		var move dto.Move
		err = resultQuery.Scan(&move.Id, &move.GameId, &move.MoveNumber, &move.From, &move.To, &move.FigureId, &move.KilledFigureId, &move.NewFigureId, &move.IsCheckWhite, &move.WhiteKingCell, &move.IsCheckBlack, &move.BlackKingCell)
		if err != nil {
			return []dto.Move{}, err
		}
		moves = append(moves, move)
	}

	return moves, nil
}

func (m *MovesRepository) AddMove(gameId, from, to, figureId, killedFigureId int, isCheckWhite, isCheckBlack move_service.IsCheck, tx *sql.Tx) error {

	err := tx.QueryRow(`
		insert into moves (gameId, moveNumber,from_id,to_id,figureId, killedFigureId, isCheckWhite , whiteKingCell, isCheckBlack, blackKingCell)
			values ($1, (SELECT MAX(moveNumber)+1 FROM moves),  $2, $3, $4, $5, $6, $7, $8,$9)
			RETURNING *
		`,
		gameId,
		from,
		to,
		figureId,
		killedFigureId,
		isCheckWhite.IsItCheck,
		isCheckWhite.KingGameID,
		isCheckBlack.IsItCheck,
		isCheckBlack.KingGameID,
	).Err()

	return err
}
