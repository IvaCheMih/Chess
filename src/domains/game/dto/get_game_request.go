package dto

type GetGameRequest struct {
	GameId int `json:"game_id"`
}

type GetGameResponse struct {
	GameId             int    `json:"game_id"`
	WhiteUserId        int    `json:"white_user_id"`
	BlackUserId        int    `json:"black_user_id"`
	IsStarted          bool   `json:"is_started"`
	IsEnded            bool   `json:"is_ended"`
	EndReason          string `json:"end_reason"`
	IsCheckWhite       bool   `json:"is_check_white"`
	WhiteKingCastling  bool   `json:"white_king_castling"`
	WhiteRookACastling bool   `json:"white_rook_acastling"`
	WhiteRookHCastling bool   `json:"white_rook_hcastling"`
	BlackKingCastling  bool   `json:"black_king_castling"`
	BlackRookACastling bool   `json:"black_rook_acastling"`
	BlackRookHCastling bool   `json:"black_rook_hcastling"`
	IsCheckBlack       bool   `json:"is_check_black"`
	LastPawnMove       *int   `json:"last_pawn_move"`
	LastLoss           int    `json:"last_loss"`
	Side               bool   `json:"side"`
}
