package telegram

import (
	"errors"
	"github.com/IvaCheMih/chess/src/domains/game/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func newFigures() map[int]string {
	var figures = make(map[int]string)

	figures[1] = "♖"
	figures[2] = "♘"
	figures[3] = "♗"
	figures[4] = "♕"
	figures[5] = "♔"
	figures[6] = "♙"
	figures[7] = "♖"

	figures[8] = "♜"
	figures[9] = "♞"
	figures[10] = "♝"
	figures[11] = "♛"
	figures[12] = "♚"
	figures[13] = "♟"
	figures[14] = "♜"

	return figures
}

func (b *TelegramService) makeBoardTemplate(board []dto.BoardCellEntity) tgbotapi.InlineKeyboardMarkup {
	var row1, row2, row3, row4, row5, row6, row7, row8 []tgbotapi.InlineKeyboardButton

	for i, cell := range board {

		figure := " "

		if f, ok := b.figures[cell.FigureId]; ok {
			figure = f
		}

		switch i / 8 {
		case 0:
			row1 = append(row1, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		case 1:
			row2 = append(row2, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		case 2:
			row3 = append(row3, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		case 3:
			row4 = append(row4, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		case 4:
			row5 = append(row5, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		case 5:
			row6 = append(row6, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		case 6:
			row7 = append(row7, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		case 7:
			row8 = append(row8, tgbotapi.NewInlineKeyboardButtonData(figure, makeData(cell.IndexCell, figure)))
		}
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		row1,
		row2,
		row3,
		row4,
		row5,
		row6,
		row7,
		row8,
	)
}

func makeData(indexCell int, figure string) string {
	return strconv.Itoa(indexCell) + "/" + figure
}

func parseData(data string) (int, error) {
	words := strings.Split(data, "/")
	if len(words) != 2 {
		return 0, errors.New("unknown data")
	}

	index, err := strconv.Atoi(words[0])
	if err != nil {
		return 0, errors.New("wrong data")
	}

	if index < 0 || index > 63 {
		return 0, errors.New("wrong data")
	}

	return index, nil
}
