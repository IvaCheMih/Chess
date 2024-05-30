package game

import (
	"database/sql"
	"fmt"
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
	var board_cells = []models.BoardCell{}

	for _, cell := range startField {
		c := models.BoardCell{
			GameId:    gameId,
			IndexCell: cell[0],
			FigureId:  cell[1],
		}
		board_cells = append(board_cells, c)
	}

	err := tx.Create(board_cells).Error

	return err
}

func (b *BoardCellsRepository) Find(gameId int) (models.Board, error) {
	var board_cell = []models.BoardCell{}

	res := b.db.Find(&board_cell).Where("game_id=?", gameId)
	if res.Error != nil {
		return models.Board{}, res.Error
	}

	rows, err := res.Rows()
	if err != nil {
		return models.Board{}, res.Error
	}

	cells, err := RowsToCells(rows)

	return models.Board{Cells: cells}, err
}

func (b *BoardCellsRepository) Update(id int, to int, tx *gorm.DB) error {

	err := tx.Model(&models.BoardCell{}).Where("id=?", id).Updates(map[string]interface{}{"index_cell": to}).Error
	fmt.Println(err)

	return err
}

func (b *BoardCellsRepository) Delete(id int, tx *gorm.DB) error {

	err := tx.Delete(&models.BoardCell{}, id).Error

	fmt.Println(err)

	return err
}

func RowsToCells(rows *sql.Rows) (map[int]*models.BoardCell, error) {
	var cells = map[int]*models.BoardCell{}

	for rows.Next() {
		var cell = models.BoardCell{}

		err := rows.Scan(&cell.Id, &cell.GameId, &cell.IndexCell, &cell.FigureId)
		if err != nil {
			return map[int]*models.BoardCell{}, err
		}

		cells[cell.IndexCell] = &cell
	}

	return cells, nil
}
