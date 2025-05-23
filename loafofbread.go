package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) {
	// Count non-space characters
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}
	if count < 5 {
		printStr("Invalid Output\n")
		return
	}

	letterCount := 0
	skip := false
	word := ""

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skip {
			skip = false
			continue
		}
		word += string(r)
		letterCount++
		if letterCount == 5 {
			printStr(word)
			z01.PrintRune(' ')
			word = ""
			letterCount = 0
			skip = true
		}
	}

	// Τυπώνουμε newline
	z01.PrintRune('\n')
}

// Helper function to print a string with z01.PrintRune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
