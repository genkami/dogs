package dogs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/genkami/dogs"
)

func TestDerivePairSemigroup(t *testing.T) {
	s := dogs.DerivePairSemigroup(intSemigroup, stringSemigroup)
	assert.Equal(
		t,
		s.Combine(dogs.Pair[int, string]{1, "hoge"}, dogs.Pair[int, string]{2, "fuga"}),
		dogs.Pair[int, string]{3, "hogefuga"},
	)
}

func TestDerivePtrPairSemigroup(t *testing.T) {
	s := dogs.DerivePtrPairSemigroup(intSemigroup, stringSemigroup)
	assert.Equal(
		t,
		s.Combine(&dogs.Pair[int, string]{1, "hoge"}, &dogs.Pair[int, string]{2, "fuga"}),
		&dogs.Pair[int, string]{3, "hogefuga"},
	)
}

func TestDerivePairMonoid(t *testing.T) {
	m := dogs.DerivePairMonoid(intMonoid, stringMonoid)
	assert.Equal(
		t,
		m.Combine(dogs.Pair[int, string]{1, "hoge"}, dogs.Pair[int, string]{2, "fuga"}),
		dogs.Pair[int, string]{3, "hogefuga"},
	)
	assert.Equal(
		t,
		m.Combine(m.Empty(), dogs.Pair[int, string]{123, "foo"}),
		dogs.Pair[int, string]{123, "foo"},
	)
}

func TestDerivePtrPairMonoid(t *testing.T) {
	m := dogs.DerivePtrPairMonoid(intMonoid, stringMonoid)
	assert.Equal(
		t,
		m.Combine(&dogs.Pair[int, string]{1, "hoge"}, &dogs.Pair[int, string]{2, "fuga"}),
		&dogs.Pair[int, string]{3, "hogefuga"},
	)
	assert.Equal(
		t,
		m.Combine(m.Empty(), &dogs.Pair[int, string]{123, "foo"}),
		&dogs.Pair[int, string]{123, "foo"},
	)
}