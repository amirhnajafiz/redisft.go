package redisft

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

func (c client) create(ctx context.Context) (*redis.Cmd, error) {
	args := []string{
		"FT.CREATE",
		c.query.index,
		"SCHEMA",
	}

	if len(c.query.params) == 0 {
		return nil, errors.New("not enough parameters")
	}

	for key := range c.query.params {
		args = append(args, key, c.query.params[key])
	}

	return c.conn.Do(ctx, args), nil
}

func (c client) add(ctx context.Context) (*redis.Cmd, error) {
	args := []string{
		"FT.ADD",
		c.query.index,
	}

	if len(c.query.params) == 0 {
		return nil, errors.New("not enough parameters")
	}

	for key := range c.query.params {
		args = append(args, key, c.query.params[key])
	}

	return c.conn.Do(ctx, args), nil
}

func (c client) search(ctx context.Context) (*redis.Cmd, error) {
	args := []string{
		"FT.SEARCH",
		c.query.index,
		c.query.params["query"],
		"FUZZY",
	}

	if len(c.query.params) == 0 {
		return nil, errors.New("not enough parameters")
	}

	return c.conn.Do(ctx, args), nil
}
