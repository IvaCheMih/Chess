package dto

import "github.com/IvaCheMih/chess/server/domains/game/models"

type CreateNewFigureBody struct {
	FigureId int
}

type CreateNewFigureResponse struct {
	models.Game
}
