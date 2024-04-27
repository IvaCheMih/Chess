package game

import (
	"database/sql"
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/models"
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

func (g *GamesRepository) CreateGame(userId any, color bool, tx *sql.Tx) (models.Game, error) {
	var row *sql.Row

	if color {
		row = tx.QueryRow(`
		insert into games (whiteUserId)
			values ($1)
			RETURNING *
		`,
			userId,
		)
	} else {
		row = tx.QueryRow(`
		insert into games (blackUserId)
			values ($1)
			RETURNING *
		`,
			userId,
		)
	}

	var requestCreateGame models.Game

	err := RowToGame(row, &requestCreateGame)

	return requestCreateGame, err
}

func (g *GamesRepository) FindNotStartedGame(color bool, tx *sql.Tx) (models.Game, error) {
	var row *sql.Row

	if !color {
		row = tx.QueryRow(`
		SELECT * FROM games
		    where blackUserId = 0
		    LIMIT 1 
		`,
		)
	} else {
		row = tx.QueryRow(`
		SELECT * FROM games
		    where whiteUserId = 0
		    LIMIT 1 
		`,
		)
	}

	fmt.Println(201)

	if row.Err() != nil && row.Err().Error() == "sql: no rows in result set" {
		return models.Game{}, row.Err()
	}

	fmt.Println(202)

	var requestCreateGame models.Game

	err := RowToGame(row, &requestCreateGame)

	return requestCreateGame, err

}

func (g *GamesRepository) JoinToGame(gameId any, color bool, userId any, tx *sql.Tx) (models.Game, error) {
	var row *sql.Row

	if !color {
		row = tx.QueryRow(`
		update games
		set blackUserId = $1, isStarted = true
			where id = $2
		RETURNING *
		`,
			userId,
			gameId,
		)
	} else {
		row = tx.QueryRow(`
		update games
		set whiteUserId = $1, isStarted = true
			where id = $2
		RETURNING *
		`,
			userId,
			gameId,
		)
	}

	if row.Err() != nil {
		return models.Game{}, row.Err()
	}

	var game models.Game

	err := RowToGame(row, &game)

	return game, err
}

func (g *GamesRepository) GetById(gameId int) (models.Game, error) {
	row := g.db.QueryRow(`
		SELECT * FROM games
		    where id = $1
		`,
		gameId,
	)

	var requestCreateGame models.Game

	err := RowToGame(row, &requestCreateGame)

	return requestCreateGame, err
}

func (g *GamesRepository) UpdateGame(gameId int, isCheckWhite, isCheckBlack move_service.IsCheck, side int, tx *sql.Tx) error {
	_, err := tx.Exec(`
		update games
		set isCheckWhite = $1, whiteKingCell = $2,isCheckBlack = $3,blackKingCell=$4, side =$5 
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

func (g *GamesRepository) Update(game models.Game, tx *sql.Tx) (models.Game, error) {
	row := tx.QueryRow(`
		update games
		set whiteUserId = $1, blackUserId =$2,isStarted =$3, isEnded= $4, isCheckWhite = $5, whiteKingCell = $6,isCheckBlack = $7,blackKingCell=$8, side =$9 
			where id = $10
		RETURNING *
		`,
		game.WhiteUserId,
		game.BlackUserId,
		game.IsStarted,
		game.IsEnded,
		game.IsCheckWhite,
		game.WhiteKingCell,
		game.IsCheckBlack,
		game.BlackKingCell,
		game.Side,
		game.GameId,
	)

	var modelGame models.Game

	err := RowToGame(row, &modelGame)

	return modelGame, err
}

func RowToGame(row *sql.Row, requestCreateGame *models.Game) error {
	return row.Scan(
		&requestCreateGame.GameId,

		&requestCreateGame.WhiteUserId,
		&requestCreateGame.BlackUserId,

		&requestCreateGame.IsStarted,
		&requestCreateGame.IsEnded,

		&requestCreateGame.IsCheckWhite,
		&requestCreateGame.WhiteKingCell,

		&requestCreateGame.IsCheckBlack,
		&requestCreateGame.BlackKingCell,

		&requestCreateGame.Side,
	)
}
