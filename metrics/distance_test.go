package metrics

import (
	"fmt"
	"goir/types"
	"math"
	"testing"
)

func TestEuclidean_CompareEmpty(t *testing.T) {
	wordSet1 := types.NewHashSet()
	wordSet2 := types.NewHashSet()
	distance := SetEuclidean(wordSet1, wordSet2)

	if distance != 0 {
		t.Error("it should be 0")
	}

}

func TestEuclidean_CompareWithWords(t *testing.T) {
	const expectedNWordDifferences = 3
	wordSet1 := types.NewHashSet()
	wordSet1.Put("I")
	wordSet1.Put("am")
	wordSet1.Put("inevitable")

	wordSet2 := types.NewHashSet()
	wordSet2.Put("I")
	wordSet2.Put("am")
	wordSet2.Put("Iron")
	wordSet2.Put("Man")

	distance := SetEuclidean(wordSet1, wordSet2)

	fmt.Println("distance between word set: ", distance)
	if distance != math.Sqrt(expectedNWordDifferences) {
		t.Error("distance and expected n word differences is not matched!")
	}

}

func TestWvCosine_SameVectors(t *testing.T) {
	wv1 := types.NewHashWordVector()
	wv1.Put("I", 9938)
	wv1.Put("am", 283)
	wv2 := types.NewHashWordVector()
	wv2.Put("I", 9938)
	wv2.Put("am", 283)

	distance := WvCosineSimilarity(wv1, wv2)
	fmt.Println("distance between 2 wv: ", distance)
	if math.Abs(1-distance) > 0.0001 {
		t.Error("The 2 vectors similarity should be 1")
	}
}

func TestWvCosine_DifferentVectors(t *testing.T) {
	wv1 := types.NewHashWordVector()
	wv1.Put("last", 300)
	wv1.Put("tears", 3764)

	wv2 := types.NewHashWordVector()
	wv2.Put("next", 9834)
	wv2.Put("day", 584)

	distance := WvCosineSimilarity(wv1, wv2)
	fmt.Println("distance between 2 wv: ", distance)
	if math.Abs(1-distance) < 0.01 {
		t.Error("The 2 vectors similarity is 0")
	}

}
