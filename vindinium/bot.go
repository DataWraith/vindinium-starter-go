package vindinium

// Bot is the interface that allows your bot to be plugged into the client to
// actually play Vindinium.
type Bot interface {
	Move(State) Direction
}
