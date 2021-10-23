package iterator

import (
	"constraints"
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/classes/cmp"
	"github.com/genkami/dogs/types/pair"
)

// Iterable iterates over some set of elements.
type Iterator[T any] interface {
	// Next returns the next element in this Iterable and advances its state.
	// The second return value is false if and only if there are no elements to return.
	Next() (T, bool)
}

//go:generate gotip run ../../cmd/gen-functions -template Monad -pkg iterator -name Iterator -out zz_generated.monad.go
//go:generate gotip fmt ./zz_generated.monad.go

// Range returns an Iterator that returns start, start+1, ..., end-1, end, sequentially.
// The returned Iterator does not return any valeus if end is smaller than start.
func Range[T constraints.Integer](start, end T) Iterator[T] {
	return &rangeIterator[T]{
		next: start,
		end:  end,
	}
}

type rangeIterator[T constraints.Integer] struct {
	next, end T
}

func (it *rangeIterator[T]) Next() (T, bool) {
	if it.end < it.next {
		var zero T
		return zero, false
	}
	v := it.next
	it.next++
	return v, true
}

// Find returns a first element in `it` that satisfies the given predicate `fn`.
// It returns `false` as a second return value if no elements are found.
func Find[T any](it Iterator[T], fn func(T) bool) (T, bool) {
	for {
		x, ok := it.Next()
		if !ok {
			var zero T
			return zero, false
		}
		if fn(x) {
			return x, true
		}
	}
}

// FindIndex returns a first index of an element in `it` that satisfies the given predicate `fn`.
// It returns negative value if no elements are found.
func FindIndex[T any](it Iterator[T], fn func(T) bool) int {
	for i := 0; ; i++ {
		x, ok := it.Next()
		if !ok {
			return -1
		}
		if fn(x) {
			return i
		}
	}
}

// FindElem returns a first element in `it` that equals to `e` in the sense of given `Eq`.
// It returns `false` as a second return value if no elements are found.
func FindElem[T any](it Iterator[T], e T, eq cmp.Eq[T]) (T, bool) {
	return Find[T](it, func(x T) bool { return eq.Equal(x, e) })
}

// FindElemIndex returns a first index of an element in `it` that equals to `e` in the sense of given `Eq`.
// It returns negative value if no elements are found.
func FindElemIndex[T any](it Iterator[T], e T, eq cmp.Eq[T]) int {
	return FindIndex[T](it, func(x T) bool { return eq.Equal(x, e) })
}

// Filter returns an Iterator that only returns elements that satisfies given predicate.
func Filter[T any](it Iterator[T], fn func(T) bool) Iterator[T] {
	return &filterIterator[T]{
		it: it,
		fn: fn,
	}
}

type filterIterator[T any] struct {
	it       Iterator[T]
	fn       func(T) bool
	finished bool
}

func (it *filterIterator[T]) Next() (T, bool) {
	var zero T
	if it.finished {
		return zero, false
	}
	for {
		x, ok := it.it.Next()
		if !ok {
			it.finished = true
			return zero, false
		}
		if it.fn(x) {
			return x, true
		}
	}
}

// Taks takes the first n elements in `it`.
func Take[T any](it Iterator[T], n int) Iterator[T] {
	return &takeIterator[T]{
		it: it,
		n:  n,
		i:  0,
	}
}

type takeIterator[T any] struct {
	it   Iterator[T]
	n, i int
}

func (it *takeIterator[T]) Next() (T, bool) {
	var zero T
	if it.n <= it.i {
		return zero, false
	}
	it.i++
	x, ok := it.it.Next()
	if !ok {
		return zero, false
	}
	return x, true
}

// TODO: Drop(it, n)

// Map returns an iterator that applies fn to each element of it.
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

// FlatMap applies fn to each element in it and then joins them.
func FlatMap[T, U any](it Iterator[T], fn func(T) Iterator[U]) Iterator[U] {
	return &flatMapIterator[T, U]{
		it: it,
		fn: fn,
	}
}

type flatMapIterator[T, U any] struct {
	it       Iterator[T]
	cur      Iterator[U]
	fn       func(T) Iterator[U]
	finished bool
}

func (it *flatMapIterator[T, U]) Next() (U, bool) {
	var zero U
	if it.finished {
		return zero, false
	}
	if it.cur != nil {
		x, ok := it.cur.Next()
		if ok {
			return x, true
		}
	}
	// cur == nil or cur is finished
	for {
		next, ok := it.it.Next()
		if !ok {
			it.finished = true
			return zero, false
		}
		it.cur = it.fn(next)
		x, ok := it.cur.Next()
		if ok {
			return x, true
		}
	}
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

// ForEach applies fn to each element in it.
func ForEach[T any](it Iterator[T], fn func(T)) {
	for {
		x, ok := it.Next()
		if !ok {
			return
		}
		fn(x)
	}
}

// Zip combines two Iterators into one that yields pairs of corresponding elements.
func Zip[T, U any](a Iterator[T], b Iterator[U]) Iterator[pair.Pair[T, U]] {
	return &zipIterator[T, U]{
		a: a,
		b: b,
	}
}

type zipIterator[T, U any] struct {
	a Iterator[T]
	b Iterator[U]
}

func (it *zipIterator[T, U]) Next() (pair.Pair[T, U], bool) {
	x, ok := it.a.Next()
	if !ok {
		return pair.Pair[T, U]{}, false
	}
	y, ok := it.b.Next()
	if !ok {
		return pair.Pair[T, U]{}, false
	}
	return pair.Pair[T, U]{x, y}, true
}

// TODO: ZipWith

// Unfold returns an Iterator `it` that has an initial state `init` and updating function `step`.
// On each call to `it.Next()`, it updates its internal state by applying `step` and return the second return value.
// If the third return value of `step` is `false`, `it.Next()` stops iterating and returns `(<zero value>, false)`.
func Unfold[T, U any](init T, step func(T) (T, U, bool)) Iterator[U] {
	return &unfoldIterator[T, U]{
		state:    init,
		step:     step,
		finished: false,
	}
}

type unfoldIterator[T, U any] struct {
	state    T
	step     func(T) (T, U, bool)
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
func SumWithInit[T any](s algebra.Semigroup[T]) func(init T, it Iterator[T]) T {
	return func(init T, it Iterator[T]) T {
		return Fold[T, T](init, it, s.Combine)
	}
}

// Sum sums up all values in `it`.
// It returns `Empty()` when `it` is empty.
func Sum[T any](m algebra.Monoid[T]) func(it Iterator[T]) T {
	return func(it Iterator[T]) T {
		var s algebra.Semigroup[T] = m
		return SumWithInit[T](s)(m.Empty(), it)
	}
}

// TODO: Min
// TODO: MinBy
// TODO: Max
// TODO: MaxBy

// Pure returns an Iterator that contains a single element x.
func Pure[T any](x T) Iterator[T] {
	return &pureIterator[T]{
		x: x,
	}
}

type pureIterator[T any] struct {
	x        T
	finished bool
}

func (it *pureIterator[T]) Next() (T, bool) {
	var zero T
	if it.finished {
		return zero, false
	}
	it.finished = true
	return it.x, true
}

// AndThen composes a monadic value x and action fn.
func AndThen[T, U any](x Iterator[T], fn func(T) Iterator[U]) Iterator[U] {
	return FlatMap(x, fn)
}
