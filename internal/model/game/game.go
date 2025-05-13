package game

type GameStatus int

const (
	NotDecidedYet GameStatus = iota
	BlackWon
	WhiteWon
	Draw
)

type Game struct {
	ID          int
	Board       *Board
	CurrentTurn CellState
	Status      GameStatus
}

type GameOption func(*Game)
type GameOptions struct {
	Size int
}

func WithSize(size int) GameOption {
	return func(g *Game) {
		g.Board = NewBoard(size)
	}
}

func NewGame(opts ...GameOption) *Game {
	game := &Game{
		CurrentTurn: Black,
		Status:      NotDecidedYet,
	}

	for _, opt := range opts {
		opt(game)
	}

	// Initialize the board with the default size if not provided
	if game.Board == nil {
		game.Board = NewBoard(19)
	}

	return game
}

func (g *Game) SwitchTurn() {
	if g.CurrentTurn == Black {
		g.CurrentTurn = White
	} else {
		g.CurrentTurn = Black
	}
}

func (g *Game) GetCurrentTurn() CellState {
	return g.CurrentTurn
}

func (g *Game) GetStatus() GameStatus {
	return g.Status
}

func (g *Game) SetStatus(status GameStatus) {
	g.Status = status
}

func (g *Game) GetBoard() *Board {
	return g.Board
}

func (g *Game) IsOver() bool {
	return g.Status != NotDecidedYet
}

func (g *Game) IsCurrentTurn(turn CellState) bool {
	return g.CurrentTurn == turn
}

func (g *Game) IsCurrentTurnBlack() bool {
	return g.IsCurrentTurn(Black)
}

func (g *Game) IsCurrentTurnWhite() bool {
	return g.IsCurrentTurn(White)
}
