package cmp

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Eq defines equality of type T.
type Eq[T any] interface {
	// Equal returns true if and only if given two arguments are the same.
	Equal(T, T) bool
}

// DefaultEq is Eq with default implementations.
type DefaultEq[T any] struct {
	EqualImpl func(T, T) bool
}

func (eq *DefaultEq[T]) Equal(x, y T) bool {
	return eq.EqualImpl(x, y)
}

// DeriveEq derives Eq using standard `==` operator.
func DeriveEq[T comparable]() Eq[T] {
	return derivedEq[T]{}
}

type derivedEq[T comparable] struct{}

func (derivedEq[T]) Equal(x, y T) bool {
	return x == y
}

// Ord defines order of type T.
type Ord[T any] interface {
	// Compare(x, y) returns:
	//     LT if x < y
	//     EQ if x == y
	//     GT if x > y
	Compare(T, T) Ordering

	// Lt(x, y) means x < y.
	Lt(T, T) bool

	// Le(x, y) means x <= y.
	Le(T, T) bool

	// Gt(x, y) means x > y.
	Gt(T, T) bool

	// Ge(x, y) means x >= y.
	Ge(T, T) bool

	// Eq(x, y) means x == y.
	Eq(T, T) bool

	// Ne(x, y) means x != y.
	Ne(T, T) bool
}

// DeriveOrd derives Ord using `<`, `<=`, `>`, `>=`, `==`, and `!=`.
func DeriveOrd[T constraints.Ordered]() Ord[T] {
	return derivedOrd[T]{}
}

type derivedOrd[T constraints.Ordered] struct{}

func (derivedOrd[T]) Compare(x, y T) Ordering {
	if x < y {
		return LT
	} else if x == y {
		return EQ
	} else {
		return GT
	}
}

func (derivedOrd[T]) Lt(x, y T) bool {
	return x < y
}

func (derivedOrd[T]) Le(x, y T) bool {
	return x <= y
}

func (derivedOrd[T]) Gt(x, y T) bool {
	return x > y
}

func (derivedOrd[T]) Ge(x, y T) bool {
	return x >= y
}

func (derivedOrd[T]) Eq(x, y T) bool {
	return x == y
}

func (derivedOrd[T]) Ne(x, y T) bool {
	return x != y
}

// DefaultOrd is Ord with default implementations.
type DefaultOrd[T any] struct {
	CompareImpl func(T, T) Ordering
}

func (o *DefaultOrd[T]) Compare(x, y T) Ordering {
	return o.CompareImpl(x, y)
}

func (o *DefaultOrd[T]) Lt(x, y T) bool {
	result := o.Compare(x, y)
	return result == LT
}

func (o *DefaultOrd[T]) Le(x, y T) bool {
	result := o.Compare(x, y)
	return result == LT || result == EQ
}

func (o *DefaultOrd[T]) Gt(x, y T) bool {
	result := o.Compare(x, y)
	return result == GT
}

func (o *DefaultOrd[T]) Ge(x, y T) bool {
	result := o.Compare(x, y)
	return result == GT || result == EQ
}

func (o *DefaultOrd[T]) Eq(x, y T) bool {
	result := o.Compare(x, y)
	return result == EQ
}

func (o *DefaultOrd[T]) Ne(x, y T) bool {
	result := o.Compare(x, y)
	return result != EQ
}

type Ordering int

const (
	LT Ordering = iota
	EQ
	GT
)

func (o Ordering) GoString() string {
	switch o {
	case LT:
		return "LT"
	case EQ:
		return "EQ"
	case GT:
		return "GT"
	default:
		return fmt.Sprintf("<unknown Ordering (%d)>", o)
	}
}
