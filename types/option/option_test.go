package option_test

import (
	"github.com/genkami/dogs/types/option"
	"github.com/genkami/dogs/types/slice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSome(t *testing.T) {
	x := option.Some[int](123)
	assert.True(t, option.IsSome(x))
	assert.Equal(t, option.Unwrap(x), 123)
}

func TestNone(t *testing.T) {
	x := option.None[int]()
	assert.False(t, option.IsSome(x))
	assert.Panics(t, func() { option.Unwrap(x) })
}

func TestEqual(t *testing.T) {
	assert.True(t, option.Equal(option.None[int](), option.None[int]()))
	assert.False(t, option.Equal(option.Some[int](123), option.None[int]()))
	assert.False(t, option.Equal(option.None[int](), option.Some[int](123)))
	assert.True(t, option.Equal(option.Some[int](123), option.Some[int](123)))
	assert.False(t, option.Equal(option.Some[int](123), option.Some[int](456)))
}

func TestUnwrapOr(t *testing.T) {
	assert.Equal(t, option.UnwrapOr(option.Some[int](123), 456), 123)
	assert.Equal(t, option.UnwrapOr(option.None[int](), 456), 456)
}

func TestUnwrapOrElse(t *testing.T) {
	i := 0
	fn := func() int {
		i++
		return i
	}
	assert.Equal(t, option.UnwrapOrElse(option.Some[int](123), fn), 123)
	assert.Equal(t, option.UnwrapOrElse(option.None[int](), fn), 1)
	assert.Equal(t, option.UnwrapOrElse(option.None[int](), fn), 2)
	assert.Equal(t, option.UnwrapOrElse(option.None[int](), fn), 3)
}

func TestFromIterator(t *testing.T) {
	subject := func(xs ...int) option.Option[int] {
		it := slice.Slice[int](xs).Iter()
		return option.FromIterator(it)
	}
	assert.True(t, option.Equal(subject(), option.None[int]()))
}

func TestIter(t *testing.T) {
	some := func(x int) []int {
		it := option.Some[int](x).Iter()
		return ([]int)(slice.FromIterator(it))
	}
	none := func() []int {
		it := option.None[int]().Iter()
		return ([]int)(slice.FromIterator(it))
	}
	assert.Equal(t, none(), []int{})
	assert.Equal(t, some(1), []int{1})
	assert.Equal(t, some(2), []int{2})
	assert.Equal(t, some(3), []int{3})
}
