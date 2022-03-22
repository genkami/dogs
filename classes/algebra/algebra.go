package algebra

import "golang.org/x/exp/constraints"

// Additive is a type that can use `+` operator.
type Additive interface {
	constraints.Integer | constraints.Float | constraints.Complex | ~string
}

// Multiplicative is a type that can use `*` operator.
type Multiplicative interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Semigroup is a set of type `T` and its associative binary operation `Combine(T, T) T`
type Semigroup[T any] interface {
	Combine(T, T) T
}

// DeriveAdditiveSemigroup derives Semigroup using `+` operator.
func DeriveAdditiveSemigroup[T Additive]() Semigroup[T] {
	return additiveSemigroup[T]{}
}

type additiveSemigroup[T Additive] struct{}

func (additiveSemigroup[T]) Combine(x, y T) T {
	return x + y
}

// DeriveMultiplicativeSemigroup derives Semigroup using `*` operator.
func DeriveMultiplicativeSemigroup[T Multiplicative]() Semigroup[T] {
	return multiplicativeSemigroup[T]{}
}

type multiplicativeSemigroup[T Multiplicative] struct{}

func (multiplicativeSemigroup[T]) Combine(x, y T) T {
	return x * y
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

// DeriveAdditiveMonoid derives Monoid using `+` and zero value.
func DeriveAdditiveMonoid[T Additive]() Monoid[T] {
	return additiveMonoid[T]{}
}

type additiveMonoid[T Additive] struct{}

func (additiveMonoid[T]) Combine(x, y T) T {
	return x + y
}

func (additiveMonoid[T]) Empty() (zero T) {
	return
}

// DeriveMultiplicativeMonoid derives Monoid using `*` and `1`.
func DeriveMultiplicativeMonoid[T Multiplicative]() Monoid[T] {
	return multiplicativeMonoid[T]{}
}

type multiplicativeMonoid[T Multiplicative] struct{}

func (multiplicativeMonoid[T]) Combine(x, y T) T {
	return x * y
}

func (multiplicativeMonoid[T]) Empty() T {
	return 1
}

// DefaultMonoid is a default implementation of Monoid.
type DefaultMonoid[T any] struct {
	Semigroup[T]
	EmptyImpl func() T
}

func (m *DefaultMonoid[T]) Empty() T {
	return m.EmptyImpl()
}
