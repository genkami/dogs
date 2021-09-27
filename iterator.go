package dogs

type Iterator[T any] interface {
	Next() (T, bool)
}

type SliceIterator[T any] struct {
	xs []T
	next int
}

func NewSliceIterator[T any](xs []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		xs: xs,
		next: 0,
	}
}

func (it *SliceIterator[T]) Next() (T, bool) {
	if len(it.xs) <= it.next {
		var zero T
		return zero, false
	}
	i := it.next
	it.next++
	return it.xs[i], true
}