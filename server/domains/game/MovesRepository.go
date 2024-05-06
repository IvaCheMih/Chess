package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"github.com/IvaCheMih/chess/server/domains/game/services/move_service"
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

	res := m.db.Find(&moves).Where("id=?", gameId)

	if res.Error != nil {
		return []models.Move{}, res.Error
	}

	rows, err := res.Rows()
	if err != nil {
		return []models.Move{}, err
	}

	err = RowsToMove(rows, &moves)

	return moves, nil
}

func (m *MovesRepository) AddMove(gameId, from, to int, board models.Board, isCheckWhite, isCheckBlack move_service.IsCheck, maxNumber int, tx *gorm.DB) (models.Move, error) {
	killedFigureId := 0
	if board.Cells[to] != nil {
		killedFigureId = board.Cells[to].FigureId
	}

	var move = models.Move{
		GameId:         gameId,
		MoveNumber:     maxNumber,
		From_id:        from,
		To_id:          to,
		FigureId:       board.Cells[from].FigureId,
		KilledFigureId: killedFigureId,
		NewFigureId:    0,
		IsCheckWhite:   isCheckWhite.IsItCheck,
		IsCheckBlack:   isCheckBlack.IsItCheck,
	}

	res := tx.Create(&move)

	if res.Error != nil {
		return models.Move{}, res.Error
	}

	err := FromRowToMove(res.Row(), &move)

	return move, err
}

func (m *MovesRepository) FindMaxMoveNumber(gameId int) (int, error) {
	var maxNumber int64

	res := m.db.Model(&models.Move{}).Where("game_id = ?", gameId).Count(&maxNumber)

	if res.Error != nil {
		return 0, res.Error
	}

	return int(maxNumber), res.Error
}

func RowsToMove(rows *sql.Rows, movesOut *[]models.Move) error {
	var moves []models.Move

	for rows.Next() {
		var move models.Move
		err := rows.Scan(
			&move.Id,
			&move.GameId,
			&move.MoveNumber,
			&move.From_id,
			&move.To_id,
			&move.FigureId,
			&move.KilledFigureId,
			&move.NewFigureId,
			&move.IsCheckWhite,
			&move.IsCheckBlack,
		)
		if err != nil {
			return err
		}
		moves = append(moves, move)
	}

	movesOut = &moves

	return nil
}

func FromRowToMove(row *sql.Row, move *models.Move) error {
	err := row.Scan(&move.Id,
		&move.GameId,
		&move.MoveNumber,
		&move.From_id,
		&move.To_id,
		&move.FigureId,
		&move.KilledFigureId,
		&move.NewFigureId,
		&move.IsCheckWhite,
		&move.IsCheckBlack,
	)
	return err
}
