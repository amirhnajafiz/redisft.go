package client

import net "github.com/subchord/go-sse"

type (
	Client struct {
		Broker *net.Broker
	}
)

func (c *Client) Connect(baseURL string) error {
	con, err := c.Broker.Connect(baseURL, nil, nil)

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
