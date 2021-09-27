package dogs

type Pair[T, U any] struct {
	First T
	Second U
}

func DerivePairSemigroup[T, U any](st *Semigroup[T], su *Semigroup[U]) *Semigroup[Pair[T, U]] {
	return &Semigroup[Pair[T, U]]{
		Combine: func(p, q Pair[T, U]) Pair[T, U] {
			return Pair[T, U]{
				First: st.Combine(p.First, q.First),
				Second: su.Combine(p.Second, q.Second),
			}
		},
	}
}

func DerivePtrPairSemigroup[T, U any](st *Semigroup[T], su *Semigroup[U]) *Semigroup[*Pair[T, U]] {
	return &Semigroup[*Pair[T, U]]{
		Combine: func(p, q *Pair[T, U]) *Pair[T, U] {
			return &Pair[T, U]{
				First: st.Combine(p.First, q.First),
				Second: su.Combine(p.Second, q.Second),
			}
		},
	}
}

func DerivePairMonoid[T, U any](mt *Monoid[T], mu *Monoid[U]) *Monoid[Pair[T, U]] {
	return &Monoid[Pair[T, U]]{
		Semigroup: *DerivePairSemigroup(&mt.Semigroup, &mu.Semigroup),
		Empty: func() Pair[T, U] {
			return Pair[T, U]{
				First: mt.Empty(),
				Second: mu.Empty(),
			}
		},
	}
}

func DerivePtrPairMonoid[T, U any](mt *Monoid[T], mu *Monoid[U]) *Monoid[*Pair[T, U]] {
	return &Monoid[*Pair[T, U]]{
		Semigroup: *DerivePtrPairSemigroup(&mt.Semigroup, &mu.Semigroup),
		Empty: func() *Pair[T, U]{
			return &Pair[T, U]{
				First: mt.Empty(),
				Second: mu.Empty(),
			}
		},
	}
}