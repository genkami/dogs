package channel_test

import (
	"github.com/genkami/dogs/types/channel"
	"github.com/genkami/dogs/types/iterator"
	"github.com/genkami/dogs/types/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromIterator(t *testing.T) {
	elems := slice.Slice[int]{1, 2, 3}
	ch := channel.FromIterator[int](elems.Iter())

	for _, actual := range elems {
		expected, ok := <-ch
		assert.True(t, ok)
		assert.Equal(t, expected, actual)
	}
}

func TestChannel_Iter(t *testing.T) {
	subject := func(xs []string) iterator.Iterator[string] {
		ch := make(chan string)
		go func() {
			for _, x := range xs {
				ch <- x
			}
			close(ch)
		}()
		return channel.Chan[string](ch).Iter()
	}

	t.Run("empty", func(t *testing.T) {
		it := subject([]string{})
		_, ok := it.Next()
		assert.False(t, ok)
	})

	t.Run("singleton", func(t *testing.T) {
		it := subject([]string{"hoge"})
		x, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, x, "hoge")
		_, ok = it.Next()
		assert.False(t, ok)
	})

	t.Run("multiple elements", func(t *testing.T) {
		it := subject([]string{"hoge", "fuga", "foo"})
		x, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, x, "hoge")
		x, ok = it.Next()
		assert.True(t, ok)
		assert.Equal(t, x, "fuga")
		x, ok = it.Next()
		assert.True(t, ok)
		assert.Equal(t, x, "foo")
		_, ok = it.Next()
		assert.False(t, ok)
	})
}
