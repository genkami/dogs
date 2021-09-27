package dogs

import "fmt"

type Eq[T any] struct {
	Equal func(T, T) bool
}

type Ord[T any] struct {
	Compare func(T, T) Ordering
}

func (o *Ord[T]) Lt(x, y T) bool {
	result := o.Compare(x, y)
	return result == LT
}

func (o *Ord[T]) Le(x, y T) bool {
	result := o.Compare(x, y)
	return result == LT || result == EQ
}

func (o *Ord[T]) Gt(x, y T) bool {
	result := o.Compare(x, y)
	return result == GT
}

func (o *Ord[T]) Ge(x, y T) bool {
	result := o.Compare(x, y)
	return result == GT || result == EQ
}

func (o *Ord[T]) Eq(x, y T) bool {
	result := o.Compare(x, y)
	return result == EQ
}

func (o *Ord[T]) Ne(x, y T) bool {
	result := o.Compare(x, y)
	return result != EQ
}

type Ordering int

const (
	LT Ordering = iota
	EQ
	GT
)

func (o Ordering) GoString() string {
	switch o {
	case LT:
		return "LT"
	case EQ:
		return "EQ"
	case GT:
		return "GT"
	default:
		return fmt.Sprintf("<unknown Ordering (%d)>", o)
	}
}