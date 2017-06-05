// Modify the echo program to print the index and value of each of its arguments,
// one per line.Modify the echo prog ram to print the index and value of each of its arguments,
// one per line.

package main

import (
	"fmt"
	"os"
)

func main() {

	for i, args := range os.Args {

		fmt.Printf("%d: %s \n", i, args)
	}

}
