package examples_test

import (
	"fmt"
	"github.com/genkami/dogs/types/iterator"
	"github.com/genkami/dogs/types/pair"
)

func ExampleFibonacci_pureGo() {
	a, b := 1, 1
	for i := 0; i < 5; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
	// Output: 1
	// 1
	// 2
	// 3
	// 5
}

func ExampleFibonacci_unfold() {
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
	// Output: 1
	// 1
	// 2
	// 3
	// 5
}
