package vindinium

// State represents the current state of a game of Vindinium. This is a
// direct representation of the JSON that is returned by the server.
type State struct {
	Game    Game
	Hero    Hero
	Token   string
	ViewURL string
	PlayURL string
}
