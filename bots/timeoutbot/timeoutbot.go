// Package timeoutbot implements a simple bot that times out immediately.
package timeoutbot

import (
	"fmt"
	"time"

	v "../../vindinium"
)

type Bot struct{}

func (b *Bot) Move(s *v.State) v.Direction {
	time.Sleep(2 * time.Second)
	return v.Stay
}

func (b *Bot) EndOfGame(err error, s *v.State) {
	gameID := s.Game.ID

	if err != nil {
		fmt.Printf("Game %s aborted due to error: %v\n", gameID, err)
		return
	}

	fmt.Printf("Game %s finished:\n", gameID)
	for _, h := range s.Game.Heroes {
		fmt.Printf("%4d %s\n", h.Gold, h.Name)
	}
}
