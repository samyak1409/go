package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 1. Constants Variables and Basic Data Types

	x := 1000000000000000000 // by default int = int64 if 64-bit machine
	fmt.Println(x * 10)      // int overflows and without giving any error becomes -ve, pay attention!

	// uint = 2x of int since only +ve vals

	// `float` doesn't exist, specify `float64` or `float32`

	// if we want to store rgb, (256), uint8 is the best option in terms of memory

	i1, i2 := 3, 2
	fmt.Println(i1 / i2)                   // 1
	fmt.Println(float64(i1 / i2))          // 1
	fmt.Println(float64(i1) / float64(i2)) // 1.5

	fmt.Println(`Hello
	World`) // outputs exactly the same string including line breaks and whitespace chars

	fmt.Println(len("some string")) // not the string len, but the bytes
	// 1 char = 1 byte (8 bits), but not with non-ascii chars! 1 char could be = 2 bytes
	// solution? builtin package "unicode/utf8" (`utf8.RuneCountInString()`)

	// x := some_func()
	// Better for readability:
	// var x string = some_func()
	// We instantly known the type without hovering over.

	const fixed = 1 // can't be changed once set, somewhat like tuples in Python
	fmt.Println(fixed)

	// 2. Functions and Control Structures

	// `and` -> &&
	// `or` -> ||
	// e.g. `if 1==1 || 2==2 {...}`

	// 3. Arrays, Slices, Maps and Loops

	var arr1 [5]int32 // 4*5 = 20 bytes of memory is allocated
	fmt.Println(arr1)
	// Print memory address (hexadecimal):
	fmt.Println(&arr1[0]) // 0x140000a8030
	fmt.Println(&arr1[2]) // 0x140000a8038
	fmt.Println(&arr1[3]) // 0x140000a803c
	fmt.Println(&arr1[4]) // 0x140000a8040

	arr2 := [...]int{1, 2, 3} // ... -> 3, still fixed sized!
	// compiler just adds the number depending on the num of elements
	fmt.Println(arr2)

	sl := []int{10, 11, 31}
	fmt.Println(sl, len(sl), cap(sl))
	sl = append(sl, 45, 3, 3, 3)
	fmt.Println(sl, len(sl), cap(sl))
	// extra space is allocated based on the number of elements which are being appended
	// e.g. above, first cap is 3, then appending, space is not there, so 3*2 = 6
	// now 6 is still less and can't fit all the elements in the appending
	// so (3+1)*2 is done = 8
	// so basically, (current cap + 1 + 1 + 1...) until (current cap + 1) * 2 can allocate

	// Python `extend()`: `append(sl1, sl2...)`

	// We can also use make() to initialize slices, for preallocation.
	// Benefit? See screenshot `ss1.png`.

	// Go maps doesn't preserve order.

	// 4. Strings, Runes and Bytes

	str1 := "résumé"
	fmt.Println(str1, str1[0], string(str1[0])) // résumé 114 r
	// just like c, str is by default printed as underlying ascii code

	// but on iterating:
	for i, v := range str1 {
		fmt.Println(i, v)
	}
	// outputs:
	// 0 114
	// 1 233
	// 3 115
	// 4 117
	// 5 109
	// 6 233
	// 2 skipped?
	// Also:
	fmt.Println(len(str1)) // 8 instead of 6
	// You guessed right. Non-ascii chars takes > 1 bytes because they are split up in two (or more) different memory blocks of fixed bytes,
	// é - takes two blocks of 1 bytes each
	// Go uses `uint8` (8bits=1byte) to store a str char.
	// Explanation: https://youtu.be/8uiZC0l4Ajw?t=1639

	// To avoid this index skip situation, we can use `rune` type:
	str2 := []rune("résumé")
	fmt.Println(str2, str2[0], string(str2[0]))
	// Note that: rune is an alias for int32 and is equivalent to int32 in all ways.
	// It is used, by convention, to distinguish character values from integer values.
	fmt.Printf("%T\n", str2) // []int32

	// Strings are immutable. So, concatenating is O(n) every time.
	// So, just like Java, we use string builder:
	str_b := strings.Builder{} // declare / define empty
	for i := range 5 {
		// str_b.WriteString(string(i)) // type conversion like this doesn't work
		str_b.WriteString(strconv.Itoa(i))
	}
	fmt.Println(str_b.String())
	// dynamic arr is used internally.

	// Advanced topics:
	structsAndInterfaces()
	pointers()
	goroutines()
	channels()
	generics()

}
