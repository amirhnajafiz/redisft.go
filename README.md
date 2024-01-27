<p align="center">
  <a href="https://skillicons.dev">
    <img width="200" src="https://skillicons.dev/icons?i=redis" />
  </a>
</p>

<p align="center">
    <b>redisft.go</b><br />
    Integrating go-redis with fuzzy search
</p>

A __Golang__ SDK for integrating [```go-redis```](https://github.com/redis/go-redis) with redis [fuzzy search feature](https://redis.com/blog/what-is-fuzzy-matching/).

## Redis

Redis is an in-memory data structure store often used as a database, cache, or message broker. Fuzzy search, on the other hand, is a technique used to match results that are approximately equal to a given pattern.

### Fuzzy Search

Redis itself doesn't inherently support fuzzy search out-of-the-box in the same way that some dedicated search engines or databases might. However, developers have implemented fuzzy search functionality using Redis by combining it with other tools or techniques.
One common approach to implementing fuzzy search with Redis involves using a combination of Redis data structures like sets, sorted sets, or hashes along with algorithms like Levenshtein distance or Jaccard similarity to perform approximate matching of strings or values stored in Redis.

For example, one might store indexed tokens or shingles of strings in Redis sorted sets, where each token or shingle is scored based on its relevance or similarity to the original string. Then, by querying these sorted sets with a search term and using Redis commands to retrieve relevant tokens or shingles, one can approximate a fuzzy search.

It's important to note that while Redis can be used to implement fuzzy search, it might not be as efficient or feature-rich as dedicated search solutions like Elasticsearch or Apache Solr, which are specifically designed for full-text search and offer more advanced capabilities out of the box. However, for certain use cases or scenarios where Redis is already being used in the stack, implementing fuzzy search with Redis can be a viable option.

### Query Syntax

Redis FT (Redis Full Text) is a module introduced in Redis version 6.0, which adds full-text search capabilities to Redis. It allows you to index and search textual content efficiently within Redis data structures. Redis FT uses the powerful search engine library, RediSearch, under the hood.

```redis
FT.CREATE my_index SCHEMA title TEXT body TEXT
FT.ADD my_index doc1 1.0 FIELDS title "Redis Introduction" body "Redis is an open-source, in-memory data structure store."
FT.ADD my_index doc2 1.0 FIELDS title "Redis Commands" body "Redis provides various commands for data manipulation and retrieval."
FT.SEARCH my_index "@title:(Redis)"
```

A list of the available commands can be found at [query syntax docs](https://redis.io/docs/interact/search-and-query/advanced-concepts/query_syntax/).

## redisft.go

With this software development kit (SDK), you can utilize fuzzy search capabilities alongside the features provided by the go-redis SDK.

### connection

The RedisFT tool provides two options for establishing a Redis connection to utilize its fuzzy search feature. You can either utilize an existing connection or initiate a new connection.

```go
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

// use existing
ft := redisft.NewClientFromExistingConnection(rdb)

// create new connection
ft = redisft.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})
```

### schema

Create a new schema:

```go
res, err := ft.Do(
    redisft.NewCreateQuery("default").
    AddField("cindex", redisft.FTText).
    Build(), 
)
```

### add

Add items to an index:

```go
res, err := ft.Do(
    redisft.NewAddQuery("default").
    AddValue("cindex", "value").
    Build(),
)
```

### search

Search in an index:

```go
res, err := ft.Do(
    redisft.NewSearchQuery("default").
    WithQuery("*all*").
    Build(),
)
```
