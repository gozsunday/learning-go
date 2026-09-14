package main

import "fmt"

func main() {
	// In Go, there's a 'zero value' for every type. It is usually assigned when a variable is declared but not assigned a value.

	literals()
	types()
	typeConversion()
	assignVars()
	unusedVarsAndNamingVars()
}

func literals() {
	//* Literals
	// A literal is a number, character or string. There are five types of literals (the book is mentioning only four now).

	// 1. Integer Literal
	// 1234, 0b10101, 0o1723, 0x1a2b5, 10_000
	iL := 0x1a2b5

	// 2. Floating Literal
	// 1.234, 1.2e-34, 0x12.34p5 (e is used to denote exponent in decimal, while p is used to do the same in hexadecimal).
	fL := 1.234

	// 3. Rune Literal
	// A single character denoted using a single quote.
	// 'a', '\141', '\x61', '\u0061', '\U00000061' (Unicode char, 8-bit octal number, 8-bit hexadecimal number, 16-bit hexadecimal number, 32-bit Unicode numbers).
	// Characters like '\n' (new line), '\t' (tab), '\'' (single quote), and '\\' (backslash) are runes.
	rL := 'a'

	// 4. String Literal
	// Basically a string. Use backslashes to add a new line, tab, double quotes, another backslash in the string literal.
	// "what the fuck is going on, \n\"Carlos"?"
	sL := "what the \tfuck is going on, \n\"Carlos\"?"

	// There are also raw string literals and you don't need backslashes to indicate the characters mentioned above in a raw string literal.
	rSL := `what the        fuck is going on, 
"Carlos"?`

	fmt.Println(iL, fL, rL, sL, rSL)
}

func types() {
	//* Booleans
	// true or false. its zero value is false.
	var flag bool

	//* Numeric Types
	// Integers
	// There are signed and unsigned integers. Signed ones can be negative. Unsigned is always positive. There are special names for some numeric types.
	// A byte is an alias for the uint8. An int is an alias for int32 on 32-bit computers and int64 on 64-bit computers. Just avoid using int in your code.
	// There's also a uint, which follows the same rules as an int except the numbers are zero or positive.

	// Integer Operations
	// +, -, *, /, % (modulus)
	// If an integer is divided by an integer, the result is always an integer (strange)
	// To get a floating-point result from an integer division, the integers must be converted to floating-point numbers.
	// Don't divide by 0. it causes a panic (whatever tf this is).
	// any operator can be combined with `=` to modify the variable.

	var x int = 10
	x *= 3 // this multiplies x by 3 and reassigns the answer to x.

	// ==, !=, <, >, <=, >= (comparison operators)
	// bit-manipulation operators — << (shift left), >> (shift right), & (bitwise AND), | (bitwise OR), ^ (bitwise XOR), &^ (bitwise AND NOT). I don't know what these are for.

	// Floating-point
	// float32 and float64. (Always use float64 unless specifically required to use float32).
	// zero value is 0 (obvs).
	// never use floats to handle money or values that require accurate decimal representation, (i think there are fns that help with representing numbers with decimals).
	// only use them where approximate values are acceptable.
	// all operators used for ints except % (modulus) can be used for floats.
	// dividing a non-zero float by zero returns +Inf or -Inf (+ve or -ve infinity). dividing a float set to zero by zero returns NaN.
	// dont ever compare two floats using the == pr != operators. if you really need to compare, define a max allowed variance (epsilon) and check if the diff btw the two floats is greater than epsilon.
	// to really compare two floats, use the Go equivalent of the function `nearlyEqual` shown at "https://floating-point-gui.de/errors/comparison/".

	//* Strings and Runes
	// the zero value of a string is an empty string.
	// big difference from js land — strings in Go are immutable. you can reassign another string to the variable but you cant modify the string already assigned to it.
	// the `rune` type is an alias for int32

	var myFirstInitial = 'C'

	fmt.Println(flag, myFirstInitial)
}

func typeConversion() {
	//* Explicit Type Conversion
	// Go doesnt allow automatic type conversion like TS. Type conversion is done explicitly.
	// When it comes to type comversion, even numbers of different sizes need to be converted to the same size to interact.

	var x int = 10
	var y float64 = 30.5
	var sumIF float64 = float64(x) + y
	var sumFI int = x + int(y)

	// this same shit happens with different-sized integer types as well.

	var b byte = 100
	var sumIB = b + byte(x)
	var sumBI = int(b) + x

	fmt.Println(sumIF, sumFI, sumIB, sumBI)

	// because of the strictness around type conversion, you cant treat another value like 0, or "" as a boolean. You'll always have to compare using comaprison operators.
	// you can't even convert other types to boolean. thats how strict and explicit it is.

	//! literals are untyped! it is only typed when a type is assigned to it. thats why you can do something like:
	var uTX float64 = 10
	var uTY float64 = 32 * 2.5

	fmt.Println(uTX, uTY)
}

func assignVars() {
	//* var Versus :=

	var x int = 10 // the most verbose way. use `var`, an explicit type and the assignment
	var y = 20.23  // implicit typing

	var zV int // we can assign the zero value to a variable by dropping the assignment

	// we can also declare multiple variables at once using `var`
	var m1, m2 = 10, "hey"
	var mZV1, mZV2 float64

	fmt.Println(x, y, zV, m1, m2, mZV1, mZV2)

	// there's something called a declaration list
	var (
		dV1           int
		dV2           = 20
		dV3, dV4      = 30, "wtf"
		dV5, dV6, dV7 byte
	)

	fmt.Println(dV1, dV2, dV3, dV4, dV5, dV6, dV7)

	//* using the := operator
	// when using the := operator, the type is inferred.
	// `:=` is only valid inside a function. when declaring variables at the package level, the `var` must be used.

	eO1 := 100
	eO2, eO3 := true, "wagwan!"

	fmt.Println(eO1, eO2, eO3)

	// when declaring variables inside a func, always favour the `:=` operator.
	// only use `var` when trying to assign the zero value to a variable.
	// also use the `var` when trying to explicitly apply the non-default type of a value to a variable.

	// do this:
	var dT1 byte = 40
	// instead of this:
	dT2 := byte(40)

	fmt.Println(dT1, dT2)

	// never declare variables in the package block.
	// only declare constants!
}

func constantVars() {
	//* Using `const`
	// `const` in Go are a way to give a name to literals. There's no way to declare that a variable is immutable.
	// Constants in Go are quite unremarkable, unlike JS.

	const cG = "hello"

	// declaration list
	const (
		cV1      int = 10
		cV2, cV3     = "bye!", false
	)
}

func unusedVarsAndNamingVars() {
	//* Unused Variables
	// In Go, every declared local variable must be read. A compile-time error is thrown when a local variable is declared and not read.
	// the variable just has to be read once, even if it means another value is assigned to the variable without it being read again. This is valid:

	x := 100
	x = 1000
	fmt.Println(x)
	x = 400

	//* Naming a Variable
	// names of variables in Go must start with a letter or an underscore
	// underscores are rarely used in a variable name cause idiomatic Go doesnt use snake case (index_counter)
	// within a function, try to use short variable names. "The smaller the scope, the shorter the variable name should be"
	// in the package block, outside a function, use more descriptive names.
	// Go doesnt use uppercase letters with underscores for constants like in other languages, because the case of the first letter of a package level
	// declaration determines if that item is accessible outside the package or not.
}

func exercise() {
	// 1
	i1 := 20
	var f1 float64 = float64(i1)

	fmt.Println(i1, f1)

	// 2
	const value = 100
	var i2 int = value
	var f2 float64 = value

	fmt.Println(i2, f2)

	// 3
	// var b byte = 257
	// var smallI int32 = 2147483648
	// var bigI uint64 = 18446744073709551616

	// fmt.Println(b, smallI, bigI)
}
