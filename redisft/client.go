package redisft

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type client struct {
	conn  *redis.Client
	query *query
}

func (c client) Do(query query) (*redis.Cmd, error) {
	ctx := context.Background()

	c.query = &query

	switch query.queryType {
	case QuerySearch:
		return c.search(ctx)
	case QueryCreate:
		return c.create(ctx)
	case QueryAdd:
		return c.add(ctx)
	}

	return nil, ErrQueryType
}

func (c client) GetClient() *redis.Client {
	return c.conn
}
