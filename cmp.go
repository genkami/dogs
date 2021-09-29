package dogs

import "fmt"

// Eq defines equality of type T.
type Eq[T any] interface {
	// Equal returns true if and only if given two arguments are the same.
	Equal(T, T) bool
}

// DefaultEq is a default implementation of Eq.
type DefaultEq[T any] struct {
	EqualImpl func(T, T) bool
}

func (eq *DefaultEq[T]) Equal(x, y T) bool {
	return eq.EqualImpl(x, y)
}

// TODO: add DeriveEq

// Ord defines order of type T.
type Ord[T any] struct {
	// Compare(x, y) returns:
	//     LT if x < y
	//     EQ if x == y
	//     GT if x > y
	Compare func(T, T) Ordering
}

// Lt(x, y) means x < y.
func (o *Ord[T]) Lt(x, y T) bool {
	result := o.Compare(x, y)
	return result == LT
}

// Le(x, y) means x <= y.
func (o *Ord[T]) Le(x, y T) bool {
	result := o.Compare(x, y)
	return result == LT || result == EQ
}

// Gt(x, y) means x > y.
func (o *Ord[T]) Gt(x, y T) bool {
	result := o.Compare(x, y)
	return result == GT
}

// Ge(x, y) means x >= y.
func (o *Ord[T]) Ge(x, y T) bool {
	result := o.Compare(x, y)
	return result == GT || result == EQ
}

// Eq(x, y) means x == y.
func (o *Ord[T]) Eq(x, y T) bool {
	result := o.Compare(x, y)
	return result == EQ
}

// Ne(x, y) means x != y.
func (o *Ord[T]) Ne(x, y T) bool {
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