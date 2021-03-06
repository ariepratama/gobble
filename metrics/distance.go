package metrics

import (
	"goir/types"
	"math"
)

type DistanceFn int

const (
	EUCLIDEAN DistanceFn = iota
	COSINE
)

// SetEuclidean calculate SetEuclidean distance between 2 word sets.
// this function will not take into account word counts
func SetEuclidean(wordSet1 types.Set, wordSet2 types.Set) float64 {
	return math.Sqrt(float64(wordSet1.Difference(wordSet2).Length()))
}

func WvEuclidean(wv1 types.WordVector, wv2 types.WordVector) float64 {
	return SetEuclidean(wv1.WordSet(), wv2.WordSet())
}

// WvCosineSimilarity calculate Cosine distance
func WvCosineSimilarity(wv1 types.WordVector, wv2 types.WordVector) float64 {
	return wv1.Dot(wv2) / (wv1.VectorLength() * wv2.VectorLength())
}
