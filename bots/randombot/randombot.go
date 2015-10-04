// Package randombot implements a simple bot that moves into a random valid
// direction. That is, it avoids making moves that don't do anything (such as
// walking into a wall).
//
// It also prints out part of the game state at the end of the game to
// demonstrate the use of the EndOfGame function.
package randombot

import (
	"fmt"
	"math/rand"

	v "../../vindinium"
)

type Bot struct{}

func (b *Bot) Move(s v.State) v.Direction {
	for i := 0; i < 10; i++ {
		// Pick a random direction
		randDir := v.NESW[rand.Intn(4)]

		newPos := s.Game.Board.To(s.Hero.Pos, randDir)
		tile := s.Game.Board.TileAt(newPos)

		// We can't walk through a wall
		if tile == v.WallTile {
			continue
		}

		// We can't walk through a hero
		if tile == v.HeroTile {
			continue
		}

		// We can only enter a Tavern if we have enough gold
		if tile == v.TavernTile && s.Hero.Gold < 2 {
			continue
		}

		// We can only conquer a Mine if we don't own it
		if tile == v.MineTile && s.Game.Board.MineOwner[newPos] == s.Hero.ID {
			continue
		}

		// We can walk in the chosen direction
		return randDir
	}

	// We tried to walk into a random direction 10 times and it didn't work, so
	// I guess we're staying right where we are.
	return v.Stay
}

func (b *Bot) EndOfGame(err error, s v.State) {
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
