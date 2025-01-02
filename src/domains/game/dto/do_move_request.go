package dto

type DoMoveBody struct {
	From      string
	To        string
	NewFigure byte
}

type DoMoveResponse struct {
	BoardCells []BoardCellEntity
}
