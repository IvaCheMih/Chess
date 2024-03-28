package game

import (
	"github.com/IvaCheMih/chess/server/domains/game/dto"
)

type GamesService struct {
	boardRepo  *BoardCellsRepository
	figureRepo *FiguresRepository
	gamesRepo  *GamesRepository
	movesRepo  *MovesRepository
}

func CreateGamesService(boardRepo *BoardCellsRepository, figureRepo *FiguresRepository, gamesRepo *GamesRepository, movesRepo *MovesRepository) GamesService {
	return GamesService{
		boardRepo:  boardRepo,
		figureRepo: figureRepo,
		gamesRepo:  gamesRepo,
		movesRepo:  movesRepo,
	}
}

func (g *GamesService) CreateGame(userId int, userRequestedColor dto.RequestedColor) (dto.ResponseGetGame, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	defer tx.Rollback()

	var requestCreateGame dto.ResponseGetGame

	if userRequestedColor.IsWhite {
		requestCreateGame, err = g.gamesRepo.CreateGame(userId, tx)
	} else {
		requestCreateGame, err = g.gamesRepo.FindNotStartedGame(tx)
		if err != nil {
			return dto.ResponseGetGame{}, err
		}

		err = g.gamesRepo.JoinBlackToGame(requestCreateGame.GameId, userId, tx)
	}

	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	err = g.boardRepo.CreateNewBoardCells(requestCreateGame.GameId, tx)
	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	err = tx.Commit()

	return requestCreateGame, err
}

func (g *GamesService) GetBoard(gameId int, userId int) (dto.ResponseGetBoard, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.ResponseGetBoard{}, err
	}

	defer tx.Rollback()

	responseGetGame, err := g.gamesRepo.GetGame(gameId, tx)

	if userId != responseGetGame.WhiteUserId || userId != responseGetGame.BlackUserId {
		return dto.ResponseGetBoard{}, err
	}

	boardCells, err := g.boardRepo.GetBoardCells(gameId, tx)
	if err != nil {
		return dto.ResponseGetBoard{}, err
	}

	err = tx.Commit()

	responseBoard := make([]dto.BoardCell, 64)

	for _, boardCell := range boardCells {
		responseBoard[boardCell.IndexCell] = dto.BoardCell{boardCell.IndexCell, boardCell.FigureId}

	}

	var responseGetBoard = dto.ResponseGetBoard{
		BoardCells: responseBoard,
	}

	return responseGetBoard, err
}

var startField = [][]int{
	{0, 7}, {1, 8}, {2, 9}, {3, 10}, {4, 11}, {5, 9}, {6, 8}, {7, 7},
	{8, 12}, {9, 12}, {10, 12}, {11, 12}, {12, 12}, {13, 12}, {14, 12}, {15, 12},
	{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 1},
}
