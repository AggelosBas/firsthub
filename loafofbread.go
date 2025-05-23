package piscine

import "github.com/01-edu/z01"

func LoafOfBread(str string) {
	// Αφαίρεσε τα κενά
	clean := ""
	for _, r := range str {
		if r != ' ' {
			clean += string(r)
		}
	}

	if len(clean) < 5 {
		msg := "Invalid Output\n"
		for _, r := range msg {
			z01.PrintRune(r)
		}
		return
	}

	i := 0
	count := 0
	for i+5 <= len(clean) {
		for _, r := range clean[i : i+5] {
			z01.PrintRune(r)
		}
		count++
		i += 6 // 5 χαρακτήρες + 1 skip
		if i+5 <= len(clean) {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
