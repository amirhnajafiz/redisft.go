package main

import (
	"github.com/amirhnajafiz/redisft.go/redisft"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ft := redisft.NewClientFromExistingConnection(rdb)

	if err := ft.Do(redisft.NewQuery(redisft.QueryAdd).
		WithIndex("myindex").WithParam("docID", "text"),
	).Err(); err != nil {
		panic(err)
	}
}
