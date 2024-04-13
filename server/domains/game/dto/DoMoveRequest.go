package dto

type DoMoveBody struct {
	From string
	To   string
}

type DoMoveResponse struct {
	BoardCells []BoardCellEntity
}
