package broker

import (
	"fmt"

	"github.com/amirhnajafiz/event-man/internal/logger"
	net "github.com/subchord/go-sse"
)

func New() *net.Broker {
	n := net.NewBroker(GetConfigs())

	n.SetDisconnectCallback(DisconnectCallback)

	return n
}

func DisconnectCallback(clientId string, sessionId string) {
	logger.Fatal(fmt.Errorf("[%s] Lost connection with ID: %s", sessionId, clientId))
}
