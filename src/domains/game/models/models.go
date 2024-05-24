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
	From_id        int
	To_id          int
	FigureId       int
	KilledFigureId int
	NewFigureId    int
	IsCheckWhite   bool
	IsCheckBlack   bool
}

type Game struct {
	Id                 int
	WhiteUserId        int
	BlackUserId        int
	IsStarted          bool
	IsEnded            bool
	IsCheckWhite       bool
	WhiteKingCastling  bool
	WhiteRookACastling bool
	WhiteRookHCastling bool
	IsCheckBlack       bool
	BlackKingCastling  bool
	BlackRookACastling bool
	BlackRookHCastling bool
	LastPawnMove       *int
	Side               bool
}
