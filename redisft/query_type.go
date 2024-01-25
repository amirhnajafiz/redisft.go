package redisft

type QueryType string

const (
	queryCreate QueryType = "CREATE"
	queryAdd    QueryType = "ADD"
	querySearch QueryType = "SEARCH"
)
