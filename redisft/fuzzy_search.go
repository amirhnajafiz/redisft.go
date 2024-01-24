package redisft

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (c client) create(ctx context.Context) *redis.Cmd {
	args := []string{
		c.query.getCommand(),
		c.query.Index,
		"SCHEMA",
	}

	for key := range c.query.Params {
		args = append(args, key, c.query.Params[key])
	}

	return c.conn.Do(ctx, args)
}

func (c client) add(ctx context.Context) *redis.Cmd {
	args := []string{
		c.query.getCommand(),
		c.query.Index,
	}

	for key := range c.query.Params {
		args = append(args, key, c.query.Params[key])
	}

	return c.conn.Do(ctx, args)
}

func (c client) search(ctx context.Context) *redis.Cmd {
	args := []string{
		c.query.getCommand(),
		c.query.Index,
		c.query.Params["query"],
		"FUZZY",
	}

	return c.conn.Do(ctx, args)
}
