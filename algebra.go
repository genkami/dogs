package dogs

import "constraints"

// Additive is a type that can use `+` operator.
type Additive interface {
	constraints.Integer | constraints.Float | constraints.Complex | ~string
}

// Semigroup is a set of type `T` and its associative binary operation `Combine(T, T) T`
type Semigroup[T any] interface {
	Combine(T, T) T
}

// DeriveAdditiveSemigroup derives Semigroup using `+` operator.
func DeriveAdditiveSemigroup[T Additive]() Semigroup[T] {
	return additiveSemigroup[T]{}
}

type additiveSemigroup[T Additive] struct {}

func (additiveSemigroup[T]) Combine(x, y T) T {
	return x + y
}

// DefaultSemigroup is a default implementation of Semigroup.
type DefaultSemigroup[T any] struct {
	CombineImpl func(T, T) T
}

func (s *DefaultSemigroup[T]) Combine(x, y T) T {
	return s.CombineImpl(x, y)
}

// Monoid is a Semigroup with identity.
type Monoid[T any] interface {
	Semigroup[T]
	Empty() T
}

// DefaultMonoid is a default implementation of Monoid.
type DefaultMonoid[T any] struct {
	Semigroup[T]
	EmptyImpl func() T
}

func (m *DefaultMonoid[T]) Empty() T {
	return m.EmptyImpl()
}