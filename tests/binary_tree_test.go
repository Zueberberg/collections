package tests

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Zueberberg/collections"
)

func assertNode[T int](n *collections.BinaryTree[T], item T) bool {
	if n == nil || n.Value == nil {
		return false
	} else if *n.Value != item {
		return false
	}
	return true
}

func Test_print(t *testing.T) {
	nums := []int{10, 5, 2, 11, 15, 7, 8, 6}
	tree := &collections.BinaryTree[int]{}
	for _, num := range nums {
		tree.Add(num)
	}
	tree.Print(0)
}

func Test_append_lower(t *testing.T) {
	nums := []int{10, 5, 2}
	tree := &collections.BinaryTree[int]{}
	for _, num := range nums {
		tree.Add(num)
	}

	next := tree
	for _, n := range nums {
		ok := assertNode(next, n)
		if !ok {
			t.Errorf("Node value missmatched: %#v, %d\n", next, n)
			return
		}
		next = next.Left
	}
}

func Test_append_higher(t *testing.T) {
	nums := []int{2, 5, 10}
	tree := &collections.BinaryTree[int]{}
	for _, num := range nums {
		tree.Add(num)
	}

	next := tree
	for _, n := range nums {
		ok := assertNode(next, n)
		if !ok {
			t.Errorf("Node value missmatched: %#v, %d\n", next, n)
			return
		}
		next = next.Right
	}
}

func Test_find(t *testing.T) {
	nums := []int{10, 5, 2, 11, 15, 7, 8, 6}
	tree := &collections.BinaryTree[int]{}
	for _, num := range nums {
		tree.Add(num)
	}
	for _, num := range nums {
		_, ok := tree.Find(num)
		if !ok {
			t.Error("Correct num not founded in tree:", num)
		}
	}
	incorrect_nums := []int{101, 100, 102}
	for _, num := range incorrect_nums {
		_, ok := tree.Find(num)
		if ok {
			t.Error("Incorrect num founded in tree:", num)
		}
	}
}

func Test_duplicate_add(t *testing.T) {
	nums := []int{5, 2, 10}
	tree := &collections.BinaryTree[int]{}
	for _, num := range nums {
		_, ok := tree.Add(num)
		if !ok {
			t.Error("Normal num not added:", num)
		}
	}
	for _, num := range nums {
		_, ok := tree.Add(num)
		if ok {
			t.Error("Duplicate num added:", num)
		}
	}
}

func Test_perfomance(t *testing.T) {
	elems := 1_000_000
	t.Log("Using elems:", elems)
	nums := make([]int, elems)
	tree := &collections.BinaryTree[int]{}

	for i := 0; i < elems; i++ {
		rnd := rand.Intn(100_000)
		nums[i] = rnd
		tree.Add(rnd)
	}

	sliceFind := func(arr []int, elem int) bool {
		for _, num := range arr {
			if num == elem {
				return true
			}
		}
		return false
	}

	var start time.Time
	var tend time.Duration
	var send time.Duration

	for i := 0; i < 3; i++ {
		elem := nums[rand.Intn(len(nums))]
		t.Log("Using random num:", elem)

		start = time.Now()
		sliceFind(nums, elem)
		send = time.Since(start)
		t.Log("Slice time spended:", send)

		start = time.Now()
		tree.Find(elem)
		tend = time.Since(start)
		t.Log("BinaryTree time spended:", tend)
		t.Log("BinaryTree faster when Slice:", tend < send)
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < 3; i++ {
		elem := nums[len(nums)-1]
		t.Log("Using last elem of slice:", elem)

		start = time.Now()
		sliceFind(nums, elem)
		send = time.Since(start)
		t.Log("Slice time spended:", send)

		start = time.Now()
		tree.Find(elem)
		tend = time.Since(start)
		t.Log("BinaryTree time spended:", tend)
		t.Log("BinaryTree faster when Slice:", tend < send)
		fmt.Println()
	}
}
