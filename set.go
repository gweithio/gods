package gods

type Set[T comparable] struct {
	set map[T]struct{}
}

func MakeSet[T comparable]() Set[T] {
	set := Set[T]{
		set: make(map[T]struct{}),
	}

	return set
}

func (gs Set[T]) Add(value T) {
	gs.set[value] = struct{}{}
}

func (gs Set[T]) Remove(value T) {
	delete(gs.set, value)
}

func (gs Set[T]) Contains(value T) bool {
	_, ok := gs.set[value]
	return ok
}

type Queue[T comparable] struct {
	values  []T
	start   int
	end     int
	full    bool
	maxSize int
	Size    int
}

func MakeQueue[T comparable](maxSize int) Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}

	q := Queue[T]{}

	q.values = make([]T, q.maxSize, q.maxSize)
	q.start = 0
	q.end = 0
	q.full = false
	q.Size = 0

	return q
}

func (q Queue[T]) Enqueue(value T) {
	if q.Full() {
		q.Dequeue()
	}

	q.values[q.end] = value
	q.end = q.end + 1

	if q.end >= q.maxSize {
		q.end = 0
	}

	if q.end == q.start {
		q.full = true
	}
}

func (q Queue[T]) Dequeue() (value T, ok bool) {
	var zero T
	if q.Size == 0 {
		return zero, false
	}

	value, ok = q.values[q.start], true

	if value != zero {
		q.values[q.start] = zero
		q.start = q.start + 1
		if q.start >= q.maxSize {
			q.start = 0
		}
		q.full = false
	}

	q.Size = q.Size - 1

	return
}

func (q Queue[T]) Full() bool {
	return q.Size == q.maxSize
}
