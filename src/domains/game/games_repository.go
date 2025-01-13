package game

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/IvaCheMih/chess/src/domains/game/services/move"
	"github.com/IvaCheMih/chess/src/domains/services/redis"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type GamesRepository struct {
	db    *gorm.DB
	redis redis.RedisService
}

func CreateGamesRepository(db *gorm.DB, redisService redis.RedisService) GamesRepository {
	return GamesRepository{
		redis: redisService,
		db:    db,
	}
}

func (g *GamesRepository) CreateGame(tx *gorm.DB, game models.Game) (models.Game, error) {
	err := tx.Table(`games`).
		Create(&game).
		Error
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (g *GamesRepository) FindNotStartedGame(userColorId string) (models.Game, error) {
	var game models.Game

	err := g.db.Table(`games`).
		Where(map[string]interface{}{userColorId: 0}).
		Where(map[string]interface{}{"status": models.Created}).
		Take(&game).
		Error
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (g *GamesRepository) UpdateColorUserIdByColor(tx *gorm.DB, gameId int, userColorId string, gameSide bool, userId int) error {
	return tx.Table(`games`).
		Where("id=?", gameId).
		Updates(map[string]interface{}{userColorId: userId, "side": gameSide, "status": models.Active}).
		Error
}

func (g *GamesRepository) GetByIdTx(tx *gorm.DB, gameId int) (models.Game, error) {
	var game models.Game

	err := tx.Table(`games`).
		Where("id=?", gameId).
		Take(&game).
		Error
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (g *GamesRepository) GetById(gameId int) (models.Game, error) {
	var game models.Game

	err := g.db.Table(`games`).
		Where("id=?", gameId).
		Take(&game).
		Error
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (g *GamesRepository) UpdateGame(tx *gorm.DB, gameId int, game move.Game, isEnd bool, reason models.EndgameReason, userId int) error {
	var values = map[string]interface{}{
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
	}

	if isEnd {
		values["status"] = models.Ended
		values["end_reason"] = reason

		if reason == models.Mate {
			values["winner_user_id"] = userId
		}
	}

	if game.KilledFigure == 0 {
		values["last_loss"] = game.LastLoss + 1
	} else {
		values["last_loss"] = 0
	}

	return tx.Table(`games`).
		Where("id=?", gameId).
		Updates(values).
		Error
}

func (g *GamesRepository) UpdateIsEnded(winnerUserId int, gameId int, reason models.EndgameReason) error {
	return g.db.Table(`games`).
		Where("id=?", gameId).
		Updates(map[string]interface{}{"status": models.Ended}).
		Updates(map[string]interface{}{"end_reason": reason.ToString()}).
		Updates(map[string]interface{}{"winner_user_id": winnerUserId}).
		Error
}

func (g *GamesRepository) UpdateIsCancelled(gameId int) error {
	return g.db.Table(`games`).
		Where("id=?", gameId).
		Updates(map[string]interface{}{"status": models.Cancelled}).
		Error
}
