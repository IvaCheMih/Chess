package dto

type BoardCellEntity struct {
	IndexCell int `json:"index"`
	FigureId  int `json:"figureId"`
}

type GetBoardResponse struct {
	BoardCells []BoardCellEntity `json:"boardCells"`
}
