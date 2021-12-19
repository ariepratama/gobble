package queries

import (
	"goir/types"
)

type Query interface {
	Apply(termToDocId map[string]types.Set, resultTermToDocId *map[string]types.Set) map[string]types.Set
	GetTerm() string
}

type MatchingQuery struct {
	queryStr string
}

func NewMatchingQuery(queryStr string) Query {
	return MatchingQuery{
		queryStr,
	}
}

func (m MatchingQuery) Apply(termToDocId map[string]types.Set, resultTermToDocId *map[string]types.Set) map[string]types.Set {
	docIds, ok := termToDocId[m.queryStr]

	if ok {
		(*resultTermToDocId)[m.queryStr] = docIds
	}

	return *resultTermToDocId
}

func (m MatchingQuery) GetTerm() string {
	return m.queryStr
}
