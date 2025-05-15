package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Ταξινόμηση σε ASCII σειρά με απλό bubble sort
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-1-i; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Εκτύπωση κάθε argument χαρακτήρα-χαρακτήρα
	for _, word := range args {
		for _, ch := range word {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
