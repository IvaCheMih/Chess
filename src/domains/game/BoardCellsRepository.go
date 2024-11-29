package game

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type BoardCellsRepository struct {
	db *gorm.DB
}

func CreateBoardCellsRepository(db *gorm.DB) BoardCellsRepository {
	return BoardCellsRepository{
		db: db,
	}
}

func (b *BoardCellsRepository) CreateNewBoardCells(gameId int, tx *gorm.DB) error {
	var boardCells = make([]models.BoardCell, len(startField))

	for i, cell := range startField {
		boardCells[i] = models.BoardCell{
			GameId:    gameId,
			IndexCell: cell[0],
			FigureId:  cell[1],
		}
	}

	return tx.Table(`board_cells`).
		Create(boardCells).
		Error
}

func (b *BoardCellsRepository) Find(gameId int) (models.Board, error) {
	var boardCell = []models.BoardCell{}

	err := b.db.Table(`board_cells`).
		Find(&boardCell).
		Where("game_id=?", gameId).
		Error
	if err != nil {
		return models.Board{}, err
	}

	res := b.db.Find(&boardCell).Where("game_id=?", gameId)
	if res.Error != nil {
		return models.Board{}, res.Error
	}

	var cells = map[int]*models.BoardCell{}

	for i := range boardCell {
		cells[boardCell[i].IndexCell] = &boardCell[i]
	}

	return models.Board{Cells: cells}, err
}

func (b *BoardCellsRepository) Update(id int, to int, tx *gorm.DB) error {
	return tx.Table(`board_cells`).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"index_cell": to,
		}).
		Error
}

func (b *BoardCellsRepository) UpdateNewFigure(id int, to int, newFigureId int, tx *gorm.DB) error {
	return tx.Table(`board_cells`).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"index_cell": to, "figure_id": newFigureId,
		}).
		Error
}

func (b *BoardCellsRepository) Delete(id int, tx *gorm.DB) error {
	return tx.Table(`board_cells`).
		Where("id=?", id).
		Delete(&models.BoardCell{}).
		Error
}
