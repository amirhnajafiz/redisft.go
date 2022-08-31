package main

import (
	"log"

	net "github.com/subchord/go-sse"
)

func main() {
	feed, err := net.ConnectWithSSEFeed("http://localhost:8080/connect", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	sub, err := feed.Subscribe("message")
	if err != nil {
		return
	}

	defer sub.Close()
	defer feed.Close()

	for {
		select {
		case evt := <-sub.Feed():
			log.Print(evt)
		case err := <-sub.ErrFeed():
			log.Fatal(err)

			return
		}
	}
}
