package dogs

type List[T any] struct {
	Head T
	Tail *List[T]
}

func NewList[T any](xs ...T) *List[T] {
	var list *List[T] = nil
	for i := len(xs) - 1; i >= 0; i-- {
		list = &List[T]{xs[i], list}
	}
	return list
}

func (xs *List[T]) Iter() *Iterator[T] {
	it := &listIterable[T]{
		cur: xs,
	}
	return NewIterator[T](it)
}

type listIterable[T any] struct {
	cur *List[T]
}

func (it *listIterable[T]) Next() (T, bool) {
	if it.cur == nil {
		var zero T
		return zero, false
	}
	cur := it.cur
	it.cur = it.cur.Tail
	return cur.Head, true
}