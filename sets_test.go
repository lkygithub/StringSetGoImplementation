package sets

import (
    "testing"
)

func TestInsert(t *testing.T) {
	ss := NewString()
	
	// Insert 1 item\
	ss.Insert("a")
	if !ss.Has("a") {
		t.Errorf("Insert failed.")
	}

	// Insert several items
	ss.Insert("a", "bb", "ccc", "dddd")
	if !ss.HasAll("a", "bb", "ccc", "dddd") || ss.Len() != 4 {
		t.Errorf("Insert failed.")
	}

}

func TestDelete(t *testing.T) {
	ss := NewString("a", "b", "c", "d")

	// Delete 1 item
	ss.Delete("a")
	if ss.Has("a") {
		t.Errorf("Delete failed.")
	}

	// Delete several items
	ss.Delete("b", "c", "d")
	if ss.HasAny("b", "c", "d") || ss.Len() != 0 {
		t.Errorf("Delete failed.")
	}
}

func TestHas(t *testing.T) {
	s := "abc"
	ss := NewString(s)
	_, res1 := ss[s]
	res2 := ss.Has(s)
	if (res1 && !res2) || (!res1 && res2) {
		t.Errorf("Has Check failed.")
	}
}

func TestHasAny(t *testing.T) {
	ss := NewString("a", "b", "c", "d")
	_, resA := ss["a"]
	_, resB := ss["b"]
	res1 := resA || resB
	res2 := ss.HasAny("a", "b")
	if (res1 && !res2) || (!res1 && res2) {
		t.Errorf("HasAny Check failed.")
	}

}

func TestHasAll(t *testing.T) {
	ss := NewString("a", "b", "c")
	_, resA := ss["a"]
	_, resB := ss["b"]
	_, resC := ss["c"]
	res1 := resA && resB && resC
	res2 := ss.HasAll("a", "b", "c")
	if (res1 && !res2) || (!res1 && res2) {
		t.Errorf("HasAll Check failed.")
	}
}

func TestDifference(t *testing.T) {
	ss1 := NewString("a", "b")
	ss2 := NewString("a", "c")
	res := ss1.Difference(ss2)
	if !res.Has("b") {
		t.Errorf("Difference Check failed.")
	}
}

func TestUnion(t *testing.T) {
	ss1 := NewString("a")
	ss2 := NewString("b")
	res := ss1.Union(ss2)
	if !res.HasAll("a", "b") {
		t.Errorf("Union Check failed.")
	}
}

func TestIntersection(t *testing.T) {
	ss1 := NewString("a", "b")
	ss2 := NewString("a", "c")
	res := ss1.Intersection(ss2)
	if !res.Has("a") {
		t.Errorf("Intersection Check failed.")
	}
}

func TestLen(t *testing.T) {
	ss1 := NewString("a", "b")
	if ss1.Len() != 2 {
		t.Errorf("Len Check failed.")
	}
}