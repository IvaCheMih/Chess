package game

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"gorm.io/gorm"
)

type MovesRepository struct {
	db *gorm.DB
}

func CreateMovesRepository(db *gorm.DB) MovesRepository {
	return MovesRepository{
		db: db,
	}
}

func (m *MovesRepository) Find(gameId int) ([]models.Move, error) {
	var moves []models.Move

	err := m.db.Table(`moves`).
		Find(&moves).
		Where("game_id=?", gameId).
		Error
	if err != nil {
		return nil, err
	}

	return moves, nil
}

func (m *MovesRepository) AddMove(tx *gorm.DB, move models.Move) (models.Move, error) {
	err := tx.Table(`moves`).Create(&move).Error
	if err != nil {
		return models.Move{}, err
	}

	return move, nil
}

func (m *MovesRepository) FindMaxMoveNumber(gameId int) (int, error) {
	var maxNumber int64

	err := m.db.Table(`moves`).
		Where("game_id=?", gameId).
		Count(&maxNumber).
		Error
	if err != nil {
		return 0, err
	}

	return int(maxNumber), nil
}
