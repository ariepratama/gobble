package core

import (
	"goir/preprocessors"
	"goir/types"
)

type Document interface {
	Id() int64
	Terms() types.Set
	WordVector() types.WordVector
}

type SimpleDocument struct {
	id int64
	wv types.WordVector
}

func NewSimpleDocument(id int64, sentence string) Document {
	terms := preprocessors.StandardPreProcessors(sentence)
	return SimpleDocument{
		id,
		types.NewHashWordVectorFromToken(terms),
	}
}

func (s SimpleDocument) Id() int64 {
	return s.id
}

func (s SimpleDocument) Terms() types.Set {
	return types.NewHashSetFromWords(s.wv.Words())
}

func (s SimpleDocument) WordVector() types.WordVector {
	return s.wv
}
