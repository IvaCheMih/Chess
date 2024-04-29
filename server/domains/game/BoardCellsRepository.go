package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/models"
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
	var cells []models.Cell

	for _, cell := range startField {
		c := models.Cell{
			Id:        gameId,
			IndexCell: cell[0],
			FigureId:  cell[1],
		}
		cells = append(cells, c)
	}

	err := tx.Create(&cells).Error

	return err
}

func (b *BoardCellsRepository) Find(gameId int) (models.Board, error) {
	resultQuery, err := b.db.Query(`
		SELECT id, indexCell, figureId FROM boardCells
		    where gameId = $1 ORDER BY indexCell
		`,
		gameId,
	)

	defer resultQuery.Close()

	if err != nil {
		return models.Board{}, err
	}

	var board = models.Board{
		Cells: map[int]*models.Cell{},
	}

	for resultQuery.Next() {
		var cell models.Cell

		err = resultQuery.Scan(&cell.Id, &cell.IndexCell, &cell.FigureId)
		if err != nil {
			return models.Board{}, err
		}

		board.Cells[cell.IndexCell] = &cell
	}

	return board, nil
}

func (b *BoardCellsRepository) Update(id, to int, tx *sql.Tx) error {
	_, err := tx.Exec(`
		update boardCells
		set indexCell = $1
			where (id = $2)
		`,
		to,
		id,
	)

	return err
}

func (b *BoardCellsRepository) Delete(id int, tx *sql.Tx) error {
	_, err := tx.Exec(`
		DELETE FROM boardCells
		       WHERE id = $1
		`,
		id,
	)

	return err
}

func (b *BoardCellsRepository) AddCell(gameId, tx *sql.Tx) error {

	err := tx.QueryRow(`
			INSERT INTO boardCells
				(gameId, indexCell, figureId)
				values ($1, $2, $3)
				`,
		gameId,
	).Err()
	return err
}
