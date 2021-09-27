package dogs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genkami/dogs"
)

var (
	intSemigroup = &dogs.Semigroup[int]{
		Combine: func(x, y int) int { return x + y },
	}
	intMonoid = &dogs.Monoid[int]{
		Semigroup: *intSemigroup,
		Empty: func() int { return 0 },
	}
)

func TestSemigroup_SumWithInit(t *testing.T) {
	subject := func(x int, xs []int) int {
		return intSemigroup.SumWithInit(x, xs)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{}), 0)
		assert.Equal(t, subject(1, []int{}), 1)
		assert.Equal(t, subject(999, []int{}), 999)	
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1}), 1)
		assert.Equal(t, subject(1, []int{2}), 3)
		assert.Equal(t, subject(123, []int{456}), 579)
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1, 2, 3}), 6)
		assert.Equal(t, subject(1, []int{10, 100, 1000}), 1111)
	})
}

func TestMonoid_Sum(t *testing.T) {
	subject := func(xs []int) int {
		return intMonoid.Sum(xs)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject([]int{}), 0)
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject([]int{0}), 0)
		assert.Equal(t, subject([]int{1}), 1)
		assert.Equal(t, subject([]int{999}), 999)
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject([]int{1, 2}), 3)
		assert.Equal(t, subject([]int{1, 2, 3}), 6)
		assert.Equal(t, subject([]int{1, 10, 100, 1000}), 1111)
	})
}
