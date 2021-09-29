package dogs

// Slice is a slice with extra methods.
type Slice[T any] []T

// Iter returns an Iterator that iterates over given slice.
func (xs Slice[T]) Iter() Iterator[T] {
	return &sliceIterator[T]{
		xs: ([]T)(xs),
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