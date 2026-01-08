package main

import "fmt"

func pointers() {
	var p *int
	// zero val of a pointer val is nil:
	fmt.Println(p) // <nil>
	// accessing the value the pointer is pointing to:
	// fmt.Println(*p) // panic: runtime error: invalid memory address or nil pointer dereference

	i := 10
	p = &i          // &i = address(i) = address of i
	fmt.Println(*p) // 10
	*p = 20
	fmt.Println(i) // 20

	// Copy data (no reference):
	j := i
	j = 30            // so if we update `j`, `i` is unchanged
	fmt.Println(j, i) // 30, 20

	// Unlike python, nested data is copied as well (no reference) in case of nested DS array:
	arr := [3]int{1, 2, 1}
	arr1 := arr
	arr1[0] = 4
	fmt.Println(arr1, arr) // [4 2 1] [1 2 1]

	// But not with maps, slices:
	hm := map[int]int{1: 1, 2: 2}
	hm1 := hm
	hm1[1] = 3
	fmt.Println(hm1, hm) // map[1:3 2:2] map[1:3 2:2]
	sl := []int{1, 2, 1}
	sl1 := sl
	sl1[0] = 4
	fmt.Println(sl1, sl) // [4 2 1] [4 2 1]

	// Note that: By default, values are passed in functions, not pointers.
	// Pass the pointers (to avoid duplication of data when not needed and update the same var),
	// using `&arg` instead of `arg`, and `param *<type>` instead of `param <type>`,
	// and update/access using `*param` instead of `param`.
}
