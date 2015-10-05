package vindinium

// Game contains all information relevant for the current game.
type Game struct {
	ID       string  // The ID of the game (an 8-character string)
	Turn     int     // The current Turn. Each bot has its own turn, so you only see every 4th turn
	MaxTurns int     // The maximum number of turns
	Heroes   [4]Hero // Information about the 4 heroes in the game
	Board    Board   // The current game board
	Finished bool    // Whether the game is ongoing or finished
}
