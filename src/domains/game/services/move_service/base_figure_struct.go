package move_service

type BaseFigure struct {
	IsWhite         bool
	Type            byte
	CellCoordinates []int
}

func (figure *BaseFigure) IsItWhite() bool {
	return figure.IsWhite
}

func (figure *BaseFigure) GetType() byte {
	return figure.Type
}

func (figure *BaseFigure) ChangeCoordinates(newCoordinates []int) {
	figure.CellCoordinates = newCoordinates
}

func (figure *BaseFigure) GetCoordinates() []int {
	return figure.CellCoordinates
}

func (figure *BaseFigure) Delete() {
	figure = nil
}
