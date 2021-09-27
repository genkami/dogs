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

	stringSemigroup = &dogs.Semigroup[string]{
		Combine: func(x, y string) string { return x + y },
	}
	stringMonoid = &dogs.Monoid[string]{
		Semigroup: *stringSemigroup,
		Empty: func() string { return "" },
	}
)