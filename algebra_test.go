package dogs_test

import (
	"github.com/genkami/dogs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeriveAdditiveSemigroup(t *testing.T) {
	s := dogs.DeriveAdditiveSemigroup[int]()

	assert.Equal(t, s.Combine(1, 2), 3)
}
