package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"./vindinium"
)

// Client is responsible for interacting with the bots and executing their moves
// on the server.
type Client struct {
	Server    string
	Key       string
	Bot       vindinium.Bot
	ArenaMode bool

	state vindinium.State
}

// Play plays a game using the given client.
func (c *Client) Play() {
	var err error

	err = c.startGame()
	if err == nil {
		err = c.playGame()
	}
	c.Bot.EndOfGame(err, &c.state)
}

// startGame makes the initial request to the vindinium server
func (c *Client) startGame() error {
	startURL := c.getStartURL()
	formValues := make(url.Values)
	formValues.Set("key", c.Key)

	resp, err := http.PostForm(startURL, formValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		err = c.unmarshalState(resp.Body)
		return err
	}

	return c.formatResponseError(resp)
}

// playGame uses the bot to make all moves of a game
func (c *Client) playGame() error {
	for {
		if c.state.Game.Finished {
			return nil
		}

		dir := c.Bot.Move(&c.state)

		formValues := make(url.Values)
		formValues.Set("key", c.Key)
		formValues.Set("dir", string(dir))

		resp, err := http.PostForm(c.state.PlayURL, formValues)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			err = c.unmarshalState(resp.Body)
			if err != nil {
				return err
			}
			continue
		}

		return c.formatResponseError(resp)
	}
}

// getStartURL returns the URL used to enqueue in a game
func (c Client) getStartURL() string {
	if c.ArenaMode {
		return c.Server + "/api/arena"
	}
	return c.Server + "/api/training"
}

// unmarshalState extracts the gamestate from a server response
func (c *Client) unmarshalState(body io.ReadCloser) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error while reading server response: %s", err.Error())
	}

	err = json.Unmarshal(data, &c.state)
	return err
}

// formatResponseError is invoked when something went wrong and formats an error
// message according to the HTTP status code the serve returned.
func (c Client) formatResponseError(resp *http.Response) error {
	reason, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading server response: %s", err.Error())
	}

	if resp.StatusCode >= 500 {
		return fmt.Errorf("Server error (status %d): %q", resp.StatusCode, string(reason))
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Request error (status %d): %q", resp.StatusCode, string(reason))
	}

	return fmt.Errorf("received unexpected status code from server: %d", resp.StatusCode)
}
