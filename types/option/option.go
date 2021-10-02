package option

// TODO: type Option[T any]
// TODO: func Some[T any](x T) Option[T]
// TODO: func None[T any]() Option[T]

// TODO: func FromIterator[T any](it iterator.Iterator[T]) Option[T]
// TODO: func (x Option[T]) Iter() iterator.Iterator[T]

// TODO: func Pure[T any](x T) Option[T]
// TODO: func AndThen[T, U any](x T, fn func(T) Option[U]) Option[U]

// TODO: func IsSome[T any](x Option[T]) bool
// TODO: func Unwrap[T any](x Option[T]) T
// TODO: func UnwrapOr[T any](x Option[T], default T) T
// TODO: func UnwrapOrElse[T any](x Option[T], default func() T) T

// TODO: func MapOr[T, U any](x Option[T] fn func(T) U, default U) U
// TODO: func MapOrElse[T, U any](x Option[T], fn func(T) U, default func() U) U
// TODO: func Switch[T any](x Option[T], ifSome func(T) ifNone func())

// TODO: func DeriveEq[T any](eq cmp.Eq[T]) cmp.Eq[Option[T]]
// TODO: func DeriveOrd[T any](ord cmp.Ord[T]) cmp.Ord[Option[T]]
// TODO: func DeriveSemigroup[T any](s algebra.Semigroup[T]) algebra.Semigroup[Option[T]]
// TODO: func DeriveMonoid[T any](s algebra.Semigroup[T]) algebra.Monoid[Option[T]] (note that T doesn't need to be Monoid)
