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

func TestDefaultOrd_Lt(t *testing.T) {
	subject := intOrd.Lt
	assert.True(t, subject(123, 124))
	assert.False(t, subject(123, 123))
	assert.False(t, subject(123, 122))
}

func TestDefaultOrd_Le(t *testing.T) {
	subject := intOrd.Le
	assert.True(t, subject(123, 124))
	assert.True(t, subject(123, 123))
	assert.False(t, subject(123, 122))
}

func TestDefaultOrd_Gt(t *testing.T) {
	subject := intOrd.Gt
	assert.False(t, subject(123, 124))
	assert.False(t, subject(123, 123))
	assert.True(t, subject(123, 122))
}

func TestDefaultOrd_Ge(t *testing.T) {
	subject := intOrd.Ge
	assert.False(t, subject(123, 124))
	assert.True(t, subject(123, 123))
	assert.True(t, subject(123, 122))
}

func TestDefaultOrd_Eq(t *testing.T) {
	subject := intOrd.Eq
	assert.False(t, subject(123, 124))
	assert.True(t, subject(123, 123))
	assert.False(t, subject(123, 122))
}

func TestDefaultOrd_Ne(t *testing.T) {
	subject := intOrd.Ne
	assert.True(t, subject(123, 124))
	assert.False(t, subject(123, 123))
	assert.True(t, subject(123, 122))
}
