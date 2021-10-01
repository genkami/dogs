package list

import "github.com/genkami/dogs/types/iterator"

// List is a linked list.
// An empty list is represented as nil.
type List[T any] struct {
	Head T
	Tail *List[T]
}

//go:generate gotip run ../../cmd/gen-collection -pkg list -name *List -out zz_generated.list.go
//go:generate gotip fmt ./zz_generated.list.go

func NewList[T any](xs ...T) *List[T] {
	var list *List[T] = nil
	for i := len(xs) - 1; i >= 0; i-- {
		list = &List[T]{xs[i], list}
	}
	return list
}

func FromIterator[T any](it iterator.Iterator[T]) *List[T] {
	var list *List[T] = nil
	for {
		x, ok := it.Next()
		if !ok {
			break
		}
		list = &List[T]{x, list}
	}
	return list
}

func (xs *List[T]) Iter() iterator.Iterator[T] {
	return &listIterator[T]{
		cur: xs,
	}
}

type listIterator[T any] struct {
	cur *List[T]
}

func (it *listIterator[T]) Next() (T, bool) {
	if it.cur == nil {
		var zero T
		return zero, false
	}
	cur := it.cur
	it.cur = it.cur.Tail
	return cur.Head, true
}
