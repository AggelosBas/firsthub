package student

func LoafOfBread(str string) string {
	// Αν η είσοδος έχει λιγότερους από 5 μη-space χαρακτήρες
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}
	if count < 5 {
		return "Invalid Output\n"
	}

	result := ""
	word := ""
	skipNext := false
	nonSpaceCount := 0

	for _, r := range str {
		if skipNext {
			skipNext = false
			continue
		}

		if r == ' ' {
			continue
		}

		word += string(r)
		nonSpaceCount++

		if nonSpaceCount == 5 {
			result += word + " "
			word = ""
			nonSpaceCount = 0
			skipNext = true
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
