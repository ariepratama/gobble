package types

type Set interface {
	Keys() []string
	Intersect(other Set) Set
	Put(key string) Set
	Contains(key string) bool
	Union(other Set) Set
	Difference(other Set) Set
	Length() int
	Copy() Set
}

type HashSet struct {
	s map[string]bool
}

func NewHashSet() Set {
	return HashSet{
		make(map[string]bool),
	}
}

func NewHashSetFromWords(words []string) Set {
	newSet := NewHashSet()
	for _, word := range words {
		newSet.Put(word)
	}
	return newSet
}

func (s HashSet) Keys() []string {
	keys := make([]string, len(s.s))
	i := 0
	for key := range s.s {
		keys[i] = key
		i++
	}
	return keys
}

func (s HashSet) Put(key string) Set {
	s.s[key] = true
	return s
}

func (s HashSet) Contains(key string) bool {
	_, ok := s.s[key]
	return ok
}

func (s HashSet) Intersect(other Set) Set {
	var intersectedSet = NewHashSet()
	for _, otherKey := range other.Keys() {
		if s.Contains(otherKey) {
			intersectedSet.Put(otherKey)
		}
	}
	return intersectedSet
}

func (s HashSet) Union(other Set) Set {
	var copiedSet = s.Copy()
	for _, k := range other.Keys() {
		copiedSet.Put(k)
	}
	return copiedSet
}

func (s HashSet) Difference(other Set) Set {
	difference := NewHashSet()
	union := s.Union(other)
	intersected := s.Intersect(other)

	for _, k := range union.Keys() {
		if !intersected.Contains(k) {
			difference.Put(k)
		}
	}

	return difference
}

func (s HashSet) Length() int {
	return len(s.s)
}

func (s HashSet) Copy() Set {
	var newSet = NewHashSet()
	for k, _ := range s.s {
		newSet.Put(k)
	}
	return newSet
}
