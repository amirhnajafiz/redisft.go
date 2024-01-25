package redisft

type builder struct {
	index     string
	params    map[string]string
	queryType QueryType
}

func newQuery(index string, queryType QueryType) builder {
	return builder{
		index:     index,
		queryType: queryType,
		params:    make(map[string]string),
	}
}

func (q builder) Build() builder {
	return q
}

func (q builder) AddField(field, fieldType string) CreateQuery {
	q.withParam(field, fieldType)

	return q
}

func (q builder) AddValue(key, value string) AddQuery {
	q.withParam(key, value)

	return q
}

func (q builder) WithQuery(query string) SearchQuery {
	q.withParam("query", query)

	return q
}

func (q builder) withParam(key, value string) builder {
	q.params[key] = value

	return q
}
