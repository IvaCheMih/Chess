package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/IvaCheMih/chess/src/domains/game/services/move_service"
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

	if color {
		game.WhiteUserId = userId
	} else {
		game.BlackUserId = userId
	}

	result := tx.Create(&game)

	if result.Error != nil {
		return models.Game{}, result.Error
	}

	return RowToGame(result.Row())
}

func (g *GamesRepository) FindNotStartedGame(userColorId string) (models.Game, error) {
	var game models.Game
	var res *gorm.DB

	res = g.db.Take(&game, map[string]interface{}{userColorId: 0})

	if res.Error != nil {
		return models.Game{}, res.Error
	}

	return RowToGame(res.Row())
}

func (g *GamesRepository) UpdateColorUserIdByColor(gameId int, userColorId string, gameSide bool, userId int, tx *gorm.DB) (models.Game, error) {
	var game models.Game
	var res *gorm.DB

	res = tx.Model(&game).Where("id=?", gameId).Updates(map[string]interface{}{userColorId: userId, "side": gameSide, "is_started": true})

	if res.Error != nil {
		return models.Game{}, res.Error
	}

	return RowToGame(res.Row())

}

func (g *GamesRepository) GetById(gameId int) (models.Game, error) {
	var game models.Game

	res := g.db.Take(&game, gameId)
	if res.Error != nil {
		return models.Game{}, res.Error
	}

	return RowToGame(res.Row())
}

func (g *GamesRepository) UpdateGame(gameId int, game move_service.Game, tx *gorm.DB) error {
	err := tx.Model(&models.Game{}).Where("id=?", gameId).Updates(map[string]interface{}{
		"is_check_white":        game.IsCheckWhite.IsItCheck,
		"white_king_castling":   game.WhiteCastling.KingCastling,
		"white_rook_a_castling": game.WhiteCastling.RookACastling,
		"white_rook_h_castling": game.WhiteCastling.RookHCastling,
		"is_check_black":        game.IsCheckBlack.IsItCheck,
		"black_king_castling":   game.BlackCastling.KingCastling,
		"black_rook_a_castling": game.BlackCastling.RookACastling,
		"black_rook_h_castling": game.BlackCastling.RookHCastling,
		"last_pawn_move":        game.LastPawnMove,
		"side":                  game.Side,
	}).Error

	return err
}

func (g *GamesRepository) UpdateIsEnded(gameId int) (models.Game, error) {
	res := g.db.Model(&models.Game{}).Where("id=?", gameId).Updates(map[string]interface{}{"is_ended": true})
	if res.Error != nil {
		return models.Game{}, res.Error
	}

	return RowToGame(res.Row())
}

func RowToGame(row *sql.Row) (models.Game, error) {
	var requestCreateGame models.Game

	err := row.Scan(
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

	return requestCreateGame, err
}
