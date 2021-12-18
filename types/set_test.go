package types

import (
	"testing"
)

func TestHashSet_Keys(t *testing.T) {
	const expectedKeyLen = 3
	set := NewHashSet()
	set.Put("c")
	set.Put("a")
	set.Put("a")
	set.Put("b")
	keys := set.Keys()

	if !set.Contains("a") || !set.Contains("b") || !set.Contains("c") {
		t.Error("one of the expected key is missing!")
	}

	if len(keys) != expectedKeyLen {
		t.Errorf("key is more than %d, content=%wv", expectedKeyLen, keys)
	}

}

func TestHashSet_Intersect(t *testing.T) {
	set1 := NewHashSet()
	set1.Put("a")
	set1.Put("b")
	set1.Put("c")

	set2 := NewHashSet()
	set2.Put("b")
	set2.Put("c")

	intersectedSet := set1.Intersect(set2)

	if intersectedSet.Contains("a") {
		t.Errorf("%wv should not in intersectedSet", "a")
	}

	for _, val := range []string{"b", "c"} {
		if !intersectedSet.Contains(val) {
			t.Errorf("%wv should be in intersectedSet", val)
		}
	}
}

func TestHashSet_Union(t *testing.T) {
	set1 := NewHashSet()
	set1.Put("a")
	set1.Put("b")
	set1.Put("c")

	set2 := NewHashSet()
	set2.Put("d")
	set2.Put("e")

	union := set1.Union(set2)
	if len(union.Keys()) < 5 {
		t.Error("lack of keys")
	}
}

func TestHashSet_Difference(t *testing.T) {
	set1 := NewHashSet()
	set1.Put("a")
	set1.Put("b")
	set1.Put("c")

	set2 := NewHashSet()
	set2.Put("d")
	set2.Put("c")

	difference := set1.Difference(set2)

	if difference.Length() != 3 {
		t.Error("wrong difference length")
	}

	for _, k := range []string{"a", "b", "d"} {
		if !difference.Contains(k) {
			t.Errorf("should have contain %wv", k)
		}
	}

}

func TestHashSet_Length(t *testing.T) {
	set := NewHashSet()
	set.Put("a")
	set.Put("b")

	if set.Length() != 2 {
		t.Error("length is not correct")
	}
}
