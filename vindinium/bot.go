package vindinium

// Bot is the interface that allows your bot to be plugged into the client to
// actually play Vindinium. The Move function is called to retrieve the move
// your bot wants to play and EndOfGame is called at the end of a game and given
// the error that aborted the game (if any) and the last received gamestate.
type Bot interface {
	Move(*State) Direction
	EndOfGame(error, *State)
}
