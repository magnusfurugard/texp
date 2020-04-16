package parse

import (
	"sort"
	"strings"
)

// TokensToMap parses key-value pairs (k=v) into a keymap.
func TokensToMap(tokens *[]string) map[string][]string {
	m := map[string][]string{}
	for _, token := range *tokens {
		kvs := strings.SplitN(token, "=", 2)
		if len(kvs) != 2 {
			panic("failed to unpack key-value pair: `" + token + "` (is it key=value?)")
		}
		m[kvs[0]] = append(m[kvs[0]], kvs[1])
	}
	return m
}

// ReplaceToken replaces a single token instance, duplicating the query as needed.
func ReplaceToken(queries []string, key string, values []string) []string {
	newQueries := []string{}
	for _, value := range values {
		for _, q := range queries {
			q := strings.ReplaceAll(q, key, value)
			newQueries = append(newQueries, q)
		}
	}
	return newQueries
}

// UniqueQueries returns the unique queries, deduplicating as required.
// TODO: Respect incoming sort order. Right now alphabetical to make it deterministic for tests.
func UniqueQueries(queries []string) []string {
	m := map[string]int{}
	result := []string{}
	for _, v := range queries {
		if m[v] == 0 {
			result = append(result, v)
			m[v]++
		}
	}

	sort.Strings(result)
	return result
}

// Tokens parses an incoming query and a slice of tokens into the number of result queries.
// Returns the tokens and the result queries.
func Tokens(query string, tokens *[]string) (map[string][]string, []string) {
	m := TokensToMap(tokens)
	queries := []string{query}
	for k, arr := range m {
		for range queries {
			queries = ReplaceToken(queries, k, arr)
		}
	}
	return m, UniqueQueries(queries)
}
