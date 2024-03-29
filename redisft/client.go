package redisft

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type client struct {
	conn  *redis.Client
	query *builder
}

func (c client) Do(query builder) (*redis.Cmd, error) {
	ctx := context.Background()

	c.query = &query

	if len(c.query.index) == 0 {
		return nil, ErrEmptyIndex
	}

	switch query.queryType {
	case querySearch:
		return c.search(ctx)
	case queryCreate:
		return c.create(ctx)
	case queryAdd:
		return c.add(ctx)
	}

	return nil, ErrQueryType
}

func (c client) Client() *redis.Client {
	return c.conn
}
