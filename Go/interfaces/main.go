package main

import (
	"fmt"
)

type Lang struct {
	name  string
	speed string
}

func (lang Lang) describe() string {
	fmt.Printf("%s is %s\n", lang.name, lang.speed)
	return "OK"
}

// An interface type is defined as a set of method signatures.
type I interface {
	describe() string // if method doesn't have a return type (value), leave the type empty here
} // Unlike Python, here functions can return no value (in Python, `None` is returned).
// If we call a function with no return value and try to print, it raises no value error.

func main() {
	// A value (variable) of interface type (here `x`) can hold any value that implements those methods (here `l`).
	l := Lang{"Go", "Fast"}
	var x I = l
	x.describe()

	// A type implements an interface by implementing its methods.
	// In current file, `Lang` implements `I`.
	// There is no explicit declaration of intent, no "implements" keyword.
	// Implicit interfaces decouple the definition of an interface from its implementation (which could then appear in any package without prearrangement).
	// https://go.dev/tour/methods/10

}
