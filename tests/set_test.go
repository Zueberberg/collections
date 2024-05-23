package tests

import (
	"testing"

	"github.com/Zueberberg/collections"
)

func Test_basic_set(t *testing.T) {
	set := collections.Set[int]{}
	added := set.Add(1)
	if !added {
		t.Error("Value 1 not added to empty set:", set)
	}
	exists := set.Exists(1)
	if !exists {
		t.Error("Value 1 not exists in set {1}", set)
	}
	added = set.Add(1)
	if added {
		t.Error("Value 1 added to set {1}", set)
	}
	v, removed := set.Remove(1)
	if !removed {
		t.Error("Value 1 not removed from {1}", set)
	}
	if v != 1 {
		t.Error("Removed value 1 not equal to 1", v)
	}
	_, removed = set.Remove(1000)
	if removed {
		t.Error("Non-exists value 1000 removed from set", set)
	}
}

func Test_set_many_items(t *testing.T) {
	set := collections.Set[int]{}
	for i := 0; i < 1_000_000; i++ {
		set.Add(i)
	}
	if len(set) != 1_000_000 {
		t.Error("set.Add() << Set length is broken")
	}
	arr := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		arr[i] = i
	}
	s1 := collections.Set[int]{}.FromVals(arr...)
	if len(s1) != 1_000_000 {
		t.Error("set.FromVals() << Set length is broken")
	}
	s2 := collections.Set[int]{}.FromSlice(arr)
	if len(s2) != 1_000_000 {
		t.Error("set.FromSlice() << Set length is broken")
	}
}

func Test_set_intersection(t *testing.T) {
	var (
		s1 collections.Set[int]
		s2 collections.Set[int]
		r  collections.Set[int]
	)
	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{3, 4, 5})
	r = s1.Intersect(s2)
	t.Logf("Intersect: %s & %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 1 || !r.Exists(3) {
		t.Errorf("Intersect failed: %s & %s", s1.ToStr(), s2.ToStr())
	}

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{4, 4, 5})
	r = s1.Intersect(s2)
	t.Logf("Intersect: %s & %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 0 {
		t.Errorf("Intersect failed: %s & %s", s1.ToStr(), s2.ToStr())
	}

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3, 4, 5, 6})
	s2 = collections.Set[int]{}.FromSlice([]int{4, 5, 6, 7, 8, 9, 10})
	r = s1.Intersect(s2)
	t.Logf("Intersect: %s & %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 3 || !r.Exists(4) && !r.Exists(5) && !r.Exists(6) {
		t.Errorf("Intersect failed: %s & %s", s1.ToStr(), s2.ToStr())
	}
}

func Test_set_union(t *testing.T) {
	var (
		s1 collections.Set[int]
		s2 collections.Set[int]
		r  collections.Set[int]
	)

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{2, 3, 4})
	r = s1.Union(s2)
	t.Logf("Union: %s | %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 4 {
		t.Errorf("Union failed: %s | %s", s1.ToStr(), s2.ToStr())
	}

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{4, 5, 6})
	r = s1.Union(s2)
	t.Logf("Union: %s | %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 6 {
		t.Errorf("Union failed: %s | %s", s1.ToStr(), s2.ToStr())
	}
}

func Test_set_subtraction(t *testing.T) {
	var (
		s1 collections.Set[int]
		s2 collections.Set[int]
		r  collections.Set[int]
	)

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{2, 3, 4})
	r = s1.Subtract(s2)
	t.Logf("Subtract: %s - %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 1 {
		t.Errorf("Subtract failed: %s - %s", s1.ToStr(), s2.ToStr())
	}

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{4, 5})
	r = s1.Subtract(s2)
	t.Logf("Subtract: %s - %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 3 {
		t.Errorf("Subtract failed: %s - %s", s1.ToStr(), s2.ToStr())
	}
}

func Test_set_XOR(t *testing.T) {
	var (
		s1 collections.Set[int]
		s2 collections.Set[int]
		r  collections.Set[int]
	)

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{2, 3, 4})
	r = s1.XOR(s2)
	t.Logf("XOR: %s ^ %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 2 {
		t.Errorf("XOR failed: %s ^ %s", s1.ToStr(), s2.ToStr())
	}

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{3, 4, 5})
	r = s1.XOR(s2)
	t.Logf("XOR: %s ^ %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 4 {
		t.Errorf("XOR failed: %s ^ %s", s1.ToStr(), s2.ToStr())
	}

	s1 = collections.Set[int]{}.FromSlice([]int{1, 2, 3})
	s2 = collections.Set[int]{}.FromSlice([]int{6, 4, 5})
	r = s1.XOR(s2)
	t.Logf("XOR: %s ^ %s = %s", s1.ToStr(), s2.ToStr(), r.ToStr())
	if len(r) != 6 {
		t.Errorf("XOR failed: %s ^ %s", s1.ToStr(), s2.ToStr())
	}
}
