package collections

/*
Deque - bidirectional queue

Recommended declarations:

	q := &Deque[int]{}
	// q = &{[]} - Empty deque of int type

	q := &Deque[int]{[]int{1, 2, 3}}
	// q = &{[1, 2, 3]} - Deque of int type [1, 2, 3]
*/
type Deque[T comparable] struct {
	Items []T
}

/*
Add element to the start of deque

Method accept argument item of type T,
which is indicated in deque definition.
Method returns index of inserted item.

Example:

	q := &Deque[int]{[]int{2, 3}}
	q.AddLeft(1) // q = [1, 2, 3].
*/
func (q *Deque[T]) AddLeft(item T) int {
	newItems := make([]T, len(q.Items)+1)
	newItems[0] = item
	for index, value := range q.Items {
		newItems[index+1] = value
	}
	q.Items = newItems
	return len(newItems)
}

/*
Pop item from the start of deque.

If deque is non empty, method will be
returning first item from deque and true.

If deque is empty, method will be returning
default value if type [T] and false

Example:

	q := &Deque[int]{[]int{1, 2}}
	i := q.PopLeft()
	i = 1, q = [2]
*/
func (q *Deque[T]) PopLeft() (item T, ok bool) {
	if len(q.Items) == 0 {
		return
	}
	ok = true
	item = q.Items[0]
	q.Items = q.Items[1:]
	return
}

/*
Add item to the end of deque.
It can be also called as 'AddRight'.

Method accept argument item of type T,
which is indicated in deque definition.
Method returns index of inserted item.

Example:

	q := &Deque[int]{[]int{1, 2}}
	q.Add(3) // q = [1, 2, 3]
*/
func (q *Deque[T]) Add(item T) int {
	q.Items = append(q.Items, item)
	return len(q.Items) - 1
}

/*
Pop item from the end of deque.
It can ge also called as 'PopRight'

If deque is non empty, method will be
returning last item from deque and true.

If deque is empty, method will be returning
default value if type [T] and false

Example:

	q := &Deque[int]{[]int{1, 2}}
	i := q.Pop()
	i = 2, q = [1]
*/
func (q *Deque[T]) Pop() (item T, ok bool) {
	if len(q.Items) == 0 {
		return
	}
	ok = true
	item = q.Items[len(q.Items)-1]
	q.Items = q.Items[:len(q.Items)-1]
	return
}

/*
Get deque length
*/
func (q *Deque[T]) Len() int {
	return len(q.Items)
}
