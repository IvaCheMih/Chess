package dto

import "github.com/IvaCheMih/chess/server/domains/game/models"

type GetHistoryResponse struct {
	Moves []models.Move
}
