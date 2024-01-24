package redisft

type Query struct {
	Index  string
	Params map[string]string
	Type   QueryType
}

func NewQuery(queryType QueryType) Query {
	return Query{
		Type: queryType,
	}
}

func (q Query) WithIndex(index string) Query {
	q.Index = index

	return q
}

func (q Query) WithParams(params map[string]string) Query {
	q.Params = params

	return q
}
