package set

import "github.com/genkami/dogs/types/iterator"

// Set is a set of type T.
// https://golang.org/ref/spec#Comparison_operators
type Set[T comparable] map[T]struct{}

//go:generate go run ../../cmd/gen-functions -template Collection -pkg set -name Set -constraint comparable -out zz_generated.collection.go
//go:generate go fmt .

// New creates a new Set.
func New[T comparable](elems ...T) Set[T] {
	s := Set[T]{}
	for _, e := range elems {
		Add[T](s, e)
	}
	return s
}

// Has returns true if and only if s has e.
func Has[T comparable](s Set[T], e T) bool {
	_, ok := s[e]
	return ok
}

// Len returns the number of elements in s.
func Len[T comparable](s Set[T]) int {
	return len(s)
}

// Subset returns true if and only if s is a subset of t.
func Subset[T comparable](s, t Set[T]) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// Equal returns true if and only if the two sets s and t have exactly the same elements.
func Equal[T comparable](s, t Set[T]) bool {
	return Subset(s, t) && Subset(t, s)
}

// Add adds an element e to set s.
func Add[T comparable](s Set[T], e T) {
	s[e] = struct{}{}
}

// Remove removes an element e from set s.
// If returns false if and only if s doesn't have e.
func Remove[T comparable](s Set[T], e T) bool {
	if _, ok := s[e]; !ok {
		return false
	}
	delete(s, e)
	return true
}

// FromIterator returns a Set from given Iterator.
func FromIterator[T comparable](it iterator.Iterator[T]) Set[T] {
	s := Set[T]{}
	for {
		x, ok := it.Next()
		if !ok {
			break
		}
		Add[T](s, x)
	}
	return s
}

// Iter returns an Iterator that iterates over s.
func (s Set[T]) Iter() iterator.Iterator[T] {
	keys := make([]T, 0, len(s))
	for k, _ := range s {
		keys = append(keys, k)
	}
	return iterator.Unfold[int, T](0, func(i int) (int, T, bool) {
		if len(keys) <= i {
			var zero T
			return 0, zero, false
		}
		return i + 1, keys[i], true
	})
}

// TODO: DeriveEq[T] Eq[Set[T]]
// TODO: Elems[T](s Set[T]) []T
// TODO: Merge[T](s, t Set[T]) Set[T]
// TODO: Union[T](s, t Set[T]) Set[T]
// TODO: Intersection[T](s, t Set[T]) Set[T]
