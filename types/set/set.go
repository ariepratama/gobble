package set

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

func New() HashSet {
	return HashSet{
		make(map[string]bool),
	}
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
	var intersectedSet = New()
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
	difference := New()
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
	var newSet = New()
	for k, v := range s.s {
		newSet.s[k] = v
	}
	return newSet
}
