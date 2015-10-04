package randombot

import (
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
