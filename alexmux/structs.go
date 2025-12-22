package main

import "fmt"

// start with type keyword because we're defining a user-defined type
// (just like defining functions start with func)
type ice struct {
	mileage uint8 // (let's save some memory)
	litres  float32
	// ownerInfo owner // -> car0.ownerInfo.name, car0.ownerInfo.age
	//owner // -> car0.name, car0.age
}

type ev struct {
	mileage uint8 // km per kwh
	kwh     float32
}

type vehicle interface {
	rangeLeft() float32
}

type owner struct {
	name string
	age  uint8
}

func structsAndInterfaces() {

	// 5. Structs and Interfaces

	var car0 ice                    // declare only (no define -> zero vals)
	var car1 ice = ice{mileage: 27} // declare & define with partial data
	car2 := ice{24, 12.5}           // type infer declare & define with complete data

	car2.mileage = 27

	fmt.Println(car0, car1, car2)

	// Anonymous structs
	// (in-place declaration, so not re-usable, can't be referenced, since no var store that)
	owner0 := struct {
		name string
		age  uint8
	}{"SJ", 24}
	fmt.Println(owner0, owner0.name)

	// Methods:
	fmt.Println(car2.rangeLeft())

	// Now suppose we've a function `canGo`:
	canGo(car2, 300)
	// And suppose now we've one more type `ev` as well!
	// So, now we would need to make one more `canGo` which could take type `ev`,
	// even when we've the *SAME* `canGo` definition.
	// There comes interface.

	canGo(ev{7, 79}, 500)
}

func (car ice) rangeLeft() float32 {
	return float32(car.mileage) * car.litres
}

func (car ev) rangeLeft() float32 {
	return float32(car.mileage) * car.kwh
} // this func needs to be separate since there's at least one change in definition

// func canGo(car ice, kms float32) { // without interface
func canGo(car vehicle, kms float32) {
	// `vehicle` type can take any object which have all the methods (here, only `rangeLeft`)
	// present in the interface.
	if kms <= car.rangeLeft() {
		fmt.Println("Would go.")
	} else {
		fmt.Println("Won't go.")
	}
}
