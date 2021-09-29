package dogs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/genkami/dogs"
)

func TestSlice_Iter(t *testing.T) {
	subject := func(xs []string) dogs.Iterator[string] {
		return dogs.Slice[string](xs).Iter()
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

