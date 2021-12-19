package indexes

import (
	"goir/core"
	"goir/metrics"
	"goir/queries"
)

type Index interface {
	Add(document core.Document) Index
	Delete(document core.Document) Index
	Persist()
	SearchTopK(topK int, queries []queries.Query, distFn metrics.DistanceFn) []core.Document
}

type InMemoryTermToDocIndex struct {
	termToDocId map[string][]int64
	docIdToTerm map[int64][]string
}

func NewInMemoryTermToDocIndex() InMemoryTermToDocIndex {
	return InMemoryTermToDocIndex{
		make(map[string][]int64),
		make(map[int64][]string),
	}
}

func (index InMemoryTermToDocIndex) Add(document core.Document) Index {
	for _, term := range document.Terms().Keys() {
		if _, ok := index.termToDocId[term]; !ok {
			index.termToDocId[term] = make([]int64, 0)
		}
		index.termToDocId[term] = append(index.termToDocId[term], document.Id())

		if _, ok := index.docIdToTerm[document.Id()]; !ok {
			index.docIdToTerm[document.Id()] = make([]string, 0)
		}
		index.docIdToTerm[document.Id()] = append(index.docIdToTerm[document.Id()], term)
	}
	return index
}

func (index InMemoryTermToDocIndex) Delete(document core.Document) Index {
	panic("implement me")
}

func (index InMemoryTermToDocIndex) Persist() {
	panic("implement me")
}

func (index InMemoryTermToDocIndex) SearchTopK(topK int, queries []queries.Query, distFn metrics.DistanceFn) []core.Document {
	panic("implement me")
}
