package game

import (
	"database/sql"
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	_ "github.com/lib/pq"
)

type BoardCellsRepository struct {
	db *sql.DB
}

func CreateBoardCellsRepository(db *sql.DB) BoardCellsRepository {
	return BoardCellsRepository{
		db: db,
	}
}

func (b *BoardCellsRepository) CreateNewBoardCells(gameId int, tx *sql.Tx) error {
	var err error

	/*
		baseParams = []{};

		baseInsertQuery := `
			INSERT INTO boardCells (gameId, indexCell, figureId)
			values ($1, $2, $3)
		`;

		for index, boardCell := range boardCells {
			if index == 0 {
				continue
			}

			baseInsertQuery += "($index*2, $index*2+1, $index*2+2)"
			baseParams = append(baseParams, gameId, boardCell...)
		}

		err = tx.QueryRow(baseInsertQuery, baseParams...)
	*/

	err = tx.QueryRow(`
		INSERT INTO boardCells (gameId, indexCell, figureId)
			values ($1,$2,$3), ($4,$5,$6), ($7,$8,$9), ($10,$11,$12), ($13,$14,$15), ($16,$17,$18), ($19,$20,$21), ($22,$23,$24), ($25,$26,$27), ($28,$29,$30),
					($31,$32,$33), ($34,$35,$36), ($37,$38,$39), ($40,$41,$42), ($43,$44,$45), ($46,$47,$48), ($49,$50,$51), ($52,$53,$54), ($55,$56,$57), ($58,$59,$60),
			   		  ($61,$62,$63), ($64,$65,$66), ($67,$68,$69), ($70,$71,$72), ($73,$74,$75), ($76,$77,$78), ($79,$80,$81), ($82,$83,$84), ($85,$86,$87), ($88,$89,$90),
						($91,$92,$93), ($94,$95,$96)
			   		  `,
		gameId, startField[0][0], startField[0][1], gameId, startField[1][0], startField[1][1], gameId, startField[2][0], startField[2][1], gameId, startField[3][0], startField[3][1], gameId, startField[4][0], startField[4][1],
		gameId, startField[5][0], startField[5][1], gameId, startField[6][0], startField[6][1], gameId, startField[7][0], startField[7][1], gameId, startField[8][0], startField[8][1], gameId, startField[9][0], startField[9][1],
		gameId, startField[10][0], startField[10][1], gameId, startField[11][0], startField[11][1], gameId, startField[12][0], startField[12][1], gameId, startField[13][0], startField[13][1], gameId, startField[14][0], startField[14][1],
		gameId, startField[15][0], startField[15][1], gameId, startField[16][0], startField[16][1], gameId, startField[17][0], startField[17][1], gameId, startField[18][0], startField[18][1], gameId, startField[19][0], startField[19][1],
		gameId, startField[20][0], startField[20][1], gameId, startField[21][0], startField[21][1], gameId, startField[22][0], startField[22][1], gameId, startField[23][0], startField[23][1], gameId, startField[24][0], startField[24][1],
		gameId, startField[25][0], startField[25][1], gameId, startField[26][0], startField[26][1], gameId, startField[27][0], startField[27][1], gameId, startField[28][0], startField[28][1], gameId, startField[29][0], startField[29][1],
		gameId, startField[30][0], startField[30][1], gameId, startField[31][0], startField[31][1],
	).Err()

	if err != nil {
		return err
	}

	fmt.Println(1)
	return err
}

func (b *BoardCellsRepository) GetBoardCells(gameId int, tx *sql.Tx) ([]dto.BoardCell, error) {
	resultQuery, err := tx.Query(`
		SELECT indexCell, figureId FROM boardCells
		    where gameId = $1 ORDER BY indexCell
		`,
		gameId,
	)

	if err != nil {
		return []dto.BoardCell{}, err
	}

	var cells []dto.BoardCell

	for resultQuery.Next() {
		var cell dto.BoardCell
		err = resultQuery.Scan(&cell.IndexCell, &cell.FigureId)
		if err != nil {
			return []dto.BoardCell{}, err
		}
		cells = append(cells, cell)
	}

	return cells, nil
}

func GetCellsFromRows(rows *sql.Rows) ([]dto.BoardCell, error) {
	var cells []dto.BoardCell

	for rows.Next() {
		var cell dto.BoardCell
		err := rows.Scan(&cell)
		if err != nil {
			return []dto.BoardCell{}, err
		}
		cells = append(cells, cell)
	}

	return cells, nil
}
