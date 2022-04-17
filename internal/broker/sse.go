package broker

import net "github.com/subchord/go-sse"

func New() *net.Broker {
	return net.NewBroker(GetConfigs())
}
