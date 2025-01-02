package move

type BaseFigure struct {
	IsWhite         bool
	Type            byte
	CellCoordinates [2]int
}

func (figure *BaseFigure) IsItWhite() bool {
	return figure.IsWhite
}

func (figure *BaseFigure) GetType() byte {
	return figure.Type
}

func (figure *BaseFigure) ChangeCoordinates(newCoordinates [2]int) {
	figure.CellCoordinates = newCoordinates
}

func (figure *BaseFigure) GetCoordinates() [2]int {
	return figure.CellCoordinates
}

func (figure *BaseFigure) Delete() {
	figure = nil //nolint:ineffassign
}

func (figure *BaseFigure) ChangeType(newType byte) {
	figure.Type = newType
}
