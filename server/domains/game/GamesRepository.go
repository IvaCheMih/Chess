package game

import (
	"database/sql"
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

func (g *GamesRepository) CreateGame(userId any, tx *sql.Tx) (models.CreateGameResponse, error) {
	row := tx.QueryRow(`
		insert into games (whiteUserId)
			values ($1)
			RETURNING *
		`,
		userId,
	)

	var requestCreateGame models.CreateGameResponse

	err := RowToGame(row, &requestCreateGame)

	return requestCreateGame, err
}

func (g *GamesRepository) FindNotStartedGame(tx *sql.Tx) (models.CreateGameResponse, error) {
	row := tx.QueryRow(`
		SELECT * FROM games
		    where blackUserId = 0
		    LIMIT 1 
		`,
	)

	var requestCreateGame models.CreateGameResponse

	err := RowToGame(row, &requestCreateGame)

	return requestCreateGame, err

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

func (g *GamesRepository) GetById(gameId int, tx *sql.Tx) (models.CreateGameResponse, error) {
	row := tx.QueryRow(`
		SELECT * FROM games
		    where id = $1
		`,
		gameId,
	)

	var requestCreateGame models.CreateGameResponse

	err := RowToGame(row, &requestCreateGame)

	return requestCreateGame, err
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

func RowToGame(row *sql.Row, requestCreateGame *models.CreateGameResponse) error {
	return row.Scan(&requestCreateGame.GameId, &requestCreateGame.WhiteUserId, &requestCreateGame.BlackUserId, &requestCreateGame.IsStarted, &requestCreateGame.IsEnded)
}
