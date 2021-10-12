package option

import (
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/types/iterator"
)

// Option is an optional value.
type Option[T any] struct {
	some bool
	v    T
}

//go:generate gotip run ../../cmd/gen-functions -template Collection -pkg option -name Option -exclude FindElemIndex,FindIndex,Zip -out zz_generated.collection.go
//go:generate gotip fmt ./zz_generated.collection.go

// Some returns an Option that consists of x.
func Some[T any](x T) Option[T] {
	return Option[T]{
		some: true,
		v:    x,
	}
}

// None returns an Option that has no value.
func None[T any]() Option[T] {
	return Option[T]{}
}

// IsSome returns true if and only if x has a value.
func IsSome[T any](x Option[T]) bool {
	return x.some
}

// Equal returns true if and only if the two value x and y are the same, that is,
// both x and y are nil or both have the same value in the sense of ==.
func Equal[T comparable](x, y Option[T]) bool {
	if x.some && y.some {
		return x.v == y.v
	}
	if x.some || y.some {
		return false
	}
	return true
}

// Unwrap returns a value that x has.
// It panics if x has no value.
func Unwrap[T any](x Option[T]) T {
	if !x.some {
		panic("option.Unwrap: None")
	}
	return x.v
}

// UnwrapOr returns a value that x has.
// It returns def if x has no value.
func UnwrapOr[T any](x Option[T], def T) T {
	if !x.some {
		return def
	}
	return x.v
}

// UnwrapOrElse returns a value that x has.
// It returns def() if x has no value.
func UnwrapOrElse[T any](x Option[T], def func() T) T {
	if !x.some {
		return def()
	}
	return x.v
}

// FromIterator returns an Option that has the first element in the given Iterator.
// It returns None() if the iterator has no value.
func FromIterator[T any](it iterator.Iterator[T]) Option[T] {
	if x, ok := it.Next(); ok {
		return Some(x)
	}
	return None[T]()
}

// Iter returns an Iterator that returns a single value that x has, or no value if x is nil.
func (x Option[T]) Iter() iterator.Iterator[T] {
	return &optionIterator[T]{
		opt:      x,
		finished: false,
	}
}

type optionIterator[T any] struct {
	opt      Option[T]
	finished bool
}

func (it *optionIterator[T]) Next() (T, bool) {
	var zero T
	if it.finished {
		return zero, false
	}
	it.finished = true
	if it.opt.some {
		return it.opt.v, true
	}
	return zero, false
}

// TODO: func FromIterator[T any](it iterator.Iterator[T]) Option[T]
// TODO: func (x Option[T]) Iter() iterator.Iterator[T]

// TODO: func Pure[T any](x T) Option[T]
// TODO: func AndThen[T, U any](x T, fn func(T) Option[U]) Option[U]

// TODO: func MapOr[T, U any](x Option[T] fn func(T) U, default U) U
// TODO: func MapOrElse[T, U any](x Option[T], fn func(T) U, default func() U) U
// TODO: func Switch[T any](x Option[T], ifSome func(T) ifNone func())

// TODO: func DeriveEq[T any](eq cmp.Eq[T]) cmp.Eq[Option[T]]
// TODO: func DeriveOrd[T any](ord cmp.Ord[T]) cmp.Ord[Option[T]]

// DeriveSemigroup derives Semigroup[Option[T]] from Semigroup[T]
func DeriveSemigroup[T any](s algebra.Semigroup[T]) algebra.Semigroup[Option[T]] {
	return &algebra.DefaultSemigroup[Option[T]]{
		CombineImpl: func(x, y Option[T]) Option[T] {
			if !IsSome(x) {
				return y
			}
			if !IsSome(y) {
				return x
			}
			return Some(s.Combine(Unwrap(x), Unwrap(y)))
		},
	}
}

// DeriveMonoid derives Monoid[Option[T]] from Semigroup[T].
// T doesn't need to be Monoid since the None() is the identity element of the monoid.
func DeriveMonoid[T any](s algebra.Semigroup[T]) algebra.Monoid[Option[T]] {
	semi := DeriveSemigroup[T](s)
	return &algebra.DefaultMonoid[Option[T]]{
		Semigroup: semi,
		EmptyImpl: func() Option[T] {
			return None[T]()
		},
	}
}
