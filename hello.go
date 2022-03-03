// This is our package name.
package main

// Here we have imports.
// The 'fmt' package gives us formatting functions like Print.
import (
	"fmt"

	"rsc.io/quote" // This is an external package.
)

// When you run a package, it's main function will be called.
func main() {
	// We're calling Go() from the quote package.
	fmt.Println(quote.Go())
}
