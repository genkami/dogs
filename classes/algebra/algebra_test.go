package algebra_test

import (
	"github.com/genkami/dogs/classes/algebra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeriveAdditiveSemigroup(t *testing.T) {
	s1 := algebra.DeriveAdditiveSemigroup[int]()
	assert.Equal(t, s1.Combine(1, 2), 3)

	s2 := algebra.DeriveAdditiveSemigroup[string]()
	assert.Equal(t, s2.Combine("a", "b"), "ab")
}

func TestDeriveMultiplicativeSemigroup(t *testing.T) {
	s1 := algebra.DeriveMultiplicativeSemigroup[int]()
	assert.Equal(t, s1.Combine(2, 3), 6)

	s2 := algebra.DeriveMultiplicativeSemigroup[int32]()
	assert.Equal(t, s2.Combine(4, 5), int32(20))
}

func TestDeriveAdditiveMonoid(t *testing.T) {
	s1 := algebra.DeriveAdditiveMonoid[int]()
	assert.Equal(t, s1.Combine(1, 2), 3)
	assert.Equal(t, s1.Combine(s1.Empty(), 4), 4)
	assert.Equal(t, s1.Combine(5, s1.Empty()), 5)

	s2 := algebra.DeriveAdditiveMonoid[string]()
	assert.Equal(t, s2.Combine("a", "b"), "ab")
	assert.Equal(t, s2.Combine(s2.Empty(), "c"), "c")
	assert.Equal(t, s2.Combine("d", s2.Empty()), "d")
}
func TestDeriveMultiplicativeMonoid(t *testing.T) {
	s1 := algebra.DeriveMultiplicativeMonoid[int]()
	assert.Equal(t, s1.Combine(2, 3), 6)
	assert.Equal(t, s1.Combine(s1.Empty(), 4), 4)
	assert.Equal(t, s1.Combine(5, s1.Empty()), 5)

	s2 := algebra.DeriveMultiplicativeMonoid[int32]()
	assert.Equal(t, s2.Combine(6, 7), int32(42))
	assert.Equal(t, s2.Combine(s2.Empty(), 8), int32(8))
	assert.Equal(t, s2.Combine(9, s2.Empty()), int32(9))
}
