package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Composite Types — Arrays, Slices, Maps and Structs.

	array()
	slicesInGo()
	stringsRunesBytes()
	mapsInGo()
	structs()
	exercises()
}

func array() {
	//* ARRAYS
	// Arrays in Go are too rigid. They're rarely used.

	// To declare an array, you can specify the size of the array and the type of the elements in the array.
	// Once an array has been declared, the size and the type cannot be changed. Like I said, it's quite rigid.
	// When declared this way, the elements of the array, will have the zero value of the assigned type.

	var a1 [3]int // [0, 0, 0]

	// The value for the arrays can also be specified when declaring them.

	a2 := [5]int{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]

	// If some elements need to have the zero value of the assigned type, the nonzero elements can be specified.

	a3 := [5]int{1, 2: 24, 4: 100} // [1, 0, 24, 0, 100]

	// when declaring an array with array literals, the size of the array can be replaced with `...`. The size is implied from the array literal.

	a4 := [...]float64{1.2, 3.4, 5.6} // [1.2, 3.4, 5.6]

	// Two arrays are the same if they're of the same type, size and have the same values.

	// To create multidimensional arrays, we can do:

	var mA1 [3][4]int // An array of 3 arrays of 4 strings.

	// Array elements in Go are read and written using the bracket notation.
	// You can't read or write past the end of the array or use a negative index.

	a5 := mA1[2]

	// Arrays in Go have some limitations. The size of the array is considered part of the type of the array. This has certain consequences. You can't use variables to specify the size of an array. An array of [3]int is diff from an array of [4]int. They cannot even be type converted into one another. You also cannot assign arrays of diff sizes to the same variable.

	fmt.Println(a1, a2, a3, a4, mA1, a5)
}

func slicesInGo() {
	//* SLICES
	// Slices are different from arrays because the size isn't a part of it's type. It can also grow/shrink as needed.

	s1 := []int{1, 2, 3} // [1, 2, 3]

	// Specific indices can also be specified as nonzero just like arrays.

	s2 := []float64{1.2, 4: 22.5, 10, 8: 9.81} // [1.2, 0, 0, 0, 22.5, 10, 0, 0, 9.81]

	// We can also make multidimensional slices just like in arrays

	var mS1 [][]int

	// Elements in a slice is also read or written using the bracket notation.
	// Just as in arrays, you can't read or write past the end of the slice or use a negative index.

	s2[1] = 5.8 // [1.2, 5.8, 0, 0, 22.5, 10, 0, 0, 9.81]

	// Slices are also diff from Arrays in certain ways

	// The zero value of a slice declared without a literal (ie. an empty slice) is `nil`.

	var eS []string

	// Two slices cannot be compared to each other. Doing it will throw an error.
	// A slice can only be compared with `nil`.
	// A slice with a value of `nil` means the slice is empty.

	// We can compare slices using the `slices` package from the standard library.
	// If the slices both have the same value, type and size, the `slices.Equal` func will return true.

	s3 := []int{1, 2, 3}
	sIE := slices.Equal(s1, s3)

	// The length of a slice can be gotten using the func `len()`. Passing a `nil` slice to len returns 0.

	var s4 []int

	fmt.Println(s1, s2, mS1, eS == nil, sIE, len(s3), len(s4))

	//* append
	// We can append elements of a valid type to a slice of that type.

	s4 = append(s4, 10)

	// we can append more than one element at a time

	s4 = append(s4, 20, 30)

	// we can also append elements from another slice by using `...`.

	s4 = append(s4, s3...)

	//* Capacity
	// Each element in a slice is assigned to consecutive memory locations. Every slice has a capacity. The capacity of a slice is the number of consecutive memory locations reserved.
	// We can check the capacity of a slice using `cap()`. The capacity of an array is always equal to the length of that array. The capacity of a slice can be greater than its length.

	fmt.Println("s4 value, length and capacity:", s4, len(s4), cap(s4))

	//* make()
	// There might be times when you know the number of things you want to put into a slice. While, its a good thing that slices can grow, its far more efficient to create them with the correct initial capacity.
	// We can do that using `make()`

	sLC1 := make([]int, 5) // this creates a slice of length 5. The zero value isn't `nil`. [0, 0, 0, 0, 0]

	sLC2 := make([]float64, 4, 10) // a slice of length 4, and capacity 10.

	// NOTE: if you make a slice using `make`, appending elements will append them to the end of the already created slice. Appending always increase the length.

	fmt.Println(sLC1, sLC2)

	//* Emptying a Slice
	// We can turn all the elements of a slice to the corresponding zero value using the `clear()` function.

	clear(s4)
	fmt.Println("emptied slice", s4)

	//* Slicing Slices
	// We can slice out Slices from other slices, using the notation slice[a:b], where a is the index of the first element of the new slice and b is the index + 1 of the last element of the new slice.
	// a or/and b can be left out. When a is removed [:b], the new slice is taken from the first element to the element just before the element with index of b. When b is removed [a:], the new slice is taken from the element of index a to the last element of the old slice. When both are removed [:], the new slice is a copy of the old slice.

	s5 := []string{"a", "b", "c", "d"}

	s5S1 := s5[1:3]
	s5S2 := s5[:2]
	s5S3 := s5[1:]
	s5S4 := s5[:]

	fmt.Println(s5, s5S1, s5S2, s5S3, s5S4)

	// When you create a slice from another slice, a copy isn't being made. Rather, you get two variables that are sharing the same memory. This means that a change to an original slice, will affect all slices that share the same memory (was sliced out from that original slice).

	s6X := []string{"a", "b", "c", "d"}
	s6Y := s6X[:2] // ["a", "b"]
	s6Z := s6X[1:] // ["b", "c", "d"]
	s6X[1] = "y"
	s6Y[0] = "x"
	s6Z[1] = "z"

	fmt.Println("X:", s6X) // ["x", "y", "z", "d"]
	fmt.Println("Y:", s6Y) // ["x", "y"]
	fmt.Println("Z:", s6Z) // ["y", "z", "d"]

	//* copy
	// this is used to create a new slice that's independent from the original slice being copied. it takes in the destination slice and the source slice, and returns the number of elements successfully copied. The number of elements copied is restricted by the length of the smaller slice.

	s7a := []int{1, 2, 9, 8, 5, 7}
	s7b := make([]int, 4)
	noOfCopiedEl := copy(s7b, s7a) // 4

	// a subslice of a slice can also be copied to a slice
	s7c := make([]int, 5)
	copy(s7c, s7a[3:])

	fmt.Println("No. of copied elements", noOfCopiedEl)
	fmt.Println(s7c)

	// you can use `copy()` with arrays by taking a slice of the array

	s7d := []int{1, 2, 3, 4}
	a1a := [4]int{9, 8, 7, 6}
	copy(s7d, a1a[:])
	copy(a1a[:], s7d)

	//* Convert Arrays to Slices
	// we can convert arrays to slices by just taking a slice of the array.

	a2a := [...]int{1, 2, 3}
	s8a := a2a[:]

	// NOTE: since we're taking a slice from the array, note that the appropriate elements will share the same memory. to prevent that, use `copy()`

	a2b := [...]int{8, 9}
	s8b := []int{1, 2}
	copy(s8b, a2b[:])

	fmt.Println(s8a, s8b)

	//* Converting Slices into Arrays
	// type conversion can be used to convert a slice into an array. this copies the data to a new memory.

	s9a := []int{2, 4, 7}
	a3a := [3]int(s9a)

	fmt.Println(a3a)
}

func stringsRunesBytes() {
	// slicing notation can be used with strings to create substrings. A single rune can also be gotten from a string using the bracket notation just like in arrays and slices. be careful be with slicing strings tho. some characters can be more than 1 byte long. since its usuallly utf-8 (except in 0x), the number of bytes in a character in a string can range from 1 to 4. this is commonly seen in emojis.
	// the `len()` func can be used to find out the number of bytes (not characters) in a string.
}

func mapsInGo() {
	//* MAPS
	// slices are for sequential data (same thing you would use arrays for in JS land).
	// maps on the other hand are for when you want to associate one value with another. it's still kind of like an slice but instead of just values, we have key-value pairs here.

	// the zero value for a map is nil. you can only read a nil map. trying to write to a nil map causes a panic.

	var nilMap map[string]int
	var m1a = map[string]int{
		"wtf":   3,
		"hello": 5,
	}

	fmt.Println("Nil Map:", nilMap)
	fmt.Println(m1a)

	// empty maps can be created using the `:=` operator and a map literal. it's not the same as a nil map.
	// reading and writing to an empty map is allowed.

	m2a := map[int]string{}
	m2b := map[string]bool{
		"isBeaut": true,
		"isTall":  false,
	}

	fmt.Println("empty map:", m2a)
	fmt.Println(m2b)

	// `make()` can also be used to make a map. apparently, maps made with `make()` still have a length of 0, and can be grown past the specified size, so i dont see why I'd use this honestly.

	m3a := make(map[int]string, 5) // this will make an empty map, I think

	fmt.Println(m3a)

	// passing a map to `len()` returns the number of key-value pair in that map.

	//* Reading and Writing a Map
	totalWins := map[string]int{}
	totalWins["Anomander"] = 12
	totalWins["Christian"] = 7
	totalWins["Whiskeyjack"] = 9
	totalWins["Laseen"] = 6

	fmt.Println("Anomander score:", totalWins["Anomander"])
	fmt.Println("Christian score:", totalWins["Christian"])

	totalWins["Laseen"]++
	fmt.Println("Laseen score:", totalWins["Laseen"])

	// i think it's all pretty explanatory. now, a little note here: if you try to get the value assigned to a map key that was never set, you'll get the zero value of the map's value type. it's pretty neat.

	fmt.Println("Apsalar score:", totalWins["Apsalar"]) // will return 0

	// there are times however, when we need to know if a key is in a map. we can do that using the "comma-ok" idiom.

	m4a := map[string]int{
		"one": 1,
		"two": 0,
	}

	v, ok := m4a["one"]
	fmt.Println(v, ok) // 1, true

	v, ok = m4a["two"]
	fmt.Println(v, ok) // 0, true

	v, ok = m4a["three"]
	fmt.Println(v, ok) // 0, false

	//* Deleting from Maps
	// key-value pairs can be removed using the `delete()` func. it doesnt return anything.

	delM := map[string]int{
		"one": 1,
		"two": 2,
	}

	delete(delM, "one")
	fmt.Println(delM)

	//* Emptying a Map
	// the `clear()` func is used to empty a map, just as it is used in slices. unlike slices tho, in maps, it sets the length to 0.

	clM := map[int]string{
		1: "one",
		2: "two",
	}
	clear(clM)
	fmt.Println(clM)

	//* Comparing Maps
	// Go provides a `maps` package that contain helper funcs for comparing two maps

	comM1 := map[int]string{
		1: "one",
	}
	comM2 := map[int]string{
		1: "one",
	}

	fmt.Println(maps.Equal(comM1, comM2))

	//* Using Maps as Sets
	// Go doesnt have inbuilt sets but we can use maps to emulate sets.

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet)) // 11, 8
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])

	// even though vals is 11 in length, len(intSet) returned 8, because you cant have duplicate items in a set.

	if intSet[1] {
		fmt.Println("100 is in the set!")
	}
}

func structs() {
	//* STRUCTS
	// Maps are cool for some stuff but they have limitations. We can't constrain a map to allow only certain keys. Also, all values in a map must be of the same type. For these reasons, maps are not an ideal way to pass data from func to func.

	// Structs do not have these limitations. They are the true equivalent of objects in JS land.

	type Person struct {
		name string
		age  int
		pet  string
	}

	var joe Person
	max := Person{}

	fmt.Println("zero-value structs:", joe, max)

	// when creating a non-empty struct, there are two styles that can be used.
	// in the first, the struct literal is a comma-separated list of values. here, every field must have a value, and they must be arranged the way they're arranged in the struct type.

	jane := Person{
		"Jane",
		26,
		"none",
	}

	// in the second style, the names of the fields are used to specify the values. we can omit fields and arrange them as we see fit. i prefer this.

	eve := Person{
		age:  25,
		name: "Eve",
	}

	// a field in a struct is read and written to using dot notation (duhhh!).

	eve.name = "Evelyn"

	fmt.Println(jane)
	fmt.Println(eve)

	//* Anon Structs
	// we can declare a variable as a struct without first giving the struct type a name.

	var city struct {
		name    string
		country string
		popl    int
	}

	city.name = "Lagos"
	city.country = "Nigeria"
	city.popl = 20_000_000

	// i prefer this.

	pet := struct {
		name string
		kind string
	}{
		name: "Baby",
		kind: "Dog",
	}

	fmt.Println(city.name, city.country, city.popl)
	fmt.Println(pet)
}

func exercises() {
	exerciseNo1()
	exerciseNo3()
}

func exerciseNo1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	gSl1 := greetings[:2]
	gSl2 := greetings[1:4]
	gSl3 := greetings[3:]

	fmt.Println(greetings, gSl1, gSl2, gSl3)
}

func exerciseNo3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	john := Employee{
		"John",
		"Obi",
		12,
	}

	esther := Employee{
		firstName: "Esther",
		lastName:  "Jachi",
		id:        01,
	}

	var nnenna Employee

	nnenna.firstName = "Nnenna"
	nnenna.lastName = "Igwe"
	nnenna.id = 23

	fmt.Println(john, esther, nnenna)
}
