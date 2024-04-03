package game

import (
	"errors"
	"fmt"
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

func (g *GamesService) CreateGame(userId any, userRequestedColor bool) (dto.ResponseGetGame, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	defer tx.Rollback()

	var requestCreateGame dto.ResponseGetGame

	if userRequestedColor {
		requestCreateGame, err = g.gamesRepo.CreateGame(userId, tx)
	} else {
		requestCreateGame, err = g.gamesRepo.FindNotStartedGame(tx)
		if err != nil {
			return dto.ResponseGetGame{}, err
		}

		err = g.gamesRepo.JoinBlackToGame(requestCreateGame.GameId, userId, tx)
	}

	if err != nil {
		fmt.Println(222)
		return dto.ResponseGetGame{}, err
	}

	if userRequestedColor {
		err = g.boardRepo.CreateNewBoardCells(requestCreateGame.GameId, tx)
	}
	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	err = tx.Commit()

	return requestCreateGame, err
}

func (g *GamesService) GetBoard(gameId int, userId any) (dto.GetBoardResponse, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	defer tx.Rollback()

	game, err := g.gamesRepo.GetById(gameId, tx)

	if userId != game.WhiteUserId && userId != game.BlackUserId {
		return dto.GetBoardResponse{}, err
	}

	boardCells, err := g.boardRepo.Find(gameId, tx)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	err = tx.Commit()

	responseBoard := make([]dto.BoardCellEntity, 64)

	for _, boardCell := range boardCells {
		responseBoard[boardCell.IndexCell] = dto.BoardCellEntity{boardCell.IndexCell, boardCell.FigureId}
	}

	getBoardResponse := dto.GetBoardResponse{
		BoardCells: responseBoard,
	}

	return getBoardResponse, err
}

func (g *GamesService) GetHistory(gameId int, userId any) (dto.ResponseGetHistory, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.ResponseGetHistory{}, err
	}

	defer tx.Rollback()

	responseGetGame, err := g.gamesRepo.GetById(gameId, tx)
	if err != nil {
		return dto.ResponseGetHistory{}, err
	}

	if userId != responseGetGame.WhiteUserId && userId != responseGetGame.BlackUserId {
		return dto.ResponseGetHistory{}, errors.New("This is not your game")
	}

	moves, err := g.movesRepo.Find(gameId, tx)
	if err != nil {
		return dto.ResponseGetHistory{}, err
	}

	var responseGetHistory = dto.ResponseGetHistory{
		Moves: moves,
	}

	return responseGetHistory, err
}

func (g *GamesService) DoMove(gameId int, userId any, requestFromTo dto.RequestDoMove) (dto.ResponseDoMove, error) {
	if !dto.CheckCorrectRequest(requestFromTo.From, requestFromTo.To) {
		return dto.ResponseDoMove{}, errors.New("Move is not correct")
	}

	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.ResponseDoMove{}, err
	}

	defer tx.Rollback()

	responseGetGame, err := g.gamesRepo.GetById(gameId, tx)
	if err != nil {
		return dto.ResponseDoMove{}, err
	}

	moves, err := g.movesRepo.Find(gameId, tx)
	if err != nil {
		return dto.ResponseDoMove{}, err
	}

	if err = CheckCorrectRequestSideUser(userId, responseGetGame, moves); err != nil {
		return dto.ResponseDoMove{}, err
	}

	//boardCells, err := g.boardRepo.Find(gameId, tx)
	//if err != nil {
	//	return dto.ResponseDoMove{}, err
	//}
	//
	//game := move_service.CreateGameStruct(responseGetGame, boardCells, moves[len(moves)])

	var responseDoMove dto.ResponseDoMove

	return responseDoMove, err
}

func CheckCorrectRequestSideUser(userId any, responseGetGame dto.ResponseGetGame, moves []dto.Move) error {
	if userId != responseGetGame.WhiteUserId && userId != responseGetGame.BlackUserId {
		return errors.New("This is not your game")
	}

	if !responseGetGame.IsStarted || responseGetGame.IsEnded {
		return errors.New("Game is not active")
	}

	if len(moves)%2 == 0 && userId != responseGetGame.WhiteUserId {
		return errors.New("Its not your move now")
	}

	if len(moves)%2 != 0 && userId != responseGetGame.BlackUserId {
		return errors.New("Its not your move now")
	}
	return nil
}

var startField = [][]int{
	{0, 7}, {1, 8}, {2, 9}, {3, 10}, {4, 11}, {5, 9}, {6, 8}, {7, 7},
	{8, 12}, {9, 12}, {10, 12}, {11, 12}, {12, 12}, {13, 12}, {14, 12}, {15, 12},
	{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 1},
}
