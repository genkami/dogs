package pair

import (
	"fmt"
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/classes/cmp"
)

// Pair is a pair of two values.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Values returns its values.
func (p Pair[T, U]) Values() (T, U) {
	return p.First, p.Second
}

// DerivePairEq derives Eq[Pair[T, U]] from Eq[T] and Eq[U].
func DerivePairEq[T, U any](et cmp.Eq[T], eu cmp.Eq[U]) cmp.Eq[Pair[T, U]] {
	return &cmp.DefaultEq[Pair[T, U]]{
		EqualImpl: func(p, q Pair[T, U]) bool {
			return et.Equal(p.First, q.First) && eu.Equal(p.Second, q.Second)
		},
	}
}

// DerivePtrPairEq derives Eq[*Pair[T, U]] from Eq[T] and Eq[U].
func DerivePtrPairEq[T, U any](et cmp.Eq[T], eu cmp.Eq[U]) cmp.Eq[*Pair[T, U]] {
	return &cmp.DefaultEq[*Pair[T, U]]{
		EqualImpl: func(p, q *Pair[T, U]) bool {
			return et.Equal(p.First, q.First) && eu.Equal(p.Second, q.Second)
		},
	}
}

// DerivePairOrd derives Ord[Pair[T, U]] from Ord[T] and Ord[U].
func DerivePairOrd[T, U any](ot cmp.Ord[T], ou cmp.Ord[U]) cmp.Ord[Pair[T, U]] {
	return &cmp.DefaultOrd[Pair[T, U]]{
		CompareImpl: func(p, q Pair[T, U]) cmp.Ordering {
			v := ot.Compare(p.First, q.First)
			switch v {
			case cmp.LT:
				return cmp.LT
			case cmp.EQ:
				return ou.Compare(p.Second, q.Second)
			case cmp.GT:
				return cmp.GT
			}
			panic(fmt.Errorf("unknown Ordering: %d", v))
		},
	}
}

// DerivePtrPairOrd derives Ord[*Pair[T, U]] from Ord[T] and Ord[U].
func DerivePtrPairOrd[T, U any](ot cmp.Ord[T], ou cmp.Ord[U]) cmp.Ord[*Pair[T, U]] {
	return &cmp.DefaultOrd[*Pair[T, U]]{
		CompareImpl: func(p, q *Pair[T, U]) cmp.Ordering {
			v := ot.Compare(p.First, q.First)
			switch v {
			case cmp.LT:
				return cmp.LT
			case cmp.EQ:
				return ou.Compare(p.Second, q.Second)
			case cmp.GT:
				return cmp.GT
			}
			panic(fmt.Errorf("unknown Ordering: %d", v))
		},
	}
}

// DerivePairSemigroup derives algebra.Semigroup[Pair[T, U]] from algebra.Semigroup[T] and algebra.Semigroup[U].
func DerivePairSemigroup[T, U any](st algebra.Semigroup[T], su algebra.Semigroup[U]) algebra.Semigroup[Pair[T, U]] {
	return &algebra.DefaultSemigroup[Pair[T, U]]{
		CombineImpl: func(p, q Pair[T, U]) Pair[T, U] {
			return Pair[T, U]{
				First:  st.Combine(p.First, q.First),
				Second: su.Combine(p.Second, q.Second),
			}
		},
	}
}

// DerivePtrPairSemigroup derives algebra.Semigroup[*Pair[T, U]] from algebra.Semigroup[T] and algebra.Semigroup[U].
func DerivePtrPairSemigroup[T, U any](st algebra.Semigroup[T], su algebra.Semigroup[U]) algebra.Semigroup[*Pair[T, U]] {
	return &algebra.DefaultSemigroup[*Pair[T, U]]{
		CombineImpl: func(p, q *Pair[T, U]) *Pair[T, U] {
			return &Pair[T, U]{
				First:  st.Combine(p.First, q.First),
				Second: su.Combine(p.Second, q.Second),
			}
		},
	}
}

// DerivePairMonoid derives Monoid[Pair[T, U]] from Monoid[T] and Monoid[U].
func DerivePairMonoid[T, U any](mt algebra.Monoid[T], mu algebra.Monoid[U]) algebra.Monoid[Pair[T, U]] {
	var st algebra.Semigroup[T] = mt
	var su algebra.Semigroup[U] = mu
	return &algebra.DefaultMonoid[Pair[T, U]]{
		Semigroup: DerivePairSemigroup[T, U](st, su),
		EmptyImpl: func() Pair[T, U] {
			return Pair[T, U]{
				First:  mt.Empty(),
				Second: mu.Empty(),
			}
		},
	}
}

// DerivePtrPairMonoid derives Monoid[Pair[T, U]] from Monoid[T] and Monoid[U].
func DerivePtrPairMonoid[T, U any](mt algebra.Monoid[T], mu algebra.Monoid[U]) algebra.Monoid[*Pair[T, U]] {
	var st algebra.Semigroup[T] = mt
	var su algebra.Semigroup[U] = mu
	return &algebra.DefaultMonoid[*Pair[T, U]]{
		Semigroup: DerivePtrPairSemigroup[T, U](st, su),
		EmptyImpl: func() *Pair[T, U] {
			return &Pair[T, U]{
				First:  mt.Empty(),
				Second: mu.Empty(),
			}
		},
	}
}
