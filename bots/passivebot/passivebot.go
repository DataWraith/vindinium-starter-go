package passivebot

import "../../vindinium/"

type Bot struct{}

func (b *Bot) Move(state vindinium.State) vindinium.Direction {
	return vindinium.Stay
}
