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

	arenaMode  bool
	shouldExit bool
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

	// Check that the options are correct
	if key == "" {
		fmt.Println("You must provide an API key with the -k command-line option.")
		os.Exit(1)
	}

	if numParallel < 1 {
		fmt.Printf("The -j command-line option (number of parallel instances) must be at least 1. You provided: %d.\n", numParallel)
		os.Exit(1)
	}

	if numGames < 0 {
		fmt.Printf("The -c command-line option (number of games to play) must be at least 1. You provided: %d.\n", numGames)
		os.Exit(1)
	}

	if mode == "training" {
		numParallel = 1
	} else if mode == "arena" {
		arenaMode = true
	} else {
		fmt.Printf("Unrecognized mode: %q. Mode must be \"arena\" or \"training\".\n", mode)
		os.Exit(1)
	}

	bot, ok := BotRegistry[botname]
	if !ok {
		fmt.Printf("The bot name %q is not registered. To register a new bot, edit `bog_registry.go` and recompile the client with `go build`.\n", botname)
		os.Exit(1)
	}

	gameChan := make(chan struct{})
	doneChan := make(chan struct{})

	// Start numParallel instances of the bot
	for i := 0; i < numParallel; i++ {
		c := &Client{
			Server:    server,
			Key:       key,
			Bot:       bot,
			ArenaMode: arenaMode,
		}

		go func(c *Client) {
			for _ = range gameChan {
				// Extra check because in most situations we already have the next game
				// buffered inside gameChan, and we don't want to start that if the user
				// wants us to quit.
				if !shouldExit {
					c.Play()
				}
			}
			doneChan <- struct{}{}
		}(c)
	}

	if numGames == 0 {
		// Continuous mode
		for !shouldExit {
			gameChan <- struct{}{}
		}
	} else {
		// Play numGames games
		for i := 0; i < numGames; i++ {
			if shouldExit {
				break
			}

			gameChan <- struct{}{}
		}
	}

	close(gameChan)

	// Wait for the running games to finish
	for i := 0; i < numParallel; i++ {
		<-doneChan
	}
}
