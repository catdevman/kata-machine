package katamachine

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := new(Queue[int])

	q.enqueue(5)
	q.enqueue(7)
	q.enqueue(9)

	if q.deque() != 5 {

	}

	if q.length != 2 {

	}

	q.enqueue(11)

	if q.deque() != 7 {
	}
	if q.deque() != 9 {
	}
	if q.peek() != 11 {
	}
	if q.deque() != 11 {
	}
	if q.length != 0 {
	}

	// just wanted to make sure that I could not blow up myself when i remove
	// everything
	q.enqueue(69)
	if q.peek() != 69 {
	}
	if q.length != 1 {
	}
}
