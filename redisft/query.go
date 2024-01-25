package redisft

type CreateQuery interface {
	AddField(field, fieldType string) CreateQuery
}

type AddQuery interface {
	AddValue(key, value string) AddQuery
}

type SearchQuery interface {
	WithQuery(query string) SearchQuery
}

func NewCreateQuery(index string) CreateQuery {
	return newQuery(index, QueryCreate)
}

func NewSearchQuery(index string) CreateQuery {
	return newQuery(index, QuerySearch)
}

func NewAddQuery(index string) CreateQuery {
	return newQuery(index, QueryAdd)
}

type query struct {
	index     string
	params    map[string]string
	queryType QueryType
}

func newQuery(index string, queryType QueryType) query {
	return query{
		index:     index,
		queryType: queryType,
		params:    make(map[string]string),
	}
}

func (q query) AddField(field, fieldType string) CreateQuery {
	q.withParam(field, fieldType)

	return q
}

func (q query) AddValue(key, value string) AddQuery {
	q.withParam(key, value)

	return q
}

func (q query) WithQuery(query string) SearchQuery {
	q.withParam("query", query)

	return q
}

func (q query) withParam(key, value string) query {
	q.params[key] = value

	return q
}
