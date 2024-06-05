package collections

import "fmt"

type BinaryTree[T ordered] struct {
	Value *T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func (t *BinaryTree[T]) Add(item T) (bt *BinaryTree[T], ok bool) {
	switch {
	case t == nil:
		t = &BinaryTree[T]{Value: &item}
		return t, true

	case t.Value == nil:
		t.Value = &item
		return t, true

	case item < *t.Value:
		t.Left, ok = t.Left.Add(item)
		return t, ok

	case item > *t.Value:
		t.Right, ok = t.Right.Add(item)
		return t, ok

	default:
		return t, false
	}
}

func (t *BinaryTree[T]) Find(item T) (bt *BinaryTree[T], ok bool) {
	switch {
	case t == nil || t.Value == nil:
		return nil, false

	case item < *t.Value:
		return t.Left.Find(item)

	case item > *t.Value:
		return t.Right.Find(item)

	default:
		return t, true
	}
}

func (t *BinaryTree[T]) Print(count int) {
	if t.Left != nil {
		t.Left.Print(count + 1)
	}
	for i := 0; i < count; i++ {
		fmt.Print("\t")
	}
	if t == nil || t.Value == nil {
		fmt.Println(nil)
	} else {
		fmt.Println(*t.Value)
	}
	if t.Right != nil {
		t.Right.Print(count + 1)
	}
}
