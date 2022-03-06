// This is our package name.
package main

// Here we have imports.
import (
	// The 'fmt' package supports formatting functions like Print.
	"fmt"
	// The 'log' package supports log messages and configs.
	"log"
	// greetings is currently stored on our local disk.
	// For development purposes we can create an alias that will
	// point "example.com/greetings" at the local package instead
	// of trying to search the URL-type address for it.
	// We acheive this with;
	//   $ go mod edit -replace example.com/greetings=../greetings
	// assuming the greetings package is at ../greetings.
	// Then sync;
	//   $ go mod tidy
	"example.com/greetings"
)

// When you run a package, it's main function will be called.
func main() {
	// Set a logging prefix.
	log.SetPrefix("greetings: ")
	// Set a flag to disable the time being printed.
	log.SetFlags(0)
	// Create a slice of string names using the slice shorthand.
	names := []string{"Alex", "Jake", "Sarah"}
	// We're calling Hellos() from the greetings package.
	// And we supply the names slice as an argument.
	// We'll also capture any errors returned.
	message, err := greetings.Hellos(names)
	// If an error is returned we log it with Fatal so the program exits.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
