package move_service

type BaseFigure struct {
	IsItWhite      bool
	Type           byte
	CellCoordinate []int
}

func (figure *BaseFigure) IsWhite() bool {
	return figure.IsItWhite
}

func (figure *BaseFigure) GetType() byte {
	return figure.Type
}

func (figure *BaseFigure) ChangeCoordinates(newCoordinate []int) {
	figure.CellCoordinate = newCoordinate
}

func (figure *BaseFigure) GetCoordinates() []int {
	return figure.CellCoordinate
}

func (figure *BaseFigure) Delete() {
	figure = nil
}
