package main

import (
	"errors"
	"fmt"
)

func main() {
	Generics()
	Exercises()
}

func Generics() {
	//* Generics in Structs
	// Generics is like a type params. They allow us to write functions or types that work with various types without sacrificing type safety or resorting to `any`.

	// Generics in Structs are placed in square brackets immediately after the Struct name, and then used as a type in the Struct fields. The letter `T` as the name of the type is commonly used.

	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	// intStack.Push("wtf") // this will panic because intStack was initialized using int as T.
	v, ok := intStack.Pop()
	fmt.Println(v, ok)
	fmt.Println(intStack.Contains(10))
	fmt.Println(intStack.Contains(5))

	//* Generics in Functions
	// Generics in functions are placed in square brackets immediately after the func name, and then used as a type in the func. The letter `T` as the name of the type is commonly used.

	words := []string{"One", "Potato", "Two", "Potato", "Three", "Potato"}

	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)

	lengths := Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)

	sum := Reduce(lengths, 0, func(acc, val int) int {
		return acc + val
	})
	fmt.Println(sum)

	//* Generics and Interfaces
	// Any interface can be used as a type constraint in a generic.
	// For example, say we want to make a type that holds any two values of the same type, as
	// long as the type implements fmt.Stringer.

	type Pair[T fmt.Stringer] struct {
		Val1 T
		Val2 T
	}

	// We can also pass type params to interfaces. Eg, here's an interface with a method that
	// compares against a value of the specified type and returns a `float64`. It also embeds
	// `fmt.Stringer`.

	type Differ[T any] interface {
		fmt.Stringer
		Diff(T) float64
	}

	//* Type Elements
	// We can compose one or more type terms within an interface and use it as a type param to
	// denote the types that are allowed

	//? Go to the `Integer` interface

	var aUint uint = 18_446_744_073_709_551_615
	var bUint uint = 9_223_372_036_854_775_808
	fmt.Println(DivAndRemainder(aUint, bUint))

	// Now, if we try to pass a user-defined type thats still an Integer underneath to the above
	// function, a panic will be triggered. To prevent this, put a `~` before the type terms.

	type MyInt int
	var myA MyInt = 10
	var myB MyInt = 20
	fmt.Println(DivAndRemainder(myA, myB))

	// Type elements also constrain which constants/literal can be assigned to variables of the
	// generic type. The literal has to be valid for all the type terms in the type element.

	// For example, trying to add 1000 to a variable with a type of `Integer` causes an error
	// because it cant be assigned to `int8` which is one of the type terms.

	fmt.Println(PlusOneHundred(1))
}

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		// to return the zero value of whatever type is passed as T, just declare a variable with
		// type T and no assingment. this will always return the zero value of that type.
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:(len(s.vals) - 1)]
	return top, true
}
func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false

	// can all be replaced by `slices.Contains(s.vals, val)`
}

// Map items in a slice from one type to another using a mapping function
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))

	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce a slice into a single value using a reduce function
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer

	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// FIlter values from a slice using a filter function
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T

	for _, v := range s {
		if b := f(v); b {
			r = append(r, v)
		}
	}
	return r
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func DivAndRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

//! INVALID!
// This will cause and error because `in` could be of type `int8` and 1_000 cannot be
// assigned to `int8`.

// func PlusOneThousand[T Integer](in T) T {
//   return in + 1_000
// }

//* VALID
// This will work since 100 can eb assigned to an 8-bit integer.

func PlusOneHundred[T Integer](in T) T {
	return in + 100
}
