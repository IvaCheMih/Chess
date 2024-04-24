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

func (figure *BaseFigure) ChangeGameIndex(toCoordinate []int) {
	figure.CellCoordinate = toCoordinate
}

func (figure *BaseFigure) GetGameIndex() []int {
	return figure.CellCoordinate
}

func (figure *BaseFigure) Delete() {
	figure = nil
}
