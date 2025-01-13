package models

import (
	"errors"
)

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
	Status             GameStatus
	EndReason          EndgameReason
	WinnerUserId       int
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
	Draw       EndgameReason = "Draw"
	GiveUp     EndgameReason = "GiveUp"
)

func (e *EndgameReason) ToString() string {
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
	case Draw:
		return "Draw"
	case GiveUp:
		return "GiveUp"
	}

	return ""
}

func (e *EndgameReason) FromString(reason string) error {
	switch reason {
	case string(Mate):
		*e = Mate
		return nil
	case string(Pat):
		*e = Pat
		return nil
	case string(Repetition):
		*e = Repetition
		return nil
	case string(NoLosses):
		*e = NoLosses
		return nil
	case string(NotEndgame):
		*e = NotEndgame
		return nil
	case string(Draw):
		*e = Draw
		return nil
	case string(GiveUp):
		*e = GiveUp
		return nil
	default:
		return errors.New("invalid endgame reason")
	}
}

type GameStatus string

const (
	Created   GameStatus = "Created"
	Cancelled GameStatus = "Cancelled"
	Active    GameStatus = "Active"
	Ended     GameStatus = "Ended"
)

func (e *GameStatus) ToString() string {
	switch *e {
	case Created:
		return "Created"
	case Cancelled:
		return "Cancelled"
	case Active:
		return "Active"
	case Ended:
		return "Ended"
	}

	return ""
}

func (e *GameStatus) FromString(reason string) error {
	switch reason {
	case string(Created):
		*e = Created
		return nil
	case string(Cancelled):
		*e = Cancelled
		return nil
	case string(Active):
		*e = Active
		return nil
	case string(Ended):
		*e = Ended
		return nil
	default:
		return errors.New("invalid game status")
	}
}

func (g *Game) IsActive() bool {
	switch g.Status {
	case Active:
		return true
	default:
		return false
	}
}

func (g *Game) CanBeCancelled() bool {
	switch g.Status {
	case Created:
		return true
	default:
		return false
	}
}
