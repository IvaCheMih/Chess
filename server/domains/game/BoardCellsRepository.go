package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
)

type BoardCellsRepository struct {
	db *sql.DB
}

func CreateBoardCellsRepository(db *sql.DB) BoardCellsRepository {
	return BoardCellsRepository{
		db: db,
	}
}

func (b *BoardCellsRepository) CreateNewBoardCells(gameId int, tx *sql.Tx) bool {
	startField := [][]int{
		{0, 7}, {1, 8}, {2, 9}, {3, 10}, {4, 11}, {5, 9}, {6, 8}, {7, 7},
		{8, 12}, {9, 12}, {10, 12}, {11, 12}, {12, 12}, {13, 12}, {14, 12}, {15, 12},
		{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
		{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 1},
	}

	for _, v := range startField {
		err := tx.QueryRow(`
			INSERT INTO boardCells 
				values (gameId, indexCell, figureId)
			`,
			gameId,
			v[0],
			v[1],
		).Err()

		if err != nil {
			return false
		}
	}

	return true
}

func (b *BoardCellsRepository) GetBoardCells(gameId int, tx *sql.Tx) ([]dto.BoardCell, error) {
	resultQuery, err := tx.Query(`
		SELECT (indexCell, figureId) FROM boardCells
		    where gameId = $1
		`,
		gameId,
	)

	if err != nil {
		return []dto.BoardCell{}, err
	}

	boardCells, err := GetCellsFromRows(resultQuery)

	return boardCells, err

}

func GetCellsFromRows(rows *sql.Rows) ([]dto.BoardCell, error) {
	var cells []dto.BoardCell

	for rows.Next() {
		var cell dto.BoardCell
		err := rows.Scan(&cell.IndexCell, &cell.FigureId)
		if err != nil {
			return []dto.BoardCell{}, err
		}
		cells = append(cells, cell)
	}

	return cells, nil
}
