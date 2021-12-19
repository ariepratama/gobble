package indexes

import (
	"goir/core"
	"goir/metrics"
	"goir/queries"
	"goir/repositories"
	"goir/types"
	"testing"
)

type RepositoryStub struct{}

func newRepositoryStub() repositories.DocumentRepository {
	return RepositoryStub{}
}

func (r RepositoryStub) Put(document core.Document) {
	panic("implement me")
}
func (r RepositoryStub) Contains(docId int64) bool {
	panic("implement me")
}

func (r RepositoryStub) Get(docId int64) core.Document {
	panic("implement me")
}

func (r RepositoryStub) GetByIds(docIds types.Set) []core.Document {
	panic("implement me")
}

func TestInMemoryTermToDocIndex_Add(t *testing.T) {
	repository := newRepositoryStub()
	index := NewInMemoryTermToDocIndex(repository)
	doc1 := core.NewSimpleDocument(1, "bohemian rhapsody")
	doc2 := core.NewSimpleDocument(2, "love of my life")

	index.Add(doc1)
	index.Add(doc2)

	if index.Length() != 2 {
		t.Error("Length should be ")
	}
}

func TestInMemoryTermToDocIndex_SearchTopK(t *testing.T) {
	repository := repositories.NewInMemoryDocumentRepository()
	index := NewInMemoryTermToDocIndex(repository)
	doc1 := core.NewSimpleDocument(1, "bohemian rhapsody")
	doc2 := core.NewSimpleDocument(2, "love of my life")
	doc3 := core.NewSimpleDocument(2, "love of bohemian")

	index.Add(doc1)
	index.Add(doc2)
	index.Add(doc3)

	q := make([]queries.Query, 0)
	q = append(q, queries.NewMatchingQuery("bohemian"))
	result := index.SearchTopK(10, q, metrics.COSINE)

	if len(result) != 2 {
		t.Error("result should be 2")
	}

}
