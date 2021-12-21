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

type ConjunctiveQuery struct {
	queries []Query
}

type DisjunctiveQuery struct {
	queries []Query
}

type merger func(types.Set, types.Set) types.Set

func And(queries ...Query) Query {
	return ConjunctiveQuery{
		queries,
	}
}

func Or(queries ...Query) Query {
	return DisjunctiveQuery{
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

func (c ConjunctiveQuery) And(other Query) Query {
	c.queries = append(c.queries, other)
	return c
}

func (c ConjunctiveQuery) Apply(termToDocId map[string]types.Set, resultTermToDocId *map[string]types.Set) map[string]types.Set {
	return applyQuery(c.queries, termToDocId, resultTermToDocId, func(s1 types.Set, s2 types.Set) types.Set { return s1.Intersect(s2) })
}

func (c ConjunctiveQuery) GetTerm() string {
	return getTerm(c.queries)
}

func (d DisjunctiveQuery) Apply(termToDocId map[string]types.Set, resultTermToDocId *map[string]types.Set) map[string]types.Set {
	return applyQuery(d.queries, termToDocId, resultTermToDocId, func(s1 types.Set, s2 types.Set) types.Set { return s1.Union(s2) })
}

func (d DisjunctiveQuery) GetTerm() string {
	return getTerm(d.queries)
}

func getTerm(queries []Query) string {
	var buffer bytes.Buffer

	for _, internalQuery := range queries {
		buffer.WriteString(internalQuery.GetTerm())
		buffer.WriteString(" ")
	}

	return buffer.String()

}

func applyQuery(queries []Query, termToDocId map[string]types.Set, resultTermToDocId *map[string]types.Set, mergeFn merger) map[string]types.Set {
	for _, internalQuery := range queries {
		x := make(map[string]types.Set)
		x = internalQuery.Apply(termToDocId, &x)
		for term, docSet := range x {
			if _, ok := (*resultTermToDocId)[term]; !ok {
				(*resultTermToDocId)[term] = types.NewHashSet()
			}

			(*resultTermToDocId)[term] = (*resultTermToDocId)[term].Union(docSet)
		}
	}

	var resultDocSet types.Set
	for _, docSet := range *resultTermToDocId {
		if resultDocSet == nil {
			resultDocSet = docSet
			continue
		}

		resultDocSet = mergeFn(resultDocSet, docSet)
	}

	for term := range *resultTermToDocId {
		(*resultTermToDocId)[term] = resultDocSet
	}

	return *resultTermToDocId
}
