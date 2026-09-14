package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("Hello, my name is %s!\n", "Gozman")

	// `go build` creates an executable.
	// `go fmt` formats the code. Go always uses a tab for indentation. What you set it in your code editor is your cup of tea!
	// `go vet` catches syntactically valid but likely incorrect code.
}
