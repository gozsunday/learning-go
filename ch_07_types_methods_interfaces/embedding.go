package main

import "fmt"

// ———— METHOD DEFINITIONS START ————

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return m.Reports
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

// ———— METHOD DEFINITIONS END ————

func embedding() {
	// while Go doesnt have inheritance, it has embeddings for composition and promotion.
	// Note that `Manager` has a field of type `Employee`, but no name is assigned to that field.
	// This makes `Employee` an embedded field.
	// Any fields or methods belonging to the type of an embedded field are promoted to the
	// containing struct and can be invoked directly on it.

	m := Manager{
		Employee: Employee{
			Name: "Saint Jaeromi",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID)            // Manager gets the ID field of Employee
	fmt.Println(m.Description()) // and the Description method of Employee

	// any type can be embedded within a struct.

	// if the embedded struct has a field or method with the same name as a field or method of
	// the containing struct, the embedded struct type will then have to be used to refer to the shadowed fields or methods.

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}

	fmt.Println(o.X)       // prints 20
	fmt.Println(o.Inner.X) // prints 10

	//! NOTE: Embedding isn't Inheritance. You cannot pass an embedded string where a containing
	//! struct is expected.

	// You can't assign a variable of type Outer to that of type Inner.
	// To access the Inner field in Outer, you need to pass it explicitly.

	// var innerV Inner = o // compile error
	var innerV Inner = o.Inner // OK

	fmt.Println(innerV)
}
