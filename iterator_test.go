package dogs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genkami/dogs"
)

func TestSliceIterator_Next(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		it := dogs.NewSliceIterator[string]([]string{})
		_, ok := it.Next()
		assert.False(t, ok)
	})

	t.Run("singleton", func(t *testing.T) {
		it := dogs.NewSliceIterator[string]([]string{"hoge"})
		x, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, x, "hoge")
		_, ok = it.Next()
		assert.False(t, ok)
	})

	t.Run("multiple elements", func(t *testing.T) {
		it := dogs.NewSliceIterator[string]([]string{"hoge", "fuga", "foo"})
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

func TestIterator_Fold(t *testing.T) {
	add := func(x, y int) int { return x + y }
	subject := func(x int, xs []int) int {
		it := dogs.NewSliceIterator(xs)
		return it.Fold(x, add)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{}), 0)
		assert.Equal(t, subject(1, []int{}), 1)
		assert.Equal(t, subject(999, []int{}), 999)
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1}), 1)
		assert.Equal(t, subject(1, []int{2}), 3)
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1, 2, 3}), 6)
		assert.Equal(t, subject(1, []int{10, 100, 1000}), 1111)
	})
}
