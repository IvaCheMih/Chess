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

func (g *GamesRepository) FindNotStartedGame(userColorId string) (models.Game, error) {
	var game models.Game
	var res *gorm.DB

	res = g.db.Take(&game, map[string]interface{}{userColorId: 0})

	if res.Error != nil {
		return models.Game{}, res.Error
	}

	err := RowToGame(res.Row(), &game)

	return game, err
}

func (g *GamesRepository) UpdateColorUserIdByColor(gameId int, userColorId string, userId int, tx *gorm.DB) (models.Game, error) {
	var game models.Game
	var res *gorm.DB

	res = tx.Model(&game).Where("id=?", gameId).Updates(map[string]interface{}{userColorId: userId, "is_started": true})

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

func (g *GamesRepository) UpdateGame(gameId int, isCheckWhite, isCheckBlack move_service.IsCheck, side int, tx *gorm.DB) error {
	err := g.db.Model(&models.Game{}).Where("id=?", gameId).Updates(map[string]interface{}{
		"is_check_white":  isCheckWhite.IsItCheck,
		"white_king_cell": isCheckWhite.KingGameID,
		"is_check_black":  isCheckBlack.IsItCheck,
		"black_king_cell": isCheckBlack.KingGameID,
		"side":            side,
	}).Error

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
		&requestCreateGame.Id,

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
