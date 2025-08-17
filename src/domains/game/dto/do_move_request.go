package dto

import "github.com/IvaCheMih/chess/src/domains/game/models"

type DoMoveBody struct {
	From      string
	To        string
	NewFigure byte
}

type DoMoveResponse struct {
	models.Move
	End bool
}
