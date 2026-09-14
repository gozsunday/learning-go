package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// Every Go program starts from a main function. It doesnt take in any parameters or return values.

func main() {
	functions()
	calc()
	outer()
	printAnon()
	closures()
	funcAsParams()

	// Here, we pass the base to the `makeMult` func to assign to the variables, a unique func that depends on the base.
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	// We can then pass the factor to this returned func.
	for i := range 3 {
		fmt.Println(twoBase(i), threeBase(i))
	}

	// myCat()
	deferExample()
}

func functions() {
	fmt.Println(multiply(2, 5))
	fmt.Println(divNum(5, 2))

	PrintPerson(PrintPersonOpts{
		FirstName: "Gozman",
		LastName:  "Faraday",
		Age:       25,
	})

	fmt.Println(addTo(5))
	fmt.Println(addTo(5, 2, 8, 4, 3, 9, 7))
	slA := []int{1, 2, 4}
	fmt.Println(addTo(5, slA...))
	fmt.Println(addTo(9, []int{2, 5, 7}...))

	result1, remainder1, err := divAndRemainderOne(7, 4)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	fmt.Println(result1, remainder1)

	result2, remainder2, err := divAndRemainderTwo(3, 4)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	fmt.Println(result2, remainder2)

	result3, remainder3, err := divAndRemainderTwo(5, 2)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	fmt.Println(result3, remainder3)

	i := 2
	s := "Hello"
	p := person{}

	modifyFails(i, s, p)
	fmt.Println("call by value: ", i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println("maps are not call by value: ", m)

	sl := []int{1, 2, 3}
	modSlice(sl)
	fmt.Println("slices are not call by value: ", sl)
}

// Other functions can accept parameters and return values

func multiply(num int, denom int) int {
	return num * denom
}

// if two or more consecutive parameters have the same type, you can specify the type for all of them after declaring them.

func divNum(num, denom float64) float64 {
	if denom == 0 {
		return 0
	}
	return num / denom
}

// Go does not have Named and Optional parameters. But it can be imitated using a struct.

type PrintPersonOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func PrintPerson(opts PrintPersonOpts) {
	fmt.Println("first name:", opts.FirstName)
	fmt.Println("last name:", opts.LastName)
	fmt.Println("age:", opts.Age)
}

//* Variadic Input Parameters and Slices
// Functions in Go can be configured to allow any number of input parameters. We call this variadic Input Parameters. It is indicated with three dots (...) before the type, and its the last or only parameter in the function. The value of the variadic parameter inside the function is a slice of the specified type.

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out

	// this function can then called in various ways.
	// addTo(1, 2, 3, 4, 5, 6)
	// a := []int{1, 2, 4}
	// addTo(5, a...)
	// addTo(9, []int{2, 5, 7}...)
}

//* Multiple Return Values
// unlike other languages, a function in Go can return multiple values.
// This shit is amazing. Every return instance in the function must return all the values specified.
// This is commonly used in functions where an error has to be returned if an error occured. If no error occured, `nil` is returned.
// By convention, the error is always the last value (or only) value returned from a function.

func divAndRemainderOne(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

// If a function in Go returns multiple values, the returned values must be assigned to multiple variables. Trying to assign all the values to a single variable causes a compile-time error.
// If one or more of the returned values needs to be ignored, you can use the `_` since Go doesnt allow unused variables.

//* Named Return Values
// Go also allows you to name the return values. The names given to the return values, can be used as variables within the function. Named return values are initialized to their zero values when created.
// The name thats used for a named return value remains local to the function.

func divAndRemainderTwo(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}

	result, remainder = num/denom, num%denom
	return result, remainder, err // zero value for an error is nil obvs!
}

// However, using named return values can cause issues. For one, they can be shadowed. Ensure you're assigning to the return value and not to a shadow from outside the block.
// Another issue is that you don't have to return them.
// For example, passing 5 and 2 to the fn `divAndRemainderThree` will return 2 and 1 as the result and remainder instead of 20 and 30, even though they were assinged to the named return values.

func divAndRemainderThree(num, denom int) (result int, remainder int, err error) {
	// assign some values
	result, remainder = 20, 30

	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

//* Blank/Naked Returns
// When named return values are used, a feature called blank returns can also be used. Here, the keyword return can just be written without specifying the values being returned. This will return the last value assigned to the named return values.
//! NEVER USE BLANK RETURNS

func divAndRemainderFour(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

//* Functions Are Values
// Functions in Go are values, and the type of a function is built out of the `func` keyword and the types of the parameters and return values. This combination is called the signature of the function.
// Since functions are values, you can declare a function variable.

var myFuncVariable func(string) int

// The default zero value for a func value is nil.
// Having functions as values allows us to do some clever things. For example, a simple calculator using functions as values.

// exercise 1: adding division by zero error handling
func add(i, j int) (int, error) { return i + j, nil }
func sub(i, j int) (int, error) { return i - j, nil }
func mul(i, j int) (int, error) { return i * j, nil }
func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

func calc() {
	opMap := map[string]func(int, int) (int, error){
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		// As you can see, the function gotten from opMap is passed to opFunc, because FUNCTIONS ARE VALUES.
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}
}

//* Function Type Declarations
// We can use the `type` keyword to define a function, the same way we can use it to define a struct.

type opFuncType func(int, int) int

//* Anonymous Functions
// We can define Anon functions by putting the input parenthesis immediately following the func keyword. They can then be assigned to variables. This is the only way a function can be defined inside another function (also called a closure).

func outer() {
	inner := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := range 5 {
		inner(i)
	}
}

// We dont need to assign an anon func to a variable to use it. We can write them inline and call them immediately.

func printAnon() {
	for i := range 5 {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}

//* Closures
// Closures are nested functions that can access & modify variables declared in the outer functions.

func closures() {
	a := 20
	f := func() {
		fmt.Println(a)
		a = 30
	}
	f()
	fmt.Println(a)
}

//* Passing Functions as Parameters
// In Go, functions can be passed into other functions. It is used several times in the standard lib.
// An example is the sort.Slice function that takes a slice and a function that is used to sort the slice that is passed in.

func funcAsParams() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Gozman", "Faraday", 25},
		{"Christian", "Xander", 34},
		{"Saint", "Jaeromi", 18},
	}

	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}

//* Returning Functions from Functions
// Nested functions (also called closures!), can also be returned from functions.

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

//* defer
// Programs often create temp resources, like files or network connections that needs to be cleaned up. This cleanup has to happen no matter how many exit points a function has, or whether a function completed successfully or not. The keyword `defer` is used to perform this cleanup.

func myCat() {
	// check if a filename was specified, and if the argument to the program was provided.
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	// try to open the file in read-only with os.Open
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// once we're certain there's an existing file, we close the file.
	// to ensure that it runs, we use the `defer` keyword followed by the method call
	// closing the file.
	// normally, a function or method call runs immediately, but the `defer` keyword delays it
	// until the surrounding function is exited.
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

// functions, methods and closures can be used with defer.
// Multiple functions can be defered in a Go functions. They run in a last-in, first-out (LIFO) order, ie. the last declared defer runs first.
// The code in the defer function runs after the return statement.
// If input params are passed to a function being deferred, the values of the params are evaluated immediately and stored until the function runs.

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("existing:", a)
	return a
}

// The above func will print:
// existing: 30
// second: 20
// first: 10

// Now, defer functions can be used to modify the return value of the surrounding function.
// This is the best reason to use named return values.
// An example is a defer func handling DB transactions cleanup.

//

func DoSomeInserts(ctx context.Context, db *sql.DB, v1, v2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", v1)
	if err != nil {
		return err
	}

	// use tx to do more db inserts here

	return nil
}

//* Go is Call by Value
// This means that when a variable is passed as the parameter to a function, Go always makes a copy of the value of the variable.

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

// You'll see that for primitive types, and most other types like Structs, functions can't modify the values of variables passed as parameters into it.
// The behaviour is a little bit different for slices and maps.

// the reason why maps and slices act differently is because they're both implemented with pointers.

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

// exercise 2: using defer
func fileLen(n string) (int, error) {
	// try to open the file in read-only
	f, err := os.Open(n)
	if err != nil {
		return 0, err
	}
	// close file after function finishes running
	defer f.Close()

	// get file stats
	info, err := f.Stat()
	return int(info.Size()), nil
}

// exercise 3
func prefixer(prefix string) func(string) string {
	return func(base string) string {
		return prefix + " " + base
	}
}
