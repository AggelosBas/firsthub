package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) {
	// Καθαρίζουμε τα spaces
	clean := ""
	for _, r := range str {
		if r != ' ' {
			clean += string(r)
		}
	}

	if len(clean) < 5 {
		for _, r := range "Invalid Output\n" {
			z01.PrintRune(r)
		}
		return
	}

	// Επεξεργασία 5 χαρακτήρων και skip 1
	for i := 0; i+5 <= len(clean); i += 6 {
		word := clean[i : i+5]
		for _, r := range word {
			z01.PrintRune(r)
		}
		// Αν υπάρχει άλλος διαθέσιμος χαρακτήρας για νέο block, βάλε space
		if i+6+4 < len(clean) {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
