package redisft

import "errors"

var (
	ErrParamsNumber = errors.New("not enough parameters given")
	ErrQueryType    = errors.New("query type not supported")
)
