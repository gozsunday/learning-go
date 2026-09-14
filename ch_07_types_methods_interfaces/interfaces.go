package main

import (
	"fmt"
	"io"
	"math"
)

// ———— METHODS & INTERFACES DEFINITIONS START ————

type Shaper interface {
	Area() float64
	Perimeter() float64
}

// Rectangle Type Definition
type Rectangle struct {
	Width, Height float64
}

// Rectangle Methods
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle type Definition
type Circle struct {
	Radius float64
}

// Circle Methods
func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}
func (r Circle) Perimeter() float64 {
	return 2 * math.Pi * r.Radius
}

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

// ———— METHODS & INTERFACES DEFINITIONS END ————

func printShapeInfo(s Shaper) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func printAnything(v any) {
	fmt.Println(v)
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func interfaces() {
	// Interfaces in Go define a set of method signatures. Any type that implements all these
	// methods automatically satisfies the interface. no explicit declaration required.
	// This is called implicit satisfaction.

	// Interfaces are usually named with words ending with "-er".

	// From the types and their respective methods, we can see that both the Rectangle and Circle
	// types, satisfy the Shaper interface.

	printShapeInfo(Rectangle{
		Width:  12.54,
		Height: 9.38,
	})
	printShapeInfo(Circle{
		Radius: 9.99,
	})

	var s Shaper = Rectangle{Width: 12, Height: 3}

	fmt.Printf("%T\n", s)

	//* The empty interface
	// `any` is an alias for `interface{}`, and it has no methods. This interface will accept every type.
	// `any` is simply saying "I dont know what this type is yet". Any variable withe the `any`
	// interface still needs to go through a type assertion before it can be used for anything useful.

	printAnything("wtf")
	printAnything(920)
	printAnything(Circle{Radius: 31})

	//* Embedding and Interfaces
	// Intrerfaces can also be embedded into other interfaces,

	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Closer interface {
		Close() error
	}

	type ReadCloser interface {
		Reader
		Closer
	}

	//* Accept Interfaces, return Structs

	//* Interfaces and nil
	// nil can be used to represent the zero value for an interface, but its not as straightforward
	// as concrete types.

	// In Go, interfaces are implemented as a struct with two pointer fields. One for the value and
	// the other for the type of the value. As long as the type field is non-nil, the interface is
	// non-nil.
	// Therefore, for an interface to be considered nil, the value field and the type field have
	// to be nil.

	// If an interface variable is nil, invoking any method on it will trigger a panic.
	// If the interface variable is not nil, you can invoke methods on it. However, note that if
	// the value is nil, and the methods of the assigned type doesnt properly handle nil, a panic
	// can still be triggered.

	//* Interfaces are Comparable
	// Interfaces with the same value and type and equal to each other.
	// However, this isnt the case, when the values are equal but the types are not.

	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}

	DoublerCompare(&di, &di2)
	DoublerCompare(&di, dis)
	// DoublerCompare(dis, dis2) // this causes a panic

	fmt.Println(dis2)

	//* Empty Interface
	// Like we've established before, the empty interface `interface{}` is used to declare a
	// variable that can accept a value of any type.
	// The keyword `any` can also be used to do the say thing since its an alias of `interface{}`.

	var i interface{} // `any` can be used instead

	i = 20
	i = "wagwan"
	i = map[string]int{
		"emeka":  12,
		"nkechi": 14,
	}
	i = struct {
		FName string
		LName string
	}{
		FName: "Chidubem",
		LName: "Chukwu",
	}

	fmt.Println(i)

	// because an empty interface doesnt say anything about the value it represents, it cannot be
	// used for much.
	// one common use of `any` is as a placeholder for data of uncertain schema thats read from
	// an external source.

	err := typeAssertionsAndSwitches()
	fmt.Println(err)
}

type MyInt int

func typeAssertionsAndSwitches() error {
	//* Type Assertions and Type Switches
	// There are two ways to check if a variable of an interface type has a specific concrete type,
	// or if the concrete type implements another interface.

	// A type assertion names the concrete type of the value that implemeted the interface, or
	// names another interface that is also implemented by the concrete type whose value is stored
	// in the interface.

	var num1 any // interface{}
	var mine MyInt = 20
	num1 = mine
	num2 := num1.(MyInt)

	// in the code above, we're asserting that the variable `num1` is of the type `MyInt`. if
	// it is truly of that type, it assigns the value of num1 to num2. if it isnt, a panic is
	// triggered.

	// a safer way to do this assertion is to use the "comma ok" idiom.

	num3, ok := num1.(MyInt)

	if !ok {
		return fmt.Errorf("unexpected type for %v", num1)
	}
	fmt.Println(num3 + 1)

	fmt.Println(num2, num3)

	// when an interface could be one of multiple possible types, use a Type Switch instead.

	doThings(num1)

	//! Type Assertions and Switches are to be used sparingly.

	return nil
}

func doThings(i any) {
	switch i := i.(type) {
	case nil:
	case int:
		fmt.Printf("The value is of type %T", i)
	case MyInt:
		fmt.Printf("The value is of type %T", i)
	case io.Reader:
		fmt.Printf("The value is of type %T", i)
	case string:
		fmt.Printf("The value is of type %T", i)
	case bool, rune:
		fmt.Printf("The value is of type %T", i)
	default:
		fmt.Println("I have no freaking idea what the type of the value is")
	}
}
