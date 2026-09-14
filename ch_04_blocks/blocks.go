package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Blocks, Shadows, and Control Structures

	blocks()
	ifElseStatements()
	forStatements()
	switchStatements()
	exercises()
}

func blocks() {
	//* BLOCKS
	// A block is any place a declaration can be made.
	// Within a function, every set of braces ({}) is a block.

	//* Shadowing Variables
	// A variable is shadowed when a variable in a particular block has the same name as a variable
	// in the containing block.

	x1 := 10
	if x1 > 5 {
		fmt.Println(x1) // 10
		x1 := 5
		fmt.Println(x1) // 5
	}
	fmt.Println(x1) // 10

	// When doing multiple assignments using ":=", we can reassign a value to an already declared
	// variable if a new variable is also being created. However, this only works if the variable
	// being reassigned was declared in the same block. If it was declared in a containing block,
	// then it becomes a shadow.

	x2 := 10
	if x2 > 5 {
		x2, y2 := 5, 20 // This second x2 is shadowing the first x2, instead of being a reassignment.
		// In clearer terms, the assignment to x2 inside the block only applies inside the block.
		fmt.Println(x2, y2) // 5, 20
	}
	fmt.Println(x2) // 10

	// Imported packages can also be shadowed
}

func ifElseStatements() {
	//* if and else
	// the biggest diff between Go and other langs in if else statements is that Go doesnt have
	// a parentheses around the condition.

	ranN := rand.Intn(10)
	if ranN == 0 {
		fmt.Println("That's too low")
	} else if ranN > 5 {
		fmt.Println("That's too big:", ranN)
	} else {
		fmt.Println("That's a good number:", ranN)
	}

	// another diff is that in Go, you can declare a variable in the condition part of the if
	// else statement and it would be scoped to just that if else statement. this is very handy.

	if ranN := rand.Intn(10); ranN == 0 {
		fmt.Println("That's too low")
	} else if ranN > 5 {
		fmt.Println("That's too big:", ranN)
	} else {
		fmt.Println("That's a good number:", ranN)
	}

	// In the example above, note that we shadowed the first ranN.
}

func forStatements() {
	//* for
	// In Go, `for` is the only looping keyword there is. it can be used in four formats

	// 1. Complete `for` Statement
	// this is the classic `for` loop.

	// this initializes the variable `i` scoped to the `for` statement, states the condition where the loop will run if `i` is smaller than 10, and then increments `i` by 1.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// the initialization can be left out if the `for` loop is based on an already declared value. the two semicolons are preserved.

	i1 := 0
	for ; i1 > 10; i1++ {
		fmt.Println(i1)
	}

	// or the increment is left out because you have a more complicated increment rule inside the loop. the two semicolons are also preserved.

	for i := 0; i > 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}

	// 2. Condition-Only `for` statement
	// here, the initialization and the increment is left out, and the semicolons are also removed.
	// this will leave behind a `for` statement that functions like a `while` loop in other
	// languages.

	i2 := 1
	for i2 < 100 {
		fmt.Println(i2)
		i2 *= 2
	}

	// 3. Infinite `for` Statement
	// here, the initialization, condition and increment are all done away with. this creates a loop
	// that will run infintely until a break keyword is used. this is Go's equivalent of the classic
	// `do while` loop.

	i3 := 0
	for {
		fmt.Println(i3)
		i3++

		if i3 > 100 {
			break
		}
	}

	// continue
	// the `continue` keyword is used to skip that particular iteration of the loop and go to the
	// next iteration.

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz!", i)
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz!", i)
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz!", i)
			continue
		}
		fmt.Println(i)
	}

	// 4. `for-range` Statement
	// this is used to iterate over elements in some of Go's built-in types. It can be used to
	// iterate over the elements/parts in strings, arrays, slices and maps. it can even be used
	// with numbers

	// with a number
	for i := range 20 {
		fmt.Println(i)
	}

	// with a slice
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println("position:", i, "value:", v)
	}

	// when iterating over numbers, arrays and slices using for range, `i` is usually used for index.
	// when iterating over maps, `k` is usually used for key.

	// if we need to access the value, without accessing the index or key, we can denote it
	// using an underscore. this tells Go to ignore the value.

	for _, v := range evenVals {
		fmt.Println("value:", v)
	}

	// if we need to access the key/index and not the value, we can simply leave it off.

	uniqueNames := map[string]bool{"steve": false, "chiugo": true, "kamharida": true, "mary": false}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	// we can also iterate over a string. i will be the index, but v will be the numeric value of
	// the letter. to get the letter, we can do `string(v)`

	sSamples := []string{"hello", "apple_π!"}
	for _, sample := range sSamples {
		for i, v := range sample {
			fmt.Println(i, v, string(v))
		}
	}

	// NOTE: in `for range` the value gotten from the compound type variable and assigned to v is
	// copied, so modifying v wont affect the compound type variable.

	oddVals := []int{1, 3, 5, 7, 9, 11, 13}
	for _, v := range oddVals {
		v *= 2
	}
	fmt.Println(oddVals) // [1 3 5 7 9 11 13]

	// Labeling `for` Statements
	// normally, the `continue` keyword applies to the current `for` loop it currently in. however it
	// is possible to use a `continue` in a for loop, and make it applying to an outer loop. this
	// can be done by labeling the loop.

	sMSamples := []string{"wtf!", "wagwan?!"}
outer:
	for _, sample := range sMSamples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'g' {
				continue outer
			}
		}
		fmt.Println()
	}
}

func switchStatements() {
	//* switch
	// just like in the if else statements, no parenthesis are put around the value being compared.
	// also like the if else statements, a variable can be declared in the switch statement and itll be scoped to all the cases of the switch statement.

	// all case clauses are contained inside a bracket. the default clause is optional. each case is a block of its own and can have multiple lines of code. no brackets are used to contain cases.

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the rigth length", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	// if a case is empty, then nothing happens. thats why in the sample above, the words "gopher", and "octopus" are not printed to the console.

	// the switch statement does have a `fallthrough` keyword but it should never be used at all. restructure your code if you find yourself needing to use it.

	// the `break` keyword can also be used inside a case of a switch, but should be used sparingly. you should only use a `break` in a case of a switch, if you want to break out of a containing `for` loop. in that scenario, the for loop must be labelled, else Go will asssume you want to just break out of the case and not the loop.

loop:
	for i := range 10 {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even!")
		case 3:
			fmt.Println(i, "is divisble by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	// Blank Switches
	// switches can also be written in such a way that the value being compared against isnt specified. its called a blank switch. a regular switch only allows you to check a value for equality. a blank switch allows you to use any boolean comarison for each case.

	names := []string{"Emeka", "Uche", "Chinero", "Ada", "Chiagoziem", "Kamharida"}

	for _, name := range names {
		switch nameLen := len(name); {
		case nameLen < 5:
			fmt.Println(name, "is a short name!")
		case nameLen > 7:
			fmt.Println(name, "is a long name!")
		default:
			fmt.Println(name, "is exactly the right length.")
		}
	}
}

func exercises() {
	// exercise 1
	intSl := []int{}
	for len(intSl) <= 100 {
		randNo := rand.Intn(100)
		intSl = append(intSl, randNo)
	}
	fmt.Println("Int Slice:", intSl)

	// exercise 2
	for _, v := range intSl {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println(v, "Six!")
			continue
		}

		if v%2 == 0 {
			fmt.Println(v, "Two!")
			continue
		}

		if v%3 == 0 {
			fmt.Println(v, "Three!")
			continue
		}

		fmt.Println(v, "Never mind!")
	}
}
