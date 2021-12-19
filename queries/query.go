package queries

import "goir/core"

type Query interface {
	Apply(document core.Document) bool
}

type MatchingQuery struct {
	queryStr string
}

func NewMatchingQuery(queryStr string) Query {
	return MatchingQuery{
		queryStr,
	}
}

func (m MatchingQuery) Apply(document core.Document) bool {
	return document.Terms().Contains(m.queryStr)
}
