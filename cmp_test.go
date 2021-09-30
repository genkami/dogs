package dogs_test

import (
	"github.com/genkami/dogs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeriveEq(t *testing.T) {
	var eq dogs.Eq[int] = dogs.DeriveEq[int]()

	assert.True(t, eq.Equal(1, 1))
	assert.False(t, eq.Equal(1, 2))
}

func TestDeriveOrd(t *testing.T) {
	ord := dogs.DeriveOrd[int]()
	testIntOrd(t, ord)
}

func TestDefaultOrd(t *testing.T) {
	var ord dogs.Ord[int] = &dogs.DefaultOrd[int]{
		CompareImpl: func(x, y int) dogs.Ordering {
			if x < y {
				return dogs.LT
			} else if x == y {
				return dogs.EQ
			} else {
				return dogs.GT
			}
		},
	}
	testIntOrd(t, ord)
}

func testIntOrd(t *testing.T, intOrd dogs.Ord[int]) {
	t.Run("Compare", func(t *testing.T) {
		subject := intOrd.Compare
		assert.Equal(t, subject(1, 1), dogs.EQ)
		assert.Equal(t, subject(1, 2), dogs.LT)
		assert.Equal(t, subject(1, 0), dogs.GT)
	})

	t.Run("Lt", func(t *testing.T) {
		subject := intOrd.Lt
		assert.True(t, subject(123, 124))
		assert.False(t, subject(123, 123))
		assert.False(t, subject(123, 122))
	})

	t.Run("Le", func(t *testing.T) {
		subject := intOrd.Le
		assert.True(t, subject(123, 124))
		assert.True(t, subject(123, 123))
		assert.False(t, subject(123, 122))
	})

	t.Run("Gt", func(t *testing.T) {
		subject := intOrd.Gt
		assert.False(t, subject(123, 124))
		assert.False(t, subject(123, 123))
		assert.True(t, subject(123, 122))
	})

	t.Run("Ge", func(t *testing.T) {
		subject := intOrd.Ge
		assert.False(t, subject(123, 124))
		assert.True(t, subject(123, 123))
		assert.True(t, subject(123, 122))
	})

	t.Run("Eq", func(t *testing.T) {
		subject := intOrd.Eq
		assert.False(t, subject(123, 124))
		assert.True(t, subject(123, 123))
		assert.False(t, subject(123, 122))
	})

	t.Run("Ne", func(t *testing.T) {
		subject := intOrd.Ne
		assert.True(t, subject(123, 124))
		assert.False(t, subject(123, 123))
		assert.True(t, subject(123, 122))
	})
}
