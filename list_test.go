package dogs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genkami/dogs"
)

func TestNewList(t *testing.T) {
	assert.Equal(t, dogs.NewList[int](), (*dogs.List[int])(nil))
	assert.Equal(t, dogs.NewList[int](1), &dogs.List[int]{1, nil})
	assert.Equal(t, dogs.NewList[int](1, 2), &dogs.List[int]{1, &dogs.List[int]{2, nil}})
	assert.Equal(t, dogs.NewList[int](1, 2, 3), &dogs.List[int]{1, &dogs.List[int]{2, &dogs.List[int]{3, nil}}})
}

func TestList_Iter(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		it := dogs.NewList[string]().Iter()
		_, ok := it.Next()
		assert.False(t, ok)
	})

	t.Run("singleton", func(t *testing.T) {
		it := dogs.NewList[string]("hoge").Iter()
		x, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, x, "hoge")
		_, ok = it.Next()
		assert.False(t, ok)
	})

	t.Run("multiple elements", func(t *testing.T) {
		it := dogs.NewList[string]("hoge", "fuga", "foo").Iter()
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