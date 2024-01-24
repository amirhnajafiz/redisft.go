package redisft

type Query struct {
	index     string
	params    map[string]string
	queryType QueryType
}

func NewQuery(queryType QueryType) Query {
	return Query{
		queryType: queryType,
	}
}

func NewAddQuery() Query {

}

func (q Query) WithIndex(index string) Query {
	q.index = index

	return q
}

func (q Query) WithParam(key, value string) Query {
	q.params[key] = value

	return q
}
