package dto

type CreateGameRequest struct {
	IsWhite bool
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

type RequestGetBoard struct {
	GameId int
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

type GetHistoryResponse struct {
	Moves []Move
}

type DoMoveRequest struct {
	From string
	To   string
}

type DoMoveResponse struct {
	BoardCells []BoardCellEntity
}
