package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"github.com/IvaCheMih/chess/server/domains/game/move_service"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
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

func (m *MovesRepository) AddMove(gameId, from, to int, board models.Board, isCheckWhite, isCheckBlack move_service.IsCheck, tx *gorm.DB) (models.Move, error) {
	killedFigureId := 0
	if board.Cells[to] != nil {
		killedFigureId = board.Cells[to].FigureId
	}

	moveNumQueryString := "(SELECT COUNT(*) FROM moves WHERE gameId = "
	moveNumQueryString += strconv.Itoa(gameId)
	moveNumQueryString += ")+1"

	res := tx.Model(&models.Move{}).Create(map[string]interface{}{
		"gameId":         gameId,
		"moveNumber":     clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{moveNumQueryString}},
		"from_id":        from,
		"to_id":          to,
		"figureId":       board.Cells[from].FigureId,
		"killedFigureId": killedFigureId,
		"newFigureId":    0,
		"isCheckWhite":   isCheckWhite.IsItCheck,
		"whiteKingCell":  isCheckWhite.KingGameID,
		"isCheckBlack":   isCheckBlack.IsItCheck,
		"blackKingCell":  isCheckBlack.KingGameID,
	})

	if res.Error != nil {
		return models.Move{}, res.Error
	}

	var move models.Move

	err := FromRowToMove(res.Row(), &move)

	return move, err
}

func RowsToMove(rows *sql.Rows, movesOut *[]models.Move) error {
	var moves []models.Move

	for rows.Next() {
		var move models.Move
		err := rows.Scan(
			&move.Id,
			&move.GameId,
			&move.MoveNumber,
			&move.From,
			&move.To,
			&move.FigureId,
			&move.KilledFigureId,
			&move.NewFigureId,
			&move.IsCheckWhite,
			&move.WhiteKingCell,
			&move.IsCheckBlack,
			&move.BlackKingCell)
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
		&move.From,
		&move.To,
		&move.FigureId,
		&move.KilledFigureId,
		&move.NewFigureId,
		&move.IsCheckWhite,
		&move.WhiteKingCell,
		&move.IsCheckBlack,
		&move.BlackKingCell)
	return err
}
