package move_service

//type Figure interface {
//	IsWhite() bool
//	ToString() string
//	GetType() byte
//	GetPossibleMoves(*Game) *TheoryMoves
//	ChangeGameIndex(int)
//	GetGameIndex() int
//}
//
//func CreateDefaultField(cells []dto.BoardCell) []*Figure {
//	//startField := []byte{
//	//	'r', 'k', 'b', 'q', 'K', 'b', 'k', 'r',
//	//	'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p',
//	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
//	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
//	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
//	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
//	//	'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p',
//	//	'r', 'k', 'b', 'q', 'K', 'b', 'k', 'r',
//	//}
//	field := []*Figure{}
//
//	for i, r := range cells {
//
//		if i <= 15 {
//			field = append(field, CreateFigure(r, false, i))
//		}
//		if i >= 48 {
//			field = append(field, CreateFigure(r, true, i))
//		}
//		if i < 48 && i > 15 {
//			field = append(field, nil)
//		}
//
//	}
//
//	return field
//}
//
//func FigureToString(figure *Figure) string {
//	if figure == nil {
//		return "0"
//	}
//
//	return (*figure).ToString()
//}

//func CreateFigure(_type byte, isWhite bool, index int) *Figure {
//	figure := CreateFigure1(_type, isWhite, index)
//
//	if figure == nil {
//		return nil
//	}
//
//	return &figure
//}
//
//func CreateFigure1(_type byte, isWhite bool, index int) Figure {
//	var bf = BaseFigure{isWhite, _type, index}
//	//var tm = TheoryMoves{nil, nil, nil, nil, nil, nil, nil, nil, nil}
//	switch _type {
//	case 'p':
//		return &FigurePawn{bf}
//	case 'r':
//		return &FigureRook{bf}
//	case 'k':
//		return &FigureKnight{bf}
//	case 'b':
//		return &FigureBishop{bf}
//	case 'q':
//		return &FigureQueen{bf}
//	case 'K':
//		return &FigureKing{bf}
//	}
//	return nil
//}
