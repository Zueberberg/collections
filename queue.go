package collections

/*
Basic queue interface
*/
type Queue[T any] interface {
	Add(item T) int
	Pop() (item T, ok bool)
	Len() int
}

/*
FIFO (First in, first out) queue

Recommended declarations:

	q := &FIFOqueue[int]{}
	// q = &{[]} - Empty queue of int type

	q := &FIFOqueue[int]{[]int{1, 2, 3}}
	// q = &{[1, 2, 3]} - Queue of int type [1, 2, 3]
*/
type FIFOqueue[T any] struct {
	Items []T
}

/*
Add item to queue.

Method accept argument item of type T,
which is indicated in queue definition.
Method returns index of inserted item.

Example:

	q := &FIFOqueue[int]{[]int{1, 2}}
	q.Add(3) // q = [1, 2, 3]
*/
func (q *FIFOqueue[T]) Add(item T) int {
	q.Items = append(q.Items, item)
	return len(q.Items) - 1
}

/*
Pop item from queue.

If queue is non empty, method will be
returning first item from queue and true.

If queue is empty, method will be returning
default value if type [T] and false

Example:

	q := &FIFOqueue[int]{[]int{1, 2}}
	i := q.Pop()
	i = 1, q = [2]
*/
func (q *FIFOqueue[T]) Pop() (item T, ok bool) {
	if len(q.Items) == 0 {
		return
	}
	ok = true
	item = q.Items[0]
	q.Items = q.Items[1:]
	return
}

/*
Get queue length
*/
func (q *FIFOqueue[T]) Len() int {
	return len(q.Items)
}

/*
LIFO (Last in - first out) queue

Recommended declarations:

	q := &LIFOqueue[int]{}
	// q = &{[]} - Empty queue of int type

	q := &LIFOqueue[int]{[]int{1, 2, 3}}
	// q = &{[1, 2, 3]} - Queue of int type [1, 2, 3]
*/
type LIFOqueue[T any] struct {
	Items []T
}

/*
Add item to queue.

Method accept argument item of type T,
which is indicated in queue definition.
Method returns index of inserted item.

Example:

	q := &LIFOqueue[int]{[]int{1, 2}}
	q.Add(3) // q = [1, 2, 3]
*/
func (q *LIFOqueue[T]) Add(item T) int {
	q.Items = append(q.Items, item)
	return len(q.Items) - 1
}

/*
Pop item from queue.

If queue is non empty, method will be
returning last item from queue and true.

If queue is empty, method will be returning
default value if type [T] and false

Example:

	q := &LIFOqueue[int]{[]int{1, 2}}
	i := q.Pop()
	i = 2, q = [1]
*/
func (q *LIFOqueue[T]) Pop() (item T, ok bool) {
	if len(q.Items) == 0 {
		return
	}
	ok = true
	item = q.Items[len(q.Items)-1]
	q.Items = q.Items[:len(q.Items)-1]
	return
}

/*
Get queue length
*/
func (q *LIFOqueue[T]) Len() int {
	return len(q.Items)
}
