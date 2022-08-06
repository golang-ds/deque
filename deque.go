package deque

import "github.com/golang-ds/linkedlist/doubly"

type Deque[T any] struct {
	data doubly.LinkedList[T]
}

// New constructs and returns an empty deque.
// time-complexity: O(1)
func New[T any]() Deque[T] {
	return Deque[T]{data: doubly.New[T]()}
}

// Size returns the number of the elements in the deque.
// time-complexity: O(1)
func (d *Deque[T]) Size() int {
	return d.data.Size
}

// IsEmpty returns true if the deque doesn't have any elements.
// time-complexity: O(1)
func (d *Deque[T]) IsEmpty() bool {
	return d.Size() == 0
}

// First returns the first element in the deque. It returns false if the deque is empty.
// time-complexity: O(1)
func (d *Deque[T]) First() (T, bool) {
	return d.data.First()
}

// Last returns the last element in the deque. It returns false if the deque is empty.
// time-complexity: O(1)
func (d *Deque[T]) Last() (T, bool) {
	return d.data.Last()
}

// AddFirst adds an element to the front of the deque.
// time-complexity: O(1)
func (d *Deque[T]) AddFirst(data T) {
	d.data.AddFirst(data)
}

// AddLast adds an element to the end of the deque.
// time-complexity: O(1)
func (d *Deque[T]) AddLast(data T) {
	d.data.AddLast(data)
}

// RemoveFirst removes the first element from the deque and returns it. It returns false if the deque is empty.
// time-complexity: O(1)
func (d *Deque[T]) RemoveFirst() (T, bool) {
	return d.data.RemoveFirst()
}

// RemoveLast removes the last element from the deque and returns it. It returns false if the deque is empty.
// time-complexity: O(1)
func (d *Deque[T]) RemoveLast() (T, bool) {
	return d.data.RemoveLast()
}

// String returns the string representation of the deque.
// time-complexity: O(n)
func (d *Deque[T]) String() string {
	return d.data.String()
}
