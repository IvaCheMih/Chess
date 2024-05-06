package dto

type CreateGameBody struct {
	IsWhite bool
}

type CreateGameResponse struct {
	GameId             int
	WhiteUserId        int
	BlackUserId        int
	IsStarted          bool
	IsEnded            bool
	IsCheckWhite       bool
	WhiteKingCastling  bool
	WhiteRookACastling bool
	WhiteRookHCastling bool
	BlackKingCastling  bool
	BlackRookACastling bool
	BlackRookHCastling bool
	IsCheckBlack       bool
	LastPawnMove       int
	Side               int
}
