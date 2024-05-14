package dto

type BoardCellEntity struct {
	IndexCell int
	FigureId  int
}

type GetBoardResponse struct {
	BoardCells []BoardCellEntity
}
