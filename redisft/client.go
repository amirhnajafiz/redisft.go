package redisft

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type client struct {
	conn  *redis.Client
	query *Query
}

func (c client) Do(query Query) error {
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

	return errors.New("query type not found")
}
