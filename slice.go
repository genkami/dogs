package dogs

import "github.com/genkami/dogs/types/iterator"

// Slice is a slice with extra methods.
type Slice[T any] []T

// Iter returns an Iterator that iterates over given slice.
func (xs Slice[T]) Iter() iterator.Iterator[T] {
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

// Sort sorts `xs` using `Ord`.
// TODO: ./slice_test.go:58:54: internal compiler error: NewMethodType with type parameters in signature FUNC-method(*struct {}) func(dogs.T₆₄, dogs.T₆₄) bool
// func (xs Slice[T]) Sort(o Ord[T]) {
// 	sort.Slice(([]T)(xs), func(i, j int) bool {
// 		return o.Lt(xs[i], xs[j])
// 	})
// }