package cmp_test

import (
	"github.com/genkami/dogs/classes/cmp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeriveEq(t *testing.T) {
	var eq cmp.Eq[int] = cmp.DeriveEq[int]()

	assert.True(t, eq.Equal(1, 1))
	assert.False(t, eq.Equal(1, 2))
}

func TestDeriveOrd(t *testing.T) {
	ord := cmp.DeriveOrd[int]()
	testIntOrd(t, ord)
}

func TestDefaultOrd(t *testing.T) {
	var ord cmp.Ord[int] = &cmp.DefaultOrd[int]{
		CompareImpl: func(x, y int) cmp.Ordering {
			if x < y {
				return cmp.LT
			} else if x == y {
				return cmp.EQ
			} else {
				return cmp.GT
			}
		},
	}
	testIntOrd(t, ord)
}

func testIntOrd(t *testing.T, intOrd cmp.Ord[int]) {
	t.Run("Compare", func(t *testing.T) {
		subject := intOrd.Compare
		assert.Equal(t, subject(1, 1), cmp.EQ)
		assert.Equal(t, subject(1, 2), cmp.LT)
		assert.Equal(t, subject(1, 0), cmp.GT)
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
