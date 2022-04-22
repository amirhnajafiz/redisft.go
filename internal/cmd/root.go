package cmd

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/amirhnajafiz/event-man/internal/cmd/server"
	"github.com/amirhnajafiz/event-man/pkg/logger"
)

func Execute() {
	server.New()

	httpMux := reflect.ValueOf(http.DefaultServeMux).Elem()
	finList := httpMux.FieldByIndex([]int{1})
	fmt.Println("Routes:")
	for _, key := range finList.MapKeys() {
		logger.Info("ROUTE", fmt.Sprintf("\t 0.0.0.0:8080%s\n", key))
	}

	logger.Info("INFO", "Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
