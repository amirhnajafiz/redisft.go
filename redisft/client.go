package redisft

import "github.com/redis/go-redis/v9"

type client struct {
	conn *redis.Client
}

func (c client) Do(query Query) error {
	return nil
}
