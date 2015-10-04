package main

import "fmt"

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
