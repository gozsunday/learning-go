package main

import (
	"fmt"
	"strconv"
)

type IntOrFLoat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Exercises() {
	var i PrintInt = 12
	PrintIt(i)

	var f PrintFloat = 23.42
	PrintIt(f)
}

func Double[T IntOrFLoat](base T) T {
	return base * 2
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func PrintIt[T Printable](t T) {
	fmt.Println(t)
}
