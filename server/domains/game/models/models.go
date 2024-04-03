package models

type BoardCell struct {
	Id        int
	IndexCell int
	FigureId  int
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

type CreateGameResponse struct {
	GameId        int
	WhiteUserId   int
	BlackUserId   int
	IsStarted     bool
	IsEnded       bool
	IsCheckWhite  bool
	WhiteKingCell int
	IsCheckBlack  bool
	BlackKingCell int
	Side          int
}
