package main

import "fmt"

func generics() {

	fmt.Println("Generics:")

	// Use case: Suppose we've multiple slices but having different type of values: int, float64, etc., and we want to write a function
	// which return the sum of the values of the given slice.
	// We can't, since this is a static typed lang, we've to give a type.
	// This problem exists with all the static typed langs, one solution is Generics.
	sl1 := []int{1, 2}
	sl2 := []float64{1.3, 2.4}
	fmt.Println(getSum(sl1)) // 3
	fmt.Println(getSum(sl2)) // 3.7

	// Suppose now we want to write a function which returns whether a slice is empty or not.
	// Now, instead of doing `T int | float64 | ...` and so on, we can just do `T any`. See `isEmpty()`.
	sl3 := []bool{}
	fmt.Println(isEmpty(sl3)) // true
	fmt.Println(isEmpty(sl1)) // false

	// Note that above we're doing `getSum(sl1)`, `isEmpty(sl3)`.
	// The actual syntax is `getSum[int](sl1)`, `isEmpty[bool](sl3)`.
	// But, since Go can see the data, it's inferring the type, and suggesting to remove the types while calling.
	// There would be cases where Go won't be able to infer though, e.g. when fetching JSON from a local text file, and there are two
	// JSON struct types defined, then Go won't know. See video 54:21 - 55:08.

	// Generics can also be used with struct types. See video 55:09 - 55:46.

}

func getSum[T int | float64](sl []T) T {
	var sum T
	for _, v := range sl {
		sum += v
	}
	return sum
}

func isEmpty[T any](sl []T) bool {
	return len(sl) == 0
}
