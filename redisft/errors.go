package redisft

import "errors"

var (
	ErrEmptyIndex   = errors.New("index cannot be empty")
	ErrParamsNumber = errors.New("not enough parameters given")
	ErrQueryType    = errors.New("query type not supported")
)
