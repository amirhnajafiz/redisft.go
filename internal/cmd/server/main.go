package server

import (
	"net/http"

	"github.com/amirhnajafiz/event-man/internal/broker"
	"github.com/amirhnajafiz/event-man/internal/broker/client"
	"github.com/amirhnajafiz/event-man/internal/http/handler"
)

func New() {
	h := handler.Handler{
		Client: &client.Client{
			Broker: broker.New(),
		},
	}

	http.HandleFunc("/connect", h.ClientConnect)
	http.HandleFunc("/push", h.PushEvent)
	http.HandleFunc("/push/single", h.PushEventToSingleClient)
}
