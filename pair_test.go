package dogs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genkami/dogs"
)

func TestDerivePairEq(t *testing.T) {
	e := dogs.DerivePairEq(intEq, stringEq)
	pair := func(x int, y string) dogs.Pair[int, string] {
		return dogs.Pair[int, string]{x, y}
	}
	assert.True(t, e.Equal(pair(1, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(1, "fuga"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "fuga"), pair(1, "hoge")))
}

func TestDerivePtrPairEq(t *testing.T) {
	e := dogs.DerivePtrPairEq(intEq, stringEq)
	pair := func(x int, y string) *dogs.Pair[int, string] {
		return &dogs.Pair[int, string]{x, y}
	}
	assert.True(t, e.Equal(pair(1, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "hoge"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(1, "fuga"), pair(1, "hoge")))
	assert.False(t, e.Equal(pair(2, "fuga"), pair(1, "hoge")))
}

func TestDerivePairSemigroup(t *testing.T) {
	s := dogs.DerivePairSemigroup(intSemigroup, stringSemigroup)
	pair := func(x int, y string) dogs.Pair[int, string] {
		return dogs.Pair[int, string]{x, y}
	}
	assert.Equal(t, s.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
}

func TestDerivePtrPairSemigroup(t *testing.T) {
	s := dogs.DerivePtrPairSemigroup(intSemigroup, stringSemigroup)
	pair := func(x int, y string) *dogs.Pair[int, string] {
		return &dogs.Pair[int, string]{x, y}
	}
	assert.Equal(t, s.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
}

func TestDerivePairMonoid(t *testing.T) {
	m := dogs.DerivePairMonoid(intMonoid, stringMonoid)
	pair := func(x int, y string) dogs.Pair[int, string] {
		return dogs.Pair[int, string]{x, y}
	}
	assert.Equal(t, m.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
	assert.Equal(t, m.Combine(m.Empty(), pair(123, "foo")), pair(123, "foo"))
}

func TestDerivePtrPairMonoid(t *testing.T) {
	m := dogs.DerivePtrPairMonoid(intMonoid, stringMonoid)
	pair := func(x int, y string) *dogs.Pair[int, string] {
		return &dogs.Pair[int, string]{x, y}
	}
	assert.Equal(t, m.Combine(pair(1, "hoge"), pair(2, "fuga")), pair(3, "hogefuga"))
	assert.Equal(t, m.Combine(m.Empty(), pair(123, "foo")), pair(123, "foo"))
}