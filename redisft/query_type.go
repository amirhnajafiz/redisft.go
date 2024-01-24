package redisft

type QueryType string

const (
	QueryCreate QueryType = "CREATE"
	QueryAdd    QueryType = "ADD"
	QuerySearch QueryType = "SEARCH"
)
