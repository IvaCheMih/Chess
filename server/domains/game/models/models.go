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
	Id                 int  `json:"id"`
	WhiteUserId        int  `json:"whiteUserId"`
	BlackUserId        int  `json:"blackUserId"`
	IsStarted          bool `json:"isStarted"`
	IsEnded            bool `json:"isEnded"`
	IsCheckWhite       bool `json:"isCheckWhite"`
	WhiteKingCastling  bool `json:"whiteKingCastling"`
	WhiteRookACastling bool `json:"whiteRookCastling"`
	WhiteRookHCastling bool `json:"whiteRookHCastling"`
	IsCheckBlack       bool `json:"isCheckBlack"`
	BlackKingCastling  bool `json:"blackKingCastling"`
	BlackRookACastling bool `json:"blackRookACastling"`
	BlackRookHCastling bool `json:"blackRookHCastling"`
	LastPawnMove       int  `json:"lastPawnMove"`
	Side               int  `json:"side"`
}
