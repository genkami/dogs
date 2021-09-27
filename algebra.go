package dogs

// Semigroup is a set of type `T` and its associative binary operation `Combine(T, T) T`
type Semigroup[T any] struct {
	Combine func(T, T) T
}

func (s *Semigroup[T]) SumWithInit(init T, xs []T) T {
	acc := init
	for _, x := range xs {
		acc = s.Combine(acc, x)
	}
	return acc
}

// Monoid is a Semigroup with identity.
type Monoid[T any] struct {
	Semigroup[T]
	Empty func() T
}

func (m *Monoid[T]) Sum(xs []T) T {
	return m.SumWithInit(m.Empty(), xs)
}