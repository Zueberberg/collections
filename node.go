package collections

type Node[T any] struct {
	Left  *Node[T]
	Right *Node[T]
	Value T
}
