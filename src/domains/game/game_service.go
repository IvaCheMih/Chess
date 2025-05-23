package game

import (
	"errors"
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	moveservice "github.com/IvaCheMih/chess/src/domains/game/services/move"
	"log"
)

type GamesService struct {
	boardRepo *BoardCellsRepository
	gamesRepo *GamesRepository
	movesRepo *MovesRepository

	MoveService *moveservice.MoveService
}

func CreateGamesService(boardRepo *BoardCellsRepository, gamesRepo *GamesRepository, movesRepo *MovesRepository) GamesService {
	figureRepo := moveservice.CreateFigureRepo()
	return GamesService{
		boardRepo: boardRepo,
		gamesRepo: gamesRepo,
		movesRepo: movesRepo,

		MoveService: moveservice.NewMoveService(figureRepo),
	}
}

func (g *GamesService) GetGame(gameId int, accountId int) (dto.GetGameResponse, error) {
	game, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return dto.GetGameResponse{}, err
	}

	if accountId != game.WhiteUserId && accountId != game.BlackUserId {
		return dto.GetGameResponse{}, errors.New("you cant view this game")
	}

	return dto.GetGameResponse{
		GameId:             game.Id,
		WhiteUserId:        game.WhiteUserId,
		BlackUserId:        game.BlackUserId,
		IsStarted:          game.IsStarted,
		IsEnded:            game.IsEnded,
		EndReason:          game.EndReason.ToDTO(),
		Winner:             game.Winner,
		IsCheckWhite:       game.IsCheckWhite,
		WhiteKingCastling:  game.WhiteKingCastling,
		WhiteRookACastling: game.WhiteRookACastling,
		WhiteRookHCastling: game.WhiteRookHCastling,
		IsCheckBlack:       game.IsCheckBlack,
		BlackKingCastling:  game.BlackKingCastling,
		BlackRookACastling: game.BlackRookACastling,
		BlackRookHCastling: game.BlackRookHCastling,
		LastLoss:           game.LastLoss,
		LastPawnMove:       game.LastPawnMove,
		Side:               game.Side,
	}, nil
}

func (g *GamesService) CreateGame(userId int, userRequestedColor bool) (dto.CreateGameResponse, error) {
	var createGameResponse dto.CreateGameResponse
	createNewBoard := false

	userColor := "white_user_id"
	gameSide := false

	if !userRequestedColor {
		userColor = "black_user_id"
	}

	notStartedGame, err := g.gamesRepo.FindNotStartedGame(userColor)
	if err != nil && err.Error() != "record not found" {
		return dto.CreateGameResponse{}, err
	}

	if userColor == "black_user_id" {
		gameSide = true
	}

	tx := g.gamesRepo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return dto.CreateGameResponse{}, err
	}

	if err != nil && err.Error() == "record not found" {
		var game = models.Game{}

		if userRequestedColor {
			game.WhiteUserId = userId
		} else {
			game.BlackUserId = userId
		}

		notStartedGame, err = g.gamesRepo.CreateGame(tx, game)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}
		createNewBoard = true
	} else {
		err = g.gamesRepo.UpdateColorUserIdByColor(tx, notStartedGame.Id, userColor, gameSide, userId)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}

		notStartedGame, err = g.gamesRepo.GetByIdTx(tx, notStartedGame.Id)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}

	}

	fromModelsToDtoCreateGame(notStartedGame, &createGameResponse)

	if createNewBoard {
		boardCells := g.boardRepo.NewStartBoardCells(createGameResponse.GameId)

		err = g.boardRepo.CreateNewBoardCells(tx, boardCells)
	}
	if err != nil {
		return dto.CreateGameResponse{}, err
	}

	tx.Commit()

	return createGameResponse, nil
}

func (g *GamesService) GetBoard(gameId int, userId any) (dto.GetBoardResponse, error) {
	game, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	if userId != game.WhiteUserId && userId != game.BlackUserId {
		return dto.GetBoardResponse{}, err
	}

	board, err := g.boardRepo.Find(gameId)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	responseBoard := make([]dto.BoardCellEntity, 64)

	for i := 0; i <= 63; i++ {
		figureId := 0
		if cell, ok := board.Cells[i]; ok {
			figureId = cell.FigureId
		}
		responseBoard[i] = dto.BoardCellEntity{
			IndexCell: i,
			FigureId:  figureId,
		}
	}

	//for index, cell := range board.Cells {
	//	responseBoard[index] = dto.BoardCellEntity{
	//		IndexCell: index,
	//		FigureId:  cell.FigureId,
	//	}
	//}

	getBoardResponse := dto.GetBoardResponse{
		BoardCells: responseBoard,
	}

	return getBoardResponse, err
}

func (g *GamesService) GetHistory(gameId int, userId any) (dto.GetHistoryResponse, error) {
	responseGetGame, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return dto.GetHistoryResponse{}, err
	}

	if userId != responseGetGame.WhiteUserId && userId != responseGetGame.BlackUserId {
		return dto.GetHistoryResponse{}, errors.New("this is not your game")
	}

	moves, err := g.movesRepo.Find(gameId)
	if err != nil {
		return dto.GetHistoryResponse{}, err
	}

	var responseGetHistory = dto.GetHistoryResponse{
		Moves: moves,
	}

	return responseGetHistory, err
}

func (g *GamesService) Move(gameId int, userId any, requestFromTo dto.DoMoveBody) (models.Move, error) {
	gameModel, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Move{}, err
	}

	if userId != gameModel.WhiteUserId && userId != gameModel.BlackUserId {
		return models.Move{}, errors.New("this is not your game")
	}

	board, err := g.boardRepo.Find(gameId)
	if err != nil {
		return models.Move{}, err
	}

	if !checkCorrectRequest(requestFromTo) {
		return models.Move{}, errors.New("move is not correct")
	}

	// move logic

	if err = checkCorrectRequestSideUser(userId, gameModel); err != nil {
		log.Println(err)
		return models.Move{}, err
	}

	from := coordinatesToIndex(requestFromTo.From)
	to := coordinatesToIndex(requestFromTo.To)

	indexesToChange, game := g.MoveService.IsMoveCorrect(gameModel, board, from, to, requestFromTo.NewFigure)

	if len(indexesToChange) == 0 {
		return models.Move{}, errors.New("move is not possible (IsMoveCorrect)")
	}

	// process move and change game params
	game.DoMove(indexesToChange, requestFromTo.NewFigure)

	// player cant do this move if his king is under attack
	if game.CheckToMovingPlayer() {
		return models.Move{}, errors.New("move is not possible (Check)")
	}

	// move is correct and done. Change side to check Endgame and save game, board, move state
	game.ChangeSide()

	// end move logic

	history, err := g.movesRepo.Find(gameId)
	if err != nil {
		log.Println(err)
		return models.Move{}, err
	}

	isEnd, endReason := g.MoveService.IsItEndgame(&game, history, g.boardRepo.NewStartBoardCells(1))

	tx := g.gamesRepo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return models.Move{}, err
	}

	var move = models.Move{
		GameId:         gameId,
		MoveNumber:     len(history),
		FromId:         from,
		ToId:           to,
		FigureId:       board.Cells[from].FigureId,
		KilledFigureId: g.MoveService.GetFigureID(game.KilledFigure),
		NewFigureId:    g.MoveService.GetFigureID(requestFromTo.NewFigure),
		IsCheckWhite:   game.IsCheckWhite.IsItCheck,
		IsCheckBlack:   game.IsCheckBlack.IsItCheck,
	}

	responseMove, err := g.movesRepo.AddMove(tx, move)
	if err != nil {
		return models.Move{}, err
	}

	err = g.gamesRepo.UpdateGame(tx, gameId, game, isEnd, endReason)
	if err != nil {
		return models.Move{}, err
	}

	err = updateBoardAfterMove(tx, g, board, game.NewFigureId, indexesToChange)
	if err != nil {
		return models.Move{}, err
	}

	tx.Commit()

	responseBoard := make([]dto.BoardCellEntity, 64)

	cells, err := g.boardRepo.Find(gameId)
	if err != nil {
		return models.Move{}, err
	}

	for i := 0; i <= 63; i++ {
		figureId := 0
		if cell, ok := cells.Cells[i]; ok {
			figureId = cell.FigureId
		}
		responseBoard[i] = dto.BoardCellEntity{
			IndexCell: i,
			FigureId:  figureId,
		}
	}

	//for index, cell := range cells.Cells {
	//	responseBoard[index] = dto.BoardCellEntity{
	//		IndexCell: index,
	//		FigureId:  cell.FigureId,
	//	}
	//
	//}

	getBoardResponse := dto.GetBoardResponse{
		BoardCells: responseBoard,
	}

	for i := 0; i < 64; i++ {
		if i%8 == 0 {
			fmt.Print("\n")
		}

		if getBoardResponse.BoardCells[i].FigureId == 0 {
			fmt.Print(0)
		} else {
			fmt.Print(string(g.MoveService.GetFigureRepo()[getBoardResponse.BoardCells[i].FigureId]))
		}
	}
	fmt.Println()

	return responseMove, nil
}

func (g *GamesService) EndGame(userId int, gameId int, reasonString string) (models.Game, error) {
	var reason models.EndgameReason
	err := reason.FromString(reasonString)
	if err != nil {
		return models.Game{}, err
	}

	game, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Game{}, err
	}

	if userId != game.WhiteUserId || userId != game.BlackUserId {
		return models.Game{}, errors.New("this is not your game")
	}

	if game.IsEnded {
		return models.Game{}, errors.New("game is already ended")
	}

	if !game.IsStarted {
		return models.Game{}, errors.New("game is not started")
	}

	if reason == models.GiveUp {
		winner := 0
		if userId == game.WhiteUserId {
			winner = game.BlackUserId
		} else {
			winner = game.WhiteUserId
		}
		err = g.gamesRepo.UpdateIsEnded(winner, gameId, reason)
		if err != nil {
			return models.Game{}, err
		}
	}

	//TODO: giveup

	game, err = g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (g *GamesService) GetMoveService() *moveservice.MoveService {
	return g.MoveService
}
