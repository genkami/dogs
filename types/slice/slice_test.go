package slice_test

import (
	"github.com/genkami/dogs/classes/cmp"
	"github.com/genkami/dogs/types/iterator"
	"github.com/genkami/dogs/types/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromIterator(t *testing.T) {
	subject := func(xs []int) []int {
		return slice.FromIterator(slice.Slice[int](xs).Iter())
	}
	assert.Equal(t, subject([]int{}), []int{})
	assert.Equal(t, subject([]int{1}), []int{1})
	assert.Equal(t, subject([]int{1, 2, 3}), []int{1, 2, 3})
}

func TestSlice_Iter(t *testing.T) {
	subject := func(xs []string) iterator.Iterator[string] {
		return slice.Slice[string](xs).Iter()
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

func TestSlice_Sort(t *testing.T) {
	ord := cmp.DeriveOrd[int]()
	subject := func(xs []int) []int {
		slice.Sort[int](slice.Slice[int](xs), ord)
		return xs
	}

	assert.Equal(t, subject([]int{}), []int{})
	assert.Equal(t, subject([]int{1}), []int{1})

	assert.Equal(t, subject([]int{1, 2}), []int{1, 2})
	assert.Equal(t, subject([]int{2, 1}), []int{1, 2})

	assert.Equal(t, subject([]int{3, 5, 2, 1, 4}), []int{1, 2, 3, 4, 5})
}
