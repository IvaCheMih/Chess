package game

import (
	"errors"
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	moveservice "github.com/IvaCheMih/chess/src/domains/game/services/move"
	"gorm.io/gorm"
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

	FromModelsToDtoCreateGame(notStartedGame, &createGameResponse)

	if createNewBoard {
		boardCells := g.boardRepo.NewStartBoardCells(createGameResponse.GameId)

		err = g.boardRepo.CreateNewBoardCells(tx, boardCells)
	}
	if err != nil {
		return dto.CreateGameResponse{}, err
	}

	tx.Commit()

	return createGameResponse, err
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
		return dto.GetHistoryResponse{}, errors.New("This is not your game")
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
	board, err := g.boardRepo.Find(gameId)
	if err != nil {
		return models.Move{}, err
	}

	if !CheckCorrectRequest(requestFromTo) {
		return models.Move{}, errors.New("Move is not correct")
	}

	gameModel, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Move{}, err
	}

	// move logic

	if err = CheckCorrectRequestSideUser(userId, gameModel); err != nil {
		log.Println(err)
		return models.Move{}, err
	}

	from := CoordinatesToIndex(requestFromTo.From)
	to := CoordinatesToIndex(requestFromTo.To)

	indexesToChange, game := g.MoveService.IsMoveCorrect(gameModel, board, from, to, requestFromTo.NewFigure)

	if len(indexesToChange) == 0 {
		return models.Move{}, errors.New("Move is not possible (IsMoveCorrect)")
	}

	// process move and change game params
	game.DoMove(indexesToChange, requestFromTo.NewFigure)

	// player cant do this move if his king is under attack
	if game.CheckToMovingPlayer() {
		return models.Move{}, errors.New("Move is not possible (Check)")
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

	err = UpdateBoardAfterMove(tx, g, board, game.NewFigureId, indexesToChange)
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

	return responseMove, err
}

func (g *GamesService) GiveUp(gameId int) (models.Game, error) {
	err := g.gamesRepo.UpdateIsEnded(gameId)
	if err != nil {
		return models.Game{}, err
	}

	game, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Game{}, err
	}

	return game, err

}

func CheckCorrectRequestSideUser(userId any, game models.Game) error {
	if userId != game.WhiteUserId && userId != game.BlackUserId {
		return errors.New("This is not your game")
	}

	if !game.IsStarted || game.IsEnded {
		return errors.New("Game is not active")
	}

	if game.Side && userId != game.WhiteUserId {
		return errors.New("Its not your move now")
	}

	if !game.Side && userId != game.BlackUserId {
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

func CheckCorrectRequest(move dto.DoMoveBody) bool {
	from, to := CoordinatesToIndex(move.From), CoordinatesToIndex(move.To)

	if !CheckCellOnBoardByIndex(from) || !CheckCellOnBoardByIndex(to) {
		return false
	}

	return true
}

func UpdateBoardAfterMove(tx *gorm.DB, g *GamesService, board models.Board, newFigureId int, indexesToChange []int) error {
	var err error
	from := indexesToChange[0]
	to := indexesToChange[1]

	if board.Cells[to] != nil {
		err = g.boardRepo.Delete(tx, board.Cells[to].Id)
		if err != nil {
			return err
		}
	}

	if newFigureId != 0 {
		err = g.boardRepo.UpdateNewFigure(tx, board.Cells[from].Id, to, newFigureId)
	} else {
		err = g.boardRepo.Update(tx, board.Cells[from].Id, to)
	}

	if err != nil {
		return err
	}

	if len(indexesToChange) > 2 {
		if indexesToChange[2] == -1 {
			err = g.boardRepo.Delete(tx, board.Cells[indexesToChange[3]].Id)
		} else {
			err = g.boardRepo.Update(tx, board.Cells[indexesToChange[2]].Id, indexesToChange[3])
		}
	}

	return err
}

func FromModelsToDtoCreateGame(response models.Game, createGameResponse *dto.CreateGameResponse) {
	createGameResponse.GameId = response.Id

	createGameResponse.IsCheckWhite = response.IsCheckWhite
	createGameResponse.IsCheckBlack = response.IsCheckBlack

	createGameResponse.IsStarted = response.IsStarted
	createGameResponse.IsEnded = response.IsEnded

	createGameResponse.WhiteKingCastling = response.WhiteKingCastling
	createGameResponse.BlackKingCastling = response.BlackKingCastling

	createGameResponse.WhiteRookACastling = response.WhiteRookACastling
	createGameResponse.WhiteRookHCastling = response.WhiteRookHCastling

	createGameResponse.BlackRookACastling = response.BlackRookACastling
	createGameResponse.BlackRookHCastling = response.BlackRookHCastling

	createGameResponse.BlackUserId = response.BlackUserId
	createGameResponse.WhiteUserId = response.WhiteUserId

	createGameResponse.LastPawnMove = response.LastPawnMove
	createGameResponse.Side = response.Side
}

func (g *GamesService) GetMoveService() *moveservice.MoveService {
	return g.MoveService
}
