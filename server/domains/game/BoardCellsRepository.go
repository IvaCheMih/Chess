package game

import (
	"database/sql"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	_ "github.com/lib/pq"
	"strconv"
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
	var baseParams []any

	baseInsertQuery := "INSERT INTO boardCells (gameId, indexCell, figureId) values ($1, $2, $3)"

	for index, cell := range startField {
		baseParams = append(baseParams, gameId, cell[0], cell[1])
		if index == 0 {
			continue
		}
		baseInsertQuery += ", ($" + strconv.Itoa(index*3+1) + ",$" + strconv.Itoa(index*3+2) + ", $" + strconv.Itoa(index*3+3) + ")"

	}

	err = tx.QueryRow(baseInsertQuery, baseParams...).Err()

	//err = tx.QueryRow(`
	//	INSERT INTO boardCells (gameId, indexCell, figureId)
	//		values ($1,$2,$3), ($4,$5,$6), ($7,$8,$9), ($10,$11,$12), ($13,$14,$15), ($16,$17,$18), ($19,$20,$21), ($22,$23,$24), ($25,$26,$27), ($28,$29,$30),
	//				($31,$32,$33), ($34,$35,$36), ($37,$38,$39), ($40,$41,$42), ($43,$44,$45), ($46,$47,$48), ($49,$50,$51), ($52,$53,$54), ($55,$56,$57), ($58,$59,$60),
	//		   		  ($61,$62,$63), ($64,$65,$66), ($67,$68,$69), ($70,$71,$72), ($73,$74,$75), ($76,$77,$78), ($79,$80,$81), ($82,$83,$84), ($85,$86,$87), ($88,$89,$90),
	//					($91,$92,$93), ($94,$95,$96)
	//		   		  `,
	//	gameId, startField[0][0], startField[0][1], gameId, startField[1][0], startField[1][1], gameId, startField[2][0], startField[2][1], gameId, startField[3][0], startField[3][1], gameId, startField[4][0], startField[4][1],
	//	gameId, startField[5][0], startField[5][1], gameId, startField[6][0], startField[6][1], gameId, startField[7][0], startField[7][1], gameId, startField[8][0], startField[8][1], gameId, startField[9][0], startField[9][1],
	//	gameId, startField[10][0], startField[10][1], gameId, startField[11][0], startField[11][1], gameId, startField[12][0], startField[12][1], gameId, startField[13][0], startField[13][1], gameId, startField[14][0], startField[14][1],
	//	gameId, startField[15][0], startField[15][1], gameId, startField[16][0], startField[16][1], gameId, startField[17][0], startField[17][1], gameId, startField[18][0], startField[18][1], gameId, startField[19][0], startField[19][1],
	//	gameId, startField[20][0], startField[20][1], gameId, startField[21][0], startField[21][1], gameId, startField[22][0], startField[22][1], gameId, startField[23][0], startField[23][1], gameId, startField[24][0], startField[24][1],
	//	gameId, startField[25][0], startField[25][1], gameId, startField[26][0], startField[26][1], gameId, startField[27][0], startField[27][1], gameId, startField[28][0], startField[28][1], gameId, startField[29][0], startField[29][1],
	//	gameId, startField[30][0], startField[30][1], gameId, startField[31][0], startField[31][1],
	//).Err()

	if err != nil {
		return err
	}

	return err
}

func (b *BoardCellsRepository) Find(gameId int, tx *sql.Tx) (models.Board, error) {
	resultQuery, err := tx.Query(`
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

func GetCellsFromRows(rows *sql.Rows) (models.Board, error) {
	var board models.Board

	for rows.Next() {
		var cell models.Cell

		err := rows.Scan(&cell)
		if err != nil {
			return models.Board{}, err
		}

		board.Cells[cell.IndexCell] = &cell
	}

	return board, nil
}
