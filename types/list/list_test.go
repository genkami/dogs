package list_test

import (
	"github.com/genkami/dogs/types/list"
	"github.com/genkami/dogs/types/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.Equal(t, list.New[int](), (*list.List[int])(nil))
	assert.Equal(t, list.New[int](1), &list.List[int]{1, nil})
	assert.Equal(t, list.New[int](1, 2), &list.List[int]{1, &list.List[int]{2, nil}})
	assert.Equal(t, list.New[int](1, 2, 3), &list.List[int]{1, &list.List[int]{2, &list.List[int]{3, nil}}})
}

func TestFromIterator(t *testing.T) {
	subject := func(xs ...int) *list.List[int] {
		return list.FromIterator(slice.Slice[int](xs).Iter())
	}

	assert.Equal(t, subject(), list.New[int]())
	assert.Equal(t, subject(1), list.New[int](1))
	assert.Equal(t, subject(2, 3), list.New[int](2, 3))
	assert.Equal(t, subject(4, 5, 6), list.New[int](4, 5, 6))
}

func TestList_Iter(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		it := list.New[string]().Iter()
		_, ok := it.Next()
		assert.False(t, ok)
	})

	t.Run("singleton", func(t *testing.T) {
		it := list.New[string]("hoge").Iter()
		x, ok := it.Next()
		assert.True(t, ok)
		assert.Equal(t, x, "hoge")
		_, ok = it.Next()
		assert.False(t, ok)
	})

	t.Run("multiple elements", func(t *testing.T) {
		it := list.New[string]("hoge", "fuga", "foo").Iter()
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
