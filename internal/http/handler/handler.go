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

func generateUniqueId() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func (h Handler) ClientConnect(writer http.ResponseWriter, request *http.Request) {
	err := h.Client.Connect(strconv.FormatInt(time.Now().UnixMilli(), 10), writer, request)
	if err != nil {
		log.Println(err)

		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(FailedConnection))

		return
	}
}

func (h Handler) PushEvent(writer http.ResponseWriter, request *http.Request) {
	var r struct {
		Event string `json:"event"`
		Data  string `json:"data"`
	}

	err := json.NewDecoder(request.Body).Decode(&r)
	if err != nil {
		log.Println(err)

		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(BadRequest))

		return
	}

	h.Client.BroadcastAll(event.New(generateUniqueId(), r.Event, r.Data))

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte(SuccessfulMessage))
}

func (h Handler) PushEventToSingleClient(writer http.ResponseWriter, request *http.Request) {
	var r struct {
		Event     string `json:"event"`
		Data      string `json:"data"`
		Reference string `json:"reference"`
	}

	err := json.NewDecoder(request.Body).Decode(&r)
	if err != nil {
		log.Println(err)

		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(BadRequest))

		return
	}

	err = h.Client.BroadcastSingle(r.Reference, event.New(generateUniqueId(), r.Event, r.Data))
	if err != nil {
		log.Println(err)

		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(FailedMessage))

		return
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte(SuccessfulMessage))
}

func (h Handler) EventHistory() {
	// history
}
