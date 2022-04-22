package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amirhnajafiz/event-man/internal/cmd/server"
)

func Execute() {
	server.New()

	fmt.Println("[Info] Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
