package dogs_test

import "github.com/genkami/dogs"

var (
	intSemigroup = &dogs.Semigroup[int]{
		Combine: func(x, y int) int { return x + y },
	}
	intMonoid = &dogs.Monoid[int]{
		Semigroup: *intSemigroup,
		Empty: func() int { return 0 },
	}
	intEq dogs.Eq[int] = &dogs.DefaultEq[int]{
		EqualImpl: func(x, y int) bool { return x == y },
	}
	intOrd = &dogs.Ord[int]{
		Compare: func(x, y int) dogs.Ordering {
			if x < y {
				return dogs.LT
			} else if x == y {
				return dogs.EQ
			} else {
				return dogs.GT
			}
		},
	}

	stringSemigroup = &dogs.Semigroup[string]{
		Combine: func(x, y string) string { return x + y },
	}
	stringMonoid = &dogs.Monoid[string]{
		Semigroup: *stringSemigroup,
		Empty: func() string { return "" },
	}
	stringEq dogs.Eq[string] = &dogs.DefaultEq[string]{
		EqualImpl: func(x, y string) bool { return x == y },
	}
	stringOrd = &dogs.Ord[string]{
		Compare: func(x, y string) dogs.Ordering {
			if x < y {
				return dogs.LT
			} else if x == y {
				return dogs.EQ
			} else {
				return dogs.GT
			}
		},
	}
)