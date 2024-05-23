package tests

import (
	"testing"

	"github.com/Zueberberg/collections"
)

func Test_FIFO_queue(t *testing.T) {
	q := &collections.FIFOqueue[int]{}

	i, ok := q.Pop()
	if ok != false {
		t.Errorf("FIFOqueue: Non-Exists item not equal to default value: %d != %d", i, 0)
	}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	if q.Len() != 3 {
		t.Error("FIFOqueue length with 3 items not equal to 3", q.Items)
	}
	if q.Items[0] != 1 || q.Items[1] != 2 || q.Items[2] != 3 {
		t.Errorf("FIFOqueue [1, 2, 3] not equal to []int{1, 2, 3}: %#v != %#v", q.Items, []int{1, 2, 3})
	}

	i, _ = q.Pop()
	if i != 1 {
		t.Error("FIFOqueue [1, 2, 3].Pop() item not equal to 1:", i, q.Items)
	}
	if q.Len() != 2 || q.Items[0] != 2 || q.Items[1] != 3 {
		t.Error("FIFOqueue [1, 2, 3].Pop(); After method queue not equal to [2, 3]:", q.Items)
	}
	i, _ = q.Pop()
	if i != 2 {
		t.Error("FIFOqueue [2, 3].Pop() item not equal to 2:", i, q.Items)
	}
	if q.Len() != 1 || q.Items[0] != 3 {
		t.Error("FIFOqueue [2, 3].Pop(); After method queue not equal to [3]:", q.Items)
	}
	i, _ = q.Pop()
	if i != 3 {
		t.Error("FIFOqueue [3].Pop() item not equal to 3:", i, q.Items)
	}
	if q.Len() != 0 {
		t.Error("FIFOqueue [3].Pop(); After method queue not equal to []:", q.Items)
	}
}

func Test_LIFO_queue(t *testing.T) {
	q := &collections.LIFOqueue[int]{}

	i, ok := q.Pop()
	if ok != false {
		t.Errorf("LIFOqueue: Non-Exists item not equal to default value: %d != %d", i, 0)
	}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	if q.Len() != 3 {
		t.Error("LIFOqueue length with 3 items not equal to 3", q.Items)
	}
	if q.Items[0] != 1 || q.Items[1] != 2 || q.Items[2] != 3 {
		t.Errorf("LIFOqueue [1, 2, 3] not equal to []int{1, 2, 3}: %#v != %#v", q.Items, []int{1, 2, 3})
	}

	i, _ = q.Pop()
	if i != 3 {
		t.Error("LIFOqueue [1, 2, 3].Pop() item not equal to 3:", i, q.Items)
	}
	if q.Len() != 2 || q.Items[0] != 1 || q.Items[1] != 2 {
		t.Error("LIFOqueue [1, 2, 3].Pop(); After method queue not equal to [1, 2]:", q.Items)
	}
	i, _ = q.Pop()
	if i != 2 {
		t.Error("LIFOqueue [1, 2].Pop() item not equal to 2:", i, q.Items)
	}
	if q.Len() != 1 || q.Items[0] != 1 {
		t.Error("LIFOqueue [1, 2].Pop(); After method queue not equal to [1]:", q.Items)
	}
	i, _ = q.Pop()
	if i != 1 {
		t.Error("LIFOqueue [1].Pop() item not equal to 1:", i, q.Items)
	}
	if q.Len() != 0 {
		t.Error("FIFOqueue [1].Pop(); After method queue not equal to []:", q.Items)
	}
}

func Test_queue_interface(t *testing.T) {
	fifoQueue := &collections.FIFOqueue[int]{[]int{1, 2, 3}}
	item := foo(fifoQueue)
	if item != 1 {
		t.Error("FIFOqueue [1, 2, 3].Pop() item not equal to 1:", item, fifoQueue.Items)
	}
	lifoQueue := &collections.LIFOqueue[int]{[]int{1, 2, 3}}
	item = foo(lifoQueue)
	if item != 3 {
		t.Error("LIFOqueue [1, 2, 3].Pop() item not equal to 3:", item, lifoQueue.Items)
	}
}

func foo(q collections.Queue[int]) int {
	item, _ := q.Pop()
	return item
}
