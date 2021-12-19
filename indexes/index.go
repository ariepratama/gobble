package indexes

import (
	"fmt"
	"goir/core"
	"goir/metrics"
	"goir/queries"
	"goir/repositories"
	"goir/types"
	"goir/utils"
)

type Index interface {
	Add(document core.Document) Index
	Delete(document core.Document) Index
	TermToDocId() map[string]types.Set
	Persist()
	Length() int
	SearchTopK(topK int, queries []queries.Query, distFn metrics.DistanceFn) []core.Document
}

type InMemoryTermToDocIndex struct {
	termToDocId        map[string]types.Set
	docIdToTerm        map[int64]types.Set
	documentRepository repositories.DocumentRepository
}

func NewInMemoryTermToDocIndex(repository repositories.DocumentRepository) InMemoryTermToDocIndex {
	return InMemoryTermToDocIndex{
		make(map[string]types.Set),
		make(map[int64]types.Set),
		repository,
	}
}

func (index InMemoryTermToDocIndex) Add(document core.Document) Index {
	for _, termInt := range document.Terms().Keys() {
		term := fmt.Sprintf("%v", termInt)
		if _, ok := index.termToDocId[term]; !ok {
			index.termToDocId[term] = types.NewHashSet()
		}
		index.termToDocId[term].Put(document.Id())

		if _, ok := index.docIdToTerm[document.Id()]; !ok {
			index.docIdToTerm[document.Id()] = types.NewHashSet()
		}
		index.docIdToTerm[document.Id()].Put(term)
	}

	index.documentRepository.Put(document)
	return index
}

func (index InMemoryTermToDocIndex) Delete(document core.Document) Index {
	panic("implement me")
}

func (index InMemoryTermToDocIndex) TermToDocId() map[string]types.Set {
	return index.termToDocId
}

func (index InMemoryTermToDocIndex) Persist() {
	panic("implement me")
}

func (index InMemoryTermToDocIndex) Length() int {
	return len(index.termToDocId)
}

func (index InMemoryTermToDocIndex) SearchTopK(topK int, queries []queries.Query, distFn metrics.DistanceFn) []core.Document {
	filteredTermToDocId := make(map[string]types.Set)
	queryWv := toWordVector(queries)
	for _, query := range queries {
		filteredTermToDocId = query.Apply(index.termToDocId, &filteredTermToDocId)
	}

	docIdSet := types.NewHashSet()
	for _, docIds := range filteredTermToDocId {
		docIdSet = docIdSet.Union(docIds)
	}

	relevantDocuments := index.documentRepository.GetByIds(docIdSet)

	similarities := make([]float64, len(relevantDocuments))
	i := 0
	for _, relevantDocument := range relevantDocuments {
		similarity := float64(0)
		if distFn == metrics.COSINE {
			similarity = -1 * metrics.WvCosineSimilarity(queryWv, relevantDocument.WordVector())
		}

		if distFn == metrics.EUCLIDEAN {
			similarity = metrics.WvEuclidean(queryWv, relevantDocument.WordVector())
		}

		similarities[i] = similarity
		i++
	}

	slice := utils.NewSlice(similarities)
	relevantDocumentsCopy := make([]core.Document, len(relevantDocuments))
	j := 0
	for _, idx := range slice.Indexes() {
		relevantDocumentsCopy[j] = relevantDocuments[idx]
		j++
	}

	return relevantDocumentsCopy[:minOf(topK, len(relevantDocumentsCopy))]
}

func toWordVector(queries []queries.Query) types.WordVector {
	terms := make([]string, len(queries))
	for idx, query := range queries {
		terms[idx] = query.GetTerm()
	}

	return types.NewHashWordVectorFromToken(terms)
}

func minOf(a int, b int) int {
	if a < b {
		return a
	}

	return b

}
