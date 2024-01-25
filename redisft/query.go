package redisft

type CreateQuery interface {
	AddField(field, fieldType string) CreateQuery
	Build() builder
}

type AddQuery interface {
	AddValue(key, value string) AddQuery
	Build() builder
}

type SearchQuery interface {
	WithQuery(query string) SearchQuery
	Build() builder
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
