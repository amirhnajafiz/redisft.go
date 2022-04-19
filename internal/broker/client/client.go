package client

import net "github.com/subchord/go-sse"

type (
	Client struct {
		Broker net.Broker
	}
)

func Broadcast(e net.Event) {

}
