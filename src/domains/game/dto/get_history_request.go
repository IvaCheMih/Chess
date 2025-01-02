package dto

import "github.com/IvaCheMih/chess/src/domains/game/models"

type GetHistoryResponse struct {
	Moves []models.Move
}
