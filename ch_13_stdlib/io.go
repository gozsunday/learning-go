package main

import (
	"fmt"
	"io"
	"strings"
)

func ioInGo() {
	//* io
	//  the `io` package contain two very important interfaces. the `io.Reader` and the `io.Writer`
	// interfaces, and they both define a single method.

	type Reader interface {
		Read(p []byte) (n int, err error)
	}
	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	// the `Write` method takes in a slice of bytes and returns the number of bytes written and an
	// error if anything goes wrong.
	// the `Read` method also takes in a slice of bytes. It also returns the number of bytes
	// written, which might seem strange at first.

	s := "the quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)

	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(counts)

	// look at countLetters. you can see that to do a read, a buffer is created and then reused on
	// every call to `r.Read`. this allows for using a single memory allocation to read from a
	// potentially large data source.

	// the n value returned from `r.Read` is then used to get a subslice of the buf slice so as to
	// iterate over it and process the data that was read.

	// finally, we know that we're done reading from r, when the err returned from `r.Read` is io.
	// EOF (which really isn't an error). this means when dealing with r.Read
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func timeInGo() {

}
