package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Παίρνουμε μόνο τα arguments (χωρίς το όνομα του προγράμματος)

	for i := len(args) - 1; i >= 0; i-- { // Ανάποδα
		for _, ch := range args[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n') // Νέα γραμμή
	}
}
