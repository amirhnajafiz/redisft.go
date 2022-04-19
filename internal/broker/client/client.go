package client

import net "github.com/subchord/go-sse"

type (
	Client struct {
		Broker net.Broker
	}
)

func (c *Client) BroadcastAll(e net.Event) {
	// ...
}

func (c *Client) BroadcastSingle(e net.Event) {
	// ...
}
