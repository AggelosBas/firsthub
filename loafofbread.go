package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) {
	// Αφαιρούμε τα κενά και μετράμε τους χαρακτήρες
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}

	if count < 5 {
		for _, r := range "Invalid Output\n" {
			z01.PrintRune(r)
		}
		return
	}

	word := ""
	skipNext := false
	letterCount := 0

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skipNext {
			skipNext = false
			continue
		}

		word += string(r)
		letterCount++

		if letterCount == 5 {
			for _, ch := range word {
				z01.PrintRune(ch)
			}
			z01.PrintRune(' ')
			word = ""
			letterCount = 0
			skipNext = true
		}
	}

	z01.PrintRune('\n')
}
