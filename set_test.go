package collections_test

import (
	"testing"
)

func Test_basic_set(t *testing.T) {
	t.Log("foo")
	set := Set{}
	t.Log(set.ToStr())
	set.Add(1)
	t.Log(set.ToStr())
}
