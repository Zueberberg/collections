package collections

import (
	"fmt"
	"strings"
)

var e bool

type Set[T comparable] map[T]struct{}

func (s Set[T]) Exists(a T) bool {
	_, e = s[a]
	return e
}

func (s Set[T]) Add(a T) bool {
	_, e = s[a]
	if e {
		return false
	} else {
		s[a] = struct{}{}
		return true
	}
}

func (s Set[T]) FromVals(a ...T) Set[T] {
	r := Set[T]{}
	for _, v := range a {
		_, e := s[v]
		if !e {
			r[v] = struct{}{}
		}
	}
	return r
}

func (s Set[T]) FromSlice(a []T) Set[T] {
	return s.FromVals(a...)
}

func (s Set[T]) Remove(a T) (k T, e bool) {
	_, e = s[a]
	if !e {
		return k, e
	} else {
		delete(s, a)
		return a, e
	}
}

func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	r := Set[T]{}
	for k := range s {
		_, e := s2[k]
		if e {
			r[k] = struct{}{}
		}
	}
	return r
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	r := Set[T]{}
	for k := range s {
		r[k] = struct{}{}
	}
	for k := range s2 {
		_, e := r[k]
		if !e {
			r[k] = struct{}{}
		}
	}
	return r
}

func (s Set[T]) Subtract(s2 Set[T]) Set[T] {
	r := Set[T]{}
	for k := range s {
		_, e := s2[k]
		if !e {
			r[k] = struct{}{}
		}
	}
	return r
}

func (s Set[T]) XOR(s2 Set[T]) Set[T] {
	r := Set[T]{}
	for k := range s {
		_, e := s2[k]
		if !e {
			r[k] = struct{}{}
		}
	}
	for k := range s2 {
		_, e := s[k]
		if !e {
			r[k] = struct{}{}
		}
	}
	return r
}

func (s Set[T]) ToStr() string {
	vals := make([]string, 0, len(s))
	for k := range s {
		vals = append(vals, fmt.Sprint(k))
	}
	return fmt.Sprintf("{%s}", strings.Join(vals, ", "))
}
