package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	// Έλεγχος αν υπάρχει το flag --upper
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		n := 0

		// Μετατροπή string σε int χειροκίνητα (χωρίς strconv)
		for _, ch := range arg {
			if ch < '0' || ch > '9' {
				n = 0
				break
			}
			n = n*10 + int(ch-'0')
		}

		if n >= 1 && n <= 26 {
			var letter rune
			if upper {
				letter = rune('A' + n - 1)
			} else {
				letter = rune('a' + n - 1)
			}
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
