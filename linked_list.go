package collections

type Node[T any] struct {
	Left  *Node[T]
	Right *Node[T]
	Value T
}

type LinkedList[T comparable] struct {
	Root   *Node[T]
	Last   *Node[T]
	Length int
}

func (l *LinkedList[T]) Add(item T) *Node[T] {
	l.Length++

	if l.Root == nil {
		l.Root = &Node[T]{
			Left:  nil,
			Right: nil,
			Value: item,
		}
		l.Last = l.Root
		return l.Root
	} else {
		n := &Node[T]{
			Left:  l.Last,
			Right: nil,
			Value: item,
		}
		l.Last.Right = n
		l.Last = n
		return n
	}
}

func (l *LinkedList[T]) Find(item T) *Node[T] {
	var current *Node[T] = l.Root

	for {
		if current != nil {
			if current.Value == item {
				break
			}
			current = current.Right
		} else {
			break
		}
	}
	return current
}

func (l *LinkedList[T]) FindFunc(f func(n *Node[T]) bool) *Node[T] {
	var current *Node[T] = l.Root

	for {
		if current != nil {
			if f(current) {
				break
			}
			current = current.Right
		} else {
			break
		}
	}
	return current
}
