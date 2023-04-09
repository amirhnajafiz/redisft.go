package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/amirhnajafiz/event-man/internal/broker/client"
	"github.com/amirhnajafiz/event-man/internal/broker/event"
	"github.com/amirhnajafiz/event-man/internal/cache"
	"github.com/amirhnajafiz/event-man/internal/logger"
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
		logger.Error(err)

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
		logger.Error(err)

		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(BadRequest))

		return
	}

	e := event.New(generateUniqueId(), r.Event, r.Data)

	h.Client.BroadcastAll(e)
	cache.Put(e)

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
		logger.Error(err)

		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(BadRequest))

		return
	}

	e := event.New(generateUniqueId(), r.Event, r.Data)

	err = h.Client.BroadcastSingle(r.Reference, e)
	if err != nil {
		logger.Error(err)

		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(FailedMessage))

		return
	}

	cache.Put(e)

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte(SuccessfulMessage))
}

func (h Handler) EventHistory(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(cache.Pull())
	if err != nil {
		logger.Error(err)

		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(FailedMessage))

		return
	}
}
