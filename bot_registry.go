package main

import "./vindinium"

import (
	"./bots/passivebot"
	"./bots/randombot"
	"./bots/timeoutbot"
)

// BotRegistry maps the names of the bots to actual values
var BotRegistry = map[string]vindinium.Bot{
	"passive": &passivebot.Bot{},
	"random":  &randombot.Bot{},
	"timeout": &timeoutbot.Bot{},
}
