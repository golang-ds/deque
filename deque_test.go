package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	deque := New[int]()

	s := deque.Size()
	assert.Equal(t, 0, s)

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddFirst(3)
	deque.AddFirst(4)

	s = deque.Size()
	assert.Equal(t, 4, s)
}

func TestIsEmpty(t *testing.T) {
	deque := New[int]()

	assert.True(t, deque.IsEmpty())

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddFirst(3)
	deque.AddFirst(4)

	assert.False(t, deque.IsEmpty())
}

func TestFirst(t *testing.T) {
	deque := New[int]()

	f, ok := deque.First()
	assert.False(t, ok)
	assert.Equal(t, 0, f)

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddFirst(3)
	deque.AddFirst(4)

	f, ok = deque.First()
	assert.True(t, ok)
	assert.Equal(t, 4, f)
}

func TestLast(t *testing.T) {
	deque := New[int]()

	l, ok := deque.Last()
	assert.False(t, ok)
	assert.Equal(t, 0, l)

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddFirst(3)
	deque.AddFirst(4)

	l, ok = deque.Last()
	assert.True(t, ok)
	assert.Equal(t, 1, l)
}

func TestAddFirst(t *testing.T) {
	deque := New[int]()

	f, ok := deque.First()
	assert.False(t, ok)
	assert.Equal(t, 0, f)

	deque.AddFirst(1)
	deque.AddFirst(2)

	f, ok = deque.First()
	assert.True(t, ok)
	assert.Equal(t, 2, f)

	deque.AddFirst(3)
	deque.AddFirst(4)

	f, ok = deque.First()
	assert.True(t, ok)
	assert.Equal(t, 4, f)
}

func TestAddLast(t *testing.T) {
	deque := New[int]()

	l, ok := deque.Last()
	assert.False(t, ok)
	assert.Equal(t, 0, l)

	deque.AddLast(1)
	deque.AddLast(2)

	l, ok = deque.Last()
	assert.True(t, ok)
	assert.Equal(t, 2, l)

	deque.AddLast(3)
	deque.AddLast(4)

	l, ok = deque.Last()
	assert.True(t, ok)
	assert.Equal(t, 4, l)
}

func TestRemoveFirst(t *testing.T) {
	deque := New[int]()

	f, ok := deque.RemoveFirst()
	assert.False(t, ok)
	assert.Equal(t, 0, f)

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddFirst(3)
	deque.AddFirst(4)

	f, ok = deque.RemoveFirst()
	assert.True(t, ok)
	assert.Equal(t, 4, f)
}

func TestRemoveLast(t *testing.T) {
	deque := New[int]()

	l, ok := deque.RemoveLast()
	assert.False(t, ok)
	assert.Equal(t, 0, l)

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddFirst(3)
	deque.AddFirst(4)

	l, ok = deque.RemoveLast()
	assert.True(t, ok)
	assert.Equal(t, 1, l)
}

func TestString(t *testing.T) {
	deque := New[int]()

	s := deque.String()
	assert.Equal(t, "[ ]", s)

	deque.AddFirst(1)
	deque.AddFirst(2)
	deque.AddFirst(3)
	deque.AddLast(10)
	deque.AddLast(20)
	deque.AddFirst(4)

	s = deque.String()
	assert.Equal(t, "[ 4 3 2 1 10 20 ]", s)
}
