package set

import (
	"testing"
)

func TestHashSet_Keys(t *testing.T) {
	const expectedKeyLen = 3
	set := New()
	set.Put("c")
	set.Put("a")
	set.Put("a")
	set.Put("b")
	keys := set.Keys()

	if !set.Contains("a") || !set.Contains("b") || !set.Contains("c") {
		t.Error("one of the expected key is missing!")
	}

	if len(keys) != expectedKeyLen {
		t.Errorf("key is more than %d, content=%s", expectedKeyLen, keys)
	}

}

func TestHashSet_Intersect(t *testing.T) {
	set1 := New()
	set1.Put("a")
	set1.Put("b")
	set1.Put("c")

	set2 := New()
	set2.Put("b")
	set2.Put("c")

	intersectedSet := set1.Intersect(set2)

	if intersectedSet.Contains("a") {
		t.Errorf("%s should not in intersectedSet", "a")
	}

	for _, val := range []string{"b", "c"} {
		if !intersectedSet.Contains(val) {
			t.Errorf("%s should be in intersectedSet", val)
		}
	}
}

func TestHashSet_Union(t *testing.T) {
	set1 := New()
	set1.Put("a")
	set1.Put("b")
	set1.Put("c")

	set2 := New()
	set2.Put("d")
	set2.Put("e")

	union := set1.Union(set2)
	if len(union.Keys()) < 5 {
		t.Error("lack of keys")
	}
}

func TestHashSet_Length(t *testing.T) {
	set := New()
	set.Put("a")
	set.Put("b")

	if set.Length() != 2 {
		t.Error("length is not correct")
	}
}
