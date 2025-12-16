package main

import (
	"fmt"

	"errors"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	newFunc()
	fmt.Println(quote.Go())

	var xx int // init (with default val like 0, turned out they're called zero values https://go.dev/tour/basics/12)
	fmt.Println(xx)
	xx = 2                     // declare
	var name string = "Samyak" // init + declare
	lang := "Go"               // type infer + init (:) + declare (=)
	// Alt:
	var lang2 = "Python" // type infer + init (var) + declare (=)
	// Outside a function, every statement begins with a keyword (var, func, and so on),
	// so the := construct is not available.
	const pi = 3.14
	// Imp: Cannot declare a variable without using it -> error!
	fmt.Println(name, lang, xx, lang2, pi)

	score := 90
	if score > 90 {
		fmt.Println("Excellent")
	} else if score > 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	for n := 0; n < 4; n++ {
		fmt.Println("Count:", n)
	}

	sum := 1
	for sum < 1000 { // init and post statements are optional
		sum += sum
	}
	fmt.Println(sum)

	// Doesn't have a separate `while` loop, instead, you just use `for` with condition.
	start := 1
	for start <= 2 {
		fmt.Println("Will run 2 times.")
		start++
	}

	// Infinite loop: https://go.dev/tour/flowcontrol/4

	// `v` won't be available outside the `if` (but available in subsequent `else`s / `else if`s)
	// Similar to recently added := in Python, but there, it's not local.
	if v := 10; v < 11 {
		fmt.Println(v)
	}

	sports := []string{"Badminton", "TT", "Cricket"}
	// (above is a slice, a dynamic arr)
	for i, sport := range sports {
		fmt.Println(i, sport)
	}
	// can use `_` in place of `i` if not required
	// if only i is required, only i can be used, no need to do `i, _`

	fmt.Println(add(2, 3))
	fmt.Println(subtract(2, 3))

	res, err := divide(10, 2)
	if err == "" {
		fmt.Println("Result:", res)
	} else {
		fmt.Println("Error:", err)
	}

	values := []int{5, 8, 2, 10, 3}
	min, max, sum := getStats(values)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	fmt.Println("Sum:", sum)

	res2, err2 := divide2(10, 0)
	if err2 == nil {
		fmt.Println("Result:", res2)
	} else {
		fmt.Println("Error:", err2)
	}

	// switch: Not like C, more similar to Python (match).
	switch k := 1; k {
	case 2:
		fmt.Println("C1")
	case 3: // can also add a func call here, like: `case get_val():`
		fmt.Println("C2")
	default:
		fmt.Println("D", k)
	}

	// Switch without a condition is the same as `switch true`.
	// Long if else chains? Use https://go.dev/tour/flowcontrol/11

	// Defer
	testDefer()

	// DS

	// A struct is a collection of fields.
	type Detail struct {
		initials string
		height   int
	}
	d1 := Detail{"SJ", 180}
	d2 := Detail{initials: "SJ"} // other fields get "zero" (default) values
	fmt.Println(d1, d1.initials, d2)

	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3} // this creates the same array as above, then builds a slice that references it
	slice = append(slice, -1, -2)
	// if the cap() is less than only new array would be allocated!
	// https://go.dev/tour/moretypes/15
	// (So, just like python list's internal working.)
	fmt.Println(arr, slice, arr[1:3]) // arr[1:3] is also a slice

	// Imp: Unlike Python, here, slices are references to array,
	// so updating an element in slice updates that element in the original array.
	// (They're the same.)
	// https://go.dev/tour/moretypes/8

	fmt.Println(len(slice), cap(slice)) // 5, 6 (guess why 6)
	// len(slice): num of elements it contains
	// cap(slice): num of elements in the underlying arr, counting from the first element in the slice

	// Creating a slice with make (https://go.dev/tour/moretypes/13)
	z1 := make([]int, 5)
	printSlice("z1", z1)
	z2 := make([]int, 0, 5)
	printSlice("z2", z2)
	z3 := z2[:2]
	printSlice("z3", z3)
	z4 := z3[2:5]
	printSlice("z4", z4)

	// Maps
	var hm0 map[string]int  // init with zero value of map
	fmt.Println(hm0)
	// hm0["z"] = 100 // error: zero val of map = nil, and nil map can't add keys
	// Looks like a Go implementation problem.
	// Solution? Use `make` function.
	// https://go.dev/tour/moretypes/19

	hm := map[string]int{"a": 1, "b": 2}

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	// https://go.dev/tour/moretypes/21

	hm["c"] = 3
	hm["b"] = -2
	fmt.Println(hm, hm["d"], hm["b"])
	delete(hm, "a")
	val, exists := hm["a"]
	fmt.Println(val, exists) // 0 false
	val, exists = hm["b"]
	fmt.Println(val, exists) // -2 true

	// Functions can be passed as args just like in Python.
	// https://go.dev/tour/moretypes/24
}

func add(x int, y int) int {
	return x + y
}

// same param types
func subtract(x, y int) int {
	return x - y
}

// return two vals: res, err
func divide(x, y float64) (float64, string) {
	if y == 0 {
		return 0, "division by zero"
	}
	return x / y, ""
}

// named return vals
func getStats(values []int) (min, max, sum int) {
	if len(values) == 0 {
		return 0, 0, 0
	}
	min, max, sum = values[0], values[0], 0 // no initialization needed (`:=`)
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	return // naked return: no need to return explicitly
}

// Go's approach to error handling is different from Python's exceptions.
// Instead of try/except blocks, Go functions return errors that must be explicitly checked:
func divide2(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
} // Apparently:
// The focus on explicit error checking helps make Go code robust and readable.
// It forces you to think about what should happen when something goes wrong.

func testDefer() {
	fmt.Println("Testing Defer")
	fmt.Println("1")
	defer fmt.Println("2") // ran just before func return, in reverse (uses stack)
	fmt.Println("3")
	defer fmt.Println("4") // ran just before func return, in reverse (uses stack)
	fmt.Println("5")
	// 1, 3, 5, 4, 2

	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.

	// Side-note: Can use defer to print reverse counting like: https://go.dev/tour/flowcontrol/13
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
