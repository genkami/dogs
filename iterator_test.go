package dogs_test

import (
	"strconv"
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
	add := func(x string, y int) string { return x + strconv.FormatInt(int64(y), 10) }
	subject := func(x string, xs []int) string {
		it := dogs.NewSliceIterator(xs)
		return dogs.Fold[string, int](x, it, add)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject("", []int{}), "")
		assert.Equal(t, subject("a", []int{}), "a")
		assert.Equal(t, subject("xyzzy", []int{}), "xyzzy")
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject("", []int{1}), "1")
		assert.Equal(t, subject("a", []int{2}), "a2")
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject("", []int{1, 2, 3}), "123")
		assert.Equal(t, subject("hoge", []int{3, 2, 1, 0}), "hoge3210")
	})
}
