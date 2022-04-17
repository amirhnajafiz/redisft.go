package broker

import (
	"log"

	net "github.com/subchord/go-sse"
)

func New() *net.Broker {
	n := net.NewBroker(GetConfigs())

	n.SetDisconnectCallback(DisconnectCallback)

	return n
}

func DisconnectCallback(clientId string, sessionId string) {
	log.Fatalf("[%s] Lost connection with ID: %s", sessionId, clientId)
}
