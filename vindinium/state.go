package vindinium

// State represents the current state of a game of Vindinium.
type State struct {
	Game    Game
	Hero    Hero
	Token   string
	ViewURL string
	PlayURL string
}
