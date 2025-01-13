package test

import (
	gamedto "github.com/IvaCheMih/chess/src/domains/game/dto"
)

var moves1 = []gamedto.DoMoveBody{
	{From: "C2", To: "C4", NewFigure: 0},
	{From: "B7", To: "B5", NewFigure: 0},

	{From: "C4", To: "B5", NewFigure: 0},
	{From: "B8", To: "A6", NewFigure: 0},

	{From: "B5", To: "A6", NewFigure: 0},
	{From: "C8", To: "B7", NewFigure: 0},

	{From: "A6", To: "B7", NewFigure: 0},
	{From: "D8", To: "C8", NewFigure: 0},

	{From: "B7", To: "A8", NewFigure: 113},
	{From: "C8", To: "B8", NewFigure: 0},

	{From: "A8", To: "B8", NewFigure: 0},
}

var board1 = [][]int{
	{0, 0}, {1, 4}, {2, 0}, {3, 0}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
	{8, 13}, {9, 0}, {10, 13}, {11, 13}, {12, 13}, {13, 13}, {14, 13}, {15, 13},
	{16, 0}, {17, 0}, {18, 0}, {19, 0}, {20, 0}, {21, 0}, {22, 0}, {23, 0},
	{24, 0}, {25, 0}, {26, 0}, {27, 0}, {28, 0}, {29, 0}, {30, 0}, {31, 0},
	{32, 0}, {33, 0}, {34, 0}, {35, 0}, {36, 0}, {37, 0}, {38, 0}, {39, 0},
	{40, 0}, {41, 0}, {42, 0}, {43, 0}, {44, 0}, {45, 0}, {46, 0}, {47, 0},
	{48, 6}, {49, 6}, {50, 0}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
}

var game1 = gamedto.GetGameResponse{
	Status:    "Ended",
	EndReason: "Mate",
}

var moves2 = []gamedto.DoMoveBody{
	{From: "D2", To: "D4", NewFigure: 0},
	{From: "F7", To: "F5", NewFigure: 0},
	//
	//{From: "G3", To: "F3", NewFigure: 0},
	//{From: "B8", To: "C6", NewFigure: 0},
	//
	//{From: "D2", To: "D4", NewFigure: 0},
	//{From: "F8", To: "B4", NewFigure: 0},
	//
	//{From: "C2", To: "C3", NewFigure: 0},
	//{From: "B4", To: "A5", NewFigure: 0},
	//
	//{From: "D4", To: "D5", NewFigure: 0},
	//{From: "C6", To: "E7", NewFigure: 0},
	//
	//{From: "F3", To: "E5", NewFigure: 0},
	//{From: "G8", To: "F6", NewFigure: 0},
	//
	//{From: "C3", To: "G5", NewFigure: 0},
	//{From: "E7", To: "G6", NewFigure: 0},
}

var board2 = [][]int{
	{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
	{8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13}, {13, 0}, {14, 13}, {15, 13},
	{16, 0}, {17, 0}, {18, 0}, {19, 0}, {20, 0}, {21, 0}, {22, 0}, {23, 0},
	{24, 0}, {25, 0}, {26, 0}, {27, 0}, {28, 0}, {29, 13}, {30, 0}, {31, 0},
	{32, 0}, {33, 0}, {34, 0}, {35, 6}, {36, 0}, {37, 0}, {38, 0}, {39, 0},
	{40, 0}, {41, 0}, {42, 0}, {43, 0}, {44, 0}, {45, 0}, {46, 0}, {47, 0},
	{48, 6}, {49, 6}, {50, 6}, {51, 0}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
}

var game2 = gamedto.GetGameResponse{
	Status:    "Active",
	EndReason: "NotEndgame",
}
