package redisft

import "github.com/redis/go-redis/v9"

// Client is the main interface of redisft.go SDK.
type Client interface {
	Do(query builder) (*redis.Cmd, error)
	Client() *redis.Client
}

// NewClient returns opens a new redis connection
// and returns a redisft.go Client.
func NewClient(options *redis.Options) Client {
	rdb := redis.NewClient(options)

	return &client{
		conn: rdb,
	}
}

// NewClientFromExistingConnection only redisft.go
// Client without opening a new connection.
func NewClientFromExistingConnection(conn *redis.Client) Client {
	return &client{
		conn: conn,
	}
}
