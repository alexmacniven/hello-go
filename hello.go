// This is our package name.
package main

// Here we have imports.
// The 'fmt' package gives us formatting functions like Print.
import (
	"fmt"
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
	// We're calling Hello() from the greetings package.
	message := greetings.Hello("Alex")
	fmt.Println(message)
}
