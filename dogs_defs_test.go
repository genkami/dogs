package dogs_test

import "github.com/genkami/dogs"

var (
	intSemigroup dogs.Semigroup[int] = &dogs.DefaultSemigroup[int]{
		CombineImpl: func(x, y int) int { return x + y },
	}
	intMonoid dogs.Monoid[int] = &dogs.DefaultMonoid[int]{
		Semigroup: intSemigroup,
		EmptyImpl: func() int { return 0 },
	}

	stringSemigroup dogs.Semigroup[string] = &dogs.DefaultSemigroup[string]{
		CombineImpl: func(x, y string) string { return x + y },
	}
	stringMonoid dogs.Monoid[string] = &dogs.DefaultMonoid[string]{
		Semigroup: stringSemigroup,
		EmptyImpl: func() string { return "" },
	}
)