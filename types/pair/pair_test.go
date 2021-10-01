package pair_test

import (
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/classes/cmp"
	"github.com/genkami/dogs/types/pair"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPair_Values(t *testing.T) {
	a, b := pair.Pair[int, string]{123, "abc"}.Values()
	assert.Equal(t, a, 123)
	assert.Equal(t, b, "abc")
}

func TestDerivePairEq(t *testing.T) {
	e := pair.DerivePairEq(cmp.DeriveEq[int](), cmp.DeriveEq[string]())
	pair := func(x int, y string) pair.Pair[int, string] {
		return pair.Pair[int, string]{x, y}
	}
	assert.True(t, e.Equal(pair(1, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(1, "fuga"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "fuga"), pair(1, "hoge")))
}

func TestDerivePtrPairEq(t *testing.T) {
	e := pair.DerivePtrPairEq(cmp.DeriveEq[int](), cmp.DeriveEq[string]())
	pair := func(x int, y string) *pair.Pair[int, string] {
		return &pair.Pair[int, string]{x, y}
	}
	assert.True(t, e.Equal(pair(1, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(1, "fuga"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "fuga"), pair(1, "hoge")))
}

func TestDerivePairOrd(t *testing.T) {
	subject := pair.DerivePairOrd(cmp.DeriveOrd[int](), cmp.DeriveOrd[string]()).Compare
	pair := func(x int, y string) pair.Pair[int, string] {
		return pair.Pair[int, string]{x, y}
	}

	assert.Equal(t, subject(pair(1, "hoge"), pair(1, "hoge")), cmp.EQ)

	assert.Equal(t, subject(pair(1, "hoga"), pair(1, "hoge")), cmp.LT)
	assert.Equal(t, subject(pair(0, "hoge"), pair(1, "hoge")), cmp.LT)
	assert.Equal(t, subject(pair(0, "hoga"), pair(1, "hoge")), cmp.LT)

	assert.Equal(t, subject(pair(1, "hogz"), pair(1, "hoge")), cmp.GT)
	assert.Equal(t, subject(pair(2, "hoge"), pair(1, "hoge")), cmp.GT)
	assert.Equal(t, subject(pair(2, "hogz"), pair(1, "hoge")), cmp.GT)
}

func TestDerivePtrPairOrd(t *testing.T) {
	subject := pair.DerivePtrPairOrd(cmp.DeriveOrd[int](), cmp.DeriveOrd[string]()).Compare
	pair := func(x int, y string) *pair.Pair[int, string] {
		return &pair.Pair[int, string]{x, y}
	}

	assert.Equal(t, subject(pair(1, "hoge"), pair(1, "hoge")), cmp.EQ)

	assert.Equal(t, subject(pair(1, "hoga"), pair(1, "hoge")), cmp.LT)
	assert.Equal(t, subject(pair(0, "hoge"), pair(1, "hoge")), cmp.LT)
	assert.Equal(t, subject(pair(0, "hoga"), pair(1, "hoge")), cmp.LT)

	assert.Equal(t, subject(pair(1, "hogz"), pair(1, "hoge")), cmp.GT)
	assert.Equal(t, subject(pair(2, "hoge"), pair(1, "hoge")), cmp.GT)
	assert.Equal(t, subject(pair(2, "hogz"), pair(1, "hoge")), cmp.GT)
}

func TestDerivePairSemigroup(t *testing.T) {
	s := pair.DerivePairSemigroup(
		algebra.DeriveAdditiveSemigroup[int](),
		algebra.DeriveAdditiveSemigroup[string](),
	)
	pair := func(x int, y string) pair.Pair[int, string] {
		return pair.Pair[int, string]{x, y}
	}
	assert.Equal(t, s.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
}

func TestDerivePtrPairSemigroup(t *testing.T) {
	s := pair.DerivePtrPairSemigroup(
		algebra.DeriveAdditiveSemigroup[int](),
		algebra.DeriveAdditiveSemigroup[string](),
	)
	pair := func(x int, y string) *pair.Pair[int, string] {
		return &pair.Pair[int, string]{x, y}
	}
	assert.Equal(t, s.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
}

func TestDerivePairMonoid(t *testing.T) {
	m := pair.DerivePairMonoid(
		algebra.DeriveAdditiveMonoid[int](),
		algebra.DeriveAdditiveMonoid[string](),
	)
	pair := func(x int, y string) pair.Pair[int, string] {
		return pair.Pair[int, string]{x, y}
	}
	assert.Equal(t, m.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
	assert.Equal(t, m.Combine(m.Empty(), pair(123, "foo")), pair(123, "foo"))
}

func TestDerivePtrPairMonoid(t *testing.T) {
	m := pair.DerivePtrPairMonoid(
		algebra.DeriveAdditiveMonoid[int](),
		algebra.DeriveAdditiveMonoid[string](),
	)
	pair := func(x int, y string) *pair.Pair[int, string] {
		return &pair.Pair[int, string]{x, y}
	}
	assert.Equal(t, m.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
	assert.Equal(t, m.Combine(m.Empty(), pair(123, "foo")), pair(123, "foo"))
}
