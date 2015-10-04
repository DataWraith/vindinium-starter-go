// Package passivebot implements a bot that does nothing. It literally stands
// on spawn all game long, but it serves as a basic example of how to implement
// a bot.
package passivebot

import "../../vindinium/"

type Bot struct{}

func (b *Bot) Move(state vindinium.State) vindinium.Direction {
	return vindinium.Stay
}
