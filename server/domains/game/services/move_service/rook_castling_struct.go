package move_service

type RookCastling struct {
	RookType bool
	Castling bool
}

func (r *RookCastling) GetType() bool {
	return r.RookType
}

func (r *RookCastling) WasCastling() bool {
	return r.Castling
}
