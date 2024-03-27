package dto

type RequestedColor struct {
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

type ResponseGetBoard struct {
	BoardCells []BoardCell
}
