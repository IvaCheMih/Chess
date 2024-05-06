package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"github.com/IvaCheMih/chess/server/domains/game/services/move_service"
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
		game.Side = userId
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

func (g *GamesRepository) UpdateColorUserIdByColor(gameId int, userColorId string, gameSide int, userId int, tx *gorm.DB) (models.Game, error) {
	var game models.Game
	var res *gorm.DB

	res = tx.Model(&game).Where("id=?", gameId).Updates(map[string]interface{}{userColorId: userId, "side": gameSide, "is_started": true})

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

func (g *GamesRepository) UpdateGame(gameId int, game move_service.Game, tx *gorm.DB) error {
	err := tx.Model(&models.Game{}).Where("id=?", gameId).Updates(map[string]interface{}{
		"is_check_white":        game.IsCheckWhite.IsItCheck,
		"white_king_castling":   game.WhiteCastling.WhiteKingCastling,
		"white_rook_a_castling": game.WhiteCastling.WhiteRookACastling,
		"white_rook_h_castling": game.WhiteCastling.WhiteRookHCastling,
		"is_check_black":        game.IsCheckBlack.IsItCheck,
		"black_king_castling":   game.BlackCastling.BlackKingCastling,
		"black_rook_a_castling": game.BlackCastling.BlackRookACastling,
		"black_rook_h_castling": game.BlackCastling.BlackRookHCastling,
		"last_pawn_move":        game.LastPawnMove,
		"side":                  game.Side,
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
		&requestCreateGame.WhiteKingCastling,

		&requestCreateGame.WhiteRookACastling,
		&requestCreateGame.WhiteRookHCastling,

		&requestCreateGame.IsCheckBlack,
		&requestCreateGame.BlackKingCastling,

		&requestCreateGame.BlackRookACastling,
		&requestCreateGame.BlackRookHCastling,

		&requestCreateGame.LastPawnMove,
		&requestCreateGame.Side,
	)
}
