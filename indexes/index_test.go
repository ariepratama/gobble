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

	index.Add(doc1).Add(doc2)

	if index.DocLength() != 2 {
		t.Error("TermLength should be 2")
	}
}

func TestInMemoryTermToDocIndex_SearchTopK(t *testing.T) {
	repository := repositories.NewInMemoryDocumentRepository()
	index := NewInMemoryTermToDocIndex(repository)
	doc1 := core.NewSimpleDocument(1, "bohemian rhapsody")
	doc2 := core.NewSimpleDocument(2, "love of my life")
	doc3 := core.NewSimpleDocument(3, "love of bohemian")

	index.Add(doc1)
	index.Add(doc2)
	index.Add(doc3)

	result := index.SearchTopK(10, queries.MatchWith("bohemian"), metrics.COSINE)

	if len(result) != 2 {
		t.Error("result should be 2")
	}

}

func TestInMemoryTermToDocIndex_And(t *testing.T) {
	repository := repositories.NewInMemoryDocumentRepository()
	index := NewInMemoryTermToDocIndex(repository)
	doc1 := core.NewSimpleDocument(1, "bohemian rhapsody")
	doc2 := core.NewSimpleDocument(2, "love of my life")
	doc3 := core.NewSimpleDocument(3, "love of bohemian")

	index.Add(doc1)
	index.Add(doc2)
	index.Add(doc3)

	q := queries.And(
		queries.MatchWith("bohemian"),
		queries.MatchWith("love"))
	result := index.SearchTopK(10, q, metrics.COSINE)

	if len(result) != 1 {
		t.Error("results should be 1")
	}

	q = queries.And(
		queries.MatchWith("my"),
		queries.MatchWith("love"))
	result = index.SearchTopK(10, q, metrics.COSINE)

	if len(result) != 1 {
		t.Error("results should be 1")
	}

	q = queries.And(
		queries.MatchWith("love"),
		queries.MatchWith("of"))
	result = index.SearchTopK(10, q, metrics.COSINE)

	if len(result) != 2 {
		t.Error("results should be 2")
	}
}

func TestInMemoryTermToDocIndex_Or(t *testing.T) {
	repository := repositories.NewInMemoryDocumentRepository()
	index := NewInMemoryTermToDocIndex(repository)
	doc1 := core.NewSimpleDocument(1, "bohemian rhapsody")
	doc2 := core.NewSimpleDocument(2, "love of my life")
	doc3 := core.NewSimpleDocument(3, "love of bohemian")

	index.Add(doc1)
	index.Add(doc2)
	index.Add(doc3)

	q := queries.Or(
		queries.MatchWith("bohemian"),
		queries.MatchWith("love"))
	result := index.SearchTopK(10, q, metrics.COSINE)

	if len(result) != 3 {
		t.Error("results should be 3")
	}

	q = queries.Or(
		queries.MatchWith("of"),
		queries.MatchWith("life"))
	result = index.SearchTopK(10, q, metrics.COSINE)

	if len(result) != 2 {
		t.Error("results should be 2")
	}
}

func TestInMemoryTermToDocIndex_Complex(t *testing.T) {
	repository := repositories.NewInMemoryDocumentRepository()
	index := NewInMemoryTermToDocIndex(repository)

	index.Add(core.NewSimpleDocument(1, "bohemian rhapsody"))
	index.Add(core.NewSimpleDocument(2, "love of my life"))
	index.Add(core.NewSimpleDocument(3, "love of bohemian"))
	index.Add(core.NewSimpleDocument(4, "i want to ride my bicycle"))

	q := queries.Or(
		queries.And(queries.MatchWith("bohemian"), queries.MatchWith("rhapsody")),
		queries.And(
			queries.MatchWith("i"),
			queries.MatchWith("want"),
			queries.MatchWith("ride"),
			queries.MatchWith("my"),
			queries.MatchWith("bicycle")))
	result := index.SearchTopK(10, q, metrics.COSINE)

	if len(result) != 2 {
		t.Error("results should be 2")
	}
}
