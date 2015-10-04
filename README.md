# Vindinium starter-kit for Go

This is an alternative Go starter-kit for the [Vindinium] AI Challenge.

[Vindinium]: http://vindinium.org


## Why another starter-kit?

The official [starter-kit](https://github.com/geetarista/vindinium-starter-go)
for Go is a fine choice; however, I feel that some of the design decisions
were not optimal.


### Differences

This project has the following differences to the official starter-kit:

* Different Board representation

    The board representation is different. Instead of using `interface{}` for
    tiles, tiles are integers that represent the type of tile (Air, Hero, etc.).
    Hero identity, mine ownership information and tavern locations are available
    as seperate data structures.

* Separation of client code and bot code

    The code for the actual bot(s) is isolated in a sub-directory, instead of in
    the same directory as the starter-kit code. This comes at a cost, however,
    as you must import the starter-kit code, which means that everything is
    prefixed with `vindinium.`. That can be a little cumbersome.

* Multi-instance mode

    You can run multiple instances of a bot in parallel.

* Graceful shutdown

    You can interrupt the client with Ctrl-C (SIGTERM) at any time. It will
    finish the running games and then shut down. If you press Ctrl-C a second
    time, it will shut down immediately without finishing the current game(s).

* Does not panic

    The official starter-kit will panic (crash) if the server sends an
    unexpected response (such as a Timeout error or a truncated response).

* Less configurable

    You cannot configure the number of turns or map used for training mode.
    I consider the default settings good enough.

* Less verbose

    The client does not output anything _at all_. If you want to see progress
    updates, statistics or even the viewUrl, your bot must print that itself.


## Installation

1. `git clone https://github.com/DataWraith/vindinium-starter-go`
2. `cd vindinium-starter-go`
3. `go build`
4. `./vindinium-starter-go -h`


## Usage

### Client

The client takes the following commandline switches:

* `-k <API key>`

    Your API key. This option is required.

* `-s <Server>`

    The server to play on. Defaults to `http://vindinium.org`.

* `-m <mode>`

    The game mode. Must be either 'training' or 'arena'. Defaults to 'training'.

* `-b <bot name>`

    The name of the bot you want to run (defaults to "random"). See below for
    how to add your own bot.

* `-c <number of games>`

    How many games you want to play (defaults to 1). If you specify 0 here, the
    client will continue to play games until interrupted.

    Note that a game will be counted multiple times if more than one instance
    of the bot participates. That means that if three instances participate in
    the same game, that game will be counted as three games.

* `-j <number of instances>`

    How many instances of the bot to run in parallel. Defaults to 1.

    Make sure your machine has enough CPU power to support bots running in
    parallel, or you may run into timeout issues.

    You probably don't want to run more bots than you have CPU cores, but if
    you do, make sure that the GOMAXPROCS environment variable is set to a
    high enough value.


### Writing your own bot

#### Quickstart

The fastest way to get started is to look at the example bots provided in
the `bots/` directory. Note that you have to register your bot in the file
`bot_registry.go` in order to make it known to the client.


#### Explanation

To write your own bot, create a directory under `bots/`. This directory will
contain the source code for your bot.

Your bot's code lives in its own package. To interact with the client, you
have to import it (`import "../../vindinium"`) and then create a `struct` that
conforms to the `vindinium.Bot` interface.

    `type Bot struct {}`

To implement the interface, you have to implement two methods:

1. `func (b *Bot) Move(s vindinium.State) vindinium.Direction {}`
2. `func (b *Bot) EndOfGame(err error, s vindinium.State) {}`

The first method is where the meat of your bot is. It takes a gamestate object
and returns your move. The second method is useful for bots that can learn from
the game outcome or want to print out statistics of the match -- `err` is the
error that caused the game to be aborted (if any) and `s` is the last-received
gamestate.

## Questions

Feel free to open GitHub issues for any questions you may have.

## License

This starter-kit is distributed under the unlicense. The the LICENSE file for
more information.
