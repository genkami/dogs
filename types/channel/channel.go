package channel

import (
	"github.com/genkami/dogs/types/iterator"
)

// Chan[T] is a channel of type T.
type Chan[T any] chan T

//go:generate go run ../../cmd/gen-functions -template Collection -pkg channel -name Chan -out zz_generated.collection.go
//go:generate go run ../../cmd/gen-functions -template OrderedCollection -pkg channel -name Chan -out zz_generated.ordered_collection.go
//go:generate go fmt .

func FromIterator[T any](it iterator.Iterator[T]) Chan[T] {
	ch := make(chan T)
	go func() {
		for {
			x, ok := it.Next()
			if !ok {
				break
			}
			ch <- x
		}
		close(ch)
	}()
	return Chan[T](ch)
}

func (ch Chan[T]) Iter() iterator.Iterator[T] {
	return &chanIterator[T]{
		ch: ch,
	}
}

type chanIterator[T any] struct {
	ch Chan[T]
}

func (it *chanIterator[T]) Next() (T, bool) {
	v, ok := <-it.ch
	return v, ok
}
