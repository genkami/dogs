package slice

import (
	"github.com/genkami/dogs/classes/cmp"
	"github.com/genkami/dogs/types/iterator"
	"sort"
)

// Slice is a slice with extra methods.
type Slice[T any] []T

// FromIterator builds a Slice from given Iterator.
func FromIterator[T any](it iterator.Iterator[T]) Slice[T] {
	return Slice[T](iterator.Fold[[]T, T](
		make([]T, 0),
		it,
		func(xs []T, x T) []T { return append(xs, x) },
	))
}

//go:generate go run ../../cmd/gen-functions -template Collection -pkg slice -name Slice -out zz_generated.collection.go
//go:generate go run ../../cmd/gen-functions -template OrderedCollection -pkg slice -name Slice -out zz_generated.ordered_collection.go
//go:generate go fmt .

// Iter returns an Iterator that iterates over given slice.
func (xs Slice[T]) Iter() iterator.Iterator[T] {
	return &sliceIterator[T]{
		xs:   ([]T)(xs),
		next: 0,
	}
}

type sliceIterator[T any] struct {
	xs   []T
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

// Sort destructively sorts `xs` using `Ord`.
func Sort[T any](xs Slice[T], o cmp.Ord[T]) {
	sort.Slice(([]T)(xs), func(i, j int) bool {
		return o.Lt(xs[i], xs[j])
	})
}
