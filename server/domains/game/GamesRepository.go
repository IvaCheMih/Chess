package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/move_service"
	_ "github.com/lib/pq"
)

type GamesRepository struct {
	db *sql.DB
}

func CreateGamesRepository(db *sql.DB) GamesRepository {
	return GamesRepository{
		db: db,
	}
}

func (g *GamesRepository) CreateGame(userId any, tx *sql.Tx) (dto.ResponseGetGame, error) {
	row := tx.QueryRow(`
		insert into games (whiteUserId)
			values ($1)
			RETURNING *
		`,
		userId,
	)

	var requestCreateGame dto.ResponseGetGame

	err := RowToGame(row, &requestCreateGame)

	return requestCreateGame, err
}

func (g *GamesRepository) FindNotStartedGame(tx *sql.Tx) (dto.ResponseGetGame, error) {
	resultQuery, err := tx.Query(`
		SELECT * FROM games
		    where blackUserId = 0
		    LIMIT 1 
		`,
	)

	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	var responseCreateGame dto.ResponseGetGame

	for resultQuery.Next() {
		err = resultQuery.Scan(
			&responseCreateGame.GameId,
			&responseCreateGame.WhiteUserId,
			&responseCreateGame.BlackUserId,
			&responseCreateGame.IsStarted,
			&responseCreateGame.IsEnded,
			&responseCreateGame.IsCheckWhite,
			&responseCreateGame.WhiteKingCell,
			&responseCreateGame.IsCheckWhite,
			&responseCreateGame.BlackKingCell,
			&responseCreateGame.Side,
		)
		if err != nil {
			return dto.ResponseGetGame{}, err
		}
	}

	return responseCreateGame, nil

}

func (g *GamesRepository) JoinBlackToGame(gameId any, userId any, tx *sql.Tx) error {
	_, err := tx.Exec(`
		update games
		set blackUserId = $1, isStarted = true
			where id = $2
		`,
		userId,
		gameId,
	)

	return err
}

func (g *GamesRepository) GetById(gameId int, tx *sql.Tx) (dto.ResponseGetGame, error) {
	resultQuery, err := tx.Query(`
		SELECT * FROM games
		    where id = $1
		`,
		gameId,
	)

	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	var responseCreateGame dto.ResponseGetGame

	for resultQuery.Next() {
		err = resultQuery.Scan(
			&responseCreateGame.GameId,
			&responseCreateGame.WhiteUserId,
			&responseCreateGame.BlackUserId,
			&responseCreateGame.IsStarted,
			&responseCreateGame.IsEnded,
			&responseCreateGame.IsCheckWhite,
			&responseCreateGame.WhiteKingCell,
			&responseCreateGame.IsCheckWhite,
			&responseCreateGame.BlackKingCell,
			&responseCreateGame.Side,
		)
		if err != nil {
			return dto.ResponseGetGame{}, err
		}
	}

	return responseCreateGame, nil
}

func (g *GamesRepository) UpdateGame(gameId int, isCheckWhite, isCheckBlack move_service.IsCheck, side int, tx *sql.Tx) error {
	_, err := tx.Exec(`
		update games
		set (isCheckWhite, whiteKingCell ,isCheckBlack ,blackKingCell, side)
		 	values ($1, $2, $3, $4, $5)
			where id = $6
		`,
		isCheckWhite.IsItCheck,
		isCheckWhite.KingGameID,
		isCheckBlack.IsItCheck,
		isCheckBlack.KingGameID,
		side,
		gameId,
	)

	return err
}

func RowToGame(row *sql.Row, requestCreateGame *dto.ResponseGetGame) error {
	return row.Scan(&requestCreateGame.GameId, &requestCreateGame.WhiteUserId, &requestCreateGame.BlackUserId, &requestCreateGame.IsStarted, &requestCreateGame.IsEnded)
}
