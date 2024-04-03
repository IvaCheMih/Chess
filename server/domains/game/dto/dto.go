package dto

type RequestedCreateGame struct {
	IsWhite bool
}

type ResponseGetGame struct {
	GameId      int
	WhiteUserId int
	BlackUserId int
	IsStarted   bool
	IsEnded     bool
}

type BoardCell struct {
	IndexCell int
	FigureId  int
}

type BoardCellEntity struct {
	IndexCell int `json:"index"`
	FigureId  int `json:"figureId"`
}

type GetBoardResponse struct {
	BoardCells []BoardCellEntity `json:"boardCells"`
}

type Move struct {
	Id             int
	GameId         int
	MoveNumber     int
	From           int
	To             int
	FigureId       int
	KilledFigureId int
	NewFigureId    int
	IsCheckWhite   bool
	WhiteKingCell  int
	IsCheckBlack   bool
	BlackKingCell  int
}

type ResponseGetHistory struct {
	Moves []Move
}

type RequestGetBoard struct {
	GameId int
}

type RequestDoMove struct {
	From string
	To   string
}

type ResponseDoMove struct {
	BoardCells []BoardCell
}
