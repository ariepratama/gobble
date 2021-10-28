package metrics

import (
	"goir/types/set"
	"math"
)

func Euclidean(wordSet1 set.Set, wordSet2 set.Set) float64 {
	return math.Sqrt(float64(wordSet1.Difference(wordSet2).Length()))
}
