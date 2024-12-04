package game

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type BoardCellsRepository struct {
	db *gorm.DB

	startField [][]int
}

func CreateBoardCellsRepository(db *gorm.DB) BoardCellsRepository {
	return BoardCellsRepository{
		db:         db,
		startField: makeStartField(),
	}
}

func (b *BoardCellsRepository) CreateNewBoardCells(gameId int, tx *gorm.DB) error {
	var boardCells = make([]models.BoardCell, len(b.startField))

	for i, cell := range b.startField {
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

func makeStartField() [][]int {
	return [][]int{
		{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
		{8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13}, {13, 13}, {14, 13}, {15, 13},
		{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
		{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
	}
}
