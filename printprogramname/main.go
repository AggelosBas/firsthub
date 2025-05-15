package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]

	// Βρίσκουμε το τελευταίο '/' για να πάρουμε το basename
	lastSlash := -1
	for i := 0; i < len(arg); i++ {
		if arg[i] == '/' {
			lastSlash = i
		}
	}

	name := arg
	if lastSlash != -1 {
		name = arg[lastSlash+1:]
	}

	// Εκτυπώνουμε χαρακτήρα-χαρακτήρα
	for _, c := range name {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
