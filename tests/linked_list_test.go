package tests

import (
	"testing"

	"github.com/Zueberberg/collections"
)

var ints = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Test_linked_list_basics(t *testing.T) {
	list := &collections.LinkedList[int]{}
	list.Add(0)

	var (
		last_val int
		left_val int
	)

	for _, i := range ints {
		list.Add(i)
		last_val = list.Last.Value
		left_val = list.Last.Left.Value
		if last_val-1 != left_val {
			t.Errorf("Add() failed: last %d != left %d, %#v", last_val, left_val, list)
			return
		}
	}
	if list.Length != 11 {
		t.Error("LinkedList length expected to be 11:", list.Length, list)
	}
}

func Test_linked_list_find(t *testing.T) {
	list := &collections.LinkedList[int]{}
	for _, i := range ints {
		list.Add(i)
	}

	n := list.Find(0)
	if n != nil {
		t.Error("LinkedList.Find(0) finded Non-Exists value:", n)
	}
	n = list.Find(5)
	if n == nil {
		t.Error("LinkedList.Find(5) nil node finded with existing value 5:", list)
	} else if n.Value != 5 {
		t.Error("LinkedList.Find(5) finded incorrect node with value:", n)
	}
}

func Test_linked_list_findfunc(t *testing.T) {
	type Bar struct {
		Value int
	}
	type Foo struct {
		ID  int
		Bar Bar
	}
	list := &collections.LinkedList[Foo]{}

	for i := 1; i < 100; i++ {
		f := Foo{
			ID: i,
			Bar: Bar{
				Value: i * 100,
			},
		}
		list.Add(f)
	}

	n := list.FindFunc(func(n *collections.Node[Foo]) bool {
		if n.Value.Bar.Value == 5000 {
			return true
		}
		return false
	})
	if n == nil {
		t.Error("LinkedList.FindFunc() not finded node Foo with id 50 and value Bar.Value 5000")
	} else if n.Value.ID != 50 {
		t.Error("LinkedList().FindFunc() incorrect node finded:", n)
	} else if n.Left.Value.ID != 49 || n.Right.Value.ID != 51 {
		t.Error("LinkedList.FindFunc() node 50 bad sibling nodes:", n.Left, n.Right)
	}
}
