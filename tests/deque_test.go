package tests

import (
	"testing"

	"github.com/Zueberberg/collections"
)

func Test_AddLeft(t *testing.T) {
	q := &collections.Deque[int]{}
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range ints {
		q.AddLeft(i)
		if q.Items[0] != i {
			t.Errorf("Deque AddLeft() fail: q.Items[0] != %d; %#v", i, q)
			return
		}
	}
	if q.Len() != 10 {
		t.Error("Deque len not equal to 10:", q)
	}
}

func Test_Add(t *testing.T) {
	q := &collections.Deque[int]{}
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range ints {
		q.Add(i)
		index := q.Len() - 1
		if q.Items[index] != i {
			t.Errorf("Deque Add() fail: q.Items[%d] != %d; %#v", index, i, q)
			return
		}
	}
	if q.Len() != 10 {
		t.Error("Deque len not equal to 10:", q)
	}
}

func Test_PopLeft(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	q := &collections.Deque[int]{ints}

	for i := 0; i < len(ints); i++ {
		elem := ints[i]
		qElem, ok := q.PopLeft()
		if !ok {
			t.Error("Iteration:", i)
			t.Error("deque PopLeft() item not exists;", ints[i])
			return
		}
		if elem != qElem {
			t.Error("Iteration:", i)
			t.Error("deque PopLeft() item not equal:", qElem, elem)
			return
		}
	}
	if q.Len() != 0 {
		t.Error("deque PopLeft(): expected to deque to be empty:", q.Items)
	}
}

func Test_Pop(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	q := &collections.Deque[int]{ints}

	for i := len(ints) - 1; i >= 0; i-- {
		elem := ints[i]
		qElem, ok := q.Pop()
		if !ok {
			t.Error("Iteration:", len(ints)-i)
			t.Error("deque PopLeft() item not exists;", ints[i])
			return
		}
		if elem != qElem {
			t.Error("Iteration:", len(ints)-i)
			t.Error("deque PopLeft() item not equal:", qElem, elem)
			return
		}
	}
	if q.Len() != 0 {
		t.Error("deque PopLeft(): expected to deque to be empty:", q.Items)
	}
}
