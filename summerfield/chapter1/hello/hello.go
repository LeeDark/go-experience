// hello.go
package main

// import statement: imports three packages from the standard library.
import (
"fmt"
"os"
"strings"
)

// Go functions and methods are defined using the func keyword.
func main() {
	// short variable declaration.
	who := "World!"
	// the os.Args variable is a slice of strings and len() is the build-in function.
	if len(os.Args) > 1 {
		// we can access to slice elements using the [] index operator.
		// we set the "who" string to contain all the arguments joined up as a single string.
		who = strings.Join(os.Args[1:], " ")
	}

	// finally, we print "Hello", a space, the "who" string, and a newline.
	fmt.Println("Hello", who)
}
