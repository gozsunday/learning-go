package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	chiLeague := League{
		Name: "Chi League",
		Teams: map[string]Team{
			"Spain": {
				Name:    "Spain",
				Players: []string{"Player1", "Player2"},
			},
			"Bolivia": {
				Name:    "Bolivia",
				Players: []string{"Player1", "Player2"},
			},
			"USA": {
				Name:    "USA",
				Players: []string{"Player1", "Player2"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2"},
			},
			"Thailand": {
				Name:    "Thailand",
				Players: []string{"Player1", "Player2"},
			},
		},
		Wins: map[string]int{},
	}

	chiLeague.MatchResult("Spain", 4, "Nigeria", 2)
	chiLeague.MatchResult("Bolivia", 1, "USA", 3)
	chiLeague.MatchResult("Nigeria", 3, "Thailand", 2)
	chiLeague.MatchResult("Spain", 3, "Bolivia", 2)
	chiLeague.MatchResult("Nigeria", 1, "USA", 0)
	chiLeague.MatchResult("Thailand", 1, "Bolivia", 2)
	chiLeague.MatchResult("Bolivia", 3, "Nigeria", 4)
	chiLeague.MatchResult("USA", 2, "Spain", 3)
	chiLeague.MatchResult("Thailand", 3, "Spain", 5)
	chiLeague.MatchResult("USA", 4, "Thailand", 1)

	results := chiLeague.Ranking()
	fmt.Println(results)

	RankPrinter(&chiLeague, os.Stdout)
}

func types() {
	//* Types
	// When declaring types, the keyword `type` is written first, followed by the name of the type followed by the primitve or composite type being declared.

	type Score int
	type Converter func(string) Score
	type TeamScores map[string]Score

	type Car struct {
		Name     string
		Maker    string
		YearMade int
	}

	// Types can be defined at any block level but they must be accessed from within that scope.
}

// ———— METHOD DEFINITIONS START ————

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type User struct {
	Name string
}

func (u *User) Greet() string {
	return "Hello!"
}

func (u *User) CallName() string {
	return fmt.Sprintf("Hello, %s", u.Name)
}

func (u *User) SafeCallName() string {
	if u == nil {
		return "Hello, stranger"
	}
	return fmt.Sprintf("Hello, %s", u.Name)
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// ———— METHOD DEFINITIONS END ————

func methods() {
	//* Methods
	// The methods for a user-defined type are always defined at the package block level.
	// Method declarations look like function declarations, with one addition: the receiver specification.
	// The receiver specification appears between the func keyword and the name of the method.
	// By convention, the receiver name is a short abbreviation of the type's name, usually its first letter.

	// Methods differ from functions in the sense that they can only be defined at the package block level, while functions can be defined in any block.

	p := Person{
		FirstName: "Gozman",
		LastName:  "Faraday",
		Age:       25,
	}
	output := p.String()
	fmt.Println(output)

	// Methods can have pointer receivers or value receivers. If the method modifies the receiver,
	// a pointer receiver must be used. If the method needs to handle nil instances, then a
	// pointer receiver must be used. If the method doesn't modify the receiver, then a value
	// receiver can be used.

	// It's common practice to use pointer receivers in all methods of a type if even one of the
	// methods uses pointer receivers. It's done to be consistent.

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	//* Functions or Methods: When to choose
	// Use functions when the logic is more general or operates using multiple types.
	// Use methods when the logic is connected tightly to the type.

	//* Using Methods for `nil` Instances
	// In Go, methods can be called on nil pointers. If the method doesnt access any field of the struct the pointer is referencing, then everything will work as expected. If the method however, accesses the fields of the referenced struct, a panic will occur.

	var u *User = nil

	fmt.Println(u.Greet()) // no panic happens; prints `Hello!`

	// fmt.Println(u.CallName()) // panic happens, coz struct fields were accessed!

	// A nil check can be added to methods to prevent a panic.

	fmt.Println(u.SafeCallName()) // prints `Hello, stranger` since its a nil pointer

	//* Method Expression
	// Methods are usually called by creating an instance of the type and then invoking its method:

	myAdder := Adder{start: 2}
	fmt.Println(myAdder.AddTo(5)) // prints 7

	// The method can also be assigned to a variable or passed to a parameter of a function.

	f1 := myAdder.AddTo
	fmt.Println(f1(13)) // prints 15

	// It can be called like a closure, in which it is called a Method Expression.

	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 17)) // prints 19

	// When using a Method Expression, the receiver for the method is always the first parameter.

	//* Type Declarations Aren’t Inheritance
	// A user-defined type can also be based on another user-defined type.
	// However, when this happen, methods arw not inherited. The second user-defined type only gets
	// the data structure of the first one.

	type Score int
	type HighScore Score

	var i int = 30
	var s Score = 10
	var hs HighScore = 20

	// hs = s // compilation error!
	// s = i // compilation error!
	s = Score(i)      // ok
	hs = HighScore(s) // ok

	fmt.Println(hs)
}

func iotaInGo() {
	//* iota
	// Go doesnt have an enumeration type. Instead it has `iota`, which lets you assign an
	// increasing value to a set of constants.
	// When using the iota, the practice is to first define a type based on int that'll represent all the valid values:

	// Next, use a const block to define a set of values for the type:

	type MailCategory int

	const (
		Uncategorized MailCategory = iota
		Personal
		Social
		Advertisements
	)

	// the first constant in the const block has the type specified and its value is set to `iota`.
	// other constants in the const block do not have the type specified.
	// when the Go compiler sees this, it gives the same type and value to the subsequent constants.
	// therefore, the value that is given to all the constants is `iota`.
	// `iota` starts from 0, and increments by 1 with each constant in the const block.
	// when a new const block is created, `iota` starts from 0.

	// the value of iota increaments for each constants in the const block regardless of the value assigned to each constant.

	const (
		Field1 = 0        // 0
		Field2 = 1 + iota // 2; since iota for Field2 is 1
		Field3 = 20       // 20
		Field4            // 20; since no value is explicitly assigned and preceding constant is 20
		Field5 = iota     // 4; since iota for Field5 is 4
	)

	fmt.Println(Field1, Field2, Field3, Field4, Field5)

	// only use iota-based enumeration when you care about differentiating a set of values,
	// and dont particularly care what the value is behind the scenes.
}
