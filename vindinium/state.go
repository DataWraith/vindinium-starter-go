package vindinium

// State represents the current state of a game of Vindinium. This is a
// direct representation of the JSON that is returned by the server.
type State struct {
	Game    Game   // The current gamestate
	Hero    Hero   // Convenience accessor for information about your own hero
	Token   string // The security token needed for playing the game (you can ignore this)
	ViewURL string // The URL at which you can watch the game in your browser
	PlayURL string // The URL the client uses to play in the game (you can ignore this)
}
