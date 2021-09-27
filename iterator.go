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

// NewSliceIterator returns an Iterator that iterates over given slice.
func NewSliceIterator[T any](xs []T) Iterator[T] {
	return &sliceIterator[T]{
		xs: xs,
		next: 0,
	}
}

type sliceIterator[T any] struct {
	xs []T
	next int
}


func (it *sliceIterator[T]) Next() (T, bool) {
	if len(it.xs) <= it.next {
		var zero T
		return zero, false
	}
	i := it.next
	it.next++
	return it.xs[i], true
}
