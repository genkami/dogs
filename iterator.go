package dogs

// Iterable represents basic methods that are needed to construct Iterator.
type Iterable[T any] interface {
	// Next returns the next element in this Iterable.
	// The second return value is false if and only if there are no elements to return.
	Next() (T, bool)
}

// Iterator is an Iterable with additional utility functions.
// TODO: we don't need this anymore.
type Iterator[T any] struct {
	it Iterable[T]
}

func NewIterator[T any](it Iterable[T]) *Iterator[T] {
	return &Iterator[T]{it}
}

func (it *Iterator[T]) Next() (T, bool) {
	return it.it.Next()
}

// ToSlice converts an Iterator[T] into []T.
func (it *Iterator[T]) ToSlice() []T {
	return Fold[[]T, T](
		make([]T, 0),
		it,
		func(xs []T, x T) []T { return append(xs, x) },
	)
}

// Map(it, f) returns an iterator that applies f to each element of it.
func Map[T, U any](it *Iterator[T], fn func(T) U) *Iterator[U] {
	iterable := &mapIterable[T, U]{
		it: it,
		fn: fn,
	}
	return NewIterator[U](iterable)
}

type mapIterable[T, U any] struct {
	it *Iterator[T]
	fn func(T) U
}

func (it *mapIterable[T, U]) Next() (U, bool) {
	x, ok := it.it.Next()
	if !ok {
		var zero U
		return zero, false
	}
	return it.fn(x), true
}

// Fold accumulates every element in Iterator by applying fn.
func Fold[T, U any](init T, it *Iterator[U], fn func(T, U) T) T {
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
func Zip[T, U any](a *Iterator[T], b *Iterator[U]) *Iterator[Pair[T, U]] {
	it := &zipIterable[T, U]{
		a: a,
		b: b,
	}
	return NewIterator[Pair[T, U]](it)
}

type zipIterable[T, U any] struct {
	a *Iterator[T]
	b *Iterator[U]
}

func (it *zipIterable[T, U]) Next() (Pair[T, U], bool) {
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

type sliceIterable[T any] struct {
	xs []T
	next int
}

// NewSliceIterator returns an Iterator that iterates over given slice.
func NewSliceIterator[T any](xs []T) *Iterator[T] {
	it := &sliceIterable[T]{
		xs: xs,
		next: 0,
	}
	return NewIterator[T](it)
}

func (it *sliceIterable[T]) Next() (T, bool) {
	if len(it.xs) <= it.next {
		var zero T
		return zero, false
	}
	i := it.next
	it.next++
	return it.xs[i], true
}
