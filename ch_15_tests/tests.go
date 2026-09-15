package main

import "time"

func main() {

}

func addNumbers(x, y int) int {
	return x + y
}

func subtractNumbers(x, y int) int {
	return x - y
}

type Person struct {
	Name      string
	Age       int
	DateAdded time.Time
}

func CreatePerson(name string, age int) Person {
	return Person{
		Name:      name,
		Age:       age,
		DateAdded: time.Now(),
	}
}
