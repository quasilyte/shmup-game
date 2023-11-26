package battle

type queue[T any] struct {
	data []T
	head int
	tail int
}

func newQueue[T any](maxSize int) *queue[T] {
	return &queue[T]{
		data: make([]T, maxSize+1),
	}
}

func (q *queue[T]) Cap() int {
	return len(q.data) - 1
}

func (q *queue[T]) Len() int {
	if q.head > q.tail {
		return len(q.data) - q.head + q.tail
	}
	return q.tail - q.head
}

func (q *queue[T]) Peek() T {
	return q.data[q.head]
}

func (q *queue[T]) Push(v T) {
	if q.Len() == q.Cap() {
		panic("queue overflow")
	}

	q.data[q.tail] = v
	q.tail++
	if q.tail >= len(q.data) {
		q.tail = 0
	}
}

func (q *queue[T]) Pop() T {
	v := q.data[q.head]
	q.head++
	if q.head >= len(q.data) {
		q.head = 0
	}
	return v
}
