package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"github.com/IvaCheMih/chess/server/domains/game/move_service"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type GamesRepository struct {
	db *gorm.DB
}

func CreateGamesRepository(db *gorm.DB) GamesRepository {
	return GamesRepository{
		db: db,
	}
}

func (g *GamesRepository) CreateGame(userId int, color bool, tx *gorm.DB) (models.Game, error) {
	var game models.Game
	var err error

	if color {
		game.WhiteUserId = userId
	} else {
		game.BlackUserId = userId
	}

	result := tx.Create(&game)

	if result.Error != nil {
		return models.Game{}, result.Error
	}

	err = RowToGame(result.Row(), &game)

	return game, err
}

//func (g *GamesRepository) FindNotStartedGame(color bool, tx *sql.Tx) (models.Game, error) {
//	var row *sql.Row
//
//	if !color {
//		row = tx.QueryRow(`
//		SELECT * FROM games
//		    where blackUserId = 0
//		    LIMIT 1
//		`,
//		)
//	} else {
//		row = tx.QueryRow(`
//		SELECT * FROM games
//		    where whiteUserId = 0
//		    LIMIT 1
//		`,
//		)
//	}
//
//	if row.Err() != nil && row.Err().Error() == "sql: no rows in result set" {
//		return models.Game{}, row.Err()
//	}
//
//	var requestCreateGame models.Game
//
//	err := RowToGame(row, &requestCreateGame)
//
//	return requestCreateGame, err
//
//}

func (g *GamesRepository) FindNotStartedGame(color bool) (models.Game, error) {
	var game models.Game
	var res *gorm.DB

	if color {
		res = g.db.Take(&game, models.Game{WhiteUserId: 0})
	} else {
		res = g.db.Take(&game, models.Game{BlackUserId: 0})
	}

	if res.Error != nil {
		return models.Game{}, res.Error
	}

	err := RowToGame(res.Row(), &game)

	return game, err
}

func (g *GamesRepository) UpdateColorUserIdByColor(gameId int, color bool, userId int, tx *gorm.DB) (models.Game, error) {
	var game models.Game
	var res *gorm.DB
	userColor := "whiteUserId"

	if !color {
		userColor = "blackUserId"
	}

	res = tx.Model(&game).Where("id=?", gameId).Updates(map[string]interface{}{userColor: userId, "isStarted": true})

	if res.Error != nil {
		return models.Game{}, res.Error
	}

	err := RowToGame(res.Row(), &game)

	return game, err
}

func (g *GamesRepository) GetById(gameId int) (models.Game, error) {
	var game models.Game

	res := g.db.Take(&game, gameId)
	if res.Error != nil {
		return models.Game{}, res.Error
	}

	err := RowToGame(res.Row(), &game)

	return game, err
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

func (g *GamesRepository) UpdateIsEnded(gameId int) (models.Game, error) {

	res := g.db.Model(&models.Game{}).Where("id=?", gameId).Updates(map[string]interface{}{"isEnded": true})
	if res.Error != nil {
		return models.Game{}, res.Error
	}

	var modelGame models.Game

	err := RowToGame(res.Row(), &modelGame)

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
