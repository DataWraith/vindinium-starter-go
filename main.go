package main

import (
	"flag"
	"fmt"
)

var (
	server      string
	key         string
	mode        string
	botname     string
	numParallel int
	numGames    int

	shouldExit bool
)

func init() {
	flag.StringVar(&server, "s", "http://vindinium.org", "server")
	flag.StringVar(&key, "k", "", "API key")
	flag.StringVar(&mode, "m", "training", "game mode (\"arena\" or \"training\")")
	flag.StringVar(&botname, "b", "random", "name of the bot to use")
	flag.IntVar(&numParallel, "j", 1, "how many instances of the bot to run in parallel")
	flag.IntVar(&numGames, "c", 1, "how many games to play (0 for continuous play)")
	flag.Parse()
}

func main() {
	c := &Client{
		Server:    "http://vindinium.org",
		Key:       "3oli39f3",
		Bot:       BotRegistry["random"],
		ArenaMode: false,
	}

	result := c.Play()
	fmt.Println(result)
	fmt.Println(result.LastState.ViewURL)
}
