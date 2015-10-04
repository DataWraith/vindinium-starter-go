package main

import "./vindinium"
import "strconv"

// GameResult holds the result of a Vindinium game.
type GameResult struct {
	// Error holds the last error (if any)
	Error error

	// LastState holds the last state received before the game ended or was aborted
	// due to error
	LastState vindinium.State
}

func (g GameResult) String() string {
	if g.Error != nil {
		return "aborted due to error: " + g.Error.Error()
	}

	result := ""
	for _, h := range g.LastState.Game.Heroes {
		result += h.Name + ": " + strconv.Itoa(h.Gold) + "\n"
	}
	return result
}
