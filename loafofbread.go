package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) {
	// Αφαίρεση κενών
	cleaned := ""
	for _, r := range str {
		if r != ' ' {
			cleaned += string(r)
		}
	}

	if len(cleaned) < 5 {
		for _, r := range "Invalid Output\n" {
			z01.PrintRune(r)
		}
		return
	}

	i := 0
	for i+5 <= len(cleaned) {
		// Εκτύπωσε τους 5 χαρακτήρες
		for _, r := range cleaned[i : i+5] {
			z01.PrintRune(r)
		}
		// Εκτύπωσε space
		z01.PrintRune(' ')
		// Προχώρα 6 θέσεις (οι 5 που πήρες + 1 που αγνοείς)
		i += 6
	}

	// Αντικατάσταση του τελευταίου space με newline
	z01.PrintRune('\n')
}
