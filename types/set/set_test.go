package set_test

import (
	"github.com/genkami/dogs/types/slice"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/genkami/dogs/types/set"
)

func TestHas(t *testing.T) {
	assert.False(t, set.Has(set.New[int](), 1))
	assert.True(t, set.Has(set.New[int](1),1))

	assert.True(t, set.Has(set.New[int](1, 2, 3), 1))
	assert.True(t, set.Has(set.New[int](1, 2, 3), 2))
	assert.True(t, set.Has(set.New[int](1, 2, 3), 3))
	assert.False(t, set.Has(set.New[int](1, 2, 3), 4))
}

func TestLen(t *testing.T) {
	assert.Equal(t, set.Len(set.New[int]()), 0)
	assert.Equal(t, set.Len(set.New[int](5)), 1)
	assert.Equal(t, set.Len(set.New[int](5, 4)), 2)
	assert.Equal(t, set.Len(set.New[int](5, 4, 3)), 3)
	assert.Equal(t, set.Len(set.New[int](5, 4, 4, 3, 3, 3)), 3)
}

func TestSubset(t *testing.T) {
	assert.True(t, set.Subset(set.New[int](), set.New[int]()))

	s := set.New[int](3, 4, 5)
	assert.True(t, set.Subset(set.New[int](), s))
	assert.True(t, set.Subset(set.New[int](3), s))
	assert.True(t, set.Subset(set.New[int](3, 5), s))
	assert.True(t, set.Subset(set.New[int](4, 5), s))
	assert.True(t, set.Subset(set.New[int](3, 4, 5), s))

	assert.False(t, set.Subset(set.New[int](2), s))
	assert.False(t, set.Subset(set.New[int](3, 4, 5, 6), s))
}

func TestEqual(t *testing.T) {
	assert.True(t, set.Equal(set.New[int](), set.New[int]()))
	assert.False(t, set.Equal(set.New[int](), set.New[int](1)))
	assert.False(t, set.Equal(set.New[int](1), set.New[int]()))

	assert.True(t, set.Equal(set.New[int](1, 2, 3), set.New[int](1, 2, 3)))
	assert.False(t, set.Equal(set.New[int](1, 2, 3), set.New[int](1, 2, 4)))
	assert.False(t, set.Equal(set.New[int](1, 2, 3), set.New[int](1, 2)))
	assert.False(t, set.Equal(set.New[int](1, 2, 3), set.New[int](1, 2, 3, 4)))
}

func TestAdd(t *testing.T) {
	s := set.New[int](1, 2, 3)
	assert.True(t, set.Has(s, 1))
	assert.True(t, set.Has(s, 2))
	assert.True(t, set.Has(s, 3))
	assert.False(t, set.Has(s, 4))

	set.Add(s, 4)
	assert.True(t, set.Has(s, 1))
	assert.True(t, set.Has(s, 2))
	assert.True(t, set.Has(s, 3))
	assert.True(t, set.Has(s, 4))
}

func TestRemove(t *testing.T) {
	s := set.New[int](1, 2, 3, 4)
	assert.True(t, set.Has(s, 1))
	assert.True(t, set.Has(s, 2))
	assert.True(t, set.Has(s, 3))
	assert.True(t, set.Has(s, 4))


	ok := set.Remove(s, 4)
	assert.True(t, ok)
	assert.True(t, set.Has(s, 1))
	assert.True(t, set.Has(s, 2))
	assert.True(t, set.Has(s, 3))
	assert.False(t, set.Has(s, 4))

	ok = set.Remove(s, 4)
	assert.False(t, ok)
}

func TestFromIterator(t *testing.T) {
	subject := func(xs ...int) set.Set[int] {
		return set.FromIterator(slice.Slice[int](xs).Iter())
	}
	assert.True(t, set.Equal(subject(), set.New[int]()))
	assert.True(t, set.Equal(subject(1), set.New[int](1)))
	assert.True(t, set.Equal(subject(1, 2), set.New[int](1, 2)))
	assert.True(t, set.Equal(subject(1, 2, 3), set.New[int](1, 2, 3)))
}

func TestSet_Iter(t *testing.T) {
	subject := func(xs ...int) []int {
		return []int(slice.FromIterator(set.New[int](xs...).Iter()))
	}
	assert.ElementsMatch(t, subject(), []int{})
	assert.ElementsMatch(t, subject(1), []int{1})
	assert.ElementsMatch(t, subject(1, 2), []int{1, 2})
	assert.ElementsMatch(t, subject(1, 2, 3), []int{1, 2, 3})
}