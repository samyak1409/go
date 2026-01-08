package main

import (
	"fmt"
)

type Lang struct {
	name  string
	speed string
}

// this is a method since it's defined on a type using "receiver" arg
// https://go.dev/tour/methods/1
func (lang Lang) describe() {
	fmt.Printf("%s is %s\n", lang.name, lang.speed)
} // Though note that we can write this method as a function as well, just make the special "receiver" arg a normal arg.
// And then we can do `describe(l1)` instead of `l1.describe()`.

func main() {
	l1, l2 := Lang{"Python", "Slow"}, Lang{"Go", "Fast"}
	l1.describe()
	l2.describe()
	// We can declare a method on non-struct types, too. https://go.dev/tour/methods/3

}
