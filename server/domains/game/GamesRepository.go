package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
)

type GamesRepository struct {
	db *sql.DB
}

func CreateGamesRepository(db *sql.DB) GamesRepository {
	return GamesRepository{
		db: db,
	}
}

func (g *GamesRepository) CreateGame(userId int, tx *sql.Tx) (dto.ResponseGetGame, error) {
	row := tx.QueryRow(`
		insert into games (whiteUserId)
			values ($1)
			RETURNING *
		`,
		userId,
	)

	var requestCreateGame dto.ResponseGetGame

	err := row.Scan(&requestCreateGame)

	return requestCreateGame, err
}

func (g *GamesRepository) FindNotStartedGame(tx *sql.Tx) (dto.ResponseGetGame, error) {
	resultQuery, err := tx.Query(`
		SELECT * FROM games
		    where blackUserId = 0
		    desc limit 1
		`,
	)

	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	var responseCreateGame dto.ResponseGetGame

	for resultQuery.Next() {
		err = resultQuery.Scan(&responseCreateGame)
		if err != nil {
			return dto.ResponseGetGame{}, err
		}
	}

	return responseCreateGame, nil

}

func (g *GamesRepository) JoinBlackToGame(gameId int, userId int, tx *sql.Tx) error {
	_, err := tx.Exec(`
		update games
		set blackUserId = $1
			where id = $2`,
		userId,
		gameId,
	)

	return err
}

func (g *GamesRepository) GetGame(gameId int, tx *sql.Tx) (dto.ResponseGetGame, error) {
	resultQuery, err := tx.Query(`
		SELECT * FROM games
		    where gameId = $1
		`,
		gameId,
	)
	var responseCreateGame dto.ResponseGetGame

	for resultQuery.Next() {
		err = resultQuery.Scan(&responseCreateGame)
		if err != nil {
			return dto.ResponseGetGame{}, err
		}
	}

	return responseCreateGame, nil
}
