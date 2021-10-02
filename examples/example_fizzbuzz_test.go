package examples_test

import (
	"fmt"
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/types/iterator"
	"github.com/genkami/dogs/types/option"
)

func ExampleFizzBuzz_pureGo() {
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	// Output: 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func ExampleFizzBuzz_rangeAndMap() {
	fizzBuzz := func(i int) string {
		if i%3 == 0 && i%5 == 0 {
			return "FizzBuzz"
		} else if i%3 == 0 {
			return "Fizz"
		} else if i%5 == 0 {
			return "Buzz"
		} else {
			return fmt.Sprint(i)
		}
	}
	it := iterator.Map(iterator.Range[int](1, 15), fizzBuzz)
	iterator.ForEach(it, func(s string) { fmt.Println(s) })
	// Output: 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func ExampleFizzBuzz_monoid() {
	monoid := option.DeriveMonoid[string](algebra.DeriveAdditiveMonoid[string]())
	fizzBuzz := func(i int) string {
		fizz := option.Filter(option.Some[string]("Fizz"), func(_ string) bool { return i%3 == 0 })
		buzz := option.Filter(option.Some[string]("Buzz"), func(_ string) bool { return i%5 == 0 })
		return option.UnwrapOr(monoid.Combine(fizz, buzz), fmt.Sprint(i))
	}
	it := iterator.Map(iterator.Range[int](1, 15), fizzBuzz)
	iterator.ForEach(it, func(s string) { fmt.Println(s) })
	// Output: 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}
