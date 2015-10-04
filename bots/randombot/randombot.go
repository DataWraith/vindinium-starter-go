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
	validDirections := make([]v.Direction, 0, 4)

	for _, dir := range v.NESW {
		newPos := s.Game.Board.To(s.Hero.Pos, dir)
		tile := s.Game.Board.TileAt(newPos)

		switch tile {
		case v.WallTile:
			// We can't walk through a Wall

		case v.HeroTile:
			// We can't walk through a Hero

		case v.TavernTile:
			// We can enter a tavern if we have gold
			if s.Hero.Gold >= 2 {
				validDirections = append(validDirections, dir)
			}

		case v.MineTile:
			// We can conquer a mine if we have enough Life and it doesn't belong to us already
			if s.Hero.Life > 20 && s.Game.Board.MineOwner[newPos] != s.Hero.ID {
				validDirections = append(validDirections, dir)
			}

		case v.AirTile:
			// We can always walk through an AirTile
			validDirections = append(validDirections, dir)
		}
	}

	// If we don't have a direction to move in, stay put.
	if len(validDirections) == 0 {
		return v.Stay
	}

	// Otherwise move in a random valid direction
	return validDirections[rand.Intn(len(validDirections))]
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
