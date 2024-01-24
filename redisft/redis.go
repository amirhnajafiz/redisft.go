package redisft

import "github.com/redis/go-redis/v9"

type Client interface{}

func NewClient(options *redis.Options) Client {
	rdb := redis.NewClient(options)

	return &client{
		conn: rdb,
	}
}

func NewClientFromExistingConnection(conn *redis.Client) Client {
	return &client{
		conn: conn,
	}
}
