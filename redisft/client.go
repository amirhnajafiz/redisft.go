package redisft

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type client struct {
	conn  *redis.Client
	query *Query
}

func (c client) Do(query Query) *redis.Cmd {
	ctx := context.Background()

	c.query = &query

	switch query.Type {
	case QuerySearch:
		return c.search(ctx)
	case QueryCreate:
		return c.create(ctx)
	case QueryAdd:
		return c.add(ctx)
	}

	return nil
}

func (c client) GetClient() *redis.Client {
	return c.conn
}
