package option

// Option is an optional value.
type Option[T any] *T

// Some returns an Option that consists of x.
func Some[T any](x T) Option[T] {
	return Option[T](&x)
}

// None returns an Option that has no value.
func None[T any]() Option[T] {
	return nil
}

// IsSome returns true if and only if x has a value.
func IsSome[T any](x Option[T]) bool {
	return x != nil
}

// Equal returns true if and only if the two value x and y are the same, that is,
// both x and y are nil or both have the same value in the sense of ==.
func Equal[T comparable](x, y Option[T]) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return *x == *y
}

// Unwrap returns a value that x has.
// It panics if x has no value.
func Unwrap[T any](x Option[T]) T {
	if x == nil {
		panic("option.Unwrap: None")
	}
	return *x
}

// UnwrapOr returns a value that x has.
// It returns def if x has no value.
func UnwrapOr[T any](x Option[T], def T) T {
	if x == nil {
		return def
	}
	return *x
}

// UnwrapOrElse returns a value that x has.
// It returns def() if x has no value.
func UnwrapOrElse[T any](x Option[T], def func() T) T {
	if x == nil {
		return def()
	}
	return *x
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
// TODO: func DeriveSemigroup[T any](s algebra.Semigroup[T]) algebra.Semigroup[Option[T]]
// TODO: func DeriveMonoid[T any](s algebra.Semigroup[T]) algebra.Monoid[Option[T]] (note that T doesn't need to be Monoid)
