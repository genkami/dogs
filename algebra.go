package dogs

// Semigroup is a set of type `T` and its associative binary operation `Combine(T, T) T`
type Semigroup[T any] struct {
	Combine func(T, T) T
}

// SumWithInit sums up `init` and all values in `it`.
func (s *Semigroup[T]) SumWithInit(init T, it *Iterator[T]) T {
	return Fold[T, T](init, it, s.Combine)
}

// Monoid is a Semigroup with identity.
type Monoid[T any] struct {
	Semigroup[T]
	Empty func() T
}

// Sum sums up all values in `it`.
// It returns `Empty()` when `it` is empty.
func (m *Monoid[T]) Sum(it *Iterator[T]) T {
	return m.SumWithInit(m.Empty(), it)
}
