package game

import (
	"errors"
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/move_service"
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

	err = tx.Commit()

	var responseGetHistory = dto.ResponseGetHistory{
		Moves: moves,
	}

	return responseGetHistory, err
}

func (g *GamesService) DoMove(gameId int, userId any, requestFromTo dto.RequestDoMove) (dto.GetBoardResponse, error) {
	if !CheckCorrectRequest(requestFromTo.From, requestFromTo.To) {
		return dto.GetBoardResponse{}, errors.New("Move is not correct")
	}

	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	defer tx.Rollback()

	responseGetGame, err := g.gamesRepo.GetById(gameId, tx)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	if err = CheckCorrectRequestSideUser(userId, responseGetGame); err != nil {
		return dto.GetBoardResponse{}, err
	}

	boardCells, err := g.boardRepo.Find(gameId, tx)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	if !move_service.CheckCorrectMove(responseGetGame, boardCells, requestFromTo) {
		return dto.GetBoardResponse{}, errors.New("Move is not possible (CheckCorrectMove)")
	}

	from := CoordinatesToIndex(requestFromTo.From)
	to := CoordinatesToIndex(requestFromTo.To)

	game, check := move_service.CheckIsItCheck(responseGetGame, boardCells, from, to)

	if !check {
		return dto.GetBoardResponse{}, errors.New("Move is not possible (CheckIsItCheck)")
	}

	err = g.movesRepo.AddMove(gameId, from, to, boardCells[from].FigureId, boardCells[to].FigureId, game.IsCheckWhite, game.IsCheckBlack, tx)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	err = g.gamesRepo.UpdateGame(gameId, game.IsCheckWhite, game.IsCheckBlack, game.Side, tx)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	if boardCells[to].FigureId != 0 {
		err = g.boardRepo.Delete(boardCells[to].Id, tx)

		if err != nil {
			return dto.GetBoardResponse{}, err
		}
	}

	err = g.boardRepo.Update(boardCells[from].Id, to, tx)

	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	boardCells, err = g.boardRepo.Find(gameId, tx)
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

func CheckCorrectRequestSideUser(userId any, responseGetGame dto.ResponseGetGame) error {
	if userId != responseGetGame.WhiteUserId && userId != responseGetGame.BlackUserId {
		return errors.New("This is not your game")
	}

	if !responseGetGame.IsStarted || responseGetGame.IsEnded {
		return errors.New("Game is not active")
	}

	if responseGetGame.WhiteUserId == responseGetGame.Side && userId != responseGetGame.WhiteUserId {
		return errors.New("Its not your move now")
	}

	if responseGetGame.BlackUserId == responseGetGame.Side && userId != responseGetGame.BlackUserId {
		return errors.New("Its not your move now")
	}
	return nil
}

func IndexToCoordinates(index int) string {
	y := int('8') - (index / 8)
	x := (index % 8) + int('A')

	return string(byte(x)) + string(byte(y))
}

func CoordinatesToIndex(coordinates string) int {
	x := int(coordinates[0]) - int('A')
	y := int('8') - int(coordinates[1])

	return (y * 8) + x
}

func CheckCellOnBoardByIndex(index int) bool {
	coordinates := IndexToCoordinates(index)
	if coordinates[0] >= byte('A') && coordinates[0] <= byte('H') {
		if coordinates[1] >= byte('1') && coordinates[1] <= byte('8') {
			return true
		}
	}
	return false
}

func CheckCorrectRequest(f, t string) bool {
	from, to := CoordinatesToIndex(f), CoordinatesToIndex(t)

	if !CheckCellOnBoardByIndex(from) || !CheckCellOnBoardByIndex(to) {
		return false
	}
	return true
}

var startField = [][]int{
	{0, 7}, {1, 8}, {2, 9}, {3, 10}, {4, 11}, {5, 9}, {6, 8}, {7, 7},
	{8, 12}, {9, 12}, {10, 12}, {11, 12}, {12, 12}, {13, 12}, {14, 12}, {15, 12},
	{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 1},
}
