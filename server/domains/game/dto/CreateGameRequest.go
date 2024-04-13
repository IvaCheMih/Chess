package dto

type CreateGameBody struct {
	IsWhite bool
}

type CreateGameResponse struct {
	GameId        int
	WhiteUserId   int
	BlackUserId   int
	IsStarted     bool
	IsEnded       bool
	IsCheckWhite  bool
	WhiteKingCell int
	IsCheckBlack  bool
	BlackKingCell int
	Side          int
}
