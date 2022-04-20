package cmd

import (
	"log"
	"net/http"

	"github.com/amirhnajafiz/event-man/internal/cmd/server"
)

func Execute() {
	server.New()

	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
