package repositories

import (
	"goir/core"
	"goir/types"
)

type DocumentRepository interface {
	Put(document core.Document)
	Contains(docId int64) bool
	Get(docId int64) core.Document
	GetByIds(docIds types.Set) []core.Document
}

type InMemoryDocumentRepository struct {
	internalRepo map[int64]core.Document
}

func NewInMemoryDocumentRepository() DocumentRepository {
	return InMemoryDocumentRepository{
		make(map[int64]core.Document),
	}
}

func (i InMemoryDocumentRepository) Put(document core.Document) {
	i.internalRepo[document.Id()] = document
}

func (i InMemoryDocumentRepository) Contains(docId int64) bool {
	_, ok := i.internalRepo[docId]
	return ok
}

func (i InMemoryDocumentRepository) Get(docId int64) core.Document {
	return i.internalRepo[docId]
}

func (i InMemoryDocumentRepository) GetByIds(docIds types.Set) []core.Document {
	result := make([]core.Document, docIds.Length())
	idx := 0
	for _, k := range docIds.Keys() {
		result[idx] = i.Get(k.(int64))
		idx++
	}

	return result
}
