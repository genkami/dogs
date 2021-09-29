package dogs

// Semigroup is a set of type `T` and its associative binary operation `Combine(T, T) T`
type Semigroup[T any] interface {
	Combine(T, T) T
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