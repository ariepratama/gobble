package indexes

import (
	"goir/core"
	"goir/metrics"
	"goir/queries"
	"goir/repositories"
	"testing"
)

func TestInMemoryTermToDocIndex_Add(t *testing.T) {
	repository := repositories.NewInMemoryDocumentRepository()
	index := NewInMemoryTermToDocIndex(repository)
	doc1 := core.NewSimpleDocument(1, "bohemian rhapsody")
	doc2 := core.NewSimpleDocument(2, "love of my life")

	index.Add(doc1)
	index.Add(doc2)

	if index.DocLength() != 2 {
		t.Error("TermLength should be 2")
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
