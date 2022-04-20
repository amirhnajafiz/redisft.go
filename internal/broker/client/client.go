package client

import (
	"net/http"

	net "github.com/subchord/go-sse"
)

type (
	Client struct {
		Broker *net.Broker
	}
)

func (c *Client) Connect(baseURL string, w http.ResponseWriter, r *http.Request) error {
	con, err := c.Broker.Connect(baseURL, w, r)

	if err != nil {
		return err
	}

	<-con.Done()

	return nil
}

func (c *Client) BroadcastAll(e net.Event) {
	c.Broker.Broadcast(e)
}

func (c *Client) BroadcastSingle(id string, e net.Event) error {
	err := c.Broker.Send(id, e)

	if err != nil {
		return err
	}

	return nil
}
