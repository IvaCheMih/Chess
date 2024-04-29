package models

type BoardCell struct {
	Id        int
	GameId    int
	IndexCell int
	FigureId  int
}

type Board struct {
	Cells map[int]*BoardCell
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
	Id            int  `json:"id"`
	WhiteUserId   int  `json:"whiteUserId"`
	BlackUserId   int  `json:"blackUserId"`
	IsStarted     bool `json:"isStarted"`
	IsEnded       bool `json:"isEnded"`
	IsCheckWhite  bool `json:"isCheckWhite"`
	WhiteKingCell int  `json:"whiteKingCell"`
	IsCheckBlack  bool `json:"isCheckBlack"`
	BlackKingCell int  `json:"blackKingCell"`
	Side          int  `json:"side"`
}
