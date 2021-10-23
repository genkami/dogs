package iterator_test

import (
	"fmt"
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/classes/cmp"
	"github.com/genkami/dogs/types/iterator"
	"github.com/genkami/dogs/types/pair"
	"github.com/genkami/dogs/types/slice"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRange(t *testing.T) {
	subject := func(start, end int) []int {
		return toSlice[int](iterator.Range(start, end))
	}

	assert.Equal(t, subject(1, 3), []int{1, 2, 3})
	assert.Equal(t, subject(1, 1), []int{1})
	assert.Equal(t, subject(1, 0), []int{})
}

func TestFind(t *testing.T) {
	assertFound := func(name string, xs []int, x int, fn func(int) bool) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			found, ok := iterator.Find[int](it, fn)
			assert.True(t, ok)
			assert.Equal(t, found, x)
		})
	}
	assertNotFound := func(name string, xs []int, fn func(int) bool) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			_, ok := iterator.Find[int](it, fn)
			assert.False(t, ok)
		})
	}

	assertFound("ok", []int{1, 2, 3}, 3, func(x int) bool {
		return x == 3
	})
	assertFound("first", []int{1, 2, 3}, 1, func(x int) bool {
		return true
	})
	assertNotFound("not found", []int{1, 2, 3}, func(x int) bool {
		return x > 5
	})
	assertNotFound("empty", []int{}, func(x int) bool {
		return true
	})
}

func TestFindIndex(t *testing.T) {
	assertFound := func(name string, xs []int, i int, fn func(int) bool) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			found := iterator.FindIndex[int](it, fn)
			assert.Equal(t, found, i)
		})
	}
	assertNotFound := func(name string, xs []int, fn func(int) bool) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			found := iterator.FindIndex[int](it, fn)
			assert.True(t, found < 0)
		})
	}

	assertFound("ok", []int{1, 2, 3}, 2, func(x int) bool {
		return x == 3
	})
	assertFound("first", []int{1, 2, 3}, 0, func(x int) bool {
		return true
	})
	assertNotFound("not found", []int{1, 2, 3}, func(x int) bool {
		return x > 5
	})
	assertNotFound("empty", []int{}, func(x int) bool {
		return true
	})
}

func TestFindElem(t *testing.T) {
	eq := cmp.DeriveEq[int]()
	assertFound := func(name string, xs []int, x int) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			found, ok := iterator.FindElem[int](eq)(it, x)
			assert.True(t, ok)
			assert.Equal(t, found, x)
		})
	}
	assertNotFound := func(name string, xs []int, x int) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			_, ok := iterator.FindElem[int](eq)(it, x)
			assert.False(t, ok)
		})
	}

	assertFound("found", []int{1, 2, 3}, 2)
	assertNotFound("not found", []int{1, 2, 3}, 4)
	assertNotFound("empty", []int{}, 0)
}

func TestFindElemIndex(t *testing.T) {
	eq := cmp.DeriveEq[int]()
	assertFound := func(name string, xs []int, x int, i int) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			found := iterator.FindElemIndex[int](eq)(it, x)
			assert.Equal(t, found, i)
		})
	}
	assertNotFound := func(name string, xs []int, x int) {
		t.Run(name, func(t *testing.T) {
			it := slice.Slice[int](xs).Iter()
			found := iterator.FindElemIndex[int](eq)(it, x)
			assert.True(t, found < 0)
		})
	}

	assertFound("found", []int{1, 2, 3}, 2, 1)
	assertNotFound("not found", []int{1, 2, 3}, 4)
	assertNotFound("empty", []int{}, 0)
}

func TestFilter(t *testing.T) {
	assertEqual := func(xs []int, expected []int, fn func(int) bool) {
		it := iterator.Filter[int](slice.Slice[int](xs).Iter(), fn)
		actual := ([]int)(slice.FromIterator[int](it))
		assert.Equal(t, expected, actual)
	}

	assertEqual([]int{}, []int{}, func(_ int) bool { return true })
	assertEqual([]int{1}, []int{1}, func(_ int) bool { return true })
	assertEqual([]int{1, 2, 3}, []int{1, 2, 3}, func(_ int) bool { return true })

	assertEqual([]int{}, []int{}, func(i int) bool { return i%2 == 0 })
	assertEqual([]int{1}, []int{}, func(i int) bool { return i%2 == 0 })
	assertEqual([]int{2}, []int{2}, func(i int) bool { return i%2 == 0 })
	assertEqual([]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}, func(i int) bool { return i%2 == 0 })
	assertEqual([]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}, func(i int) bool { return i%2 == 0 })
}

func TestTake(t *testing.T) {
	subject := func(xs []int, n int) []int {
		it := iterator.Take(slice.Slice[int](xs).Iter(), n)
		return ([]int)(slice.FromIterator(it))
	}

	assert.Equal(t, subject([]int{}, 0), []int{})
	assert.Equal(t, subject([]int{}, 1), []int{})
	assert.Equal(t, subject([]int{}, 2), []int{})

	assert.Equal(t, subject([]int{3, 4, 5}, 0), []int{})
	assert.Equal(t, subject([]int{3, 4, 5}, 1), []int{3})
	assert.Equal(t, subject([]int{3, 4, 5}, 2), []int{3, 4})
	assert.Equal(t, subject([]int{3, 4, 5}, 3), []int{3, 4, 5})
	assert.Equal(t, subject([]int{3, 4, 5}, 4), []int{3, 4, 5})
}

func TestMap(t *testing.T) {
	subject := func(xs []int) []string {
		it := slice.Slice[int](xs).Iter()
		mapped := iterator.Map[int, string](it, func(x int) string {
			return strconv.FormatInt(int64(x), 10)
		})
		return toSlice(mapped)
	}

	assert.Equal(t, subject([]int{}), []string{})
	assert.Equal(t, subject([]int{1}), []string{"1"})
	assert.Equal(t, subject([]int{1, 2}), []string{"1", "2"})
	assert.Equal(t, subject([]int{1, 2, 3}), []string{"1", "2", "3"})
}

func TestFlatMap(t *testing.T) {
	subject := func(xs ...int) []string {
		it := slice.Slice[int](xs).Iter()
		mapped := iterator.FlatMap(it, func(x int) iterator.Iterator[string] {
			strs := make([]string, 0, x)
			for i := 0; i < x; i++ {
				strs = append(strs, fmt.Sprint(x))
			}
			return slice.Slice[string](strs).Iter()
		})
		return toSlice(mapped)
	}

	assert.Equal(t, subject(), []string{})
	assert.Equal(t, subject(1), []string{"1"})
	assert.Equal(t, subject(2), []string{"2", "2"})

	assert.Equal(t, subject(2, 3, 0, 1), []string{"2", "2", "3", "3", "3", "1"})
}

func TestForEach(t *testing.T) {
	iter := func(xs ...int) iterator.Iterator[int] {
		return slice.Slice[int](xs).Iter()
	}

	t.Run("empty", func(t *testing.T) {
		numCalled := 0
		args := []int{}
		iterator.ForEach(iter(), func(i int) {
			numCalled++
			args = append(args, i)
		})

		assert.Equal(t, numCalled, 0)
		assert.Equal(t, args, []int{})
	})

	t.Run("singleton", func(t *testing.T) {
		numCalled := 0
		args := []int{}
		iterator.ForEach(iter(3), func(i int) {
			numCalled++
			args = append(args, i)
		})

		assert.Equal(t, numCalled, 1)
		assert.Equal(t, args, []int{3})
	})

	t.Run("many", func(t *testing.T) {
		numCalled := 0
		args := []int{}
		iterator.ForEach(iter(1, 2, 3, 4, 5), func(i int) {
			numCalled++
			args = append(args, i)
		})

		assert.Equal(t, numCalled, 5)
		assert.Equal(t, args, []int{1, 2, 3, 4, 5})
	})
}

func TestFold(t *testing.T) {
	add := func(x string, y int) string {
		return x + strconv.FormatInt(int64(y), 10)
	}
	subject := func(x string, xs []int) string {
		it := slice.Slice[int](xs).Iter()
		return iterator.Fold[string, int](x, it, add)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject("", []int{}), "")
		assert.Equal(t, subject("a", []int{}), "a")
		assert.Equal(t, subject("xyzzy", []int{}), "xyzzy")
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject("", []int{1}), "1")
		assert.Equal(t, subject("a", []int{2}), "a2")
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject("", []int{1, 2, 3}), "123")
		assert.Equal(t, subject("hoge", []int{3, 2, 1, 0}), "hoge3210")
	})
}

func TestZip(t *testing.T) {
	type Pair = pair.Pair[int, string]
	subject := func(xs []int, ys []string) []Pair {
		xit := slice.Slice[int](xs).Iter()
		yit := slice.Slice[string](ys).Iter()
		zipped := iterator.Zip(xit, yit)
		return toSlice(zipped)
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject([]int{}, []string{}), []Pair{})
		assert.Equal(t, subject([]int{1, 2}, []string{}), []Pair{})
		assert.Equal(t, subject([]int{}, []string{"a", "b"}), []Pair{})
	})

	t.Run("same length", func(t *testing.T) {
		assert.Equal(t, subject([]int{1}, []string{"a"}), []Pair{{1, "a"}})
		assert.Equal(t, subject([]int{1, 2}, []string{"a", "b"}), []Pair{{1, "a"}, {2, "b"}})
		assert.Equal(t, subject([]int{1, 2, 3}, []string{"a", "b", "c"}), []Pair{{1, "a"}, {2, "b"}, {3, "c"}})
	})

	t.Run("different length", func(t *testing.T) {
		assert.Equal(t, subject([]int{1, 2, 3}, []string{"a", "b"}), []Pair{{1, "a"}, {2, "b"}})
		assert.Equal(t, subject([]int{1, 2}, []string{"a", "b", "c"}), []Pair{{1, "a"}, {2, "b"}})
		assert.Equal(t, subject([]int{1, 2, 3}, []string{"a"}), []Pair{{1, "a"}})
	})
}

func TestUnfold(t *testing.T) {
	subject := func(init int, step func(int) (int, int, bool)) []int {
		return toSlice[int](iterator.Unfold[int, int](init, step))
	}

	f := func(_ int) (int, int, bool) {
		return 0, 0, false
	}
	assert.Equal(t, subject(0, f), []int{})

	f = func(s int) (int, int, bool) {
		if 5 < s {
			return 0, 0, false
		}
		return s + 1, s, true
	}
	assert.Equal(t, subject(0, f), []int{0, 1, 2, 3, 4, 5})

	f = func(s int) (int, int, bool) {
		if 3 < s {
			return 0, 0, false
		}
		return s + 1, s * 2, true
	}
	assert.Equal(t, subject(0, f), []int{0, 2, 4, 6})
}

func TestSumWithInit(t *testing.T) {
	intSemigroup := algebra.DeriveAdditiveSemigroup[int]()
	subject := func(x int, xs []int) int {
		return iterator.SumWithInit(intSemigroup)(x, slice.Slice[int](xs).Iter())
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{}), 0)
		assert.Equal(t, subject(1, []int{}), 1)
		assert.Equal(t, subject(999, []int{}), 999)
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1}), 1)
		assert.Equal(t, subject(1, []int{2}), 3)
		assert.Equal(t, subject(123, []int{456}), 579)
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject(0, []int{1, 2, 3}), 6)
		assert.Equal(t, subject(1, []int{10, 100, 1000}), 1111)
	})
}

func TestSum(t *testing.T) {
	intMonoid := algebra.DeriveAdditiveMonoid[int]()
	subject := func(xs []int) int {
		return iterator.Sum(intMonoid)(slice.Slice[int](xs).Iter())
	}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, subject([]int{}), 0)
	})

	t.Run("singleton", func(t *testing.T) {
		assert.Equal(t, subject([]int{0}), 0)
		assert.Equal(t, subject([]int{1}), 1)
		assert.Equal(t, subject([]int{999}), 999)
	})

	t.Run("multiple elements", func(t *testing.T) {
		assert.Equal(t, subject([]int{1, 2}), 3)
		assert.Equal(t, subject([]int{1, 2, 3}), 6)
		assert.Equal(t, subject([]int{1, 10, 100, 1000}), 1111)
	})
}

func TestPure(t *testing.T) {
	subject := func(x int) []int {
		it := iterator.Pure(x)
		return toSlice(it)
	}

	assert.Equal(t, subject(1), []int{1})
	assert.Equal(t, subject(2), []int{2})
	assert.Equal(t, subject(3), []int{3})
}

func toSlice[T any](it iterator.Iterator[T]) []T {
	return iterator.Fold[[]T, T](
		make([]T, 0),
		it,
		func(xs []T, x T) []T { return append(xs, x) },
	)
}
