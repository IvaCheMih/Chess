package models

type Cell struct {
	Id        int
	IndexCell int
	FigureId  int
}

type Board struct {
	Cells map[int]*Cell
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

type Game struct {
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
