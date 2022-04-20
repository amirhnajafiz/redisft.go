package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/amirhnajafiz/event-man/internal/broker/client"
	"github.com/amirhnajafiz/event-man/internal/event"
)

type Handler struct {
	Client *client.Client
}

func (h Handler) PushEvent(writer http.ResponseWriter, request *http.Request) {
	err := h.Client.Connect(strconv.FormatInt(time.Now().UnixMilli(), 10), writer, request)
	if err != nil {
		log.Println(err)

		return
	}

	var r struct {
		Data string `json:"data"`
	}

	err = json.NewDecoder(request.Body).Decode(&r)
	if err != nil {
		log.Println(err)

		return
	}

	h.Client.BroadcastAll(event.New(strconv.FormatInt(time.Now().UnixMilli(), 10), "event", r.Data))

	request.Response.StatusCode = http.StatusOK
}

func (h Handler) EventHistory() {
	// history
}
