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
	FromId         int
	ToId           int
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
	EndReason          EndgameReason
	IsCheckWhite       bool
	WhiteKingCastling  bool
	WhiteRookACastling bool
	WhiteRookHCastling bool
	IsCheckBlack       bool
	BlackKingCastling  bool
	BlackRookACastling bool
	BlackRookHCastling bool
	LastLoss           int
	LastPawnMove       *int
	Side               bool
}

type EndgameReason string

const (
	NotEndgame EndgameReason = "NotEndgame"
	Mate       EndgameReason = "Mate"
	Pat        EndgameReason = "Pat"
	Repetition EndgameReason = "Repetition"
	NoLosses   EndgameReason = "NoLosses"
)

func (e *EndgameReason) ToDTO() string {
	switch *e {
	case Mate:
		return "Mate"
	case Pat:
		return "Pat"
	case Repetition:
		return "Repetition"
	case NoLosses:
		return "NoLosses"
	case NotEndgame:
		return "NotEndgame"
	}

	return ""
}
