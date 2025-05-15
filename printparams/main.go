package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Τα arguments χωρίς το όνομα του προγράμματος

	for _, arg := range args {
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n') // Για να πάει στην επόμενη γραμμή
	}
}
