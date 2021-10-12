package list

import "github.com/genkami/dogs/types/iterator"

// List is a linked list.
// An empty list is represented as nil.
type List[T any] struct {
	Head T
	Tail *List[T]
}

//go:generate gotip run ../../cmd/gen-functions -template Collection -pkg list -name *List -out zz_generated.collection.go
//go:generate gotip fmt ./zz_generated.collection.go

func New[T any](xs ...T) *List[T] {
	var list *List[T] = nil
	for i := len(xs) - 1; i >= 0; i-- {
		list = &List[T]{xs[i], list}
	}
	return list
}

func FromIterator[T any](it iterator.Iterator[T]) *List[T] {
	var head, tail *List[T]
	for {
		x, ok := it.Next()
		if !ok {
			break
		}
		if head == nil {
			head = &List[T]{x, nil}
			tail = head
		} else {
			tail.Tail = &List[T]{x, nil}
			tail = tail.Tail
		}
	}
	return head
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
