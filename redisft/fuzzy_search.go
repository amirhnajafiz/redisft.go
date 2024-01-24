package redisft

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (c client) create(ctx context.Context) *redis.Cmd {
	args := []string{
		"FT.CREATE",
		c.query.index,
		"SCHEMA",
	}

	for key := range c.query.params {
		args = append(args, key, c.query.params[key])
	}

	return c.conn.Do(ctx, args)
}

func (c client) add(ctx context.Context) *redis.Cmd {
	args := []string{
		"FT.ADD",
		c.query.index,
	}

	for key := range c.query.params {
		args = append(args, key, c.query.params[key])
	}

	return c.conn.Do(ctx, args)
}

func (c client) search(ctx context.Context) *redis.Cmd {
	args := []string{
		"FT.SEARCH",
		c.query.index,
		c.query.params["query"],
		"FUZZY",
	}

	return c.conn.Do(ctx, args)
}
