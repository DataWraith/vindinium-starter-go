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
func (c *Client) Play() GameResult {
	var err error

	err = c.startGame()
	if err == nil {
		err = c.playGame()
	}

	return GameResult{
		Error:     err,
		LastState: c.state,
	}
}

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

func (c *Client) playGame() error {
	for {
		if c.state.Game.Finished {
			fmt.Println(c.state)
			return nil
		}

		dir := c.Bot.Move(c.state)

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

func (c Client) getStartURL() string {
	if c.ArenaMode {
		return c.Server + "/api/arena"
	}
	return c.Server + "/api/training"
}

func (c *Client) unmarshalState(body io.ReadCloser) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error while reading server response: %s", err.Error())
	}

	err = json.Unmarshal(data, &c.state)
	return err
}

func (c Client) formatResponseError(resp *http.Response) error {
	if resp.StatusCode >= 500 {
		return fmt.Errorf("The server responded with status %d.", resp.StatusCode)
	}

	if resp.StatusCode >= 400 {
		reason, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error while reading server response: %s", err.Error())
		}

		return fmt.Errorf("Request error (status %d): %s", resp.StatusCode, string(reason))
	}

	return fmt.Errorf("received unexpected status code from server: %d", resp.StatusCode)
}
