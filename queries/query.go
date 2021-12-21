package queries

import (
	"bytes"
	"goir/types"
)

type Query interface {
	Apply(termToDocId map[string]types.Set, resultTermToDocId *map[string]types.Set) map[string]types.Set
	GetTerm() string
}

type MatchingQuery struct {
	queryStr string
}

type ConjunctionQuery struct {
	queries []Query
}

func And(queries ...Query) Query {
	return ConjunctionQuery{
		queries,
	}

}

func MatchWith(queryStr string) Query {
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

func (c ConjunctionQuery) And(other Query) Query {
	c.queries = append(c.queries, other)
	return c
}

func (c ConjunctionQuery) Apply(termToDocId map[string]types.Set, resultTermToDocId *map[string]types.Set) map[string]types.Set {
	for _, internalQuery := range c.queries {
		internalQuery.Apply(termToDocId, resultTermToDocId)
	}

	var resultDocSet types.Set
	for _, docSet := range *resultTermToDocId {
		if resultDocSet == nil {
			resultDocSet = docSet
			continue
		}

		resultDocSet = resultDocSet.Intersect(docSet)
	}

	for term := range *resultTermToDocId {
		(*resultTermToDocId)[term] = resultDocSet
	}

	return *resultTermToDocId
}

func (c ConjunctionQuery) GetTerm() string {
	var buffer bytes.Buffer

	for _, internalQuery := range c.queries {
		buffer.WriteString(internalQuery.GetTerm())
		buffer.WriteString(" ")
	}

	return buffer.String()
}
