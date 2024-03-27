package postgresql

type User struct {
	Id       int
	Password string
}

type Game struct {
	Id          int
	WhiteUserId int
	BlackUserId int
	IsStarted   bool
	IsEnded     bool
}

type Move struct {
	Id             int
	GameId         int
	MoveNumber     int
	From           int
	To             int
	FigureId       int
	KilledFigureId int
	NewFigureId    int
	IsCheckWhite   bool
	IsCheckBlack   bool
}

type Figure struct {
	Id      int
	Type    string
	IsWhite bool
}

type Board struct {
	Id        int
	GameId    int
	IndexCell int
	FigureId  int
}

type Data struct {
	User
	Game
	Move
	Figure
	Board
}
