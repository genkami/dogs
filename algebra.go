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
type Monoid[T any] struct {
	Semigroup[T]
	Empty func() T
}

// Sum sums up all values in `it`.
// It returns `Empty()` when `it` is empty.
// TODO: this should be a function
func (m *Monoid[T]) Sum(it Iterator[T]) T {
	return SumWithInit(m.Empty(), it, m.Semigroup)
}
