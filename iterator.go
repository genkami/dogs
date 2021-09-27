package dogs

type Iterable[T any] interface {
	Next() (T, bool)
}

type Iterator[T any] struct {
	it Iterable[T]
}

func NewIterator[T any](it Iterable[T]) *Iterator[T] {
	return &Iterator[T]{it}
}

func (it *Iterator[T]) Next() (T, bool) {
	return it.it.Next()
}

func (it *Iterator[T]) Fold(init T, fn func(T, T) T) T {
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

type sliceIterable[T any] struct {
	xs []T
	next int
}

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