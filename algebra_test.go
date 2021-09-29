package dogs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genkami/dogs"
)

func TestMonoid_Sum(t *testing.T) {
	subject := func(xs []int) int {
		return intMonoid.Sum(dogs.NewSliceIterator(xs))
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
