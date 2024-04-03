package move_service

type BaseFigure struct {
	IsItWhite bool
	Type      byte
	GameIndex int
}

func (figure *BaseFigure) IsWhite() bool {
	return figure.IsItWhite
}

func (figure *BaseFigure) GetType() byte {
	return figure.Type
}

func (figure *BaseFigure) ChangeGameIndex(toIndex int) {
	figure.GameIndex = toIndex
}

func (figure *BaseFigure) GetGameIndex() int {
	return figure.GameIndex
}
