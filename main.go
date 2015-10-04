package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
)

var (
	server      string
	key         string
	mode        string
	botname     string
	numParallel int
	numGames    int

	shouldExit   bool
	gamesPlayerd int
)

func init() {
	flag.StringVar(&server, "s", "http://vindinium.org", "server")
	flag.StringVar(&key, "k", "", "API key")
	flag.StringVar(&mode, "m", "training", "game mode (\"arena\" or \"training\")")
	flag.StringVar(&botname, "b", "random", "name of the bot to use")
	flag.IntVar(&numParallel, "j", 1, "how many instances of the bot to run in parallel (arena mode only)")
	flag.IntVar(&numGames, "c", 1, "how many games to play (0 for continuous play)")
	flag.Parse()
}

func main() {
	// Handle interrups
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	go func() {
		for _ = range interruptChan {
			if shouldExit {
				fmt.Println("Received second interrupt, exiting immediately.")
				os.Exit(1)
			}

			fmt.Println("Received interrupt. Waiting for running game(s) to end.")
			fmt.Println("Send interrupt again to exit immediately.")
			shouldExit = true
		}
	}()

	// Set up the options correctly
	if key == "" {
		fmt.Println("You must provide an API key with the -k command-line option.")
		os.Exit(1)
	}

	if mode == "training" {
		numParallel = 1
	} else if mode == "arena" {
		// Do nothing
	} else {
		fmt.Printf("Unrecognized mode: %q. Mode must be \"arena\" or \"training\".\n", mode)
		os.Exit(1)
	}

	bot, ok := BotRegistry[botname]
	if !ok {
		fmt.Printf("The bot name %q is not registered. To register a new bot, edit `bog_registry.go` and recompile the client with `go build`.\n", botname)
		os.Exit(1)
	}

	_ = bot

	c := &Client{
		Server:    "http://vindinium.org",
		Key:       "3oli39f3",
		Bot:       BotRegistry["random"],
		ArenaMode: false,
	}

	for {
		c.Play()

		numGames--
		if numGames <= 0 {
			break
		}

		if shouldExit {
			break
		}
	}
}
