package vindinium

// Game contains all information relevant for the current game.
type Game struct {
	ID       string
	Turn     int
	MaxTurns int
	Heroes   [4]Hero
	Board    Board
	Finished bool
}
