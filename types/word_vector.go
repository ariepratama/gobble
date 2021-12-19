package types

import "math"

type WordVector interface {
	Words() []string
	WordSet() Set
	Frequencies() []int
	FrequencyOf(key string) int
	Put(key string, count int) WordVector
	Inc(key string) WordVector
	Contains(key string) bool
	Length() int
	Dot(otherWv WordVector) float64
	VectorLength() float64
	Copy() WordVector
}

type HashWordVector struct {
	wv map[string]int
}

func NewHashWordVector() HashWordVector {
	return HashWordVector{
		make(map[string]int),
	}
}

func NewHashWordVectorFromToken(tokens []string) HashWordVector {
	wv := NewHashWordVector()
	for _, token := range tokens {
		wv.Inc(token)
	}
	return wv
}

func (wv HashWordVector) Words() []string {
	keys := make([]string, len(wv.wv))
	i := 0
	for key := range wv.wv {
		keys[i] = key
		i++
	}
	return keys
}

func (wv HashWordVector) Frequencies() []int {
	values := make([]int, len(wv.wv))
	i := 0
	for _, val := range wv.wv {
		values[i] = val
		i++
	}
	return values
}

func (wv HashWordVector) FrequencyOf(key string) int {
	return wv.wv[key]
}

func (wv HashWordVector) Put(key string, count int) WordVector {
	wv.wv[key] = count
	return wv
}

func (wv HashWordVector) Inc(key string) WordVector {
	if _, ok := wv.wv[key]; !ok {
		wv.wv[key] = 0
	}

	wv.wv[key] = wv.wv[key] + 1
	return wv
}

func (wv HashWordVector) Contains(key string) bool {
	_, ok := wv.wv[key]
	return ok
}

func (wv HashWordVector) Length() int {
	return len(wv.wv)
}

func (wv HashWordVector) Dot(otherWv WordVector) float64 {
	sumDotProduct := 0
	for key, val := range wv.wv {
		if otherWv.Contains(key) {
			sumDotProduct = sumDotProduct + val*otherWv.FrequencyOf(key)
		}
	}
	return float64(sumDotProduct)
}

func (wv HashWordVector) VectorLength() float64 {
	sumValues := float64(0)
	for _, val := range wv.Frequencies() {
		sumValues = sumValues + math.Pow(float64(val), 2)
	}
	return math.Sqrt(sumValues)
}

func (wv HashWordVector) Copy() WordVector {
	newWv := NewHashWordVector()
	for k, v := range wv.wv {
		newWv.Put(k, v)
	}
	return newWv
}

func (wv HashWordVector) WordSet() Set {
	return NewHashSetFromWords(wv.Words())
}
