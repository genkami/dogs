package dogs

import "fmt"

// Pair is a pair of two values.
type Pair[T, U any] struct {
	First T
	Second U
}

// DerivePairEq derives Eq[Pair[T, U]] from Eq[T] and Eq[U].
func DerivePairEq[T, U any](et *Eq[T], eu *Eq[U]) *Eq[Pair[T, U]] {
	return &Eq[Pair[T, U]]{
		Equal: func(p, q Pair[T, U]) bool {
			return et.Equal(p.First, q.First) && eu.Equal(p.Second, q.Second)
		},
	}
}

// DerivePtrPairEq derives Eq[*Pair[T, U]] from Eq[T] and Eq[U].
func DerivePtrPairEq[T, U any](et *Eq[T], eu *Eq[U]) *Eq[*Pair[T, U]] {
	return &Eq[*Pair[T, U]]{
		Equal: func(p, q *Pair[T, U]) bool {
			return et.Equal(p.First, q.First) && eu.Equal(p.Second, q.Second)
		},
	}
}

// DerivePairOrd derives Ord[Pair[T, U]] from Ord[T] and Ord[U].
func DerivePairOrd[T, U any](ot *Ord[T], ou *Ord[U]) *Ord[Pair[T, U]] {
	return &Ord[Pair[T, U]]{
		Compare: func(p, q Pair[T, U]) Ordering {
			v := ot.Compare(p.First, q.First)
			switch v {
			case LT:
				return LT
			case EQ:
				return ou.Compare(p.Second, q.Second)
			case GT:
				return GT
			}
			panic(fmt.Errorf("unknown Ordering: %d", v))
		},
	}
}

// DerivePtrPairOrd derives Ord[*Pair[T, U]] from Ord[T] and Ord[U].
func DerivePtrPairOrd[T, U any](ot *Ord[T], ou *Ord[U]) *Ord[*Pair[T, U]] {
	return &Ord[*Pair[T, U]]{
		Compare: func(p, q *Pair[T, U]) Ordering {
			v := ot.Compare(p.First, q.First)
			switch v {
			case LT:
				return LT
			case EQ:
				return ou.Compare(p.Second, q.Second)
			case GT:
				return GT
			}
			panic(fmt.Errorf("unknown Ordering: %d", v))
		},
	}
}

// DerivePairSemigroup derives Semigroup[Pair[T, U]] from Semigroup[T] and Semigroup[U].
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

// DerivePtrPairSemigroup derives Semigroup[*Pair[T, U]] from Semigroup[T] and Semigroup[U].
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

// DerivePairMonoid derives Monoid[Pair[T, U]] from Monoid[T] and Monoid[U].
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

// DerivePtrPairMonoid derives Monoid[Pair[T, U]] from Monoid[T] and Monoid[U].
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