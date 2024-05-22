package collections

import (
	"fmt"
	"strings"
)

var e bool

type Set map[any]struct{}

func (s Set) Exists(a any) bool {
	_, e = s[a]
	return e
}

func (s Set) Add(a any) bool {
	_, e = s[a]
	if e {
		return false
	} else {
		s[a] = struct{}{}
		return true
	}
}

func (s Set) Remove(a any) (any, bool) {
	v, e := s[a]
	if !e {
		return nil, e
	} else {
		delete(s, a)
		return v, e
	}
}

func (s Set) ToStr() string {
	vals := make([]string, 0, len(s))
	for key := range s {
		vals = append(vals, fmt.Sprint(key))
	}
	return fmt.Sprintf("{%s}", strings.Join(vals, ", "))
}
