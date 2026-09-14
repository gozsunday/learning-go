package main

import "fmt"

func main() {
	// pointers()
	MakePerson("Saint", "Jaeromi", 24)
	MakePersonPointer("Saint", "Jaeromi", 24)
}

func pointers() {
	//* Pointers
	// Pointers are variables that hold the location in memory where a value is stored.
	// The zero value for a pointer is nil.

	// The `&` is the address operator. it's put right in front of a value type and returns the address where the value is stored.

	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)

	// The `*` is the indirection operator. It precedes a variable of pointer type, and returns the value being pointed to.

	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	// Always make sure the pointer is non-nil. If you try to dereference a nil pointer, the program will panic.

	var nX *int
	fmt.Println(nX == nil) // true
	// fmt.Println(*nX) // will cause a panic

	// A pointer type is a type that represents a pointer. It has an * before a type name. It can be based on any type.

	// The built-in function `new` creates a pointer variable.

	p1 := new(int)         // will create a pointer variable with the zero-value instance of the base type
	fmt.Println(p1 == nil) // will print false, since the zero value of an int is 0
	fmt.Println(*p1)       // will print 0

	// For structs, using an & before the struct literal will create a pointer.
	// Doing the same thing for primitive literals and constants wont work because they dont
	// have memory addresses.

	type Foo struct {
		isFoo bool
	}
	pS1 := &Foo{} // & can be used before a struct literal

	// pP1 := &"some string" // can not be used before primitive literals, instead we can do
	pP1 := "some string"
	pointerPP1 := &pP1

	fmt.Println(pS1, pointerPP1)

	// pointer types can be used as the type of fields in a struct, but when initialising the struct,
	// if the pointer is of a primitive type, a literal cant be assigned directly to the field.

	type Person struct {
		fName string
		lName *string
	}

	var psn1 = Person{
		fName: "Gozman",
		// lName: "Sunday", // cannot use "Sunday" (untyped string constant) as *string value in struct literal
	}

	// We can use a variable to hold the value

	lastName := "Sunday"
	psn1.lName = &lastName

	// Or we can use a generic helper function that takes in a parameter of any type and returns a pointer to that type.

	psn2 := Person{
		fName: "Saint",
		lName: makePointer("Jaeromi"), // new("Jaeromi") can be used here!
	}

	fmt.Println(psn1, psn2)

	//* Mutable Parameters
	// Since Go is a call-by-value language, the values passed into functions are copies of the
	// original variables. For non pointer types like primitives, structs and arrays, this
	// means that the called function cant modify the original variable. This ensures the
	// immutability of the original.

	// However, if a pointer is passed to the function, the function gets a copy of the pointer.
	// This still points to the original, which means that the original can then be modified by
	// the called function.

	// This has some implications. One is that, if a nil pointer is passed to a function, the
	// value of the pointer cant be made non-nil. You can only reassign the value only if there
	// was a value already asssigned to the pointer.

	var nP *int // nP is nil
	failedUpdate01(nP)
	fmt.Println(nP) // prints nil

	// To actually update the original from inside a func when that func is being passed a pointer,
	// you need to dereference the pointer before assigning a new value to it.

	x1 := 10
	failedUpdate02(&x1)
	fmt.Println(x1) // prints 10
	update(&x1)
	fmt.Println(x1) // prints 20

	//! Pointers are a last resort!
	// Pointers should be used only when theyre absolutely necessary.

	// Dont do this:

	makeFoo01 := func(f *Foo) error {
		f.isFoo = false
		return nil
	}

	// Do this instead:

	makeFoo02 := func() (Foo, error) {
		f := Foo{
			isFoo: true,
		}
		return f, nil
	}

	fmt.Println(makeFoo01(&Foo{isFoo: false}))
	fmt.Println(makeFoo02())

	//* Diff between Maps and Slices
	// Maps are implemented as a pointer to a struct, so passing maps into a function means
	// you're copying a pointer. Therefore, modifications to a map passed to a function will
	// always be reflected in the original map.
	//! Because of this, you should be careful about passing maps into a function, or returning a map from a function.
	//? Rather than passing a map around, use a struct.

	// Slices are more complicated. Modifications to the contents of a slice passed to a function is always reflected in the original slice, but using append doesn't reflect on the original slice even if the capacity of the original slice is large enough.
	// The reason for this is that, a slice is implemented as a struct with three fields. an int field for its length, an int field for its capacity, and a pointer to a block of memory which contains the array.
	// Now, when a slice is copied to a diff variable or passed into a function, a copy of the capacity, length and pointer of the original slice is made. Updating the contents of this copied slice means updating the value at that pointer, which means updating the value of the original slice.
	// However, updating the length and/or the capacity of the copied slice doesnt reflect in the original slice because the slice being modified is a copy, and Go is a call-by-value language.
	// If the slice copy is appended to and there isnt enough capacity in the original slice for the new values, a new, bigger memnory block is allocated for the slice copy, the values are copied over, and the pointer, length, and capacity in the slice copy are updated.
}

func makePointer[T any](t T) *T {
	return &t
}

func failedUpdate01(g *int) {
	x := 10
	g = &x
}

func failedUpdate02(px *int) {
	x2 := 20
	px = &x2
}
func update(px *int) {
	*px = 20
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(fName, lName string, age int) Person {
	return Person{
		FirstName: fName,
		LastName:  lName,
		Age:       age,
	}
}

func MakePersonPointer(fName, lName string, age int) *Person {
	p := Person{
		FirstName: fName,
		LastName:  lName,
		Age:       age,
	}
	return &p
}
