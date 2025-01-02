package dto

import "github.com/IvaCheMih/chess/src/domains/game/models"

type CreateNewFigureBody struct {
	FigureId int
}

type CreateNewFigureResponse struct {
	models.Game
}
