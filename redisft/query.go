package redisft

type CreateQuery interface {
	AddField(field, fieldType string) CreateQuery
	Build() query
}

type AddQuery interface {
	AddValue(key, value string) AddQuery
	Build() query
}

type SearchQuery interface {
	WithQuery(query string) SearchQuery
	Build() query
}

func NewCreateQuery(index string) CreateQuery {
	return newQuery(index, queryCreate)
}

func NewSearchQuery(index string) SearchQuery {
	return newQuery(index, querySearch)
}

func NewAddQuery(index string) AddQuery {
	return newQuery(index, queryAdd)
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

func (q query) Build() query {
	return q
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
