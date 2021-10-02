# dogs

![ci status](https://github.com/genkami/dogs/workflows/Test/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/genkami/dogs.svg)](https://pkg.go.dev/github.com/genkami/dogs)

![logo](./doc/logo.png)

Make Go functional with dogs

# Caution
This is a **highly-experimental** package. Any changes will be made in a backward-incompatible manner.

This package will not compile without `gotip` since type parameters are not supported currently in any of Go releases.

Probably **you don't need this** even after type parameters becomes GA. Even if you feel you do, maybe **you shouldn't use this**. It's against Go's philosophy.

# Features

We will continue to implement more utility types and functions.

## Type classes
* Eq
* Ord
* Semigroup
* Monoid

## Data types
* Pair
* List
* Slice
* Map
* Set
* Iterator

# Examples

More examples [here](https://github.com/genkami/dogs/tree/main/examples).

## FizzBuzz

```go
func main() {
	monoid := option.DeriveMonoid[string](algebra.DeriveAdditiveSemigroup[string]())
	fizzBuzz := func(i int) string {
		fizz := option.Filter(option.Some[string]("Fizz"), func(_ string) bool { return i%3 == 0 })
		buzz := option.Filter(option.Some[string]("Buzz"), func(_ string) bool { return i%5 == 0 })
		return option.UnwrapOr(monoid.Combine(fizz, buzz), fmt.Sprint(i))
	}
	it := iterator.Map(iterator.Range[int](1, 15), fizzBuzz)
	iterator.ForEach(it, func(s string) { fmt.Println(s) })
}
```

## Fibonacci

```go
func main() {
	type Pair = pair.Pair[int, int]
	it := iterator.Unfold(
		Pair{1, 1},
		func(p Pair) (Pair, int, bool) {
			a, b := p.Values()
			return Pair{b, a + b}, a, true
		},
	)
	iterator.ForEach(
		iterator.Take(it, 5),
		func(i int) { fmt.Println(i) },
	)
}
```

# Acknowledgements
This library is inspired mainly by:

* [Haskell standard library](https://hackage.haskell.org/package/base)
* [Scala Cats](https://typelevel.org/cats/)

and many other functional languages.


# License

Distributed under the Apache License, Version 2.0. See LICENSE for more information.
