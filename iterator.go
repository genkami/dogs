package dogs

// Iterable iterates over some set of elements.
type Iterator[T any] interface {
	// Next returns the next element in this Iterable and advances its state.
	// The second return value is false if and only if there are no elements to return.
	Next() (T, bool)
}

// SliceFromIterator converts an Iterator[T] into []T.
func SliceFromIterator[T any](it Iterator[T]) []T {
	return Fold[[]T, T](
		make([]T, 0),
		it,
		func(xs []T, x T) []T { return append(xs, x) },
	)
}

// TODO: Find(it, fn)
// TODO: FindIndex(it, fn)
// TODO: FindElem(it, e, eq)
// TODO: FindElemIndex(it, e, eq)
// TODO: Take(it, n)
// TODO: Drop(it, n)

// Map(it, f) returns an iterator that applies f to each element of it.
func Map[T, U any](it Iterator[T], fn func(T) U) Iterator[U] {
	return &mapIterator[T, U]{
		it: it,
		fn: fn,
	}
}

type mapIterator[T, U any] struct {
	it Iterator[T]
	fn func(T) U
}

func (it *mapIterator[T, U]) Next() (U, bool) {
	x, ok := it.it.Next()
	if !ok {
		var zero U
		return zero, false
	}
	return it.fn(x), true
}

// Fold accumulates every element in Iterator by applying fn.
func Fold[T, U any](init T, it Iterator[U], fn func(T, U) T) T {
	var acc T = init
	for {
		x, ok := it.Next()
		if !ok {
			break
		}
		acc = fn(acc, x)
	}
	return acc
}

// Zip combines two Iterators that yields pairs of corresponding elements.
func Zip[T, U any](a Iterator[T], b Iterator[U]) Iterator[Pair[T, U]] {
	return &zipIterator[T, U]{
		a: a,
		b: b,
	}
}

type zipIterator[T, U any] struct {
	a Iterator[T]
	b Iterator[U]
}

func (it *zipIterator[T, U]) Next() (Pair[T, U], bool) {
	x, ok := it.a.Next()
	if !ok {
		return Pair[T, U]{}, false
	}
	y, ok := it.b.Next()
	if !ok {
		return Pair[T, U]{}, false
	}
	return Pair[T, U]{x, y}, true
}


// Unfold returns an Iterator `it` that has an initial state `init` and updating function `step`.
// On each call to `it.Next()`, it updates its internal state by applying `step` and return the second return value.
// If the third return value of `step` is `false`, `it.Next()` stops iterating and returns `(<zero value>, false)`.
func Unfold[T, U any](init T, step func(T) (T, U, bool)) Iterator[U] {
	return &unfoldIterator[T, U]{
		state: init,
		step: step,
		finished: false,
	}
}

type unfoldIterator[T, U any] struct {
	state T
	step func(T) (T, U, bool)
	finished bool
}

func (it *unfoldIterator[T, U]) Next() (U, bool) {
	var zero U
	if it.finished {
		return zero, false
	}
	state, next, ok := it.step(it.state)
	if !ok {
		it.finished = true
		return zero, false
	}
	it.state = state
	return next, true
}

// SumWithInit sums up `init` and all values in `it`.
func SumWithInit[T any](init T, it Iterator[T], s Semigroup[T]) T {
	return Fold[T, T](init, it, s.Combine)
}

// Sum sums up all values in `it`.
// It returns `Empty()` when `it` is empty.
func Sum[T any](it Iterator[T], m Monoid[T]) T {
	var s Semigroup[T] = m
	return SumWithInit(m.Empty(), it, s)
}

