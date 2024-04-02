package dto

type RequestedCreateGame struct {
	Id      int
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

type RequestGetBoard struct {
	GameId int
	UserId int
}
