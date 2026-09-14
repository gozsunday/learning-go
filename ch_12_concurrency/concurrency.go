package main

import (
	"context"
	"fmt"
)

func main() {
	concurrencyPracticesAndPatterns()
}

func goroutines() {
	// Goroutines are Go's way of doing things concurrently. They're lightweight threads created
	// and managed by the Go runtime rather than the OS.
	// A goroutine is a function that runs alongside the main program and it is started with the `go`
	// keyword.

	// go doSomething()

	// In Go, it is idiomatic to launch goroutines using closures that wrap the business logic. The
	// closures take care of the concurrent business.
}

func channels() {
	// Goroutines communicate using channels. They're created using the `make()` function.

	ch1 := make(chan int)

	// Like maps, channels are reference types. When you pass a channel to a function, you're actually
	// passing a pointer to the channel. The zero value of a channel is nil.

	fmt.Println(ch1) // prints something like `0x117afc782070`

	// READING, WRITING & BUFFERING

	// the operator `<-` is used to interact with a channel.

	var b int

	a := <-ch1 // reads a value from `ch` and assigns it to `a`
	ch1 <- b   // write the value in `b` to `ch`
	fmt.Println(a)

	// Each value written to a channel can be read only once. if multiple goroutines are reading
	// from the same channel, a value written to the channel will be read by only one of them.

	// By default, channels are unbuffered. Every write to an open, unbuffered channel causes the
	// writing goroutine to pause until another goroutine reads from the same channel. Likewise, a read
	// from an open, unbuffered channel causes the reading goroutine to pause until another goroutine
	// writes to the channel. This means you cannot write to or read from an unbuffered channel without at least two concurrently running goroutines.

	// Go also has buffered channels. These channels buffer a specified number of writes without
	// blocking. if the buffer fills before a read from the channel is done, subsequent writes to the
	// channel pauses the writing goroutine until the channel is read. Basically, if the buffer is
	// filled, the channel acts like an unbuffered channel. A buffered channel is made by specifying
	// the number of buffers when creating the channel.

	ch2 := make(chan int, 10)
	fmt.Println(ch2)

	// Just like slices, the built-in fns `len` and `cap` can be used with channels. `len` to find out
	// how many values are currently in the buffer and `cap` to find out the max size of the buffer.
	// The capacity of the buffer cannot be changed. Passing an unbuffered channel to len and cap both
	// return 0.

	// `for-range` loops can also be used to read from channels.

	for v := range ch2 {
		fmt.Println(v)
	}

	// CLOSING A CHANNEL

	// After writing to a channel, it can be closed using the built-in close function.

	close(ch1)

	// Trying to write to or close an already closed channel will cause a panic. However, you can read
	// from a closed channel, and if there are still unread values in the buffered channel, they'll be
	// returned. if the channel is unbuffered or there's no more values in the channel, the zero value
	// for the channel's type will be returned.

	// A reliable way to check if a channel is open or not when reading from it is using the comma-ok
	// idiom. If ok is set to false, even tho there's a value, the channel is closed.

	v, ok := <-ch1
	if ok {
		fmt.Println(v)
	}

	// Closing a channel is the responsibility of the goroutine that writes to the channel. Also, it's
	// not necessary unless there's a goroutine waiting for it to close (such as one reading from a
	// for-range loop). Go's runtime can detect open channels that are no longer reference and garbage
	// collect them.

	// There are scenarios where multiple goroutines are writing to the same channel. Now, since its
	// the repsonsibility of the writing goroutine to close the channel, this becomes complicated since
	// there are multiple writing goroutines. Attempting to close an already closed channel will cause
	// a panic, and since there are multiple writing goroutines, we could have a scenario where one of
	// these goroutines try to close a channel that has already been closed. This will cause a panic.
	// To prevent this, we use a `sync.WaitGroup`.
}

func selectInGoroutines() {
	// The `select` keyword is a little bit like `switch`. It allows a goroutine to read from or write
	// to a set of multiple channels.

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	x := 12

	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	case ch3 <- x:
		fmt.Println("wrote", x)
	case <-ch4:
		fmt.Println("got value on ch4, but ignored it")
	}

	// if multiple cases have channels that can be read from or written to, `select` picks randomly from
	// any of its cases that can go forward; order is unimportant. This is very different from how a
	// `switch` statement works (it always picks the first case that resolves to true).
}

func concurrencyPracticesAndPatterns() {
	//* Keep APIs Concurrency-Free
	// never allow go code concerning to be present in your API. unless you're making a library that has
	// a concurrency helper of course. this simply means that stuff like channels should never be
	// exported.

	//* Goroutines, `for` Loops, and Varying Variables
	// anytime a closure depends on a variabele whose value must change, you must pass the value into
	// the closure to make sure a unique copy of the variable is made. this applies for when using
	// closures for goroutines as well.

	example01()

	//* Always Clean Up Goroutines
	// whenever you launch a goroutine, you must make sure that goroutine will eventually exit.
	// an easy way to do this is to use Context tp terminate the goroutine.

	example02()
}

func example01() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for range len(a) {
		fmt.Println(<-ch)
	}
}

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range max {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	return ch
}

func example02() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
